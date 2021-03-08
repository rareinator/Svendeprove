package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	"github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
)

func (s *server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("🚀 Server is up and running!!!!"))
	}
}

func (s *server) handleJournalHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		j := &journal.JEmpty{}

		response, err := s.journalService.GetHealth(context.Background(), j)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error getting in contact with the journal service %v", err)))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.Message))
		}
	}
}

func (s *server) handleJournalSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var journal journal.Journal
		json.NewDecoder(r.Body).Decode(&journal)

		fmt.Println("should do handle save stuff here")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Test"))
	}
}

func (s *server) handleJournalRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		i, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No journal found for that id"))
			return
		}

		j := &journal.JournalRequest{
			JournalId: int32(i),
		}

		response, err := s.journalService.GetJournal(context.Background(), j)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No journal found for that id"))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleJournalUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *server) handleJournalDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *server) handleJournalByPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		patientId, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No journals found for that patient id"))
			return
		}

		pr := &journal.PatientRequest{
			PatientId: int32(patientId),
		}

		response, err := s.journalService.GetJournalsByPatient(context.Background(), pr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error getting in contact with the journal service %v", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Journals)
	}
}

func (s *server) handleAuthenticationHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		e := &authentication.AEmpty{}

		response, err := s.authenticationService.GetHealth(context.Background(), e)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error getting in contact with the authentication service %v", err)))
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

		a := &authentication.User{
			Username: login.Username,
			Password: login.Password,
		}

		response, err := s.authenticationService.LoginEmployee(context.Background(), a)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(fmt.Sprintf("Error logging in")))
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

		a := &authentication.User{
			Username: login.Username,
			Password: login.Password,
		}

		response, err := s.authenticationService.LoginPatient(context.Background(), a)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Error logging in"))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}
}
