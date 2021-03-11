package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	"github.com/rareinator/Svendeprove/Backend/services/bookingService/booking"
	"github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
	"github.com/rareinator/Svendeprove/Backend/services/patientService/patient"
	"github.com/rareinator/Svendeprove/Backend/services/useradminService/useradmin"
)

type server struct {
	router                *mux.Router
	journalService        journal.JournalServiceClient
	authenticationService authentication.AuthenticationServiceClient
	patientService        patient.PatientServiceClient
	bookingService        booking.BookingServiceClient
	useradminService      useradmin.UseradminServiceClient
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

	corsHandler := &corsHandler{
		router: s.router,
	}

	http.ListenAndServe(address, corsHandler)
}
