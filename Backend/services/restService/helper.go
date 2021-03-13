package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
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
	if len(splitToken) != 2 {
		fmt.Println("trying to acces with no token")
		return 0, fmt.Errorf("No valid token specified")
	}
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
	if len(splitToken) != 2 {
		fmt.Println("trying to acces with no token")
		return 0, fmt.Errorf("No valid token specified")
	}
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

func (s *server) getDeviceID(r *http.Request) (int32, error) {
	reqToken := r.URL.Query().Get("Key")
	if reqToken == "" {
		return 0, fmt.Errorf("No valid token specified, found %v", reqToken)
	}

	tokenRequest := authentication.TokenRequest{
		Token: reqToken,
	}

	response, err := s.authenticationService.ValidateToken(context.Background(), &tokenRequest)
	if err != nil {
		return 0, err
	}

	if !response.Valid {
		return 0, fmt.Errorf("Could not find the token")
	}

	return response.IOTDeviceId, nil
}

func (s *server) getUsername(token string) (string, error) {
	tokenRequest := authentication.TokenRequest{
		Token: token,
	}

	response, err := s.authenticationService.ValidateToken(context.Background(), &tokenRequest)
	if err != nil {
		return "", err
	}

	if !response.Valid {
		return "", fmt.Errorf("Could not find the token")
	}

	return response.Username, nil
}

func (s *server) saveFile(file multipart.File, fileName string) error {
	fmt.Println("Saving file")
	var buf bytes.Buffer
	io.Copy(&buf, file)

	filePath := fmt.Sprintf("%v/%v", s.staticFileDir, fileName)

	fmt.Printf("Saving file to %v\n\r", filePath)

	basepath := path.Dir(filePath)
	if err := os.MkdirAll(basepath, 0777); err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(buf.Bytes()); err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	fmt.Println("got done saving file")

	return nil
}
