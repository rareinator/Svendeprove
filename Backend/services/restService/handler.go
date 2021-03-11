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
	bookingService "github.com/rareinator/Svendeprove/Backend/services/bookingService/booking"
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

		employeeID, err := s.getEmployeeID(r)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		journal.CreatedBy = employeeID

		fmt.Printf("TEST!!!! %v\n\r", employeeID)

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
		if len(response.Journals) == 0 {
			response.Journals = make([]*journalService.Journal, 0)
		}
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

		fmt.Println("Inserting Journal Document")

		employeeID, err := s.getEmployeeID(r)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		journalDocument.CreatedBy = employeeID
		fmt.Printf("EmployeeID Got: %v\n\r", employeeID)

		response, err := s.journalService.CreateJournalDocument(context.Background(), &journalDocument)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		fmt.Println("Called journalService")

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

		var allowed bool
		if patientID == 0 {
			allowed = true
		} else {
			allowed = false
			if len(response.JournalDocuments) > 0 {
				allowed, err = s.patientIsAuthenticated(mssql.DBJournalDocument{}, response.JournalDocuments[0].DocumentId, patientID)
				if err != nil {
					s.returnError(w, http.StatusForbidden, err.Error())
					return
				}
			} else {
				allowed = true
			}
		}

		if allowed {
			w.WriteHeader(http.StatusOK)
			if len(response.JournalDocuments) == 0 {
				response.JournalDocuments = make([]*journalService.JournalDocument, 0)
			}
			json.NewEncoder(w).Encode(response.JournalDocuments)
		} else {
			s.returnError(w, http.StatusForbidden, "Not Allowed")
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

		response, err := s.patientService.GetPatient(context.Background(), &p)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := s.patientService.GetPatients(context.Background(), &patientService.PEmpty{})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Patients) == 0 {
			response.Patients = make([]*patientService.Patient, 0)
		}
		json.NewEncoder(w).Encode(response.Patients)
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

func (s *server) handlePatientDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		patientID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No patient found for that id")
			return
		}

		response, err := s.patientService.DeletePatient(context.Background(), &patientService.PRequest{Id: int32(patientID)})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.returnError(w, http.StatusInternalServerError, "Somethin unknown went gorribly wrong!!! ☠️☠️☠️")
	}
}

func (s *server) handleDiagnoseGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		p := patientService.PRequest{
			Id: int32(id),
		}

		response, err := s.patientService.GetDiagnose(context.Background(), &p)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleDiagnosesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := s.patientService.GetDiagnoses(context.Background(), &patientService.PEmpty{})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Diagnoses) == 0 {
			response.Diagnoses = make([]*patientService.Diagnose, 0)
		}
		json.NewEncoder(w).Encode(response.Diagnoses)
	}
}

func (s *server) handleSymptomGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No symptom found with that id")
			return
		}

		p := patientService.PRequest{
			Id: int32(id),
		}

		response, err := s.patientService.GetSymptom(context.Background(), &p)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)
	}
}

func (s *server) handleSymptomsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := s.patientService.GetSymptoms(context.Background(), &patientService.PEmpty{})
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Symptoms) == 0 {
			response.Symptoms = make([]*patientService.Symptom, 0)
		}
		json.NewEncoder(w).Encode(response.Symptoms)

	}
}

func (s *server) handlePatientDiagnoseSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		patientID, err := strconv.Atoi(vars["patientID"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "patient with that id not found")
			return
		}
		var patientDiagnose patientService.PatientDiagnose
		json.NewDecoder(r.Body).Decode(&patientDiagnose)
		patientDiagnose.PatientId = int32(patientID)

		response, err := s.patientService.CreatePatientDiagnose(context.Background(), &patientDiagnose)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientDiagnosesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		patientID, err := strconv.Atoi(vars["patientID"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "Patient for that id not found")
			return
		}

		pr := patientService.PRequest{
			Id: int32(patientID),
		}

		response, err := s.patientService.GetPatientDiagnoses(context.Background(), &pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.PatientDiagnoses) == 0 {
			response.PatientDiagnoses = make([]*patientService.PatientDiagnose, 0)
		}
		json.NewEncoder(w).Encode(response.PatientDiagnoses)
	}
}

func (s *server) handlePatientDiagnoseGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "Diagnose for that id not found")
			return
		}

		pr := patientService.PRequest{
			Id: int32(id),
		}

		response, err := s.patientService.GetPatientDiagnose(context.Background(), &pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientDiagnoseUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "Diagnose for that id not found")
			return
		}

		var patientDiagnose patientService.PatientDiagnose
		json.NewDecoder(r.Body).Decode(&patientDiagnose)

		patientDiagnose.PatientDiagnoseId = int32(ID)

		response, err := s.patientService.UpdatePatientDiagnose(context.Background(), &patientDiagnose)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientDiagnoseDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No patient Diagnosis found for that id")
			return
		}

		response, err := s.patientService.DeletePatientDiagnose(context.Background(), &patientService.PRequest{Id: int32(ID)})
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

func (s *server) handlePatientSymptomCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var diagnoseSymptom patientService.DiagnoseSymptom
		json.NewDecoder(r.Body).Decode(&diagnoseSymptom)

		response, err := s.patientService.CreateDiagnoseSymptom(context.Background(), &diagnoseSymptom)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientSymptomsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		diagnoseID, err := strconv.Atoi(vars["diagnoseID"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		pr := patientService.PRequest{
			Id: int32(diagnoseID),
		}

		response, err := s.patientService.GetDiagnoseSymptoms(context.Background(), &pr)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.DiagnoseSymptoms) == 0 {
			response.DiagnoseSymptoms = make([]*patientService.DiagnoseSymptom, 0)
		}
		json.NewEncoder(w).Encode(response.DiagnoseSymptoms)
	}
}

func (s *server) handlePatientSymptomUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		diagnoseID, err := strconv.Atoi(vars["diagnoseID"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		var newDiagnoseSymptom patientService.DiagnoseSymptom
		json.NewDecoder(r.Body).Decode(&newDiagnoseSymptom)

		oldDiagnoseSymptom := patientService.DiagnoseSymptom{
			SymptomId:         int32(ID),
			PatientDiagnoseId: int32(diagnoseID),
		}

		dsur := patientService.DiagnoseSymptomUpdateRequest{
			Old: &oldDiagnoseSymptom,
			New: &newDiagnoseSymptom,
		}

		response, err := s.patientService.UpdateDiagnoseSymptom(context.Background(), &dsur)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handlePatientSymptomDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		diagnoseID, err := strconv.Atoi(vars["diagnoseID"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		ds := patientService.DiagnoseSymptom{
			SymptomId:         int32(ID),
			PatientDiagnoseId: int32(diagnoseID),
		}

		response, err := s.patientService.DeleteDiagnoseSymptom(context.Background(), &ds)
		if err != nil {
			s.returnError(w, http.StatusNotFound, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.returnError(w, http.StatusInternalServerError, "Something unknown went horribly wrong!!! ☠️☠️☠️")
	}
}

func (s *server) handleBookingHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b := bookingService.BEmpty{}

		response, err := s.bookingService.GetHealth(context.Background(), &b)
		if err != nil {
			s.returnError(w, http.StatusServiceUnavailable, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

func (s *server) handleBookingCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var booking bookingService.Booking
		json.NewDecoder(r.Body).Decode(&booking)

		employeeID, err := s.getEmployeeID(r)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		booking.ApprovedByEmployee = employeeID

		response, err := s.bookingService.CreateBooking(context.Background(), &booking)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleBookingGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No booking found with that id")
			return
		}

		b := bookingService.BRequest{
			Id: int32(ID),
		}

		response, err := s.bookingService.GetBooking(context.Background(), &b)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleBookingUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookingID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No booking found with that id")
			return
		}

		var booking bookingService.Booking
		json.NewDecoder(r.Body).Decode(&booking)

		booking.BookingId = int32(bookingID)

		response, err := s.bookingService.UpdateBooking(context.Background(), &booking)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *server) handleBookingDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookingID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.returnError(w, http.StatusNotFound, "No booking with that id found")
			return
		}

		br := bookingService.BRequest{
			Id: int32(bookingID),
		}

		response, err := s.bookingService.DeleteBooking(context.Background(), &br)
		if err != nil {
			s.returnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.returnError(w, http.StatusInternalServerError, "Somethin unknown went gorribly wrong!!! ☠️☠️☠️")
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
