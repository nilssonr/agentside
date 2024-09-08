package customer

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type PhoneNumber struct {
	ID             string    `json:"id"`
	PhoneNumber    string    `json:"phoneNumber"`
	Type           string    `json:"type"`
	CustomerID     string    `json:"customerId"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}

type PhoneNumberService interface {
	CreatePhoneNumber(ctx context.Context, request *PhoneNumber) (*PhoneNumber, error)
	GetPhoneNumbers(ctx context.Context, customerID string) ([]*PhoneNumber, error)
	GetPhoneNumber(ctx context.Context, customerID, phoneNumberID string) (*PhoneNumber, error)
	UpdatePhoneNumber(ctx context.Context, request *PhoneNumber) (*PhoneNumber, error)
	DeletePhoneNumber(ctx context.Context, customerID, phoneNumberID string) error
}

type phoneNumberService struct {
	phoneNumberRepository PhoneNumberRepository
	logger                *zap.Logger
}

func NewPhoneNumberService(r PhoneNumberRepository, l *zap.Logger) PhoneNumberService {
	return &phoneNumberService{
		phoneNumberRepository: r,
		logger:                l,
	}
}

// CreatePhoneNumber implements PhoneNumberService.
func (s *phoneNumberService) CreatePhoneNumber(ctx context.Context, request *PhoneNumber) (*PhoneNumber, error) {
	result, err := s.phoneNumberRepository.InsertPhoneNumber(ctx, request)
	if err != nil {
		s.logger.Error("failed to create customer phone number",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetPhoneNumbers implements PhoneNumberService.
func (s *phoneNumberService) GetPhoneNumbers(ctx context.Context, customerID string) ([]*PhoneNumber, error) {
	result, err := s.phoneNumberRepository.GetPhoneNumbers(ctx, customerID)
	if err != nil {
		s.logger.Error("failed to get customer phone numbers",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetPhoneNumber implements PhoneNumberService.
func (s *phoneNumberService) GetPhoneNumber(ctx context.Context, customerID string, phoneNumberID string) (*PhoneNumber, error) {
	result, err := s.phoneNumberRepository.GetPhoneNumber(ctx, customerID, phoneNumberID)
	if err != nil {
		s.logger.Error("failed to get customer phone number",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// UpdatePhoneNumber implements PhoneNumberService.
func (s *phoneNumberService) UpdatePhoneNumber(ctx context.Context, request *PhoneNumber) (*PhoneNumber, error) {
	result, err := s.phoneNumberRepository.UpdatePhoneNumber(ctx, request)
	if err != nil {
		s.logger.Error("failed to update customer phone number",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// DeletePhoneNumber implements PhoneNumberService.
func (s *phoneNumberService) DeletePhoneNumber(ctx context.Context, customerID string, phoneNumberID string) error {
	if err := s.phoneNumberRepository.DeletePhoneNumber(ctx, customerID, phoneNumberID); err != nil {
		s.logger.Error("failed to delete customer phone number",
			zap.Error(err))
		return err
	}

	return nil
}
