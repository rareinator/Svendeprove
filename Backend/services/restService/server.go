package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	"github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
)

type server struct {
	router                *mux.Router
	journalService        journal.JournalServiceClient
	authenticationService authentication.AuthenticationServiceClient
}

func newServer() *server {
	s := &server{}
	s.router = mux.NewRouter()
	s.routes()
	return s
}

func (s *server) ServeHTTP() {
	address := os.Getenv("REST_SERVICE_ADDR")
	fmt.Printf("🚀 Listening on address: %v ", address)
	http.ListenAndServe(address, s.router)
}

func (s *server) handleHealth(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is up and running!!!!"))
}

func (s *server) handleJournalHealth(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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

func (s *server) handleJournalSave(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

}

func (s *server) handleJournalRead(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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

func (s *server) handleJournalUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}

func (s *server) handleJournalDelete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}

func (s *server) handleJournalByPatient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

}

func (s *server) handleAuthenticationHealth(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

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

func (s *server) handleAuthenticationEmployeeLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var login struct {
		Username string
		Password string
	}

	json.NewDecoder(r.Body).Decode(&login)

	a := &authentication.User{
		Username:       login.Username,
		HashedPassword: login.Password,
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
