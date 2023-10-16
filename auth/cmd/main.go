package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/harveysanders/tiketibet/auth"
	db "github.com/harveysanders/tiketibet/auth/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURL := os.Getenv("MONGO_URI")
	if mongoURL == "" {
		log.Println("MONGO_URI not set, using default")
		mongoURL = "mongodb://127.0.0.1:27017/auth"
	}

	dbOpts := options.Client().ApplyURI(mongoURL)
	dbOpts.SetConnectTimeout(time.Second * 5)
	client, err := mongo.Connect(context.TODO(), dbOpts)
	if err != nil {
		log.Fatalf("DB connect: %v\n", err)
	}

	dbName, err := db.ParseDBName(mongoURL)
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(client, dbName)

	app := auth.NewApp(store)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Printf("Starting %s server on port %s\n", app.String(), port)
	if err := http.ListenAndServe(":8888", app.Server.Routes()); err != nil {
		log.Fatal(err)
	}
}
