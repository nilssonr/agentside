package postgres

import (
	"context"

	"github.com/nilssonr/agentside/customer"
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
	arg := InsertCustomerPhoneNumberParams{
		PhoneNumber:    request.PhoneNumber,
		Type:           request.Type,
		CustomerID:     request.CustomerID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := c.db.InsertCustomerPhoneNumber(ctx, arg)
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
	rows, err := c.db.GetCustomerPhoneNumbers(ctx, customerID)
	if err != nil {
		return nil, err
	}

	result := []*customer.PhoneNumber{}
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
	arg := GetCustomerPhoneNumberParams{
		CustomerID: customerID,
		ID:         phoneNumberID,
	}

	row, err := c.db.GetCustomerPhoneNumber(ctx, arg)
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
	arg := UpdateCustomerPhoneNumberParams{
		ID:             request.ID,
		PhoneNumber:    request.PhoneNumber,
		Type:           request.Type,
		CustomerID:     request.CustomerID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := c.db.UpdateCustomerPhoneNumber(ctx, arg)
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
	arg := DeleteCustomerPhoneNumberParams{
		CustomerID: customerID,
		ID:         phoneNumberID,
	}

	if err := c.db.DeleteCustomerPhoneNumber(ctx, arg); err != nil {
		return err
	}

	return nil
}
