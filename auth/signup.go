package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/harveysanders/tiketibet/auth/resp"
)

type signUpReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}

func (a *App) signUpRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", a.handleSignUp())

	return r
}

func (a *App) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body signUpReq
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			a.logger.Error("handleSignUp: %v", err)

			render.Render(w, r, resp.ErrRender(err, http.StatusBadRequest, "invalid request body"))
			return
		}

		if err := a.validate.Struct(&body); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			a.logger.Error("handleSignUp: %v", validationErrors)

			// TODO: create a custom error response
			errMsg := validationErrors.Error()
			render.Render(w, r, resp.ErrRender(err, http.StatusBadRequest, errMsg))
			return
		}

		render.DefaultResponder(w, r, render.M{"message": "sign up"})
	}
}
