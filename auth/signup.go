package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type signUpReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *App) signUpRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", a.handleSignUp())

	return r
}

func (a *App) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"tudo bem"}`))
	}
}
