package resp

import (
	"net/http"

	"github.com/go-chi/render"
)

type Error struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}
type HTTPError interface {
	StatusCode() int
	SerializeError() []Error
	Error() string
}

type ErrResponse struct {
	StatusCode int
	Errors     []Error `json:"errors"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ErrRender(e HTTPError) render.Renderer {
	return &ErrResponse{
		Errors:     e.SerializeError(),
		StatusCode: e.StatusCode(),
	}
}

func ErrEmailAlreadyExists(msg string) render.Renderer {
	return &ErrResponse{
		Errors:     []Error{{Message: msg}},
		StatusCode: http.StatusConflict,
	}
}

func ErrInvalidRequest() render.Renderer {
	return &ErrResponse{
		Errors:     []Error{{Message: "Invalid request body"}},
		StatusCode: http.StatusBadRequest,
	}
}

func ErrServerError() render.Renderer {
	return &ErrResponse{
		Errors:     []Error{{Message: "Server error"}},
		StatusCode: http.StatusInternalServerError,
	}
}
