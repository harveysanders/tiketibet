package auth

import (
	"github.com/harveysanders/tiketibet/auth/api"
)

type store interface{}

type App struct {
	name   string
	Server *api.Server
	store  store
}

func (a *App) String() string {
	return a.name
}

func NewApp(store store) *App {
	return &App{
		name:   "auth",
		Server: api.NewServer(),
		store:  store,
	}
}
