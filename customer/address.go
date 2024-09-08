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
func (s *addressService) CreateAddress(ctx context.Context, request *Address) (*Address, error) {
	result, err := s.addressRepository.InsertAddress(ctx, request)
	if err != nil {
		s.logger.Error("failed to create customer address",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetAddresses implements AddressService.
func (s *addressService) GetAddresses(ctx context.Context, customerID string) ([]*Address, error) {
	result, err := s.addressRepository.GetAddresses(ctx, customerID)
	if err != nil {
		s.logger.Error("failed to get customer addresses",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetAddress implements AddressService.
func (s *addressService) GetAddress(ctx context.Context, customerID string, addressID string) (*Address, error) {
	result, err := s.addressRepository.GetAddress(ctx, customerID, addressID)
	if err != nil {
		s.logger.Error("failed to get customer address",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// UpdateAddress implements AddressService.
func (s *addressService) UpdateAddress(ctx context.Context, request *Address) (*Address, error) {
	result, err := s.addressRepository.UpdateAddress(ctx, request)
	if err != nil {
		s.logger.Error("failed to update customer address",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// DeleteAddress implements AddressService.
func (s *addressService) DeleteAddress(ctx context.Context, customerID string, addressID string) error {
	if err := s.addressRepository.DeleteAddress(ctx, customerID, addressID); err != nil {
		s.logger.Error("failed to delete customer address",
			zap.Error(err))
		return err
	}

	return nil
}
