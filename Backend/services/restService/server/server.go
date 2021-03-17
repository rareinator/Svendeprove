package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rareinator/Svendeprove/Backend/packages/protocol"
	"github.com/tidwall/buntdb"
)

type Server struct {
	Router           *mux.Router
	LocalDB          *buntdb.DB
	StaticFileDir    string
	JournalService   protocol.JournalServiceClient
	PatientService   protocol.PatientServiceClient
	BookingService   protocol.BookingServiceClient
	UseradminService protocol.UseradminServiceClient
	IotService       protocol.IotServiceClient
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
