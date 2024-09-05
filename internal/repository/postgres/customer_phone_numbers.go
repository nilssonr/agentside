package postgres

import (
	"context"
	"github.com/nilssonr/agentside/internal/customer"
)

type CustomerPhoneNumberRepository struct {
	db *Queries
}

func NewCustomerPhoneNumberRepository(db *Queries) customer.PhoneNumberRepository {
	return &CustomerPhoneNumberRepository{
		db: db,
	}
}

// InsertPhoneNumber implements customer.PhoneNumberRepository.
func (c *CustomerPhoneNumberRepository) InsertPhoneNumber(ctx context.Context, request *customer.PhoneNumber) (*customer.PhoneNumber, error) {
	panic("unimplemented")
}

// GetPhoneNumbers implements customer.PhoneNumberRepository.
func (c *CustomerPhoneNumberRepository) GetPhoneNumbers(ctx context.Context, customerID string) ([]*customer.PhoneNumber, error) {
	panic("unimplemented")
}

// GetPhoneNumber implements customer.PhoneNumberRepository.
func (c *CustomerPhoneNumberRepository) GetPhoneNumber(ctx context.Context, customerID string, phoneNumberID string) (*customer.PhoneNumber, error) {
	panic("unimplemented")
}

// UpdatePhoneNumber implements customer.PhoneNumberRepository.
func (c *CustomerPhoneNumberRepository) UpdatePhoneNumber(ctx context.Context, request *customer.PhoneNumber) (*customer.PhoneNumber, error) {
	panic("unimplemented")
}

// DeletePhoneNumber implements customer.PhoneNumberRepository.
func (c *CustomerPhoneNumberRepository) DeletePhoneNumber(ctx context.Context, customerID string, phoneNumberID string) error {
	panic("unimplemented")
}
