package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/nilssonr/agentside/auth"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

type AuthClientRepository struct {
	DB *sqlc.Queries
}

func NewAuthClientRepository(db *sqlc.Queries) auth.ClientRepository {
	return &AuthClientRepository{
		DB: db,
	}
}

// InsertClient implements auth.ClientRepository.
func (r *AuthClientRepository) InsertClient(ctx context.Context, request *auth.Client) (*auth.Client, error) {
	arg := sqlc.InsertAuthClientParams{
		Name:           request.Name,
		Secret:         request.Secret,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
		TenantID:       request.TenantID,
	}

	row, err := r.DB.InsertAuthClient(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &auth.Client{
		ID:             row.ID,
		Name:           row.Name,
		Secret:         row.Secret,
		TenantID:       row.TenantID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// GetClients implements auth.ClientRepository.
func (r *AuthClientRepository) GetClients(ctx context.Context, tenantID string) ([]*auth.Client, error) {
	rows, err := r.DB.GetAuthClients(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := make([]*auth.Client, 0, len(rows))
	for _, v := range rows {
		result = append(result, &auth.Client{
			ID:             v.ID,
			Name:           v.Name,
			Secret:         v.Secret,
			TenantID:       v.TenantID,
			LastModifiedAt: v.LastModifiedAt.Time,
			LastModifiedBy: v.LastModifiedBy,
		})
	}

	return result, nil
}

// GetClient implements auth.ClientRepository.
func (r *AuthClientRepository) GetClient(ctx context.Context, tenantID, clientID string) (*auth.Client, error) {
	arg := sqlc.GetAuthClientParams{
		TenantID: tenantID,
		ID:       clientID,
	}

	row, err := r.DB.GetAuthClient(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &auth.Client{
		ID:             row.ID,
		Name:           row.Name,
		Secret:         row.Secret,
		TenantID:       row.TenantID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// UpdateClient implements auth.ClientRepository.
func (r *AuthClientRepository) UpdateClient(ctx context.Context, request *auth.Client) (*auth.Client, error) {
	arg := sqlc.UpdateAuthClientParams{
		TenantID:       request.TenantID,
		ID:             request.ID,
		Name:           request.Name,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := r.DB.UpdateAuthClient(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &auth.Client{
		ID:             row.ID,
		Name:           row.Name,
		Secret:         row.Secret,
		TenantID:       row.TenantID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// DeleteClient implements auth.ClientRepository.
func (r *AuthClientRepository) DeleteClient(ctx context.Context, tenantID, clientID string) error {
	arg := sqlc.DeleteAuthClientParams{
		TenantID:  tenantID,
		ID:        clientID,
		DeletedAt: mustCreateTime(time.Now()),
	}

	if err := r.DB.DeleteAuthClient(ctx, arg); err != nil {
		return fmt.Errorf("repository: %w", err)
	}

	return nil
}
