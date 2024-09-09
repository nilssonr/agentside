package postgres

import (
	"context"
	"fmt"

	"github.com/nilssonr/agentside/customer"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

type CustomerEmailAddressRepository struct {
	DB *sqlc.Queries
}

func NewCustomerEmailAddressRepository(db *sqlc.Queries) customer.EmailAddressRepository {
	return &CustomerEmailAddressRepository{
		DB: db,
	}
}

// InsertEmailAddress implements customer.EmailAddressRepository.
func (r *CustomerEmailAddressRepository) InsertEmailAddress(ctx context.Context, request *customer.EmailAddress) (*customer.EmailAddress, error) {
	arg := sqlc.InsertCustomerEmailAddressParams{
		EmailAddress:   request.EmailAddress,
		Type:           request.Type,
		CustomerID:     request.CustomerID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := r.DB.InsertCustomerEmailAddress(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &customer.EmailAddress{
		ID:             row.ID,
		EmailAddress:   row.EmailAddress,
		Type:           row.Type,
		CustomerID:     row.CustomerID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// GetEmailAddresses implements customer.EmailAddressRepository.
func (r *CustomerEmailAddressRepository) GetEmailAddresses(ctx context.Context, customerID string) ([]*customer.EmailAddress, error) {
	rows, err := r.DB.GetCustomerEmailAddresses(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := make([]*customer.EmailAddress, 0, len(rows))
	for _, v := range rows {
		result = append(result, &customer.EmailAddress{
			ID:             v.ID,
			EmailAddress:   v.EmailAddress,
			Type:           v.Type,
			CustomerID:     v.CustomerID,
			LastModifiedAt: v.LastModifiedAt.Time,
			LastModifiedBy: v.LastModifiedBy,
		})
	}

	return result, nil
}

// GetEmailAddress implements customer.EmailAddressRepository.
func (r *CustomerEmailAddressRepository) GetEmailAddress(ctx context.Context, customerID string, emailAddressID string) (*customer.EmailAddress, error) {
	arg := sqlc.GetCustomerEmailAddressParams{
		CustomerID: customerID,
		ID:         emailAddressID,
	}

	row, err := r.DB.GetCustomerEmailAddress(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &customer.EmailAddress{
		ID:             row.ID,
		EmailAddress:   row.EmailAddress,
		Type:           row.Type,
		CustomerID:     row.CustomerID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// UpdateEmailAddress implements customer.EmailAddressRepository.
func (r *CustomerEmailAddressRepository) UpdateEmailAddress(ctx context.Context, request *customer.EmailAddress) (*customer.EmailAddress, error) {
	arg := sqlc.UpdateCustomerEmailAddressParams{
		CustomerID:     request.CustomerID,
		ID:             request.ID,
		EmailAddress:   request.EmailAddress,
		Type:           request.Type,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := r.DB.UpdateCustomerEmailAddress(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &customer.EmailAddress{
		ID:             row.ID,
		EmailAddress:   row.EmailAddress,
		Type:           row.Type,
		CustomerID:     row.CustomerID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// DeleteEmailAddress implements customer.EmailAddressRepository.
func (r *CustomerEmailAddressRepository) DeleteEmailAddress(ctx context.Context, customerID, emailAddressID string) error {
	arg := sqlc.DeleteCustomerEmailAddressParams{
		CustomerID: customerID,
		ID:         emailAddressID,
	}

	if err := r.DB.DeleteCustomerEmailAddress(ctx, arg); err != nil {
		return fmt.Errorf("repository: %w", err)
	}

	return nil
}
