package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	db "github.com/harveysanders/tiketibet/auth/mongo"
	"github.com/harveysanders/tiketibet/auth/resp"
	"golang.org/x/crypto/bcrypt"
)

type signUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=30"`
}

type signUpResponse struct {
	Email string `json:"email"`
}

func (s *Server) signUpRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", s.handleSignUp())

	return r
}

func (s *Server) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body signUpRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			log.Printf("decode: %v", err)

			render.Render(w, r, resp.ErrRender(BadRequestError{"Invalid request body"}))
			return
		}

		if err := s.validate.Struct(&body); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			log.Printf("validate Struct: %v", validationErrors)

			render.Render(w, r, resp.ErrRender(ValidationErrors(validationErrors)))
			return
		}

		existingUser, err := s.store.GetUserByEmail(r.Context(), body.Email)
		if err != nil {
			log.Printf("getUserByEmail: %v", err)
			render.Render(w, r, resp.ErrServerError())
			return
		}
		if existingUser != nil {
			render.Render(w, r, resp.ErrRender(BadRequestError{
				fmt.Sprintf("User with email %s already exists", body.Email),
			}))
			return
		}
		fmt.Printf("creating user... %+v\n", body)

		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("bcrypt.GenerateFromPassword: %v", err)
			render.Render(w, r, resp.ErrServerError())
			return
		}

		newUser := db.User{Email: body.Email, Password: string(hash)}

		created, err := s.store.CreateUser(r.Context(), &newUser)
		if err != nil {
			log.Printf("createUser: %v", err)
			render.Render(w, r, resp.ErrServerError())
			return
		}

		render.JSON(w, r, signUpResponse{Email: created.Email})
	}
}
