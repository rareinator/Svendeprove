package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
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
	w.Write([]byte("Server is op and running!!!! :D POG"))
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
