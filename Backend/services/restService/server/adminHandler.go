package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	protocol "github.com/rareinator/Svendeprove/Backend/packages/protocol"
)

func (s *Server) handleGetDoctorsInHospital() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		requestedHospitalId := vars["hospitalID"]

		_, client, err := okta.NewClient(context.Background(), okta.WithOrgUrl(os.Getenv("OKTA_URL")), okta.WithToken(os.Getenv("OKTA_SDK_TOKEN")))
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		users, _, err := client.Group.ListGroupUsers(context.Background(), "00gbttsw3ArE8GSCV5d6", &query.Params{})
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		result := make([]*doctor, 0)

		for _, user := range users {
			hospitalID := fmt.Sprintf("%v", (*user.Profile)["hospital_id"])

			if hospitalID == requestedHospitalId {
				doctor := doctor{
					Name:     fmt.Sprintf("%v %v", (*user.Profile)["firstName"], (*user.Profile)["lastName"]),
					Username: fmt.Sprintf("%v", (*user.Profile)["login"]),
					Type:     fmt.Sprintf("%v", (*user.Profile)["userType"]),
					UserId:   user.Id,
				}

				result = append(result, &doctor)
			}

		}

		if err := json.NewEncoder(w).Encode(&result); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handlePatientsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, client, err := okta.NewClient(context.Background(), okta.WithOrgUrl(os.Getenv("OKTA_URL")), okta.WithToken(os.Getenv("OKTA_SDK_TOKEN")))
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		users, _, err := client.Group.ListGroupUsers(context.Background(), "00gbqw93aeIKYWqww5d6", &query.Params{})
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		result := make([]*protocol.User, 0)

		for _, user := range users {
			age, _ := strconv.Atoi(fmt.Sprintf("%v", (*user.Profile)["age"]))

			patient := protocol.User{
				Name:       fmt.Sprintf("%v", (*user.Profile)["displayName"]),
				Address:    fmt.Sprintf("%v", (*user.Profile)["streetAddress"]),
				City:       fmt.Sprintf("%v", (*user.Profile)["city"]),
				PostalCode: fmt.Sprintf("%v", (*user.Profile)["zipCode"]),
				Country:    fmt.Sprintf("%v", (*user.Profile)["full_country"]),
				SocialIdNr: fmt.Sprintf("%v", (*user.Profile)["social_id"]),
				Username:   fmt.Sprintf("%v", (*user.Profile)["login"]),
				Age:        int32(age),
				Gender:     fmt.Sprintf("%v", (*user.Profile)["gender"]),
				UserId:     user.Id,
			}

			result = append(result, &patient)

		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleUseradminHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := protocol.Empty{}

		response, err := s.UseradminService.GetHealth(context.Background(), &u)
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

func (s *Server) handleHospitalsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := s.UseradminService.GetHospitals(context.Background(), &protocol.Empty{})
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(&response.Hospitals); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleAvailableBeds() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request protocol.BedsRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respsone, err := s.UseradminService.GetAvailableBeds(context.Background(), &request)
		if err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(&respsone.Beds); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
