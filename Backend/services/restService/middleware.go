package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rareinator/Svendeprove/Backend/packages/models"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
)

type authenticationConfig struct {
	allowedRoles        []models.UserRole
	allowedPatient      string
	allowRelatedPatient bool
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
		fmt.Printf("%v called %v\n\rresource %v\n\r", r.Host, r.Method, r.RequestURI)

		if r.Method == "POST" {
			fmt.Printf("Body: \b\r%v\n\r", r.Body)
		}
	}
}

func (s *server) authenticate(next http.HandlerFunc, config *authenticationConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authenticating")

		vars := mux.Vars(r)

		reqToken := r.Header.Get("Authorization")
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

		if config.allowedPatient != "" {
			patientID, err := strconv.Atoi(vars[config.allowedPatient])
			if err != nil {
				s.returnError(w, http.StatusForbidden, "Could not convert the id to an int")
				return
			}
			if response.PatientID == int32(patientID) {
				allowed = true
			}
		}
		if len(config.allowedRoles) > 0 {
			for _, allowedRole := range config.allowedRoles {
				fmt.Printf("role: %v\n\rallowedRole: %v\n\r", response.Role, int32(allowedRole))
				if response.Role == int32(allowedRole) {
					allowed = true
					break
				}
			}
		}

		if !allowed {
			s.returnError(w, http.StatusForbidden, "Could not succesfully authenticate you")
			return
		}

		next(w, r)
	}
}
