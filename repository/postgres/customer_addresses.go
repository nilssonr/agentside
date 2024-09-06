package postgres

import (
	"context"

	"github.com/nilssonr/agentside/customer"
)

type CustomerAddressRepository struct {
	db *Queries
}

func NewCustomerAddressRepository(db *Queries) customer.AddressRepository {
	return &CustomerAddressRepository{
		db: db,
	}
}

// InsertAddress implements customer.AddressRepository.
func (c *CustomerAddressRepository) InsertAddress(ctx context.Context, request *customer.Address) (*customer.Address, error) {
	panic("unimplemented")
}

// GetAddresses implements customer.AddressRepository.
func (c *CustomerAddressRepository) GetAddresses(ctx context.Context, customerID string) ([]*customer.Address, error) {
	panic("unimplemented")
}

// GetAddress implements customer.AddressRepository.
func (c *CustomerAddressRepository) GetAddress(ctx context.Context, customerID string, addressID string) (*customer.Address, error) {
	panic("unimplemented")
}

// UpdateAddress implements customer.AddressRepository.
func (c *CustomerAddressRepository) UpdateAddress(ctx context.Context, request *customer.Address) (*customer.Address, error) {
	panic("unimplemented")
}

// DeleteAddress implements customer.AddressRepository.
func (c *CustomerAddressRepository) DeleteAddress(ctx context.Context, customerID string, addressID string) error {
	panic("unimplemented")
}
