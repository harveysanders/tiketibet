package auth

import (
	"github.com/go-chi/chi/v5"
)

func (a *App) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/currentuser", a.currentUserRouter())
	r.Mount("/signin", a.signInRouter())
	r.Mount("/signout", a.signOutRouter())
	r.Mount("/signup", a.signUpRouter())

	return r
}
