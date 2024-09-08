package customer

import (
	"context"
	"time"
)

type Customer struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	TenantID       string    `json:"tenantId"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}

type Service interface {
	CreateCustomer(ctx context.Context, request *Customer) (*Customer, error)
	GetCustomers(ctx context.Context, tenantID string) ([]*Customer, error)
	GetCustomer(ctx context.Context, tenantID, customerID string) (*Customer, error)
	UpdateCustomer(ctx context.Context, request *Customer) (*Customer, error)
	DeleteCustomer(ctx context.Context, tenantID, customerID string) error
}

type service struct {
	customerRepository Repository
}

func NewService(r Repository) Service {
	return &service{
		customerRepository: r,
	}
}

// CreateCustomer implements Service.
func (s *service) CreateCustomer(ctx context.Context, request *Customer) (*Customer, error) {
	result, err := s.customerRepository.InsertCustomer(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCustomers implements Service.
func (s *service) GetCustomers(ctx context.Context, tenantID string) ([]*Customer, error) {
	result, err := s.customerRepository.GetCustomers(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCustomer implements Service.
func (s *service) GetCustomer(ctx context.Context, tenantID string, customerID string) (*Customer, error) {
	result, err := s.customerRepository.GetCustomer(ctx, tenantID, customerID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateCustomer implements Service.
func (s *service) UpdateCustomer(ctx context.Context, request *Customer) (*Customer, error) {
	result, err := s.customerRepository.UpdateCustomer(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteCustomer implements Service.
func (s *service) DeleteCustomer(ctx context.Context, tenantID string, customerID string) error {
	if err := s.customerRepository.DeleteCustomer(ctx, tenantID, customerID); err != nil {
		return err
	}

	return nil
}
