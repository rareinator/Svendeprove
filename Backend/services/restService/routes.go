package main

import "github.com/rareinator/Svendeprove/Backend/packages/models"

func (s *server) routes() {
	s.router.Methods("OPTIONS").Handler(s.handleCors())

	s.router.Handle("/health", s.handleHealth()).Methods("GET")

	//Journal methods
	s.router.Handle("/journal/health", s.handleJournalHealth()).Methods("GET")

	s.router.Handle("/journal", //Save journal
		s.authenticate(
			s.handleJournalSave(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			})).Methods("POST")

	s.router.Handle("/journal/{id:[0-9]+}", //Read journal
		s.authenticate(
			s.handleJournalRead(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Employee, models.Doctor, models.Nurse},
				allowedPatient: "id",
			})).Methods("GET")

	s.router.Handle("/journal/{id:[0-9]+}", //Update journal
		s.authenticate(
			s.handleJournalUpdate(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "",
			})).Methods("POST")

	s.router.Handle("/journal/{id:[0-9]+", // Delete journal
		s.authenticate(
			s.handleJournalDelete(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor},
				allowedPatient: "",
			})).Methods("DELETE")

	s.router.Handle("/journal/byPatient/{id:[0-9]+}", //Get patient journals
		s.authenticate(
			s.handleJournalByPatient(),
			&authenticationConfig{
				allowedRoles:   []models.UserRole{models.Doctor, models.Nurse},
				allowedPatient: "id",
			})).Methods("GET")

	//delete journal document, doctor and nurse
	//update journal document, doctor and nurse
	//get journal documents, doctor nurse, and patient
	//get journal document, doctor, nurse and patient

	// Authentication methods
	s.router.Handle("/authentication/health", s.handleAuthenticationHealth()).Methods("GET")

	s.router.Handle("/authentication/patient/login", s.handleAuthenticationPatientLogin()).Methods("POST")
	s.router.Handle("/authentication/employee/login", s.handleAuthenticationEmployeeLogin()).Methods("POST")

	// Patient methods
	//get all patients, doctor and nurse
	//get one patient, doctor nurse and correct patient
	//save new patient, doctor, nurse and employee
	//update patient, doctor, nurse and employee
	//
}
