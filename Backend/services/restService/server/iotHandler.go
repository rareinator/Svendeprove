package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	iotService "github.com/rareinator/Svendeprove/Backend/services/iotService/iot"
)

func (s *Server) HandleIOTHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		i := iotService.IOTEmpty{}

		response, err := s.IotService.GetHealth(context.Background(), &i)
		if err != nil {
			s.ReturnError(w, http.StatusServiceUnavailable, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

func (s *Server) HandleIOTUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Query().Get("Data")
		deviceID, err := s.getDeviceID(r)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		username := s.getUsername(r)

		//TODO: fix

		iotData := iotService.IOTData{
			Name:     username,
			SensorID: deviceID,
			Data:     data,
		}

		response, err := s.IotService.UploadData(context.Background(), &iotData)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) HandleIOTReadData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		deviceID, err := strconv.Atoi(vars["deviceID"])
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		request := iotService.IOTRequest{
			ID: int32(deviceID),
		}

		response, err := s.IotService.ReadData(context.Background(), &request)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response.IOTDatas)

	}
}

func (s *Server) HandleIOTReadDataInTimeframe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request iotService.IOTTimeframeRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			s.ReturnError(w, http.StatusBadRequest, err.Error())
			return
		}

		response, err := s.IotService.ReadDataInTimeFrame(context.Background(), &request)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response.IOTDatas)
	}
}
