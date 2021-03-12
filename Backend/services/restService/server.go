package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	"github.com/rareinator/Svendeprove/Backend/services/bookingService/booking"
	"github.com/rareinator/Svendeprove/Backend/services/iotService/iot"
	"github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
	"github.com/rareinator/Svendeprove/Backend/services/patientService/patient"
	"github.com/rareinator/Svendeprove/Backend/services/useradminService/useradmin"
)

type server struct {
	router                *mux.Router
	staticFileDir         string
	journalService        journal.JournalServiceClient
	authenticationService authentication.AuthenticationServiceClient
	patientService        patient.PatientServiceClient
	bookingService        booking.BookingServiceClient
	useradminService      useradmin.UseradminServiceClient
	iotService            iot.IotServiceClient
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

	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(s.staticFileDir))))

	http.ListenAndServe(address, corsHandler)
}
