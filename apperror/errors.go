package apperror

import (
	"net/http"

	"github.com/go-chi/render"
)

type Error struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

func (e Error) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Status)
	return nil
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrBadRequest          = &Error{Status: http.StatusBadRequest, Message: "invalid request"}
	ErrNotFound            = &Error{Status: http.StatusNotFound, Message: "resource not found"}
	ErrConflict            = &Error{Status: http.StatusConflict, Message: "the resource already exists"}
	ErrInternalServerError = &Error{Status: http.StatusInternalServerError, Message: "internal server error"}
)
