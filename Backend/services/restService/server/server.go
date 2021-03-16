package server

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
	"github.com/tidwall/buntdb"
)

type Server struct {
	Router                *mux.Router
	LocalDB               *buntdb.DB
	StaticFileDir         string
	JournalService        journal.JournalServiceClient
	AuthenticationService authentication.AuthenticationServiceClient
	PatientService        patient.PatientServiceClient
	BookingService        booking.BookingServiceClient
	UseradminService      useradmin.UseradminServiceClient
	IotService            iot.IotServiceClient
}

func NewServer() *Server {
	s := &Server{}
	s.Router = mux.NewRouter()
	s.routes()
	return s
}

func (s *Server) ServeHTTP() {
	address := os.Getenv("REST_SERVICE_ADDR")
	fmt.Printf("🚀 Listening on address: %v ", address)

	corsHandler := &corsHandler{
		router: s.Router,
	}

	s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(s.StaticFileDir))))

	http.ListenAndServe(address, corsHandler)
}
