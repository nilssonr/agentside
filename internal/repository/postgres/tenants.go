package postgres

import (
	"context"
	"time"

	"github.com/nilssonr/agentside/internal/tenant"
)

type TenantRepository struct {
	db *Queries
}

func NewTenantRepository(db *Queries) tenant.Repository {
	return &TenantRepository{
		db: db,
	}
}

// CreateTenant implements tenant.Repository.
func (tr *TenantRepository) CreateTenant(ctx context.Context, t *tenant.Tenant) (*tenant.Tenant, error) {
	arg := CreateTenantParams{
		Name:           t.Name,
		LastModifiedAt: mustCreateTime(time.Now()),
		LastModifiedBy: t.LastModifiedBy,
	}
	row, err := tr.db.CreateTenant(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result tenant.Tenant
	result.ID = row.ID
	result.Name = row.Name
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// GetTenants implements tenant.Repository.
func (tr *TenantRepository) GetTenants(ctx context.Context) ([]*tenant.Tenant, error) {
	rows, err := tr.db.GetTenants(ctx)
	if err != nil {
		return nil, err
	}

	result := []*tenant.Tenant{}
	for _, v := range rows {
		result = append(result, &tenant.Tenant{
			ID:             v.ID,
			Name:           v.Name,
			LastModifiedAt: v.LastModifiedAt.Time,
			LastModifiedBy: v.LastModifiedBy,
		})
	}

	return result, nil
}

// GetTenant implements tenant.Repository.
func (tr *TenantRepository) GetTenant(ctx context.Context, tenantID string) (*tenant.Tenant, error) {
	row, err := tr.db.GetTenant(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	var result tenant.Tenant
	result.ID = row.ID
	result.Name = row.Name
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// UpdateTenant implements tenant.Repository.
func (tr *TenantRepository) UpdateTenant(ctx context.Context, t *tenant.Tenant) (*tenant.Tenant, error) {
	arg := UpdateTenantParams{
		Name:           t.Name,
		LastModifiedAt: mustCreateTime(time.Now()),
		LastModifiedBy: t.LastModifiedBy,
	}
	row, err := tr.db.UpdateTenant(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result tenant.Tenant
	result.ID = row.ID
	result.Name = row.Name
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// DeleteTenant implements tenant.Repository.
func (tr *TenantRepository) DeleteTenant(ctx context.Context, tenantID string) error {
	arg := DeleteTenantParams{
		ID:        tenantID,
		DeletedAt: mustCreateTime(time.Now()),
	}
	if err := tr.db.DeleteTenant(ctx, arg); err != nil {
		return err
	}

	return nil
}
