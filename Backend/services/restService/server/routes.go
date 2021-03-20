package server

import (
	"github.com/rareinator/Svendeprove/Backend/packages/models"
)

func (s *Server) routes() {
	s.Router.Methods("OPTIONS").Handler(s.handleCors())

	s.Router.Handle("/health", s.handleHealth()).Methods("GET")

	//Journal methods
	s.Router.Handle("/journal/health", s.handleJournalHealth()).Methods("GET")

	s.Router.Handle("/journal", //Save journal
		s.Log(s.Authenticate(
			s.handleJournalSave(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/journal/{id:[0-9]+}", //Update journal
		s.Log(s.Authenticate(
			s.handleJournalUpdate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/journal/{id:[0-9]+}", // Delete journal
		s.Log(s.Authenticate(
			s.handleJournalDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			}))).Methods("DELETE")

	s.Router.Handle("/journal/byPatient/{userId}", //Get patient journals
		s.Log(s.Authenticate(
			s.handleJournalByPatient(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "userId",
			}))).Methods("GET")

	s.Router.Handle("/journal/document/{id:[0-9]+}", //Delete journal documents
		s.Log(s.Authenticate(
			s.handleJournalDocumentDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("DELETE")

	s.Router.Handle("/journal/document/{id:[0-9]+}", //Update journal document
		s.Log(s.Authenticate(
			s.handleJournalDocumentUpdate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/journal/document", //Create journal document
		s.Log(s.Authenticate(
			s.handleJournalDocumentSave(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/journal/document/byJournal/{id:[0-9]+}", //Get journal documents by journalID
		s.Log(s.Authenticate(
			s.handleJournalDocumentByJournal(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office, models.Patient},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/journal/document/{id:[0-9]+}", //Get journal document
		s.Log(s.Authenticate(
			s.handleJournalDocumentRead(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office, models.Patient},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/journal/ml", //Upload images to ML
		s.Log(s.Authenticate(
			s.handleJournalMLUpload(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Office, models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/journal/symptoms", //Upload symptoms to ML
		s.Log(s.Authenticate(s.handleJournalUploadSymptoms(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	// Patient methods
	s.Router.Handle("/patient/health", s.handlePatientHealth()).Methods("GET")

	s.Router.Handle("/patient", //GetPatients
		s.Log(s.Authenticate(
			s.handlePatientsGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/diagnose", //GetDiagnoses
		s.Log(s.Authenticate(
			s.handleDiagnosesGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/symptom", //GetSymptoms
		s.Log(s.Authenticate(
			s.handleSymptomsGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/patient/{userId}/diagnose", //CreatePatientDiagnose
		s.Log(s.Authenticate(
			s.handlePatientDiagnoseSave(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/patient/{userId}/diagnose", //GetDiagnoses
		s.Log(s.Authenticate(
			s.handlePatientDiagnosesGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "userId",
			}))).Methods("GET")

	s.Router.Handle("/patient/{userId}/diagnose/{id:[0-9]+}",
		s.Log(s.Authenticate(
			s.handlePatientDiagnoseDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			}))).Methods("DELETE")

	s.Router.Handle("/patient/{userId}/diagnose/{patientDiagnoseID:[0-9]+}/symptom", //CreatePatientSymptom
		s.Log(s.Authenticate(
			s.handlePatientSymptomCreate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/patient/{userId}/diagnose/{diagnoseID:[0-9]+}/symptom", //GetPatientSymptoms
		s.Log(s.Authenticate(
			s.handlePatientSymptomsGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "userId",
			}))).Methods("GET")

	s.Router.Handle("/patient/{userId}/diagnose/{diagnoseID:[0-9]+}/symptom/{id:[0-9]+}", //UpdatePatientSymptom
		s.Log(s.Authenticate(
			s.handlePatientSymptomUpdate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/patient/{userId}/diagnose/{diagnoseID:[0-9]+}/symptom/{id:[0-9]+}", //DeletePatientSymptom
		s.Log(s.Authenticate(
			s.handlePatientSymptomDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("DELETE")

	//Booking methods
	s.Router.Handle("/booking/health", s.handleBookingHealth()).Methods("GET")

	s.Router.Handle("/booking", //CreateBooking
		s.Log(s.Authenticate(
			s.handleBookingCreate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office, models.Patient},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/booking/{id:[0-9]+}", //Deletebooking
		s.Log(s.Authenticate(
			s.handleBookingDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office},
				allowedPatient: "",
			}))).Methods("DELETE")

	s.Router.Handle("/booking/byPatient/{userId}", //GetBookingsByPatient
		s.Log(s.Authenticate(
			s.handleBookingsByPatient(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office},
				allowedPatient: "userId",
			}))).Methods("GET")

	s.Router.Handle("/booking/byEmployee/{userId}", //GetBookingsByEmployee
		s.Log(s.Authenticate(
			s.handleBookingsByEmployee(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("GET")

	s.Router.Handle("/booking/availableTimesForDoctor",
		s.Log(s.Authenticate(
			s.handleAvailableTimesForDoctor(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("POST")

	// Useradmin methods
	s.Router.Handle("/useradmin/health", s.handleUseradminHealth()).Methods("GET") //GetHealth)

	s.Router.Handle("/admin/hospitals", //GetHospitals
		s.Log(s.Authenticate(
			s.handleHospitalsGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("GET")

	s.Router.Handle("/admin/availableBeds",
		s.Log(s.Authenticate(
			s.handleAvailableBeds(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("POST")

	s.Router.Handle("/admin/doctors/inHospital/{hospitalID:[0-9]+}",
		s.Log(s.Authenticate(
			s.handleGetDoctorsInHospital(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("GET")

	// IOT Methods
	s.Router.Handle("/iot/health", s.handleIOTHealth()).Methods("GET") //GetHealth

	s.Router.Handle("/iot/uploadData",
		s.Log(
			s.handleIOTUpload(),
		)).Queries("Key", "CGrtgtzxC0x5Ea6M", "SensorId", "", "Name", "", "Data", "").Methods("GET")

	s.Router.Handle("/iot/{deviceID:[0-9]+}",
		s.Log(s.Authenticate(
			s.handleIOTReadData(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("GET")

	s.Router.Handle("/iot/readDataInTimeframe",
		s.Log(s.Authenticate(
			s.handleIOTReadDataInTimeframe(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("POST")

}
