package v1

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nilssonr/agentside/tenant"
)

type TenantHandler struct {
	TenantService tenant.Service
}

func NewTenantHandler(ts tenant.Service) TenantHandler {
	return TenantHandler{
		TenantService: ts,
	}
}

// CreateTenant implements ServerInterface.
func (h TenantHandler) CreateTenant(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetTenants implements ServerInterface.
func (h TenantHandler) GetTenants(w http.ResponseWriter, r *http.Request) {
	t, err := h.TenantService.GetTenants(r.Context())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, t)
}

// GetTenant implements ServerInterface.
func (h TenantHandler) GetTenant(w http.ResponseWriter, r *http.Request, tenantId uuid.UUID) {
	t, err := h.TenantService.GetTenant(r.Context(), tenantId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, t)
}

// UpdateTenant implements ServerInterface.
func (h TenantHandler) UpdateTenant(w http.ResponseWriter, r *http.Request, tenantId uuid.UUID) {
	panic("unimplemented")
}

// DeleteTenant implements ServerInterface.
func (h TenantHandler) DeleteTenant(w http.ResponseWriter, r *http.Request, tenantId uuid.UUID) {
	panic("unimplemented")
}
