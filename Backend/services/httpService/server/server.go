package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rareinator/Svendeprove/Backend/packages/protocol"
)

type Server struct {
	Router         *mux.Router
	StaticFileDir  string
	JournalService protocol.JournalServiceClient
	PatientService protocol.PatientServiceClient
	BookingService protocol.BookingServiceClient
	AdminService   protocol.AdminServiceClient
	IotService     protocol.IotServiceClient
}

func NewServer() *Server {
	s := Server{}
	s.Router = mux.NewRouter()
	s.routes()
	return &s
}

func (s *Server) ServeHTTP() error {
	address := os.Getenv("HTTP_SERVICE_ADDR")
	fmt.Printf("🚀 Listening on address: %v ", address)

	corsHandler := corsHandler{
		router: s.Router,
	}

	s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(s.StaticFileDir))))

	if err := http.ListenAndServe(address, &corsHandler); err != nil {
		return err
	}

	return nil
}
