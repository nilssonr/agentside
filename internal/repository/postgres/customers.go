package postgres

import (
	"context"
	"github.com/nilssonr/agentside/internal/customer"
)

type CustomerRepository struct {
	db *Queries
}

func NewCustomerRepository(db *Queries) customer.Repository {
	return &CustomerRepository{
		db: db,
	}
}

// InsertCustomer implements customer.Repository.
func (c *CustomerRepository) InsertCustomer(ctx context.Context, request *customer.Customer) (*customer.Customer, error) {
	panic("unimplemented")
}

// GetCustomers implements customer.Repository.
func (c *CustomerRepository) GetCustomers(ctx context.Context, tenantID string) ([]*customer.Customer, error) {
	panic("unimplemented")
}

// GetCustomer implements customer.Repository.
func (c *CustomerRepository) GetCustomer(ctx context.Context, tenantID string, customerID string) (*customer.Customer, error) {
	panic("unimplemented")
}

// UpdateCustomer implements customer.Repository.
func (c *CustomerRepository) UpdateCustomer(ctx context.Context, request *customer.Customer) (*customer.Customer, error) {
	panic("unimplemented")
}

// DeleteCustomer implements customer.Repository.
func (c *CustomerRepository) DeleteCustomer(ctx context.Context, tenantID string, customerID string) error {
	panic("unimplemented")
}
