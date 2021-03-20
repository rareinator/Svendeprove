package server

import (
	"github.com/rareinator/Svendeprove/Backend/packages/models"
)

func (s *Server) routes() {
	s.Router.Methods("OPTIONS").Handler(s.HandleCors())

	s.Router.Handle("/health", s.HandleHealth()).Methods("GET")

	//Journal methods
	s.Router.Handle("/journal/health", s.HandleJournalHealth()).Methods("GET")

	// swagger:operation POST /journal journal Create journal
	// ---
	// summary: Creates a new journal
	// description: Creates a new journal with the the values from the request body, returns the same body with the JournalId filled out
	// parameters:
	// - name: journal
	//   description: journal to add to the list of journals
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Journal"
	// responses:
	//   "200":
	//     "$ref": "#/definitions/Journal"
	s.Router.Handle("/journal", //Save journal
		s.Log(s.Authenticate(
			s.HandleJournalSave(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			}))).Methods("POST")

	// swagger:operation GET /journal journal Reads journal
	// ---
	// summary: Læser en ny journal
	// description: Creates a new journal with the the values from the request body, returns the same body with the JournalId filled out
	// parameters:
	// - name: journal
	//   description: journal to add to the list of journals
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Journal"
	// responses:
	//   "200":
	//     "$ref": "#/definitions/Journal"
	s.Router.Handle("/journal/{id:[0-9]+}", //Read journal
		s.Log(s.Authenticate(
			s.HandleJournalRead(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Office, models.Doctor, models.Nurse, models.Patient},
				allowedPatient: "",
			}))).Methods("GET")

	// s.Router.Handle("/journal/{id:[0-9]+}", //Update journal
	// 	s.Log(s.Authenticate(
	// 		s.HandleJournalUpdate(),
	// 		&authenticationConfig{
	// 			allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
	// 			allowedPatient: "",
	// 		}))).Methods("POST")

	// s.Router.Handle("/journal/{id:[0-9]+}", // Delete journal
	// 	s.Log(s.Authenticate(
	// 		s.HandleJournalDelete(),
	// 		&authenticationConfig{
	// 			allowedRoles:   []models.UserRole{models.Doctor},
	// 			allowedPatient: "",
	// 		}))).Methods("DELETE")

	s.Router.Handle("/journal/byPatient/{userId}", //Get patient journals
		s.Log(s.Authenticate(
			s.HandleJournalByPatient(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "userId",
			}))).Methods("GET")

	s.Router.Handle("/journal/document/{id:[0-9]+}", //Delete journal documents
		s.Log(s.Authenticate(
			s.HandleJournalDocumentDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("DELETE")

	s.Router.Handle("/journal/document/{id:[0-9]+}", //Update journal document
		s.Log(s.Authenticate(
			s.HandleJournalDocumentUpdate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/journal/document", //Create journal document
		s.Log(s.Authenticate(
			s.HandleJournalDocumentSave(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/journal/document/byJournal/{id:[0-9]+}", //Get journal documents by journalID
		s.Log(s.Authenticate(
			s.HandleJournalDocumentByJournal(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office, models.Patient},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/journal/document/{id:[0-9]+}", //Get journal document
		s.Log(s.Authenticate(
			s.HandleJournalDocumentRead(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office, models.Patient},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/journal/ml", //Upload images to ML
		s.Log(s.Authenticate(
			s.HandleJournalMLUpload(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Office, models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/journal/symptoms", //Upload symptoms to ML
		s.Log(s.Authenticate(s.HandleJournalUploadSymptoms(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	// Patient methods
	s.Router.Handle("/patient/health", s.HandlePatientHealth()).Methods("GET")

	s.Router.Handle("/patient", //GetPatients
		s.Log(s.Authenticate(
			s.HandlePatientsGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/diagnose/{id:[0-9]+}", //GetDiagnose
		s.Log(s.Authenticate(
			s.HandleDiagnoseGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/diagnose", //GetDiagnoses
		s.Log(s.Authenticate(
			s.HandleDiagnosesGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/symptom/{id:[0-9]+}", //GetSymptom
		s.Log(s.Authenticate(
			s.HandleSymptomGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/symptom", //GetSymptoms
		s.Log(s.Authenticate(
			s.HandleSymptomsGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/patient/{userId}/diagnose", //CreatePatientDiagnose
		s.Log(s.Authenticate(
			s.HandlePatientDiagnoseSave(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/patient/{userId}/diagnose", //GetDiagnoses
		s.Log(s.Authenticate(
			s.HandlePatientDiagnosesGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "userId",
			}))).Methods("GET")

	s.Router.Handle("/patient/{userId}/diagnose/{id:[0-9]+}", //GetDiagnose
		s.Log(s.Authenticate(
			s.HandlePatientDiagnoseGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "userId",
			}))).Methods("GET")

	s.Router.Handle("/patient/{userId}/diagnose/{id:[0-9]+}", //UpdateDiagnose
		s.Log(s.Authenticate(
			s.HandlePatientDiagnoseUpdate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/patient/{userId}/diagnose/{id:[0-9]+}",
		s.Log(s.Authenticate(
			s.HandlePatientDiagnoseDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			}))).Methods("DELETE")

	s.Router.Handle("/patient/{userId}/diagnose/{patientDiagnoseID:[0-9]+}/symptom", //CreatePatientSymptom
		s.Log(s.Authenticate(
			s.HandlePatientSymptomCreate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/patient/{userId}/diagnose/{diagnoseID:[0-9]+}/symptom", //GetPatientSymptoms
		s.Log(s.Authenticate(
			s.HandlePatientSymptomsGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "userId",
			}))).Methods("GET")

	s.Router.Handle("/patient/{userId}/diagnose/{diagnoseID:[0-9]+}/symptom/{id:[0-9]+}", //UpdatePatientSymptom
		s.Log(s.Authenticate(
			s.HandlePatientSymptomUpdate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/patient/{userId}/diagnose/{diagnoseID:[0-9]+}/symptom/{id:[0-9]+}", //DeletePatientSymptom
		s.Log(s.Authenticate(
			s.HandlePatientSymptomDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			}))).Methods("DELETE")

	//Booking methods
	s.Router.Handle("/booking/health", s.HandleBookingHealth()).Methods("GET")

	s.Router.Handle("/booking", //CreateBooking
		s.Log(s.Authenticate(
			s.HandleBookingCreate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office, models.Patient},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/booking/{id:[0-9]+}", //GetBooking
		s.Log(s.Authenticate(
			s.HandleBookingGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office},
				allowedPatient: "",
			}))).Methods("GET")

	s.Router.Handle("/booking/{id:[0-9]+}", //UpdateBooking
		s.Log(s.Authenticate(
			s.HandleBookingUpdate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office},
				allowedPatient: "",
			}))).Methods("POST")

	s.Router.Handle("/booking/{id:[0-9]+}", //Deletebooking
		s.Log(s.Authenticate(
			s.HandleBookingDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office},
				allowedPatient: "",
			}))).Methods("DELETE")

	s.Router.Handle("/booking/byPatient/{userId}", //GetBookingsByPatient
		s.Log(s.Authenticate(
			s.HandleBookingsByPatient(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse, models.Office},
				allowedPatient: "userId",
			}))).Methods("GET")

	s.Router.Handle("/booking/byEmployee/{userId}", //GetBookingsByEmployee
		s.Log(s.Authenticate(
			s.HandleBookingsByEmployee(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("GET")

	s.Router.Handle("/booking/availableTimesForDoctor",
		s.Log(s.Authenticate(
			s.HandleAvailableTimesForDoctor(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("POST")

	// Useradmin methods
	s.Router.Handle("/useradmin/health", s.HandleUseradminHealth()).Methods("GET") //GetHealth)

	s.Router.Handle("/useradmin/{userId}", //GetUser //TODO:
		s.Log(s.Authenticate(
			s.HandleUseradminGetEmployee(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse},
				allowedPatient: "userId",
			}))).Methods("GET")

	s.Router.Handle("/admin/hospitals", //GetHospitals
		s.Log(s.Authenticate(
			s.HandleHospitalsGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("GET")

	s.Router.Handle("/admin/departments",
		s.Log(s.Authenticate(
			s.HandleDepartmentsGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("GET")

	s.Router.Handle("/admin/beds",
		s.Log(s.Authenticate(
			s.HandleBedsGet(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("GET")

	s.Router.Handle("/admin/availableBeds",
		s.Log(s.Authenticate(
			s.HandleAvailableBeds(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("POST")

	s.Router.Handle("/admin/doctors/inHospital/{hospitalID:[0-9]+}",
		s.Log(s.Authenticate(
			s.HandleGetDoctorsInHospital(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse, models.Patient},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("GET")

	// IOT Methods
	s.Router.Handle("/iot/health", s.HandleIOTHealth()).Methods("GET") //GetHealth

	s.Router.Handle("/iot/uploadData",
		s.Log(
			s.HandleIOTUpload(),
		)).Queries("Key", "CGrtgtzxC0x5Ea6M", "SensorId", "", "Name", "", "Data", "").Methods("GET")

	s.Router.Handle("/iot/{deviceID:[0-9]+}",
		s.Log(s.Authenticate(
			s.HandleIOTReadData(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("GET")

	s.Router.Handle("/iot/readDataInTimeframe",
		s.Log(s.Authenticate(
			s.HandleIOTReadDataInTimeframe(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Office, models.Nurse},
				allowedPatient: "",
				allowIOTDevice: false,
			}))).Methods("POST")

}
