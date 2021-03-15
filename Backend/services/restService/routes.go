package main

import (
	"github.com/rareinator/Svendeprove/Backend/packages/models"
)

func (s *server) routes() {
	s.router.Methods("OPTIONS").Handler(s.handleCors())

	s.router.Handle("/health", s.handleHealth()).Methods("GET")

	//oauth 2.0 implicit auth flow
	s.router.Handle("/auth", s.handleOauthAuth()).Queries("response_type", "code", "client_id", "000000", "scope", "{(\bpatient\b|\bemployee\b)}").Methods("POST")
	s.router.Handle("/token", s.handleOauthToken()).Methods("POST")

	//Journal methods
	s.router.Handle("/journal/health", s.handleJournalHealth()).Methods("GET")

	s.router.Handle("/journal", //Save journal
		s.log(s.authenticate(
			s.handleJournalSave(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			}))).Methods("POST")

	s.router.Handle("/journal/{id:[0-9]+}", //Read journal
		s.log(s.authenticate(
			s.handleJournalRead(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Employee, models.Doctor, models.Nurse},
				allowedPatient: "id",
			}))).Methods("GET")

	s.router.Handle("/journal/{id:[0-9]+}", //Update journal
		s.log(s.authenticate(
			s.handleJournalUpdate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.router.Handle("/journal/{id:[0-9]+}", // Delete journal
		s.log(s.authenticate(
			s.handleJournalDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			}))).Methods("DELETE")

	s.router.Handle("/journal/byPatient/{username}", //Get patient journals
		s.log(s.authenticate(
			s.handleJournalByPatient(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "id",
			}))).Methods("GET")

	s.router.Handle("/journal/document/{id:[0-9]+}", //Delete journal documents
		s.log(s.authenticate(
			s.handleJournalDocumentDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("DELETE")

	s.router.Handle("/journal/document/{id:[0-9]+}", //Update journal document
		s.log(s.authenticate(
			s.handleJournalDocumentUpdate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.router.Handle("/journal/document", //Create journal document
		s.log(s.authenticate(
			s.handleJournalDocumentSave(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.router.Handle("/journal/document/byJournal/{id:[0-9]+}", //Get journal documents by journalID
		s.log(s.authenticate(
			s.handleJournalDocumentByJournal(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: true,
			}))).Methods("GET")

	s.router.Handle("/journal/document/{id:[0-9]+}", //Get journal document
		s.log(s.authenticate(
			s.handleJournalDocumentRead(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("GET")

	s.router.Handle("/journal/ml", //Upload images to ML
		s.log(s.authenticate(
			s.handleJournalMLUpload(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Employee, models.Doctor, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("POST")

	s.router.Handle("/journal/symptoms", //Upload symptoms to ML
		s.log(s.authenticate(s.handleJournalUploadSymptoms(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Employee, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("POST")

	// Authentication methods
	s.router.Handle("/authentication/health", s.handleAuthenticationHealth()).Methods("GET")

	// s.router.Handle("/authentication/patient/login", s.handleAuthenticationPatientLogin()).Methods("POST")
	// s.router.Handle("/authentication/employee/login", s.handleAuthenticationEmployeeLogin()).Methods("POST")

	// Patient methods
	s.router.Handle("/patient/health", s.handlePatientHealth()).Methods("GET")

	s.router.Handle("/patient/{username}", //GetPatient //TODO:
		s.log(s.authenticate(
			s.handlePatientRead(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse, models.Employee},
				allowedPatient:      "id",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/patient", //GetPatients
		s.log(s.authenticate(
			s.handlePatientsGet(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse, models.Employee},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/diagnose/{id:[0-9]+}", //GetDiagnose
		s.log(s.authenticate(
			s.handleDiagnoseGet(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/diagnose", //GetDiagnoses
		s.log(s.authenticate(
			s.handleDiagnosesGet(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/symptom/{id:[0-9]+}", //GetSymptom
		s.log(s.authenticate(
			s.handleSymptomGet(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/symptom", //GetSymptoms
		s.log(s.authenticate(
			s.handleSymptomsGet(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/patient/{username}/diagnose", //CreatePatientDiagnose
		s.log(s.authenticate(
			s.handlePatientDiagnoseSave(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("POST")

	s.router.Handle("/patient/{username}/diagnose", //GetDiagnoses
		s.log(s.authenticate(
			s.handlePatientDiagnosesGet(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "patientID",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/patient/{username}/diagnose/{id:[0-9]+}", //GetDiagnose
		s.log(s.authenticate(
			s.handlePatientDiagnoseGet(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "patientID",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/patient/{username}/diagnose/{id:[0-9]+}", //UpdateDiagnose
		s.log(s.authenticate(
			s.handlePatientDiagnoseUpdate(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("POST")

	s.router.Handle("/patient/{username}/diagnose/{id:[0-9]+}",
		s.log(s.authenticate(
			s.handlePatientDiagnoseDelete(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("DELETE")

	s.router.Handle("/patient/{username}/diagnose/{diagnoseID:[0-9]+}/symptom", //CreatePatientSymptom
		s.log(s.authenticate(
			s.handlePatientSymptomCreate(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("POST")

	s.router.Handle("/patient/{username}/diagnose/{diagnoseID:[0-9]+}/symptom", //GetPatientSymptoms
		s.log(s.authenticate(
			s.handlePatientSymptomsGet(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "patientID",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/patient/{username}/diagnose/{diagnoseID:[0-9]+}/symptom/{id:[0-9]+}", //UpdatePatientSymptom
		s.log(s.authenticate(
			s.handlePatientSymptomUpdate(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("POST")

	s.router.Handle("/patient/{username}/diagnose/{diagnoseID:[0-9]+}/symptom/{id:[0-9]+}", //DeletePatientSymptom
		s.log(s.authenticate(
			s.handlePatientSymptomDelete(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("DELETE")

	//Booking methods
	s.router.Handle("/booking/health", s.handleBookingHealth()).Methods("GET")

	s.router.Handle("/booking", //CreateBooking
		s.log(s.authenticate(
			s.handleBookingCreate(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse, models.Employee},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("POST")

	s.router.Handle("/booking/{id:[0-9]+}", //GetBooking
		s.log(s.authenticate(
			s.handleBookingGet(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse, models.Employee},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/booking/{id:[0-9]+}", //UpdateBooking
		s.log(s.authenticate(
			s.handleBookingUpdate(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse, models.Employee},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("POST")

	s.router.Handle("/booking/{id:[0-9]+}", //Deletebooking
		s.log(s.authenticate(
			s.handleBookingDelete(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse, models.Employee},
				allowedPatient:      "",
				allowRelatedPatient: false,
			}))).Methods("DELETE")

	s.router.Handle("/booking/byPatient/{username}", //GetBookingsByPatient
		s.log(s.authenticate(
			s.handleBookingsByPatient(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Nurse, models.Employee},
				allowedPatient:      "id",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/booking/byEmployee/{username}", //GetBookingsByEmployee
		s.log(s.authenticate(
			s.handleBookingsByEmployee(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Employee, models.Nurse, models.Patient},
				allowedPatient:      "",
				allowRelatedPatient: false,
				allowIOTDevice:      false,
			}))).Methods("GET")

	s.router.Handle("/booking/availableTimesForDoctor",
		s.log(s.authenticate(
			s.handleAvailableTimesForDoctor(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Employee, models.Nurse, models.Patient},
				allowedPatient:      "",
				allowRelatedPatient: false,
				allowIOTDevice:      false,
			}))).Methods("POST")

	// Useradmin methods
	s.router.Handle("/useradmin/health", s.handleUseradminHealth()).Methods("GET") //GetHealth)

	s.router.Handle("/useradmin/{username}", //GetUser //TODO:
		s.log(s.authenticate(
			s.handleUseradminGetEmployee(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Employee, models.Nurse},
				allowedPatient:      "username",
				allowRelatedPatient: false,
			}))).Methods("GET")

	s.router.Handle("/admin/hospitals", //GetHospitals
		s.log(s.authenticate(
			s.handleHospitalsGet(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Employee, models.Nurse, models.Patient},
				allowedPatient:      "",
				allowRelatedPatient: false,
				allowIOTDevice:      false,
			}))).Methods("GET")

	s.router.Handle("/admin/doctors/inHospital/{hospitalID:[0-9]+}",
		s.log(s.authenticate(
			s.handleGetDoctorsInHospital(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Employee, models.Nurse, models.Patient},
				allowedPatient:      "",
				allowRelatedPatient: false,
				allowIOTDevice:      false,
			}))).Methods("GET")

	// IOT Methods
	s.router.Handle("/iot/health", s.handleIOTHealth()).Methods("GET") //GetHealth

	s.router.Handle("/iot/uploadData",
		s.log(s.authenticate(
			s.handleIOTUpload(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{},
				allowedPatient:      "",
				allowRelatedPatient: false,
				allowIOTDevice:      true,
			}))).Queries("Key", "", "Data", "").Methods("POST")

	s.router.Handle("/iot/{deviceID:[0-9]+}",
		s.log(s.authenticate(
			s.handleIOTReadData(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Employee, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
				allowIOTDevice:      false,
			}))).Methods("GET")

	s.router.Handle("/iot/readDataInTimeframe",
		s.log(s.authenticate(
			s.handleIOTReadDataInTimeframe(),
			&authenticationConfig{
				allowedRoles:        []models.UserRole{models.Doctor, models.Employee, models.Nurse},
				allowedPatient:      "",
				allowRelatedPatient: false,
				allowIOTDevice:      false,
			}))).Methods("POST")

}
