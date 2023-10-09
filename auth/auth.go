package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	name string
}

func (a *App) String() string {
	return a.name
}

func NewApp() *App {
	return &App{
		name: "auth",
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
