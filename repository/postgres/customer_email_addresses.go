package postgres

import (
	"context"

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
func (c *CustomerEmailAddressRepository) InsertEmailAddress(ctx context.Context, request *customer.EmailAddress) (*customer.EmailAddress, error) {
	panic("unimplemented")
}

// GetEmailAddresses implements customer.EmailAddressRepository.
func (c *CustomerEmailAddressRepository) GetEmailAddresses(ctx context.Context, customerID string) ([]*customer.EmailAddress, error) {
	panic("unimplemented")
}

// GetEmailAddress implements customer.EmailAddressRepository.
func (c *CustomerEmailAddressRepository) GetEmailAddress(ctx context.Context, customerID string, emailAddressID string) (*customer.EmailAddress, error) {
	panic("unimplemented")
}

// UpdateEmailAddress implements customer.EmailAddressRepository.
func (c *CustomerEmailAddressRepository) UpdateEmailAddress(ctx context.Context, request *customer.EmailAddress) (*customer.EmailAddress, error) {
	panic("unimplemented")
}

// DeleteEmailAddress implements customer.EmailAddressRepository.
func (c *CustomerEmailAddressRepository) DeleteEmailAddress(ctx context.Context, customerID string, emailAddressID string) error {
	panic("unimplemented")
}
