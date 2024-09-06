package postgres

import (
	"context"
	"time"

	"github.com/nilssonr/agentside/customer"
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
	arg := InsertCustomerParams{
		FirstName:      request.FirstName,
		LastName:       request.LastName,
		TenantID:       request.TenantID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := c.db.InsertCustomer(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result customer.Customer
	result.ID = row.ID
	result.FirstName = row.FirstName
	result.LastName = row.LastName
	result.TenantID = row.TenantID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// GetCustomers implements customer.Repository.
func (c *CustomerRepository) GetCustomers(ctx context.Context, tenantID string) ([]*customer.Customer, error) {
	rows, err := c.db.GetCustomers(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	result := []*customer.Customer{}
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
	arg := GetCustomerParams{
		TenantID: tenantID,
		ID:       customerID,
	}

	row, err := c.db.GetCustomer(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result customer.Customer
	result.ID = row.ID
	result.FirstName = row.FirstName
	result.LastName = row.LastName
	result.TenantID = row.TenantID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// UpdateCustomer implements customer.Repository.
func (c *CustomerRepository) UpdateCustomer(ctx context.Context, request *customer.Customer) (*customer.Customer, error) {
	arg := UpdateCustomerParams{
		ID:             request.ID,
		FirstName:      request.FirstName,
		LastName:       request.LastName,
		TenantID:       request.TenantID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := c.db.UpdateCustomer(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result customer.Customer
	result.ID = row.ID
	result.FirstName = row.FirstName
	result.LastName = row.LastName
	result.TenantID = row.TenantID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// DeleteCustomer implements customer.Repository.
func (c *CustomerRepository) DeleteCustomer(ctx context.Context, tenantID string, customerID string) error {
	arg := DeleteCustomerParams{
		TenantID:  tenantID,
		ID:        customerID,
		DeletedAt: mustCreateTime(time.Now()),
	}

	if err := c.db.DeleteCustomer(ctx, arg); err != nil {
		return err
	}

	return nil
}
