package postgres

import (
	"context"

	"github.com/nilssonr/agentside/customer"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

type CustomerAddressRepository struct {
	DB *sqlc.Queries
}

func NewCustomerAddressRepository(db *sqlc.Queries) customer.AddressRepository {
	return &CustomerAddressRepository{
		DB: db,
	}
}

// InsertAddress implements customer.AddressRepository.
func (c *CustomerAddressRepository) InsertAddress(ctx context.Context, request *customer.Address) (*customer.Address, error) {
	arg := sqlc.InsertCustomerAddressParams{
		StreetAddress:  request.StreetAddress,
		CustomerID:     request.CustomerID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	if len(request.State) > 0 {
		arg.State = mustCreateString(request.State)
	}

	if len(request.ZipCode) > 0 {
		arg.ZipCode = mustCreateString(request.ZipCode)
	}

	if len(request.Country) > 0 {
		arg.Country = mustCreateString(request.Country)
	}

	row, err := c.DB.InsertCustomerAddress(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result customer.Address
	result.ID = row.ID
	result.Country = row.Country.String
	result.CustomerID = row.CustomerID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy
	result.State = row.State.String
	result.StreetAddress = row.StreetAddress
	result.ZipCode = row.ZipCode.String

	return &result, nil
}

// GetAddresses implements customer.AddressRepository.
func (c *CustomerAddressRepository) GetAddresses(ctx context.Context, customerID string) ([]*customer.Address, error) {
	rows, err := c.DB.GetCustomerAddresses(ctx, customerID)
	if err != nil {
		return nil, err
	}

	result := []*customer.Address{}
	for _, v := range rows {
		result = append(result, &customer.Address{
			ID:             customerID,
			StreetAddress:  v.StreetAddress,
			State:          v.State.String,
			ZipCode:        v.ZipCode.String,
			Country:        v.Country.String,
			CustomerID:     customerID,
			LastModifiedAt: v.LastModifiedAt.Time,
			LastModifiedBy: v.LastModifiedBy,
		})
	}

	return result, nil
}

// GetAddress implements customer.AddressRepository.
func (c *CustomerAddressRepository) GetAddress(ctx context.Context, customerID string, addressID string) (*customer.Address, error) {
	arg := sqlc.GetCustomerAddressParams{
		CustomerID: customerID,
		ID:         addressID,
	}

	row, err := c.DB.GetCustomerAddress(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result customer.Address
	result.ID = row.ID
	result.Country = row.Country.String
	result.CustomerID = row.CustomerID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy
	result.State = row.State.String
	result.StreetAddress = row.StreetAddress
	result.ZipCode = row.ZipCode.String

	return &result, nil
}

// UpdateAddress implements customer.AddressRepository.
func (c *CustomerAddressRepository) UpdateAddress(ctx context.Context, request *customer.Address) (*customer.Address, error) {
	arg := sqlc.UpdateCustomerAddressParams{
		ID:             request.ID,
		StreetAddress:  request.StreetAddress,
		CustomerID:     request.CustomerID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	if len(request.State) > 0 {
		arg.State = mustCreateString(request.State)
	}

	if len(request.ZipCode) > 0 {
		arg.ZipCode = mustCreateString(request.ZipCode)
	}

	if len(request.Country) > 0 {
		arg.Country = mustCreateString(request.Country)
	}

	row, err := c.DB.UpdateCustomerAddress(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result customer.Address
	result.ID = row.ID
	result.Country = row.Country.String
	result.CustomerID = row.CustomerID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy
	result.State = row.State.String
	result.StreetAddress = row.StreetAddress
	result.ZipCode = row.ZipCode.String

	return &result, nil
}

// DeleteAddress implements customer.AddressRepository.
func (c *CustomerAddressRepository) DeleteAddress(ctx context.Context, customerID string, addressID string) error {
	arg := sqlc.DeleteCustomerAddressParams{
		CustomerID: customerID,
		ID:         addressID,
	}

	if err := c.DB.DeleteCustomerAddress(ctx, arg); err != nil {
		return err
	}

	return nil
}
