package main

func (s *server) routes() {
	s.router.HandleFunc("/health", s.handleHealth).Methods("GET")

	//Journal methods
	s.router.HandleFunc("/journal/health", s.handleJournalHealth).Methods("GET")
	s.router.HandleFunc("/journal", s.handleJournalSave).Methods("POST")
	s.router.HandleFunc("/journal/{id:[0-9]+}", s.handleJournalRead).Methods("GET")
	s.router.HandleFunc("/journal/{id:[0-9]+}", s.handleJournalUpdate).Methods("UPDATE")
	s.router.HandleFunc("/journal/{id:[0-9]+", s.handleJournalDelete).Methods("DELETE")
	s.router.HandleFunc("/journal/byPatient/{id:[0-9]+}", s.handleJournalByPatient).Methods("GET")

}
