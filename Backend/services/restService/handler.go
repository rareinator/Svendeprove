package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	authenticationService "github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	journalService "github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
)

func (s *server) returnError(w http.ResponseWriter, statusCode int, Message string) {
	var errorMessage struct {
		Code    int
		Message string
	}

	w.WriteHeader(statusCode)
	errorMessage.Code = statusCode
	errorMessage.Message = Message

	json.NewEncoder(w).Encode(&errorMessage)
}

func (s *server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("🚀 Server is up and running!!!!"))
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
		fmt.Println(journal.CreatedBy)

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
		patientID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journals found for that patientId")
			return
		}

		pr := &journalService.PatientRequest{
			PatientId: int32(patientID),
		}

		response, err := s.journalService.GetJournalsByPatient(context.Background(), pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the journal service %v", err))
			return
		}

		w.WriteHeader(http.StatusOK)
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
