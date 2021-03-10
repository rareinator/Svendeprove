package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
	authenticationService "github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	journalService "github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
	patientService "github.com/rareinator/Svendeprove/Backend/services/patientService/patient"
)

func (s *server) returnError(w http.ResponseWriter, statusCode int, Message string) {
	var errorMessage struct {
		Code    int
		Message string
	}

	w.WriteHeader(statusCode)
	errorMessage.Code = statusCode
	errorMessage.Message = Message

	json.NewEncoder(w).Encode(&errorMessage)
}

func (s *server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("🚀 Server is up and running!!!!"))
	}
}

func (s *server) handleJournalHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		j := &journalService.JEmpty{}

		response, err := s.journalService.GetHealth(context.Background(), j)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the journal service %v", err))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.Message))
		}
	}
}

func (s *server) handleJournalSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var journal journalService.Journal
		json.NewDecoder(r.Body).Decode(&journal)

		response, err := s.journalService.CreateJournal(context.Background(), &journal)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	}
}

func (s *server) handleJournalRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		i, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal found for that id")
			return
		}

		j := &journalService.JournalRequest{
			JournalId: int32(i),
		}

		response, err := s.journalService.GetJournal(context.Background(), j)
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal found for that id")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleJournalUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal found for that id")
			return
		}

		var journal journalService.Journal
		json.NewDecoder(r.Body).Decode(&journal)

		journal.JournalId = int32(ID)

		response, err := s.journalService.UpdateJournal(context.Background(), &journal)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleJournalDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal found for that id")
		}

		response, err := s.journalService.DeleteJournal(context.Background(), &journalService.JournalRequest{JournalId: int32(ID)})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.returnError(w, http.StatusInternalServerError, "Something unknown went horribly wrong!!! ☠️☠️☠️")
	}
}

func (s *server) handleJournalByPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		patientID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journals found for that patientId")
			return
		}

		pr := &journalService.PatientRequest{
			PatientId: int32(patientID),
		}

		response, err := s.journalService.GetJournalsByPatient(context.Background(), pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the journal service %v", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Journals)
	}
}

func (s *server) handleJournalDocumentDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal document found with that id")
			return
		}

		jdr := journalService.JournalDocumentRequest{
			JournalDocumentId: int32(ID),
		}

		response, err := s.journalService.DeleteJournalDocument(context.Background(), &jdr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if !response.Success {
			s.returnError(w, http.StatusInternalServerError, "Something went horribly wrong!!! ☠️☠️☠️")
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) handleJournalDocumentUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal document found with that id")
			return
		}

		var journalDocument journalService.JournalDocument
		json.NewDecoder(r.Body).Decode(&journalDocument)

		journalDocument.DocumentId = int32(ID)

		response, err := s.journalService.UpdateJournalDocument(context.Background(), &journalDocument)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	}
}

func (s *server) handleJournalDocumentSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var journalDocument journalService.JournalDocument
		json.NewDecoder(r.Body).Decode(&journalDocument)

		response, err := s.journalService.CreateJournalDocument(context.Background(), &journalDocument)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleJournalDocumentByJournal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		journalID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal documents found for that journal id")
			return
		}

		pr := &journalService.JournalRequest{
			JournalId: int32(journalID),
		}

		response, err := s.journalService.GetJournalDocumentsByJournal(context.Background(), pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		patientID, err := s.getPatientID(r)
		if err != nil {
			s.returnError(w, http.StatusForbidden, "")
			return
		}

		allowed, err := s.patientIsAuthenticated(mssql.DBJournalDocument{}, response.JournalDocuments[0].DocumentId, patientID)
		if err != nil {
			s.returnError(w, http.StatusForbidden, err.Error())
			return
		}

		if allowed {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response.JournalDocuments)
		} else {
			s.returnError(w, http.StatusForbidden, "")
		}
	}
}

func (s *server) handleJournalDocumentRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		journalDocumentID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No journal document found for that id")
			return
		}

		j := journalService.JournalDocumentRequest{
			JournalDocumentId: int32(journalDocumentID),
		}

		response, err := s.journalService.GetJournalDocument(context.Background(), &j)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := &patientService.PEmpty{}

		response, err := s.patientService.GetHealth(context.Background(), p)
		if err != nil {
			s.returnError(w, http.StatusAccepted, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

func (s *server) handlePatientSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var patient patientService.Patient
		json.NewDecoder(r.Body).Decode(&patient)

		response, err := s.patientService.CreatePatient(context.Background(), &patient)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		patientID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No patient found with that id")
			return
		}

		p := patientService.PRequest{
			Id: int32(patientID),
		}

		response, err := s.patientService.ReadPatient(context.Background(), &p)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		patientID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No patient found for that id")
			return
		}

		var patient patientService.Patient
		json.NewDecoder(r.Body).Decode(&patient)

		patient.PatientId = int32(patientID)

		response, err := s.patientService.UpdatePatient(context.Background(), &patient)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleAuthenticationHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		e := &authenticationService.AEmpty{}

		response, err := s.authenticationService.GetHealth(context.Background(), e)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting in contact with the authentication service %v", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

func (s *server) handleAuthenticationEmployeeLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login struct {
			Username string
			Password string
		}

		json.NewDecoder(r.Body).Decode(&login)

		a := &authenticationService.User{
			Username: login.Username,
			Password: login.Password,
		}

		response, err := s.authenticationService.LoginEmployee(context.Background(), a)
		if err != nil {
			s.returnError(w, http.StatusForbidden, "Error logging in")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleAuthenticationPatientLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login struct {
			Username string
			Password string
		}

		json.NewDecoder(r.Body).Decode(&login)

		a := &authenticationService.User{
			Username: login.Username,
			Password: login.Password,
		}

		response, err := s.authenticationService.LoginPatient(context.Background(), a)
		if err != nil {
			s.returnError(w, http.StatusForbidden, "Error logging in")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
