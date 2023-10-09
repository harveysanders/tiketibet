package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) signInRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", a.handleSignIn())

	return r
}

func (a *App) handleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"tudo bem"}`))
	}
}
