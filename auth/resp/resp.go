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
