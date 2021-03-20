package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	protocol "github.com/rareinator/Svendeprove/Backend/packages/protocol"
)

func (s *Server) handleIOTHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		i := protocol.Empty{}

		response, err := s.IotService.GetHealth(context.Background(), &i)
		if err != nil {
			s.ReturnError(w, http.StatusServiceUnavailable, err.Error())
			return
		}

		if _, err := w.Write([]byte(response.Message)); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleIOTUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Query().Get("Data")
		sensorId, _ := strconv.Atoi(r.URL.Query().Get("SensorId"))
		name := r.URL.Query().Get("Name")

		iotData := protocol.IOTData{
			Name:     name,
			SensorID: int32(sensorId),
			Data:     data,
		}

		response, err := s.IotService.UploadData(context.Background(), &iotData)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleIOTReadData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		deviceID, err := strconv.Atoi(vars["deviceID"])
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		request := protocol.Request{
			Id: int32(deviceID),
		}

		response, err := s.IotService.ReadData(context.Background(), &request)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(&response.IOTDatas); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleIOTReadDataInTimeframe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request protocol.IOTTimeframeRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			s.ReturnError(w, http.StatusBadRequest, err.Error())
			return
		}

		response, err := s.IotService.ReadDataInTimeFrame(context.Background(), &request)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(&response.IOTDatas); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
