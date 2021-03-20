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
			if _, err := w.Write([]byte(response.Message)); err != nil {
				s.ReturnError(w, http.StatusInternalServerError, err.Error())
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	}
}

func (s *Server) HandleJournalSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var journal protocol.Journal
		if err := json.NewDecoder(r.Body).Decode(&journal); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		journal.CreatedBy = s.getUserId(r)

		response, err := s.JournalService.CreateJournal(context.Background(), &journal)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
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
		if err := json.NewDecoder(r.Body).Decode(&journal); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		journal.JournalId = int32(ID)

		response, err := s.JournalService.UpdateJournal(context.Background(), &journal)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
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

		pr := protocol.UserRequest{
			UserId: vars["userId"],
		}

		response, err := s.JournalService.GetJournalsByPatient(context.Background(), &pr)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the journal service %v", err))
			return
		}

		if len(response.Journals) == 0 {
			response.Journals = make([]*protocol.Journal, 0)
		}

		if err := json.NewEncoder(w).Encode(response.Journals); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
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
		if err := json.NewDecoder(r.Body).Decode(&journalDocument); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		journalDocument.DocumentId = int32(ID)

		response, err := s.JournalService.UpdateJournalDocument(context.Background(), &journalDocument)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) HandleJournalDocumentSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var journalDocument protocol.JournalDocument
		if err := json.NewDecoder(r.Body).Decode(&journalDocument); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		journalDocument.CreatedBy = s.getUserId(r)

		response, err := s.JournalService.CreateJournalDocument(context.Background(), &journalDocument)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(response.Attachments) > 0 {
			fmt.Println("there were some journal attachments")
			for _, attachment := range response.Attachments {
				filePath := strings.ReplaceAll(*attachment.Path, "https://school.m9ssen.me/static/", "")
				err := s.saveFile(*attachment.Content, filePath)
				fmt.Printf("saving file %v\n\r", filePath)
				if err != nil {
					s.ReturnError(w, http.StatusInternalServerError, err.Error())
					return
				}
				*attachment.Content = ""
			}
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
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

		pr := protocol.JournalRequest{
			JournalId: int32(journalID),
		}

		response, err := s.JournalService.GetJournalDocumentsByJournal(context.Background(), &pr)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(response.JournalDocuments) == 0 {
			response.JournalDocuments = make([]*protocol.JournalDocument, 0)
		}

		if err := json.NewEncoder(w).Encode(response.JournalDocuments); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
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

		if err := json.NewEncoder(w).Encode(response); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
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
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if _, err := w.Write(buf); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(resp.StatusCode)
	}
}

func (s *Server) HandleJournalMLUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Attachments []string
		response := make([]*MLResponse, 0)

		if err := json.NewDecoder(r.Body).Decode(&Attachments); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		for _, attachment := range Attachments {
			fmt.Println(attachment)
			filePath := strings.ReplaceAll(attachment, "https://school.m9ssen.me/", "./")
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
			if _, err := fReader.Read(buf); err != nil {
				s.ReturnError(w, http.StatusInternalServerError, err.Error())
				return
			}

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

			data.Data = *jsonResponse
			data.Url = attachment
			response = append(response, data)
		}

		if err := json.NewEncoder(w).Encode(&response); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
