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
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
	}
}

func (s *server) authenticatePatient(next http.HandlerFunc, idKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		for _, value := range vars {
			fmt.Println(value)
		}

		patientID, err := strconv.Atoi(vars[idKey])
		if err != nil {
			s.returnError(w, http.StatusForbidden, "Could not convert the id to an int")
			return
		}

		reqToken := r.Header.Get("Authorization")
		if reqToken != "" {
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]

			tokenRequest := &authentication.TokenRequest{
				Token: reqToken,
			}

			response, err := s.authenticationService.ValidateToken(context.Background(), tokenRequest)
			if err != nil {
				s.returnError(w, http.StatusForbidden, fmt.Sprintf("%v", err))
				return
			}

			if (!response.Valid) || (response.PatientID != int32(patientID)) {
				s.returnError(w, http.StatusForbidden, "Could not log you in succesfully")
				return
			}
		}

		next(w, r)
	}
}

func (s *server) authenticateRole(next http.HandlerFunc, role models.UserRole) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		if reqToken != "" {
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]

			tokenRequest := &authentication.TokenRequest{
				Token: reqToken,
			}

			response, err := s.authenticationService.ValidateToken(context.Background(), tokenRequest)
			if err != nil {
				s.returnError(w, http.StatusForbidden, fmt.Sprintf("%v", err))
				return
			}

			if (!response.Valid) || (response.Role != int32(role)) {
				s.returnError(w, http.StatusForbidden, "Could not succesfully authenticate you")
				return
			}

			next(w, r)
		} else {
			s.returnError(w, http.StatusNotAcceptable, "No valid token specified")
		}
	}
}
