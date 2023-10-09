package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) signOutRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", a.handleSignOut())

	return r
}

func (a *App) handleSignOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"tudo bem"}`))
	}
}
