package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerTenantRouter(r chi.Router) {
	r.Route("/tenants", func(r chi.Router) {
		r.Post("/", createTenant)
		r.Get("/", getTenants)
		r.Route("/{tenantID}", func(r chi.Router) {
			r.Get("/", getTenant)
			r.Put("/", updateTenant)
			r.Delete("/", deleteTenant)
		})
	})
}

type createTenantRequest struct {
	Name string `json:"name"`
}

type updateTenantRequest struct {
	Name string `json:"name"`
}

func createTenant(w http.ResponseWriter, r *http.Request) {

}

func getTenants(w http.ResponseWriter, r *http.Request) {

}

func getTenant(w http.ResponseWriter, r *http.Request) {

}

func updateTenant(w http.ResponseWriter, r *http.Request) {

}

func deleteTenant(w http.ResponseWriter, r *http.Request) {

}
