package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) signOutRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", s.handleSignOut())

	return r
}

func (s *Server) handleSignOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"tudo bem"}`))
	}
}
