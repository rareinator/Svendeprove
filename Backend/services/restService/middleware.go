package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rareinator/Svendeprove/Backend/packages/models"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
)

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
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Error getting in contact with the authentication service %v", err)))
				return
			}

			if (!response.Valid) || (response.Role != int32(role)) {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}
}
