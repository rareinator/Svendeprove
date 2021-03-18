package server

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
)

func (s *Server) getUserId(request *http.Request) string {
	var reqToken string
	reqToken = request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	toValidate := map[string]string{}
	toValidate["aud"] = os.Getenv("OKTA_AUTH_ENDPOINT")
	toValidate["cid"] = os.Getenv("OKTA_CLIENT_ID")

	jwt := jwtverifier.JwtVerifier{
		Issuer:           fmt.Sprintf("%v/oauth2/default", os.Getenv("OKTA_URL")),
		ClaimsToValidate: toValidate,
	}

	verifier := jwt.New()

	token, _ := verifier.VerifyAccessToken(reqToken)

	return fmt.Sprintf("%v", token.Claims["uid"])

}

func (s *Server) saveFile(base64Data, fileName string) error {
	fmt.Println("Saving file")
	dec, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%v/%v", s.StaticFileDir, fileName)

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
