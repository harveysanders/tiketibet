package auth

import (
	"github.com/harveysanders/tiketibet/auth/api"
	"github.com/harveysanders/tiketibet/auth/mongo"
)

type App struct {
	name   string
	Server *api.Server
	store  *mongo.Store
}

func (a *App) String() string {
	return a.name
}

func NewApp(store *mongo.Store) *App {
	return &App{
		name:   "auth",
		Server: api.NewServer(store),
		store:  store,
	}
}
