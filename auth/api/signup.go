package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/harveysanders/tiketibet/auth/resp"
)

type signUpReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=30"`
}

func (s *Server) signUpRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", s.handleSignUp())

	return r
}

func (s *Server) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body signUpReq
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			log.Printf("handleSignUp: %v", err)

			resp := resp.ErrResponse{StatusCode: http.StatusBadRequest, Errors: []resp.Error{{Message: "Invalid request body"}}}
			resp.Render(w, r)
			return
		}

		if err := s.validate.Struct(&body); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			log.Printf("handleSignUp: %v", validationErrors)

			render.Render(w, r, resp.ErrRender(ValidationErrors(validationErrors)))
			return
		}

		fmt.Printf("creating user... %+v\n", body)

		render.DefaultResponder(w, r, render.M{})
	}
}
