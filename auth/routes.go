package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/currentUser", a.currentUser())
	r.Get("/signin", a.handleSignIn())

	return r
}

func (a *App) currentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"tudo bem"}`))
	}
}

func (a *App) handleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"tudo bem"}`))
	}
}
