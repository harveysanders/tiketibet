package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) signInRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", s.handleSignIn())

	return r
}

func (s *Server) handleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"tudo bem"}`))
	}
}
