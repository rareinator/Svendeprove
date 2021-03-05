package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
)

type server struct {
	router         *mux.Router
	journalService journal.JournalServiceClient
}

func newServer() *server {
	s := &server{}
	s.router = mux.NewRouter()
	s.routes()
	return s
}

func (s *server) ServeHTTP() {
	address := ":8080"
	fmt.Printf("Listening on address: %v ", address)
	http.ListenAndServe(":8080", s.router)
}

func (s *server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is op and running!!!!"))
}

func (s *server) handleJournalHealth(w http.ResponseWriter, r *http.Request) {
	j := &journal.Journal{
		Intro: "hello from the client",
	}

	responseJournal, err := s.journalService.GetJournal(context.Background(), j)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error getting in contact with the journal service %v", err)))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseJournal.Intro))
	}
}

func (s *server) handleJournalSave(w http.ResponseWriter, r *http.Request) {

}

func (s *server) handleJournalRead(w http.ResponseWriter, r *http.Request) {

}

func (s *server) handleJournalUpdate(w http.ResponseWriter, r *http.Request) {

}

func (s *server) handleJournalDelete(w http.ResponseWriter, r *http.Request) {

}

func (s *server) handleJournalByPatient(w http.ResponseWriter, r *http.Request) {

}
