package auth

import (
	"github.com/harveysanders/tiketibet/auth/api"
)

type App struct {
	name   string
	Server *api.Server
}

func (a *App) String() string {
	return a.name
}

func NewApp() *App {
	return &App{
		name:   "auth",
		Server: api.NewServer(),
	}
}
