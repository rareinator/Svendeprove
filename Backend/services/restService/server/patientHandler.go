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
		p := &protocol.Empty{}

		response, err := s.PatientService.GetHealth(context.Background(), p)
		if err != nil {
			s.ReturnError(w, http.StatusAccepted, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

type doctor struct {
	Username string `json:"Username"`
	Name     string `json:"Name"`
	Type     string `json:"Type"`
	UserId   string `json:"UserId"`
}

func (s *Server) HandleDiagnoseGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		p := protocol.Request{
			UserId: vars["userId"],
		}

		response, err := s.PatientService.GetDiagnose(context.Background(), &p)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
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
		json.NewEncoder(w).Encode(response.Diagnoses)
	}
}

func (s *Server) HandleSymptomGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		p := protocol.Request{
			UserId: vars["userId"],
		}

		response, err := s.PatientService.GetSymptom(context.Background(), &p)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)
	}
}

func (s *Server) HandleSymptomsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := s.PatientService.GetSymptoms(context.Background(), &protocol.Empty{})
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Symptoms) == 0 {
			response.Symptoms = make([]*protocol.Symptom, 0)
		}
		json.NewEncoder(w).Encode(response.Symptoms)

	}
}

func (s *Server) HandlePatientDiagnoseSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var patientDiagnose protocol.PatientDiagnose
		json.NewDecoder(r.Body).Decode(&patientDiagnose)
		patientDiagnose.Patient = vars["userId"]

		response, err := s.PatientService.CreatePatientDiagnose(context.Background(), &patientDiagnose)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
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

		w.WriteHeader(http.StatusOK)
		if len(response.PatientDiagnoses) == 0 {
			response.PatientDiagnoses = make([]*protocol.PatientDiagnose, 0)
		}
		json.NewEncoder(w).Encode(response.PatientDiagnoses)
	}
}

func (s *Server) HandlePatientDiagnoseGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "Diagnose for that id not found")
			return
		}

		pr := protocol.Request{
			Id: int32(id),
		}

		response, err := s.PatientService.GetPatientDiagnose(context.Background(), &pr)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) HandlePatientDiagnoseUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "Diagnose for that id not found")
			return
		}

		var patientDiagnose protocol.PatientDiagnose
		json.NewDecoder(r.Body).Decode(&patientDiagnose)

		patientDiagnose.PatientDiagnoseId = int32(ID)

		response, err := s.PatientService.UpdatePatientDiagnose(context.Background(), &patientDiagnose)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
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
		json.NewDecoder(r.Body).Decode(&diagnoseSymptom)

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

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
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
		json.NewEncoder(w).Encode(response.DiagnoseSymptoms)
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
		json.NewDecoder(r.Body).Decode(&newDiagnoseSymptom)

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

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
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
