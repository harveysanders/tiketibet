package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) currentUserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", s.handleCurrentUser())

	return r
}

func (s *Server) handleCurrentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"tudo bem"}`))
	}
}
