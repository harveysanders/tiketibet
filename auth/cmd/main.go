package main

import (
	"log"
	"net/http"
	"os"

	"github.com/harveysanders/tiketibet/auth"
)

func main() {
	app := auth.NewApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Printf("Starting %s server on port %s\n", app.String(), port)
	if err := http.ListenAndServe(":8888", app.Run()); err != nil {
		log.Fatal(err)
	}
}
