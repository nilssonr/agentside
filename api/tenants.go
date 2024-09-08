package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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

}

func (h tenantHandler) getTenants(w http.ResponseWriter, r *http.Request) {

}

func (h tenantHandler) getTenant(w http.ResponseWriter, r *http.Request) {

}

func (h tenantHandler) updateTenant(w http.ResponseWriter, r *http.Request) {

}

func (h tenantHandler) deleteTenant(w http.ResponseWriter, r *http.Request) {

}
