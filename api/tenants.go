package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/nilssonr/agentside/tenant"
)

type tenantHandler struct {
	TenantService tenant.Service
}

func (h tenantHandler) Register(r chi.Router) {
	r.Route("/tenants", func(r chi.Router) {
		r.Post("/", h.createTenant)
		r.Get("/", h.getTenants)

		r.Route("/{tenantID}", func(r chi.Router) {
			r.Get("/", h.getTenant)
			r.Put("/", h.updateTenant)
			r.Delete("/", h.deleteTenant)
		})
	})
}

type createTenantRequest struct {
	Name string `json:"name"`
}

type updateTenantRequest struct {
	Name string `json:"name"`
}

func (h tenantHandler) createTenant(w http.ResponseWriter, r *http.Request) {
	var body createTenantRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &tenant.Tenant{
		Name:           body.Name,
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	result, err := h.TenantService.CreateTenant(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h tenantHandler) getTenants(w http.ResponseWriter, r *http.Request) {
	result, err := h.TenantService.GetTenants(r.Context())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h tenantHandler) getTenant(w http.ResponseWriter, r *http.Request) {
	result, err := h.TenantService.GetTenant(r.Context(), chi.URLParam(r, "tenantID"))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h tenantHandler) updateTenant(w http.ResponseWriter, r *http.Request) {
	var body updateTenantRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &tenant.Tenant{
		ID:             chi.URLParam(r, "tenantID"),
		Name:           body.Name,
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	result, err := h.TenantService.UpdateTenant(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h tenantHandler) deleteTenant(w http.ResponseWriter, r *http.Request) {
	if err := h.TenantService.DeleteTenant(r.Context(), chi.URLParam(r, "tenantID")); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}
