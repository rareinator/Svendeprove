package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	protocol "github.com/rareinator/Svendeprove/Backend/packages/protocol"
)

func (s *Server) HandlePatientHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := s.PatientService.GetHealth(context.Background(), &protocol.Empty{})
		if err != nil {
			s.ReturnError(w, http.StatusAccepted, err.Error())
			return
		}

		if _, err := w.Write([]byte(response.Message)); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

type doctor struct {
	Username string `json:"Username"`
	Name     string `json:"Name"`
	Type     string `json:"Type"`
	UserId   string `json:"UserId"`
}

func (s *Server) HandleDiagnosesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := s.PatientService.GetDiagnoses(context.Background(), &protocol.Empty{})
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Diagnoses) == 0 {
			response.Diagnoses = make([]*protocol.Diagnose, 0)
		}
		if err := json.NewEncoder(w).Encode(response.Diagnoses); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func (s *Server) HandleSymptomsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := s.PatientService.GetSymptoms(context.Background(), &protocol.Empty{})
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(response.Symptoms) == 0 {
			response.Symptoms = make([]*protocol.Symptom, 0)
		}

		if err := json.NewEncoder(w).Encode(response.Symptoms); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)

	}
}

func (s *Server) HandlePatientDiagnoseSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var patientDiagnose protocol.PatientDiagnose
		if err := json.NewDecoder(r.Body).Decode(&patientDiagnose); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		patientDiagnose.Patient = vars["userId"]

		response, err := s.PatientService.CreatePatientDiagnose(context.Background(), &patientDiagnose)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) HandlePatientDiagnosesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		pr := protocol.Request{
			UserId: vars["userId"],
		}

		response, err := s.PatientService.GetPatientDiagnoses(context.Background(), &pr)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(response.PatientDiagnoses) == 0 {
			response.PatientDiagnoses = make([]*protocol.PatientDiagnose, 0)
		}
		if err := json.NewEncoder(w).Encode(response.PatientDiagnoses); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) HandlePatientDiagnoseDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No patient Diagnosis found for that id")
			return
		}

		response, err := s.PatientService.DeletePatientDiagnose(context.Background(), &protocol.Request{Id: int32(ID)})
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.ReturnError(w, http.StatusInternalServerError, "Something unknown went horribly wrong!!! ☠️☠️☠️")
	}
}

func (s *Server) HandlePatientSymptomCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var diagnoseSymptom protocol.DiagnoseSymptom
		if err := json.NewDecoder(r.Body).Decode(&diagnoseSymptom); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["patientDiagnoseID"])
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		diagnoseSymptom.PatientDiagnoseId = int32(ID)

		response, err := s.PatientService.CreateDiagnoseSymptom(context.Background(), &diagnoseSymptom)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) HandlePatientSymptomsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["diagnoseID"])
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		pr := protocol.Request{
			Id: int32(id),
		}

		response, err := s.PatientService.GetDiagnoseSymptoms(context.Background(), &pr)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.DiagnoseSymptoms) == 0 {
			response.DiagnoseSymptoms = make([]*protocol.DiagnoseSymptom, 0)
		}

		if err := json.NewEncoder(w).Encode(response.DiagnoseSymptoms); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func (s *Server) HandlePatientSymptomUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		diagnoseID, err := strconv.Atoi(vars["diagnoseID"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		var newDiagnoseSymptom protocol.DiagnoseSymptom
		if err := json.NewDecoder(r.Body).Decode(&newDiagnoseSymptom); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		oldDiagnoseSymptom := protocol.DiagnoseSymptom{
			SymptomId:         int32(ID),
			PatientDiagnoseId: int32(diagnoseID),
		}

		dsur := protocol.DiagnoseSymptomUpdateRequest{
			Old: &oldDiagnoseSymptom,
			New: &newDiagnoseSymptom,
		}

		response, err := s.PatientService.UpdateDiagnoseSymptom(context.Background(), &dsur)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) HandlePatientSymptomDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		diagnoseID, err := strconv.Atoi(vars["diagnoseID"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No diagnose found with that id")
			return
		}

		ds := protocol.DiagnoseSymptom{
			SymptomId:         int32(ID),
			PatientDiagnoseId: int32(diagnoseID),
		}

		response, err := s.PatientService.DeleteDiagnoseSymptom(context.Background(), &ds)
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.ReturnError(w, http.StatusInternalServerError, "Something unknown went horribly wrong!!! ☠️☠️☠️")
	}
}
