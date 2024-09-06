package tenant

import "context"

type Repository interface {
	CreateTenant(ctx context.Context, t *Tenant) (*Tenant, error)
	GetTenants(ctx context.Context) ([]*Tenant, error)
	GetTenant(ctx context.Context, tenantID string) (*Tenant, error)
	UpdateTenant(ctx context.Context, t *Tenant) (*Tenant, error)
	DeleteTenant(ctx context.Context, tenantID string) error
}
