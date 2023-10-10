package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	router   *chi.Mux
	validate *validator.Validate
}

func NewServer() *Server {
	return &Server{
		router:   chi.NewRouter(),
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (s *Server) Routes() *chi.Mux {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.URLFormat)

	s.router.Mount("/api/users", s.userRoutes())

	return s.router
}
