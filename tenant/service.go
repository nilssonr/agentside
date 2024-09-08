package tenant

import (
	"context"
)

type Service interface {
	CreateTenant(ctx context.Context, t *Tenant) (*Tenant, error)
	GetTenants(ctx context.Context) ([]*Tenant, error)
	GetTenant(ctx context.Context, tenantID string) (*Tenant, error)
	UpdateTenant(ctx context.Context, tenant *Tenant) (*Tenant, error)
	DeleteTenant(ctx context.Context, tenantID string) error
}

type service struct {
	tenantRepository Repository
}

func NewService(tr Repository) Service {
	return &service{
		tenantRepository: tr,
	}
}

// CreateTenant implements Service.
func (s *service) CreateTenant(ctx context.Context, t *Tenant) (*Tenant, error) {
	t, err := s.tenantRepository.CreateTenant(ctx, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetTenants implements Service.
func (s *service) GetTenants(ctx context.Context) ([]*Tenant, error) {
	t, err := s.tenantRepository.GetTenants(ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetTenant implements Service.
func (s *service) GetTenant(ctx context.Context, tenantID string) (*Tenant, error) {
	t, err := s.tenantRepository.GetTenant(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// UpdateTenant implements Service.
func (s *service) UpdateTenant(ctx context.Context, t *Tenant) (*Tenant, error) {
	t, err := s.tenantRepository.UpdateTenant(ctx, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// DeleteTenant implements Service.
func (s *service) DeleteTenant(ctx context.Context, tenantID string) error {
	if err := s.tenantRepository.DeleteTenant(ctx, tenantID); err != nil {
		return err
	}

	return nil
}
