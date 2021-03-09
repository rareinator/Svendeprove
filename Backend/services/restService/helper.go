package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
)

func (s *server) getPatientID(r *http.Request) (int32, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	if reqToken == "" {
		return 0, fmt.Errorf("Could not find a token")
	}

	tokenRequest := &authentication.TokenRequest{
		Token: reqToken,
	}

	response, err := s.authenticationService.ValidateToken(context.Background(), tokenRequest)
	if err != nil {
		return 0, err
	}
	if !response.Valid {
		return 0, fmt.Errorf("Could not find the token")
	}

	return response.PatientID, nil
}
