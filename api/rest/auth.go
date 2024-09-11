package rest

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/nilssonr/agentside/auth"
)

type authHandler struct {
	AuthClientService auth.ClientService
}

func (h authHandler) Register(r chi.Router) {
	r.Route("/clients", func(r chi.Router) {
		r.Post("/", h.createAuthClient)
		r.Get("/", h.getAuthClients)

		r.Route("/{clientID}", func(r chi.Router) {
			r.Get("/", h.getAuthClient)
			r.Put("/", h.updateAuthClient)
			r.Delete("/", h.deleteAuthClient)
		})
	})

	// r.Route("/auth", func(r chi.Router) {
	// 	})

	// 	r.Route("/token", func(r chi.Router) {

	// 	})
	// })
}

type createAuthClientRequest struct {
	Name string `json:"name"`
}

type updateAuthClientRequet struct {
	Name string `json:"name"`
}

func (h authHandler) createAuthClient(w http.ResponseWriter, r *http.Request) {
	var body createAuthClientRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &auth.Client{
		Name:           body.Name,
		TenantID:       tenantID(r),
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	result, err := h.AuthClientService.CreateClient(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h authHandler) getAuthClients(w http.ResponseWriter, r *http.Request) {
	result, err := h.AuthClientService.GetClients(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h authHandler) getAuthClient(w http.ResponseWriter, r *http.Request) {
	result, err := h.AuthClientService.GetClient(r.Context(), tenantID(r), chi.URLParam(r, "clientID"))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h authHandler) updateAuthClient(w http.ResponseWriter, r *http.Request) {
	var body updateAuthClientRequet
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &auth.Client{
		ID:             chi.URLParam(r, "clientID"),
		Name:           body.Name,
		TenantID:       tenantID(r),
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	result, err := h.AuthClientService.UpdateClient(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h authHandler) deleteAuthClient(w http.ResponseWriter, r *http.Request) {
	if err := h.AuthClientService.DeleteClient(r.Context(), tenantID(r), chi.URLParam(r, "clientID")); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}
