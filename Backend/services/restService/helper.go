package main

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
)

func (s *server) patientIsAuthenticated(dbType interface{}, lookupID int32, patientID int32) (bool, error) {
	rpr := authentication.RelatedPatientRequest{
		Type: reflect.TypeOf(dbType).Name(),
		Id:   lookupID,
	}
	response, err := s.authenticationService.GetRelatedPatient(context.Background(), &rpr)
	if err != nil {
		return false, err
	}

	return patientID == response.PatientId, nil
}

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

func (s *server) getEmployeeID(r *http.Request) (int32, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	if reqToken == "" {
		return 0, fmt.Errorf("Could not find a token")
	}

	fmt.Printf("getting employeeID Token: %v\n\r", reqToken)

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

	return response.EmployeeID, nil
}
