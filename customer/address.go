package customer

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Address struct {
	ID             string    `json:"id"`
	Type           string    `json:"type"`
	StreetAddress  string    `json:"streetAddress"`
	State          string    `json:"state"`
	ZipCode        string    `json:"zipCode"`
	Country        string    `json:"country"`
	CustomerID     string    `json:"customerId"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}

type AddressService interface {
	CreateAddress(ctx context.Context, request *Address) (*Address, error)
	GetAddresses(ctx context.Context, customerID string) ([]*Address, error)
	GetAddress(ctx context.Context, customerID, addressID string) (*Address, error)
	UpdateAddress(ctx context.Context, request *Address) (*Address, error)
	DeleteAddress(ctx context.Context, customerID, addressID string) error
}

type addressService struct {
	addressRepository AddressRepository
	logger            *zap.Logger
}

func NewAddressService(r AddressRepository, l *zap.Logger) AddressService {
	return &addressService{
		addressRepository: r,
		logger:            l,
	}
}

// CreateAddress implements AddressService.
func (a *addressService) CreateAddress(ctx context.Context, request *Address) (*Address, error) {
	result, err := a.addressRepository.InsertAddress(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAddresses implements AddressService.
func (a *addressService) GetAddresses(ctx context.Context, customerID string) ([]*Address, error) {
	result, err := a.addressRepository.GetAddresses(ctx, customerID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAddress implements AddressService.
func (a *addressService) GetAddress(ctx context.Context, customerID string, addressID string) (*Address, error) {
	result, err := a.addressRepository.GetAddress(ctx, customerID, addressID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateAddress implements AddressService.
func (a *addressService) UpdateAddress(ctx context.Context, request *Address) (*Address, error) {
	result, err := a.addressRepository.UpdateAddress(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteAddress implements AddressService.
func (a *addressService) DeleteAddress(ctx context.Context, customerID string, addressID string) error {
	if err := a.addressRepository.DeleteAddress(ctx, customerID, addressID); err != nil {
		return err
	}

	return nil
}
