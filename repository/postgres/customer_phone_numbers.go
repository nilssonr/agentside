package postgres

import (
	"context"

	"github.com/nilssonr/agentside/customer"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

type CustomerPhoneNumberRepository struct {
	DB *sqlc.Queries
}

func NewCustomerPhoneNumberRepository(db *sqlc.Queries) customer.PhoneNumberRepository {
	return &CustomerPhoneNumberRepository{
		DB: db,
	}
}

// InsertPhoneNumber implements customer.PhoneNumberRepository.
func (c *CustomerPhoneNumberRepository) InsertPhoneNumber(ctx context.Context, request *customer.PhoneNumber) (*customer.PhoneNumber, error) {
	arg := sqlc.InsertCustomerPhoneNumberParams{
		PhoneNumber:    request.PhoneNumber,
		Type:           request.Type,
		CustomerID:     request.CustomerID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := c.DB.InsertCustomerPhoneNumber(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result customer.PhoneNumber
	result.ID = row.ID
	result.PhoneNumber = row.PhoneNumber
	result.Type = row.Type
	result.CustomerID = row.CustomerID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// GetPhoneNumbers implements customer.PhoneNumberRepository.
func (c *CustomerPhoneNumberRepository) GetPhoneNumbers(ctx context.Context, customerID string) ([]*customer.PhoneNumber, error) {
	rows, err := c.DB.GetCustomerPhoneNumbers(ctx, customerID)
	if err != nil {
		return nil, err
	}

	result := make([]*customer.PhoneNumber, 0, len(rows))
	for _, v := range rows {
		result = append(result, &customer.PhoneNumber{
			ID:             v.ID,
			PhoneNumber:    v.PhoneNumber,
			Type:           v.Type,
			CustomerID:     v.CustomerID,
			LastModifiedAt: v.LastModifiedAt.Time,
			LastModifiedBy: v.LastModifiedBy,
		})
	}

	return result, nil
}

// GetPhoneNumber implements customer.PhoneNumberRepository.
func (c *CustomerPhoneNumberRepository) GetPhoneNumber(ctx context.Context, customerID string, phoneNumberID string) (*customer.PhoneNumber, error) {
	arg := sqlc.GetCustomerPhoneNumberParams{
		CustomerID: customerID,
		ID:         phoneNumberID,
	}

	row, err := c.DB.GetCustomerPhoneNumber(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result customer.PhoneNumber
	result.ID = row.ID
	result.PhoneNumber = row.PhoneNumber
	result.Type = row.Type
	result.CustomerID = row.CustomerID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// UpdatePhoneNumber implements customer.PhoneNumberRepository.
func (c *CustomerPhoneNumberRepository) UpdatePhoneNumber(ctx context.Context, request *customer.PhoneNumber) (*customer.PhoneNumber, error) {
	arg := sqlc.UpdateCustomerPhoneNumberParams{
		ID:             request.ID,
		PhoneNumber:    request.PhoneNumber,
		Type:           request.Type,
		CustomerID:     request.CustomerID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := c.DB.UpdateCustomerPhoneNumber(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result customer.PhoneNumber
	result.ID = row.ID
	result.PhoneNumber = row.PhoneNumber
	result.Type = row.Type
	result.CustomerID = row.CustomerID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// DeletePhoneNumber implements customer.PhoneNumberRepository.
func (c *CustomerPhoneNumberRepository) DeletePhoneNumber(ctx context.Context, customerID string, phoneNumberID string) error {
	arg := sqlc.DeleteCustomerPhoneNumberParams{
		CustomerID: customerID,
		ID:         phoneNumberID,
	}

	if err := c.DB.DeleteCustomerPhoneNumber(ctx, arg); err != nil {
		return err
	}

	return nil
}
