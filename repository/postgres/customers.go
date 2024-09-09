package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/nilssonr/agentside/customer"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

type CustomerRepository struct {
	DB *sqlc.Queries
}

func NewCustomerRepository(db *sqlc.Queries) customer.Repository {
	return &CustomerRepository{
		DB: db,
	}
}

// InsertCustomer implements customer.Repository.
func (c *CustomerRepository) InsertCustomer(ctx context.Context, request *customer.Customer) (*customer.Customer, error) {
	arg := sqlc.InsertCustomerParams{
		FirstName:      request.FirstName,
		LastName:       request.LastName,
		TenantID:       request.TenantID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := c.DB.InsertCustomer(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := &customer.Customer{
		ID:             row.ID,
		FirstName:      row.FirstName,
		LastName:       row.LastName,
		TenantID:       row.TenantID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// GetCustomers implements customer.Repository.
func (c *CustomerRepository) GetCustomers(ctx context.Context, tenantID string) ([]*customer.Customer, error) {
	rows, err := c.DB.GetCustomers(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := make([]*customer.Customer, 0, len(rows))
	for _, v := range rows {
		result = append(result, &customer.Customer{
			ID:             v.ID,
			FirstName:      v.FirstName,
			LastName:       v.LastName,
			TenantID:       v.TenantID,
			LastModifiedAt: v.LastModifiedAt.Time,
			LastModifiedBy: v.LastModifiedBy,
		})
	}

	return result, nil
}

// GetCustomer implements customer.Repository.
func (c *CustomerRepository) GetCustomer(ctx context.Context, tenantID string, customerID string) (*customer.Customer, error) {
	arg := sqlc.GetCustomerParams{
		TenantID: tenantID,
		ID:       customerID,
	}

	row, err := c.DB.GetCustomer(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &customer.Customer{
		ID:             row.ID,
		FirstName:      row.FirstName,
		LastName:       row.LastName,
		TenantID:       row.TenantID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// UpdateCustomer implements customer.Repository.
func (c *CustomerRepository) UpdateCustomer(ctx context.Context, request *customer.Customer) (*customer.Customer, error) {
	arg := sqlc.UpdateCustomerParams{
		ID:             request.ID,
		FirstName:      request.FirstName,
		LastName:       request.LastName,
		TenantID:       request.TenantID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := c.DB.UpdateCustomer(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &customer.Customer{
		ID:             row.ID,
		FirstName:      row.FirstName,
		LastName:       row.LastName,
		TenantID:       row.TenantID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// DeleteCustomer implements customer.Repository.
func (c *CustomerRepository) DeleteCustomer(ctx context.Context, tenantID string, customerID string) error {
	arg := sqlc.DeleteCustomerParams{
		TenantID:  tenantID,
		ID:        customerID,
		DeletedAt: mustCreateTime(time.Now()),
	}

	if err := c.DB.DeleteCustomer(ctx, arg); err != nil {
		return fmt.Errorf("repository: %w", err)
	}

	return nil
}
