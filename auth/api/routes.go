package api

import (
	"github.com/go-chi/chi/v5"
)

func (s *Server) userRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/currentuser", s.currentUserRouter())
	r.Mount("/signin", s.signInRouter())
	r.Mount("/signout", s.signOutRouter())
	r.Mount("/signup", s.signUpRouter())

	return r
}
