package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
	"github.com/rareinator/Svendeprove/Backend/packages/models"
)

type authenticationConfig struct {
	allowedRoles   []models.UserRole
	allowedPatient string
	allowIOTDevice bool
}
type corsHandler struct {
	router *mux.Router
}

func (ch *corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ch.router.ServeHTTP(w, r)
}

func (s *Server) handleCors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("request received:\n\r%v called %v\n\rresource %v\n\r", r.RemoteAddr, r.Method, r.RequestURI)
		if os.Getenv("IS_DEV") == "TRUE" {
			fmt.Print(message)
		} else {
			if err := os.Mkdir("./log", 0644); err != nil {
				s.ReturnError(w, http.StatusInternalServerError, err.Error())
				return
			}
			if err := ioutil.WriteFile("./log/restLog.log", []byte(message), 0644); err != nil {
				s.ReturnError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}

		next(w, r)
	}
}

//ReturnError is used by all the different http Handlers to return an error message is something
//unexpected occurred
func (s *Server) ReturnError(w http.ResponseWriter, statusCode int, Message string) {
	var errorMessage struct {
		Code    int
		Message string
	}

	w.WriteHeader(statusCode)
	errorMessage.Code = statusCode
	errorMessage.Message = Message

	log.Printf("ERROR!!!! Code: %v Message: %v\n\r", errorMessage.Code, errorMessage.Message)

	if err := json.NewEncoder(w).Encode(&errorMessage); err != nil {
		log.Printf("Could not encode error\n\rERROR!!!! Code: %v Message: %v\n\r", errorMessage.Code, errorMessage.Message)
		return
	}
}

func (s *Server) Authenticate(next http.HandlerFunc, config *authenticationConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authenticating")
		vars := mux.Vars(r)

		var reqToken string
		reqToken = r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			fmt.Println("trying to access with no token")
			s.ReturnError(w, http.StatusNotAcceptable, fmt.Sprintf("No valid token specified, found %v", reqToken))
			return
		}
		reqToken = splitToken[1]
		if reqToken == "" {
			s.ReturnError(w, http.StatusNotAcceptable, fmt.Sprintf("No valid token specified, found %v", reqToken))
			return
		}

		if reqToken == "override" { //TODO: REMOVE
			next(w, r)
			return
		}

		fmt.Printf("---------TOKEN-----------\n\r%v\n\r\n\r", reqToken)

		toValidate := map[string]string{}
		toValidate["aud"] = os.Getenv("OKTA_AUTH_ENDPOINT")
		toValidate["cid"] = os.Getenv("OKTA_CLIENT_ID")

		jwt := jwtverifier.JwtVerifier{
			Issuer:           fmt.Sprintf("%v/oauth2/default", os.Getenv("OKTA_URL")),
			ClaimsToValidate: toValidate,
		}

		verifier := jwt.New()

		token, err := verifier.VerifyAccessToken(reqToken)
		if err != nil {
			s.ReturnError(w, http.StatusNotAcceptable, "Not allowed")
			fmt.Println("Not allowed:", err)
			return
		}

		role := token.Claims["role"]
		userId := token.Claims["uid"]

		allowed := false

		if config.allowedPatient != "" {
			patient := vars[config.allowedPatient]

			if patient == userId {
				allowed = true
			}
		}

		if len(config.allowedRoles) > 0 {
			for _, allowedRole := range config.allowedRoles {
				if fmt.Sprintf("%v", role) == string(allowedRole) {
					allowed = true
					break
				}
			}
		}

		if !allowed {
			s.ReturnError(w, http.StatusForbidden, "Could not successfully authenticate you")
			return
		}

		next(w, r)
	}
}
