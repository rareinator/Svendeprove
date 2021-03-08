package main

import "github.com/rareinator/Svendeprove/Backend/packages/models"

func (s *server) routes() {
	s.router.Methods("OPTIONS").Handler(s.handleCors())

	s.router.Handle("/health", s.handleHealth()).Methods("GET")

	//Journal methods
	s.router.Handle("/journal/health", s.handleJournalHealth()).Methods("GET")
	s.router.Handle("/journal", s.authenticateRole(s.handleJournalSave(), models.Doctor)).Methods("POST")
	s.router.Handle("/journal/{id:[0-9]+}", s.authenticateRole(s.handleJournalRead(), models.Nurse)).Methods("GET")
	s.router.Handle("/journal/{id:[0-9]+}", s.handleJournalUpdate()).Methods("UPDATE")
	s.router.Handle("/journal/{id:[0-9]+", s.handleJournalDelete()).Methods("DELETE")
	s.router.Handle("/journal/byPatient/{id:[0-9]+}", s.authenticatePatient(s.handleJournalByPatient(), "id")).Methods("GET")

	s.router.Handle("/authentication/health", s.handleAuthenticationHealth()).Methods("GET")
	s.router.Handle("/authentication/patient/login", s.handleAuthenticationPatientLogin()).Methods("POST")
	s.router.Handle("/authentication/employee/login", s.handleAuthenticationEmployeeLogin()).Methods("POST")
}
