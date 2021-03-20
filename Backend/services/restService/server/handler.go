package server

import (
	"net/http"
)

func (s *Server) HandleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("🚀 Server is up and running!!!!")); err != nil {
			s.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
