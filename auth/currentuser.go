package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) currentUserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", a.handleCurrentUser())

	return r
}

func (a *App) handleCurrentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"tudo bem"}`))
	}
}
