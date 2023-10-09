package auth

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

type App struct {
	logger   *slog.Logger
	validate *validator.Validate
	name     string
}

func (a *App) String() string {
	return a.name
}

func NewApp() *App {
	return &App{
		name:     "auth",
		logger:   slog.New(slog.NewJSONHandler(os.Stderr, nil)),
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (a *App) Run() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Mount("/api/users", a.Routes())

	return r
}
