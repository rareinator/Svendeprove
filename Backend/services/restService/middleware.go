package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
	"github.com/rareinator/Svendeprove/Backend/packages/models"
)

type authenticationConfig struct {
	allowedRoles        []models.UserRole
	allowedPatient      string
	allowRelatedPatient bool
	allowIOTDevice      bool
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

func (s *server) handleCors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received:")
		fmt.Printf("%v called %v\n\rresource %v\n\r", r.RemoteAddr, r.Method, r.RequestURI)

		if r.Method == "POST" {
			fmt.Printf("Body: \b\r%v\n\r", r.Body)
		}

		next(w, r)
	}
}

func (s *server) authenticate(next http.HandlerFunc, config *authenticationConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authenticating")
		vars := mux.Vars(r)

		var reqToken string
		reqToken = r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			fmt.Println("trying to access with no token")
			s.returnError(w, http.StatusNotAcceptable, fmt.Sprintf("No valid token specified, found %v", reqToken))
			return
		}
		reqToken = splitToken[1]
		if reqToken == "" {
			s.returnError(w, http.StatusNotAcceptable, fmt.Sprintf("No valid token specified, found %v", reqToken))
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
			s.returnError(w, http.StatusNotAcceptable, "Not allowed")
			fmt.Println("Not allowed:", err)
			return
		}

		role := token.Claims["role"]
		username := token.Claims["username"]

		allowed := false

		if config.allowedPatient != "" {
			patient := vars[config.allowedPatient]

			if patient == username {
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

		// 	if config.allowIOTDevice {
		// 		// fmt.Println(response)
		// 		// fmt.Println(response.IOTDeviceId)
		// 		// if response.IOTDeviceId != 0 {
		// 		// 	allowed = true
		// 		// }

		// 		//TODO: find out what to do here properly
		// 		allowed = true
		// 	}

		if !allowed {
			s.returnError(w, http.StatusForbidden, "Could not succesfully authenticate you")
			return
		}

		next(w, r)
	}
}
