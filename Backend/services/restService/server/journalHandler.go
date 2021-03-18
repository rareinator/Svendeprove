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

	protocol "github.com/rareinator/Svendeprove/Backend/packages/protocol"
)

func (s *Server) HandleJournalHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		j := protocol.Empty{}

		response, err := s.JournalService.GetHealth(context.Background(), &j)
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
		var journal protocol.Journal
		json.NewDecoder(r.Body).Decode(&journal)

		employee := s.getUserId(r)

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

		j := &protocol.JournalRequest{
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

		var journal protocol.Journal
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

		response, err := s.JournalService.DeleteJournal(context.Background(), &protocol.JournalRequest{JournalId: int32(ID)})
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

		pr := &protocol.PatientRequest{
			Patient: vars["username"],
		}

		response, err := s.JournalService.GetJournalsByPatient(context.Background(), pr)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the journal service %v", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Journals) == 0 {
			response.Journals = make([]*protocol.Journal, 0)
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

		jdr := protocol.JournalDocumentRequest{
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

		var journalDocument protocol.JournalDocument
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
		var journalDocument protocol.JournalDocument
		json.NewDecoder(r.Body).Decode(&journalDocument)

		journalDocument.CreatedBy = s.getUserId(r)

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

func (s *Server) HandleJournalDocumentByJournal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		journalID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No journal documents found for that journal id")
			return
		}

		pr := &protocol.JournalRequest{
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
				response.JournalDocuments = make([]*protocol.JournalDocument, 0)
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

		j := protocol.JournalDocumentRequest{
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
