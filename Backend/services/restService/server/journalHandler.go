package server

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

	"github.com/gorilla/mux"

	journalService "github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
)

func (s *Server) HandleJournalHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		j := &journalService.JEmpty{}

		response, err := s.JournalService.GetHealth(context.Background(), j)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the journal service %v", err))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.Message))
		}
	}
}

func (s *Server) HandleJournalSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var journal journalService.Journal
		json.NewDecoder(r.Body).Decode(&journal)

		employee := s.getUsername(r)

		journal.CreatedBy = employee

		response, err := s.JournalService.CreateJournal(context.Background(), &journal)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	}
}

func (s *Server) HandleJournalRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		i, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No journal found for that id")
			return
		}

		j := &journalService.JournalRequest{
			JournalId: int32(i),
		}

		response, err := s.JournalService.GetJournal(context.Background(), j)
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No journal found for that id")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) HandleJournalUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No journal found for that id")
			return
		}

		var journal journalService.Journal
		json.NewDecoder(r.Body).Decode(&journal)

		journal.JournalId = int32(ID)

		response, err := s.JournalService.UpdateJournal(context.Background(), &journal)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) HandleJournalDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No journal found for that id")
		}

		response, err := s.JournalService.DeleteJournal(context.Background(), &journalService.JournalRequest{JournalId: int32(ID)})
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.ReturnError(w, http.StatusInternalServerError, "Something unknown went horribly wrong!!! ☠️☠️☠️")
	}
}

func (s *Server) HandleJournalByPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		pr := &journalService.PatientRequest{
			Patient: vars["username"],
		}

		response, err := s.JournalService.GetJournalsByPatient(context.Background(), pr)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the journal service %v", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Journals) == 0 {
			response.Journals = make([]*journalService.Journal, 0)
		}
		json.NewEncoder(w).Encode(response.Journals)
	}
}

func (s *Server) HandleJournalDocumentDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No journal document found with that id")
			return
		}

		jdr := journalService.JournalDocumentRequest{
			JournalDocumentId: int32(ID),
		}

		response, err := s.JournalService.DeleteJournalDocument(context.Background(), &jdr)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if !response.Success {
			s.ReturnError(w, http.StatusInternalServerError, "Something went horribly wrong!!! ☠️☠️☠️")
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) HandleJournalDocumentUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No journal document found with that id")
			return
		}

		var journalDocument journalService.JournalDocument
		json.NewDecoder(r.Body).Decode(&journalDocument)

		journalDocument.DocumentId = int32(ID)

		response, err := s.JournalService.UpdateJournalDocument(context.Background(), &journalDocument)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	}
}

func (s *Server) HandleJournalDocumentSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var journalDocument journalService.JournalDocument
		json.NewDecoder(r.Body).Decode(&journalDocument)

		journalDocument.CreatedBy = s.getUsername(r)

		response, err := s.JournalService.CreateJournalDocument(context.Background(), &journalDocument)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(response.Attachments) > 0 {
			fmt.Println("there were some journal attachments")
			for _, attachment := range response.Attachments {
				filePath := strings.ReplaceAll(*attachment.Path, "http://cloud.m9ssen.me:56060/static/", "")
				err := s.saveFile(*attachment.Content, filePath)
				fmt.Printf("saving file %v\n\r", filePath)
				if err != nil {
					s.ReturnError(w, http.StatusInternalServerError, err.Error())
					return
				}
				*attachment.Content = ""
			}
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) HandleDocumentUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 	vars := mux.Vars(r)
		// 	documentID, err := strconv.Atoi(vars["documentID"])
		// 	if err != nil {
		// 		s.ReturnError(w, http.StatusInternalServerError, err.Error())
		// 		return
		// 	}

		// 	file, header, err := r.FormFile("file")
		// 	if err != nil {
		// 		s.ReturnError(w, http.StatusInternalServerError, err.Error())
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

		// 	attachmentOutput, err := s.JournalService.CreateAttachment(context.Background(), &attachment)
		// 	if err != nil {
		// 		s.ReturnError(w, http.StatusInternalServerError, err.Error())
		// 		return
		// 	}

		// 	filePath := fmt.Sprintf("%v/%v.%v", storeName, fileName, fileType)
		// 	err = s.saveFile(file, filePath)
		// 	if err != nil {
		// 		s.ReturnError(w, http.StatusInternalServerError, err.Error())
		// 		return
		// 	}

		// 	path := fmt.Sprintf("http://cloud.m9ssen.me:56060/static%v/%v.%v", storeName, fileName, fileType)
		// 	fmt.Printf("path: %v\n\r", path)
		// 	attachmentOutput.Path = &path
		// 	json.NewEncoder(w).Encode(&attachmentOutput)
	}
}

func (s *Server) HandleJournalDocumentByJournal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		journalID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No journal documents found for that journal id")
			return
		}

		pr := &journalService.JournalRequest{
			JournalId: int32(journalID),
		}

		response, err := s.JournalService.GetJournalDocumentsByJournal(context.Background(), pr)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
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
			s.ReturnError(w, http.StatusForbidden, "Not Allowed")
		}
	}
}

func (s *Server) HandleJournalDocumentRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		journalDocumentID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No journal document found for that id")
			return
		}

		j := journalService.JournalDocumentRequest{
			JournalDocumentId: int32(journalDocumentID),
		}

		response, err := s.JournalService.GetJournalDocument(context.Background(), &j)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
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

func (s *Server) HandleJournalUploadSymptoms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Post(os.Getenv("ML_DIAGNOSE_ENDPOINT"), "application/json", r.Body)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer resp.Body.Close()

		buf, err := ioutil.ReadAll(resp.Body)

		w.WriteHeader(resp.StatusCode)
		w.Write(buf)
	}
}

func (s *Server) HandleJournalMLUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Attachments []string
		response := make([]*MLResponse, 0)

		json.NewDecoder(r.Body).Decode(&Attachments)

		for _, attachment := range Attachments {
			fmt.Println(attachment)
			filePath := strings.ReplaceAll(attachment, "http://cloud.m9ssen.me:56060/", "./")
			img, err := os.Open(filePath)
			if err != nil {
				s.ReturnError(w, http.StatusInternalServerError, err.Error())
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
				s.ReturnError(w, http.StatusInternalServerError, err.Error())
				return
			}

			resp, err := http.Post(os.Getenv("ML_IMAGE_ENDPOINT"), "application/json", bytes.NewBuffer(requestBody))
			if err != nil {
				s.ReturnError(w, http.StatusInternalServerError, err.Error())
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				s.ReturnError(w, http.StatusInternalServerError, err.Error())
				return
			}

			fmt.Println(string(body))
			data := new(MLResponse)
			jsonResponse := new(MLOutput)
			err = json.Unmarshal(body, &jsonResponse)
			if err != nil {
				s.ReturnError(w, http.StatusInternalServerError, err.Error())
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
