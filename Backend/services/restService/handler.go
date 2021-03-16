package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	authenticationService "github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	bookingService "github.com/rareinator/Svendeprove/Backend/services/bookingService/booking"
	iotService "github.com/rareinator/Svendeprove/Backend/services/iotService/iot"
	journalService "github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
	patientService "github.com/rareinator/Svendeprove/Backend/services/patientService/patient"
	useradminService "github.com/rareinator/Svendeprove/Backend/services/useradminService/useradmin"
	"github.com/tidwall/buntdb"
)

func (s *server) returnError(w http.ResponseWriter, statusCode int, Message string) {
	var errorMessage struct {
		Code    int
		Message string
	}

	w.WriteHeader(statusCode)
	errorMessage.Code = statusCode
	errorMessage.Message = Message

	fmt.Printf("ERROR!!!! Code: %v Message: %v\n\r", errorMessage.Code, errorMessage.Message)

	json.NewEncoder(w).Encode(&errorMessage)
}

func (s *server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("🚀 Server is up and running!!!!"))
	}
}

func (s *server) handleOauthToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestJson struct {
			Code string `json:"code"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestJson); err != nil {
			s.returnError(w, http.StatusBadRequest, err.Error())
			return
		}

		var scope string
		var role string
		var userID int32
		var username string
		var fullname string

		s.localDB.View(func(tx *buntdb.Tx) error {
			value, err := tx.Get(requestJson.Code)
			if err != nil {
				return err
			}
			fmt.Println(value)

			values := strings.Split(value, ";")
			scope = values[0]
			role = values[1]
			parsedValue, err := strconv.Atoi(values[2])
			if err != nil {
				return err
			}
			userID = int32(parsedValue)
			username = values[3]
			fullname = values[4]

			return nil
		})

		fmt.Println("stuff")
		fmt.Printf("%v;%v;%v;%v;%v\n\r", scope, role, userID, username, fullname)

		token := jwt.New()
		token.Set("scope", scope)
		token.Set("role", role)
		token.Set("userID", userID)
		token.Set("sub", username)
		token.Set("fullname", fullname)

		tokenSign, err := jwt.Sign(token, jwa.HS256, []byte("00000000"))
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		var response struct {
			AccessToken string `json:"access_token"`
			ExpiresIn   int    `json:"expires_in`
			Scope       string `json:"scope"`
			TokenType   string `json:"token_type"`
		}
		response.AccessToken = string(tokenSign)
		response.ExpiresIn = 7200
		response.Scope = scope
		response.TokenType = "Bearer"

		tokenRequest := authenticationService.TokenRequest{
			Token: response.AccessToken,
		}

		validatorResponse, err := s.authenticationService.InsertToken(context.Background(), &tokenRequest)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if !validatorResponse.Valid {
			s.returnError(w, http.StatusInternalServerError, "Something dire went wrong")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)
	}
}

func (s *server) handleOauthAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit endpoint")
		scope := r.URL.Query().Get("scope")
		var login struct {
			Username string
			Password string
		}

		json.NewDecoder(r.Body).Decode(&login)

		a := &authenticationService.User{
			Username: login.Username,
			Password: login.Password,
		}

		switch scope {
		case "employee":
			response, err := s.authenticationService.LoginEmployee(context.Background(), a)
			if err != nil {
				s.returnError(w, http.StatusForbidden, "Error logging in")
				return
			}
			code := uuid.New().String()

			s.localDB.Update(func(tx *buntdb.Tx) error {
				fmt.Printf("savingToDB: employee;%v;%v;%v;%v\n\r",
					response.Role,
					response.UserID,
					response.Username,
					response.FullName)
				_, _, err := tx.Set(code, fmt.Sprintf("employee;%v;%v;%v;%v",
					response.Role,
					response.UserID,
					response.Username,
					response.FullName), &buntdb.SetOptions{Expires: true, TTL: (time.Second * 30)})

				if err != nil {
					return err
				}
				return nil
			})

			w.WriteHeader(http.StatusOK)
			var responseJson struct {
				Code string `json:"code"`
			}
			responseJson.Code = code
			json.NewEncoder(w).Encode(&responseJson)

		case "patient":
			response, err := s.authenticationService.LoginPatient(context.Background(), a)
			if err != nil {
				s.returnError(w, http.StatusForbidden, "Error logging in")
				return
			}

			code := uuid.New().String()

			s.localDB.Update(func(tx *buntdb.Tx) error {
				_, _, err := tx.Set(code, fmt.Sprintf("patient;%v;%v;%v;%v",
					response.Role,
					response.UserID,
					response.Username,
					response.FullName), &buntdb.SetOptions{Expires: true, TTL: (time.Second * 30)})

				if err != nil {
					return err
				}
				return nil
			})

			w.WriteHeader(http.StatusOK)
			var responseJson struct {
				Code string `json:"code"`
			}
			responseJson.Code = code
			json.NewEncoder(w).Encode(&responseJson)
		}
	}
}

func (s *server) handleJournalHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		j := &journalService.JEmpty{}

		response, err := s.journalService.GetHealth(context.Background(), j)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the journal service %v", err))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.Message))
		}
	}
}

func (s *server) handleJournalSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var journal journalService.Journal
		json.NewDecoder(r.Body).Decode(&journal)

		employee := s.getUsername(r)

		journal.CreatedBy = employee

		response, err := s.journalService.CreateJournal(context.Background(), &journal)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	}
}

func (s *server) handleJournalRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		i, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal found for that id")
			return
		}

		j := &journalService.JournalRequest{
			JournalId: int32(i),
		}

		response, err := s.journalService.GetJournal(context.Background(), j)
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal found for that id")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleJournalUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal found for that id")
			return
		}

		var journal journalService.Journal
		json.NewDecoder(r.Body).Decode(&journal)

		journal.JournalId = int32(ID)

		response, err := s.journalService.UpdateJournal(context.Background(), &journal)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleJournalDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal found for that id")
		}

		response, err := s.journalService.DeleteJournal(context.Background(), &journalService.JournalRequest{JournalId: int32(ID)})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.returnError(w, http.StatusInternalServerError, "Something unknown went horribly wrong!!! ☠️☠️☠️")
	}
}

func (s *server) handleJournalByPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		pr := &journalService.PatientRequest{
			Patient: vars["username"],
		}

		response, err := s.journalService.GetJournalsByPatient(context.Background(), pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the journal service %v", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Journals) == 0 {
			response.Journals = make([]*journalService.Journal, 0)
		}
		json.NewEncoder(w).Encode(response.Journals)
	}
}

func (s *server) handleJournalDocumentDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal document found with that id")
			return
		}

		jdr := journalService.JournalDocumentRequest{
			JournalDocumentId: int32(ID),
		}

		response, err := s.journalService.DeleteJournalDocument(context.Background(), &jdr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if !response.Success {
			s.returnError(w, http.StatusInternalServerError, "Something went horribly wrong!!! ☠️☠️☠️")
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) handleJournalDocumentUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal document found with that id")
			return
		}

		var journalDocument journalService.JournalDocument
		json.NewDecoder(r.Body).Decode(&journalDocument)

		journalDocument.DocumentId = int32(ID)

		response, err := s.journalService.UpdateJournalDocument(context.Background(), &journalDocument)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	}
}

func (s *server) handleJournalDocumentSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var journalDocument journalService.JournalDocument
		json.NewDecoder(r.Body).Decode(&journalDocument)

		journalDocument.CreatedBy = s.getUsername(r)

		response, err := s.journalService.CreateJournalDocument(context.Background(), &journalDocument)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(response.Attachments) > 0 {
			fmt.Println("there were some journal attachments")
			for _, attachment := range response.Attachments {
				filePath := strings.ReplaceAll(*attachment.Path, "http://cloud.m9ssen.me:56060/static/", "")
				err := s.saveFile(*attachment.Content, filePath)
				fmt.Printf("saving file %v\n\r", filePath)
				if err != nil {
					s.returnError(w, http.StatusInternalServerError, err.Error())
					return
				}
				*attachment.Content = ""
			}
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleDocumentUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 	vars := mux.Vars(r)
		// 	documentID, err := strconv.Atoi(vars["documentID"])
		// 	if err != nil {
		// 		s.returnError(w, http.StatusInternalServerError, err.Error())
		// 		return
		// 	}

		// 	file, header, err := r.FormFile("file")
		// 	if err != nil {
		// 		s.returnError(w, http.StatusInternalServerError, err.Error())
		// 		return
		// 	}

		// 	defer file.Close()
		// 	fullFileQualifier := strings.Split(header.Filename, ".")
		// 	fileName := fullFileQualifier[0]
		// 	fileType := fullFileQualifier[1]
		// 	fmt.Printf("got file: %v.%v", fileName, fileType)

		// 	// fileType, err := j.DB.GetOrCreateFileTypeByName(fileTypeName)
		// 	// if err != nil {
		// 	// 	return nil, err
		// 	// }
		// 	// build up store name
		// 	// fmt.Println("buildUpStoreName")
		// 	// store, err := j.DB.GetOrCreateFileStoreByPath(storeName)
		// 	// if err != nil {
		// 	// 	return nil, err
		// 	// }

		// 	// dbAttachment := journalService.Attachment{
		// 	// 	FileName:    fileName,
		// 	// 	FileStoreId: store.FileStoreId,
		// 	// 	DocumentId:  dbJD.DocumentId,
		// 	// 	FileTypeId:  fileType.FileTypeId,
		// 	// }

		// 	storeName := fmt.Sprintf("/journal/document/%v", documentID)
		// 	attachment := journalService.Attachment{
		// 		FileName:   fileName,
		// 		DocumentId: int32(documentID),
		// 		FileType:   new(string),
		// 		Path:       new(string),
		// 	}
		// 	attachment.FileType = &fileType
		// 	attachment.Path = &storeName

		// 	attachmentOutput, err := s.journalService.CreateAttachment(context.Background(), &attachment)
		// 	if err != nil {
		// 		s.returnError(w, http.StatusInternalServerError, err.Error())
		// 		return
		// 	}

		// 	filePath := fmt.Sprintf("%v/%v.%v", storeName, fileName, fileType)
		// 	err = s.saveFile(file, filePath)
		// 	if err != nil {
		// 		s.returnError(w, http.StatusInternalServerError, err.Error())
		// 		return
		// 	}

		// 	path := fmt.Sprintf("http://cloud.m9ssen.me:56060/static%v/%v.%v", storeName, fileName, fileType)
		// 	fmt.Printf("path: %v\n\r", path)
		// 	attachmentOutput.Path = &path
		// 	json.NewEncoder(w).Encode(&attachmentOutput)
	}
}

func (s *server) handleJournalDocumentByJournal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		journalID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal documents found for that journal id")
			return
		}

		pr := &journalService.JournalRequest{
			JournalId: int32(journalID),
		}

		response, err := s.journalService.GetJournalDocumentsByJournal(context.Background(), pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		//TODO: oauth fix

		allowed := true

		if allowed {
			w.WriteHeader(http.StatusOK)
			if len(response.JournalDocuments) == 0 {
				response.JournalDocuments = make([]*journalService.JournalDocument, 0)
			}
			json.NewEncoder(w).Encode(response.JournalDocuments)
		} else {
			s.returnError(w, http.StatusForbidden, "Not Allowed")
		}
	}
}

func (s *server) handleJournalDocumentRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		journalDocumentID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal document found for that id")
			return
		}

		j := journalService.JournalDocumentRequest{
			JournalDocumentId: int32(journalDocumentID),
		}

		response, err := s.journalService.GetJournalDocument(context.Background(), &j)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

type Prediction struct {
	Positive string `json:"positive"`
	Negative string `json:"negative"`
}

type MLOutput struct {
	Code       int        `json:"code"`
	Prediction Prediction `json:"prediction"`
}

type MLResponse struct {
	Url  string `json:"Url"`
	Data MLOutput
}

func (s *server) handleJournalUploadSymptoms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Post(os.Getenv("ML_DIAGNOSE_ENDPOINT"), "application/json", r.Body)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer resp.Body.Close()

		buf, err := ioutil.ReadAll(resp.Body)

		w.WriteHeader(resp.StatusCode)
		w.Write(buf)
	}
}

func (s *server) handleJournalMLUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Attachments []string
		response := make([]*MLResponse, 0)

		json.NewDecoder(r.Body).Decode(&Attachments)

		for _, attachment := range Attachments {
			fmt.Println(attachment)
			filePath := strings.ReplaceAll(attachment, "http://cloud.m9ssen.me:56060/", "./")
			img, err := os.Open(filePath)
			if err != nil {
				s.returnError(w, http.StatusInternalServerError, err.Error())
				return
			}
			defer img.Close()

			fInfo, _ := img.Stat()
			var size int64 = fInfo.Size()
			buf := make([]byte, size)

			fReader := bufio.NewReader(img)
			fReader.Read(buf)
			base64Str := base64.StdEncoding.EncodeToString(buf)
			requestBody, err := json.Marshal(map[string]string{
				"scan": base64Str,
			})
			if err != nil {
				s.returnError(w, http.StatusInternalServerError, err.Error())
				return
			}

			resp, err := http.Post(os.Getenv("ML_IMAGE_ENDPOINT"), "application/json", bytes.NewBuffer(requestBody))
			if err != nil {
				s.returnError(w, http.StatusInternalServerError, err.Error())
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				s.returnError(w, http.StatusInternalServerError, err.Error())
				return
			}

			fmt.Println(string(body))
			data := new(MLResponse)
			jsonResponse := new(MLOutput)
			err = json.Unmarshal(body, &jsonResponse)
			if err != nil {
				s.returnError(w, http.StatusInternalServerError, err.Error())
				return
			}

			// fmt.Printf("got code %v\n\r", data.Data.Code)
			data.Data = *jsonResponse
			data.Url = attachment
			response = append(response, data)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)

	}
}

func (s *server) handlePatientHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := &patientService.PEmpty{}

		response, err := s.patientService.GetHealth(context.Background(), p)
		if err != nil {
			s.returnError(w, http.StatusAccepted, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

func (s *server) handlePatientSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var patient patientService.Patient
		json.NewDecoder(r.Body).Decode(&patient)

		response, err := s.patientService.CreatePatient(context.Background(), &patient)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		p := patientService.PRequest{
			Username: vars["username"],
		}

		response, err := s.patientService.GetPatient(context.Background(), &p)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

type doctor struct {
	Username string `json:"Username"`
	Name     string `json:"Name"`
}

func (s *server) handleGetDoctorsInHospital() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: fix so it also check on department ID

		_, client, err := okta.NewClient(context.Background(), okta.WithOrgUrl(os.Getenv("OKTA_URL")), okta.WithToken(os.Getenv("OKTA_SDK_TOKEN")))
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		users, _, err := client.Group.ListGroupUsers(context.Background(), "00gbrw99eH74jULM95d6", &query.Params{})

		result := make([]*doctor, 0)

		for _, user := range users {
			doctor := doctor{
				Name:     fmt.Sprintf("%v", (*user.Profile)["displayName"]),
				Username: fmt.Sprintf("%v", (*user.Profile)["login"]),
			}

			result = append(result, &doctor)

		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&result)
	}
}

func (s *server) handlePatientsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, client, err := okta.NewClient(context.Background(), okta.WithOrgUrl(os.Getenv("OKTA_URL")), okta.WithToken(os.Getenv("OKTA_SDK_TOKEN")))
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		users, _, err := client.Group.ListGroupUsers(context.Background(), "00gbqw93aeIKYWqww5d6", &query.Params{})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		result := make([]*patientService.Patient, 0)

		for _, user := range users {
			patient := patientService.Patient{
				Name:       fmt.Sprintf("%v", (*user.Profile)["displayName"]),
				Address:    fmt.Sprintf("%v", (*user.Profile)["streetAddress"]),
				City:       fmt.Sprintf("%v", (*user.Profile)["city"]),
				PostalCode: fmt.Sprintf("%v", (*user.Profile)["zipCode"]),
				Country:    fmt.Sprintf("%v", (*user.Profile)["full_country"]),
				SocialIdNr: fmt.Sprintf("%v", (*user.Profile)["social_id"]),
				Username:   fmt.Sprintf("%v", (*user.Profile)["login"]),
			}

			result = append(result, &patient)

		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (s *server) handleDiagnoseGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		p := patientService.PRequest{
			Username: vars["username"],
		}

		response, err := s.patientService.GetDiagnose(context.Background(), &p)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleDiagnosesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := s.patientService.GetDiagnoses(context.Background(), &patientService.PEmpty{})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Diagnoses) == 0 {
			response.Diagnoses = make([]*patientService.Diagnose, 0)
		}
		json.NewEncoder(w).Encode(response.Diagnoses)
	}
}

func (s *server) handleSymptomGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		p := patientService.PRequest{
			Username: vars["username"],
		}

		response, err := s.patientService.GetSymptom(context.Background(), &p)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)
	}
}

func (s *server) handleSymptomsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := s.patientService.GetSymptoms(context.Background(), &patientService.PEmpty{})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Symptoms) == 0 {
			response.Symptoms = make([]*patientService.Symptom, 0)
		}
		json.NewEncoder(w).Encode(response.Symptoms)

	}
}

func (s *server) handlePatientDiagnoseSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var patientDiagnose patientService.PatientDiagnose
		json.NewDecoder(r.Body).Decode(&patientDiagnose)
		patientDiagnose.Patient = vars["username"]

		response, err := s.patientService.CreatePatientDiagnose(context.Background(), &patientDiagnose)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientDiagnosesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		pr := patientService.PRequest{
			Username: vars["username"],
		}

		response, err := s.patientService.GetPatientDiagnoses(context.Background(), &pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.PatientDiagnoses) == 0 {
			response.PatientDiagnoses = make([]*patientService.PatientDiagnose, 0)
		}
		json.NewEncoder(w).Encode(response.PatientDiagnoses)
	}
}

func (s *server) handlePatientDiagnoseGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "Diagnose for that id not found")
			return
		}

		pr := patientService.PRequest{
			Id: int32(id),
		}

		response, err := s.patientService.GetPatientDiagnose(context.Background(), &pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientDiagnoseUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "Diagnose for that id not found")
			return
		}

		var patientDiagnose patientService.PatientDiagnose
		json.NewDecoder(r.Body).Decode(&patientDiagnose)

		patientDiagnose.PatientDiagnoseId = int32(ID)

		response, err := s.patientService.UpdatePatientDiagnose(context.Background(), &patientDiagnose)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientDiagnoseDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No patient Diagnosis found for that id")
			return
		}

		response, err := s.patientService.DeletePatientDiagnose(context.Background(), &patientService.PRequest{Id: int32(ID)})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.returnError(w, http.StatusInternalServerError, "Something unknown went horribly wrong!!! ☠️☠️☠️")
	}
}

func (s *server) handlePatientSymptomCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var diagnoseSymptom patientService.DiagnoseSymptom
		json.NewDecoder(r.Body).Decode(&diagnoseSymptom)

		response, err := s.patientService.CreateDiagnoseSymptom(context.Background(), &diagnoseSymptom)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientSymptomsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["diagnoseID"])
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		pr := patientService.PRequest{
			Id: int32(id),
		}

		response, err := s.patientService.GetDiagnoseSymptoms(context.Background(), &pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.DiagnoseSymptoms) == 0 {
			response.DiagnoseSymptoms = make([]*patientService.DiagnoseSymptom, 0)
		}
		json.NewEncoder(w).Encode(response.DiagnoseSymptoms)
	}
}

func (s *server) handlePatientSymptomUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		diagnoseID, err := strconv.Atoi(vars["diagnoseID"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		var newDiagnoseSymptom patientService.DiagnoseSymptom
		json.NewDecoder(r.Body).Decode(&newDiagnoseSymptom)

		oldDiagnoseSymptom := patientService.DiagnoseSymptom{
			SymptomId:         int32(ID),
			PatientDiagnoseId: int32(diagnoseID),
		}

		dsur := patientService.DiagnoseSymptomUpdateRequest{
			Old: &oldDiagnoseSymptom,
			New: &newDiagnoseSymptom,
		}

		response, err := s.patientService.UpdateDiagnoseSymptom(context.Background(), &dsur)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientSymptomDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		diagnoseID, err := strconv.Atoi(vars["diagnoseID"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		ds := patientService.DiagnoseSymptom{
			SymptomId:         int32(ID),
			PatientDiagnoseId: int32(diagnoseID),
		}

		response, err := s.patientService.DeleteDiagnoseSymptom(context.Background(), &ds)
		if err != nil {
			s.returnError(w, http.StatusNotFound, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.returnError(w, http.StatusInternalServerError, "Something unknown went horribly wrong!!! ☠️☠️☠️")
	}
}

func (s *server) handleBookingHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b := bookingService.BEmpty{}

		response, err := s.bookingService.GetHealth(context.Background(), &b)
		if err != nil {
			s.returnError(w, http.StatusServiceUnavailable, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

func (s *server) handleBookingCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var booking bookingService.Booking
		json.NewDecoder(r.Body).Decode(&booking)

		booking.Employee = s.getUsername(r)

		response, err := s.bookingService.CreateBooking(context.Background(), &booking)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleBookingGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No booking found with that id")
			return
		}

		b := bookingService.BRequest{
			Id: int32(ID),
		}

		response, err := s.bookingService.GetBooking(context.Background(), &b)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleBookingUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookingID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No booking found with that id")
			return
		}

		var booking bookingService.Booking
		json.NewDecoder(r.Body).Decode(&booking)

		booking.BookingId = int32(bookingID)

		response, err := s.bookingService.UpdateBooking(context.Background(), &booking)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleBookingDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookingID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No booking with that id found")
			return
		}

		br := bookingService.BRequest{
			Id: int32(bookingID),
		}

		response, err := s.bookingService.DeleteBooking(context.Background(), &br)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.returnError(w, http.StatusInternalServerError, "Somethin unknown went gorribly wrong!!! ☠️☠️☠️")
	}
}

func (s *server) handleBookingsByPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		br := bookingService.BRequest{
			Username: vars["username"],
		}

		response, err := s.bookingService.GetBookingsByPatient(context.Background(), &br)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Bookings) == 0 {
			response.Bookings = make([]*bookingService.Booking, 0)
		}
		json.NewEncoder(w).Encode(response.Bookings)
	}
}

func (s *server) handleBookingsByEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		br := bookingService.BRequest{
			Username: vars["username"],
		}

		response, err := s.bookingService.GetBookingsByEmployee(context.Background(), &br)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Bookings) == 0 {
			response.Bookings = make([]*bookingService.Booking, 0)
		}
		json.NewEncoder(w).Encode(response.Bookings)
	}
}

func (s *server) handleAvailableTimesForDoctor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request bookingService.BTimeFrameRequest

		json.NewDecoder(r.Body).Decode(&request)

		response, err := s.bookingService.GetAvailableTimesForDoctor(context.Background(), &request)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response.Strings)
	}
}

func (s *server) handleUseradminHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := useradminService.UAEmpty{}

		response, err := s.useradminService.GetHealth(context.Background(), &u)
		if err != nil {
			s.returnError(w, http.StatusServiceUnavailable, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

func (s *server) handleUseradminGetEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		er := useradminService.UserRequest{
			Username: vars["username"],
		}

		response, err := s.useradminService.GetEmployee(context.Background(), &er)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleHospitalsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := s.useradminService.GetHospitals(context.Background(), &useradminService.UAEmpty{})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response.Hospitals)
	}
}

func (s *server) handleAuthenticationHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		e := &authenticationService.AEmpty{}

		response, err := s.authenticationService.GetHealth(context.Background(), e)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the authentication service %v", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

func (s *server) handleAuthenticationEmployeeLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login struct {
			Username string
			Password string
		}

		json.NewDecoder(r.Body).Decode(&login)

		a := &authenticationService.User{
			Username: login.Username,
			Password: login.Password,
		}

		response, err := s.authenticationService.LoginEmployee(context.Background(), a)
		if err != nil {
			s.returnError(w, http.StatusForbidden, "Error logging in")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleAuthenticationPatientLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login struct {
			Username string
			Password string
		}

		json.NewDecoder(r.Body).Decode(&login)

		a := &authenticationService.User{
			Username: login.Username,
			Password: login.Password,
		}

		response, err := s.authenticationService.LoginPatient(context.Background(), a)
		if err != nil {
			s.returnError(w, http.StatusForbidden, "Error logging in")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleIOTHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		i := iotService.IOTEmpty{}

		response, err := s.iotService.GetHealth(context.Background(), &i)
		if err != nil {
			s.returnError(w, http.StatusServiceUnavailable, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

func (s *server) handleIOTUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Query().Get("Data")
		deviceID, err := s.getDeviceID(r)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		username := s.getUsername(r)

		//TODO: fix

		iotData := iotService.IOTData{
			Name:     username,
			SensorID: deviceID,
			Data:     data,
		}

		response, err := s.iotService.UploadData(context.Background(), &iotData)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleIOTReadData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		deviceID, err := strconv.Atoi(vars["deviceID"])
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		request := iotService.IOTRequest{
			ID: int32(deviceID),
		}

		response, err := s.iotService.ReadData(context.Background(), &request)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response.IOTDatas)

	}
}

func (s *server) handleIOTReadDataInTimeframe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request iotService.IOTTimeframeRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			s.returnError(w, http.StatusBadRequest, err.Error())
			return
		}

		response, err := s.iotService.ReadDataInTimeFrame(context.Background(), &request)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response.IOTDatas)
	}
}
