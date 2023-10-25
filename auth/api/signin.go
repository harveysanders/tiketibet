package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/harveysanders/tiketibet/auth/resp"
	"golang.org/x/crypto/bcrypt"
)

type signInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type signInResponse struct {
	Message string `json:"message"`
}

func (s *Server) signInRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", s.handleSignIn())

	return r
}

func (s *Server) handleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body signInRequest
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

		user, err := s.store.GetUserByEmail(r.Context(), body.Email)
		if err != nil {
			log.Printf("getUserByEmail: %v", err)
			render.Render(w, r, resp.ErrServerError())
			return
		}
		if user == nil {
			render.Render(w, r, resp.ErrRender(BadRequestError{
				"Invalid email or password",
			}))
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
			render.Render(w, r, resp.ErrRender(BadRequestError{
				"Invalid email or password",
			}))
			return
		}

		render.JSON(w, r, signInResponse{
			Message: "Sign in successful",
		})
	}
}
