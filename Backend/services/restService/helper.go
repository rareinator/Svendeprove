package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
)

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

	return 1, nil
}

func (s *server) getUsername(request *http.Request) string {
	// tokenRequest := authentication.TokenRequest{
	// 	Token: token,
	// }

	// response, err := s.authenticationService.ValidateToken(context.Background(), &tokenRequest)
	// if err != nil {
	// 	return "", err
	// }

	// if !response.Valid {
	// 	return "", fmt.Errorf("Could not find the token")
	// }

	return ""
}

func (s *server) saveFile(base64Data, fileName string) error {
	fmt.Println("Saving file")
	dec, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}

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

	if _, err := f.Write(dec); err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	fmt.Println("got done saving file")

	return nil
}
