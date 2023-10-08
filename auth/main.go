package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type App struct {
	name string
}

func newApp() *App {
	return &App{
		name: "auth",
	}
}

func (a *App) run() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/api/create/currentUser", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"tudo bem"}`))
	})

	return r
}

func main() {
	app := newApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Printf("Starting %s server on port %s\n", app.name, port)
	if err := http.ListenAndServe(":8888", app.run()); err != nil {
		log.Fatal(err)
	}
}
