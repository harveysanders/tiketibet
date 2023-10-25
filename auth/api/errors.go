package api

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/harveysanders/tiketibet/auth/resp"
)

type ValidationErrors validator.ValidationErrors

func (v ValidationErrors) StatusCode() int {
	return http.StatusBadRequest
}

func (v ValidationErrors) SerializeError() []resp.Error {
	var errors []resp.Error

	for _, err := range v {
		errors = append(errors, resp.Error{
			Message: fmt.Sprintf("invalid value '%s'", err.Value()),
			Field:   err.Tag(),
		})
	}

	return errors
}

func (v ValidationErrors) Error() string {
	return "validation errors"
}

type BadRequestError struct {
	Message string
}

func (b BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}

func (b BadRequestError) SerializeError() []resp.Error {
	return []resp.Error{{Message: b.Message}}
}

func (b BadRequestError) Error() string {
	return b.Message
}
