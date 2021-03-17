package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	bookingService "github.com/rareinator/Svendeprove/Backend/services/bookingService/booking"
)

func (s *Server) HandleBookingHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b := bookingService.BEmpty{}

		response, err := s.BookingService.GetHealth(context.Background(), &b)
		if err != nil {
			s.ReturnError(w, http.StatusServiceUnavailable, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Message))
	}
}

func (s *Server) HandleBookingCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var booking bookingService.Booking
		json.NewDecoder(r.Body).Decode(&booking)

		booking.Employee = s.getUsername(r)

		response, err := s.BookingService.CreateBooking(context.Background(), &booking)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) HandleBookingGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No booking found with that id")
			return
		}

		b := bookingService.BRequest{
			Id: int32(ID),
		}

		response, err := s.BookingService.GetBooking(context.Background(), &b)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) HandleBookingUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookingID, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.ReturnError(w, http.StatusNotFound, "No booking found with that id")
			return
		}

		var booking bookingService.Booking
		json.NewDecoder(r.Body).Decode(&booking)

		booking.BookingId = int32(bookingID)

		response, err := s.BookingService.UpdateBooking(context.Background(), &booking)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
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

		br := bookingService.BRequest{
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

		s.ReturnError(w, http.StatusInternalServerError, "Somethin unknown went gorribly wrong!!! ☠️☠️☠️")
	}
}

func (s *Server) HandleBookingsByPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		br := bookingService.BRequest{
			Username: vars["username"],
		}

		response, err := s.BookingService.GetBookingsByPatient(context.Background(), &br)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Bookings) == 0 {
			response.Bookings = make([]*bookingService.Booking, 0)
		}
		json.NewEncoder(w).Encode(response.Bookings)
	}
}

func (s *Server) HandleBookingsByEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		br := bookingService.BRequest{
			Username: vars["username"],
		}

		response, err := s.BookingService.GetBookingsByEmployee(context.Background(), &br)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		if len(response.Bookings) == 0 {
			response.Bookings = make([]*bookingService.Booking, 0)
		}
		json.NewEncoder(w).Encode(response.Bookings)
	}
}

func (s *Server) HandleAvailableTimesForDoctor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request bookingService.BTimeFrameRequest

		json.NewDecoder(r.Body).Decode(&request)

		response, err := s.BookingService.GetAvailableTimesForDoctor(context.Background(), &request)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response.Strings)
	}
}