package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/nilssonr/agentside/repository/postgres/sqlc"
	"github.com/nilssonr/agentside/tenant"
)

type TenantRepository struct {
	db *sqlc.Queries
}

func NewTenantRepository(db *sqlc.Queries) tenant.Repository {
	return &TenantRepository{
		db: db,
	}
}

// CreateTenant implements tenant.Repository.
func (tr *TenantRepository) CreateTenant(ctx context.Context, t *tenant.Tenant) (*tenant.Tenant, error) {
	arg := sqlc.InsertTenantParams{
		Name:           t.Name,
		LastModifiedAt: mustCreateTime(time.Now()),
		LastModifiedBy: t.LastModifiedBy,
	}

	row, err := tr.db.InsertTenant(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &tenant.Tenant{
		ID:             row.ID,
		Name:           row.Name,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// GetTenants implements tenant.Repository.
func (tr *TenantRepository) GetTenants(ctx context.Context) ([]*tenant.Tenant, error) {
	rows, err := tr.db.GetTenants(ctx)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := make([]*tenant.Tenant, 0, len(rows))
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
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &tenant.Tenant{
		ID:             row.ID,
		Name:           row.Name,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// UpdateTenant implements tenant.Repository.
func (tr *TenantRepository) UpdateTenant(ctx context.Context, t *tenant.Tenant) (*tenant.Tenant, error) {
	arg := sqlc.UpdateTenantParams{
		ID:             t.ID,
		Name:           t.Name,
		LastModifiedAt: mustCreateTime(time.Now()),
		LastModifiedBy: t.LastModifiedBy,
	}

	row, err := tr.db.UpdateTenant(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &tenant.Tenant{
		ID:             row.ID,
		Name:           row.Name,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// DeleteTenant implements tenant.Repository.
func (tr *TenantRepository) DeleteTenant(ctx context.Context, tenantID string) error {
	arg := sqlc.DeleteTenantParams{
		ID:        tenantID,
		DeletedAt: mustCreateTime(time.Now()),
	}

	if err := tr.db.DeleteTenant(ctx, arg); err != nil {
		return fmt.Errorf("repository: %w", err)
	}

	return nil
}
