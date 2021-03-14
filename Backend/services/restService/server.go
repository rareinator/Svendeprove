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
	"github.com/tidwall/buntdb"
)

type server struct {
	router                *mux.Router
	localDB               *buntdb.DB
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

	// Setting up oauth
	// manager := manage.NewDefaultManager()
	// manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	// manager.MustTokenStorage(store.NewMemoryTokenStore())

	// clientStore := store.NewClientStore()
	// clientStore.Set("000000", &models.Client{
	// 	ID:     "000000",
	// 	Secret: "999999",
	// 	Domain: "http://localhost:8080",
	// })
	// manager.MapClientStorage(clientStore)

	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))

	// // token, err := jwt.ParseWithClaims(access)

	// srv := oauthServer.NewDefaultServer(manager)
	// srv.SetAllowGetAccessRequest(true)
	// srv.SetClientInfoHandler(oauthServer.ClientFormHandler)

	// srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
	// 	fmt.Println("Internal error:", err.Error())
	// 	return
	// })

	// srv.SetResponseErrorHandler(func(re *errors.Response) {
	// 	fmt.Println("Response Error:", re.Error.Error())
	// })

	// srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
	// 	if username == "mni@hospi.local" && password == "P@ssw0rd" {
	// 		userID = "mni@hospi.local"
	// 	}
	// 	return
	// })

	// s.router.PathPrefix("/authorize").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	err := srv.HandleAuthorizeRequest(w, r)
	// 	if err != nil {
	// 		s.returnError(w, http.StatusBadRequest, err.Error())
	// 		return
	// 	}
	// })

	// s.router.PathPrefix("/token").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	srv.HandleTokenRequest(w, r)
	// })

	// s.router.Handle("/login", s.handleLogin()).Methods("POST")

	//hosting

	http.ListenAndServe(address, corsHandler)
}
