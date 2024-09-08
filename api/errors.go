package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/nilssonr/agentside/apperror"
)

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	switch e := err.(type) {
	case *apperror.Error:
		render.Render(w, r, e)
	default:
		render.Render(w, r, apperror.ErrInternalServerError)
	}
}
