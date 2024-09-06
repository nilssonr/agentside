package v1

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// CreateTenant implements ServerInterface.
func (ah AgentsideHandler) CreateTenant(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetTenants implements ServerInterface.
func (ah AgentsideHandler) GetTenants(w http.ResponseWriter, r *http.Request) {
	t, err := ah.TenantService.GetTenants(r.Context())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, t)
}

// GetTenant implements ServerInterface.
func (ah AgentsideHandler) GetTenant(w http.ResponseWriter, r *http.Request, tenantId uuid.UUID) {
	t, err := ah.TenantService.GetTenant(r.Context(), tenantId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, t)
}

// UpdateTenant implements ServerInterface.
func (ah AgentsideHandler) UpdateTenant(w http.ResponseWriter, r *http.Request, tenantId uuid.UUID) {
	panic("unimplemented")
}

// DeleteTenant implements ServerInterface.
func (ah AgentsideHandler) DeleteTenant(w http.ResponseWriter, r *http.Request, tenantId uuid.UUID) {
	panic("unimplemented")
}
