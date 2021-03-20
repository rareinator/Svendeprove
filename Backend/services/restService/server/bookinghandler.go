package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	protocol "github.com/rareinator/Svendeprove/Backend/packages/protocol"
)

func (s *Server) HandleBookingHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := s.BookingService.GetHealth(context.Background(), &protocol.Empty{})
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

func (s *Server) HandleBookingCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var booking protocol.Booking
		if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		response, err := s.BookingService.CreateBooking(context.Background(), &booking)
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

func (s *Server) HandleBookingDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookingID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No booking with that id found")
			return
		}

		br := protocol.Request{
			Id: int32(bookingID),
		}

		response, err := s.BookingService.DeleteBooking(context.Background(), &br)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if response.Success {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.ReturnError(w, http.StatusInternalServerError, "somethin unknown went gorribly wrong!!! ☠️☠️☠️")
	}
}

func (s *Server) HandleBookingsByPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		br := protocol.Request{
			UserId: vars["userId"],
		}

		response, err := s.BookingService.GetBookingsByPatient(context.Background(), &br)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(response.Bookings) == 0 {
			response.Bookings = make([]*protocol.Booking, 0)
		}
		if err := json.NewEncoder(w).Encode(response.Bookings); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) HandleBookingsByEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		br := protocol.Request{
			UserId: vars["userId"],
		}

		response, err := s.BookingService.GetBookingsByEmployee(context.Background(), &br)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(response.Bookings) == 0 {
			response.Bookings = make([]*protocol.Booking, 0)
		}
		if err := json.NewEncoder(w).Encode(response.Bookings); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) HandleAvailableTimesForDoctor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request protocol.TimeFrameRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		response, err := s.BookingService.GetAvailableTimesForDoctor(context.Background(), &request)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(response.Strings) == 0 {
			response.Strings = make([]string, 0)
		}
		if err := json.NewEncoder(w).Encode(&response.Strings); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
