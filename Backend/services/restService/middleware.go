package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/rareinator/Svendeprove/Backend/packages/models"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
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

		tokenRequest := &authentication.TokenRequest{
			Token: reqToken,
		}

		allowed := false
		if config.allowRelatedPatient {
			allowed = true
		}

		response, err := s.authenticationService.ValidateToken(context.Background(), tokenRequest)
		if err != nil {
			s.returnError(w, http.StatusForbidden, fmt.Sprintf("%v", err))
			return
		}
		if !response.Valid {
			s.returnError(w, http.StatusForbidden, "Could not succesfully authenticate you, response not valid")
			return
		}

		fmt.Println("1")

		jwtToken, err := jwt.Parse([]byte(reqToken))
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		fmt.Println("2")

		userID, exists := jwtToken.Get("userID")
		if !exists {
			s.returnError(w, http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println("3")

		role, exists := jwtToken.Get("role")
		if !exists {
			s.returnError(w, http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println("4")

		fmt.Println("Token")
		fmt.Println(jwtToken.Get("fullname"))

		if config.allowedPatient != "" {
			patientID, err := strconv.Atoi(vars[config.allowedPatient])
			if err != nil {
				s.returnError(w, http.StatusForbidden, "Could not convert the id to an int")
				return
			}

			if userID == int32(patientID) {
				allowed = true
			}
		}
		fmt.Println("5")
		if len(config.allowedRoles) > 0 {
			for _, allowedRole := range config.allowedRoles {
				fmt.Printf("role: %v\n\rallowedRole: %v\n\r", role, int32(allowedRole))
				roleInt, err := strconv.Atoi(role.(string))
				if err != nil {
					s.returnError(w, http.StatusInternalServerError, err.Error())
					return
				}
				if roleInt == int(allowedRole) {
					allowed = true
					break
				}
			}
		}
		fmt.Println("5")

		if config.allowIOTDevice {
			// fmt.Println(response)
			// fmt.Println(response.IOTDeviceId)
			// if response.IOTDeviceId != 0 {
			// 	allowed = true
			// }

			//TODO: find out what to do here properly
			allowed = true
		}

		if !allowed {
			s.returnError(w, http.StatusForbidden, "Could not succesfully authenticate you")
			return
		}

		next(w, r)
	}
}
