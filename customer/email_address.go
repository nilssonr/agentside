package customer

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type EmailAddress struct {
	ID             string    `json:"id"`
	EmailAddress   string    `json:"emailAddress"`
	Type           string    `json:"type"`
	CustomerID     string    `json:"customerId"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}

type EmailAddressService interface {
	CreateEmailAddress(ctx context.Context, request *EmailAddress) (*EmailAddress, error)
	GetEmailAddresses(ctx context.Context, customerID string) ([]*EmailAddress, error)
	GetEmailAddress(ctx context.Context, customerID, emailAddressID string) (*EmailAddress, error)
	UpdateEmailAddress(ctx context.Context, request *EmailAddress) (*EmailAddress, error)
	DeleteEmailAddress(ctx context.Context, customerID, emailAddressID string) error
}

type emailAddressService struct {
	emailAddressRepository EmailAddressRepository
	logger                 *zap.Logger
}

func NewEmailAddressService(r EmailAddressRepository, l *zap.Logger) EmailAddressService {
	return &emailAddressService{
		emailAddressRepository: r,
		logger:                 l,
	}
}

// CreateEmailAddress implements EmailAddressService.
func (s *emailAddressService) CreateEmailAddress(ctx context.Context, request *EmailAddress) (*EmailAddress, error) {
	result, err := s.emailAddressRepository.InsertEmailAddress(ctx, request)
	if err != nil {
		s.logger.Error("failed to create customer email address",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetEmailAddresses implements EmailAddressService.
func (s *emailAddressService) GetEmailAddresses(ctx context.Context, customerID string) ([]*EmailAddress, error) {
	result, err := s.emailAddressRepository.GetEmailAddresses(ctx, customerID)
	if err != nil {
		s.logger.Error("failed to get customer email addresses",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetEmailAddress implements EmailAddressService.
func (s *emailAddressService) GetEmailAddress(ctx context.Context, customerID, emailAddressID string) (*EmailAddress, error) {
	result, err := s.emailAddressRepository.GetEmailAddress(ctx, customerID, emailAddressID)
	if err != nil {
		s.logger.Error("failed to get customer email address",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// UpdateEmailAddress implements EmailAddressService.
func (s *emailAddressService) UpdateEmailAddress(ctx context.Context, request *EmailAddress) (*EmailAddress, error) {
	result, err := s.emailAddressRepository.UpdateEmailAddress(ctx, request)
	if err != nil {
		s.logger.Error("failed to update customer email address",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// DeleteEmailAddress implements EmailAddressService.
func (s *emailAddressService) DeleteEmailAddress(ctx context.Context, customerID, emailAddressID string) error {
	if err := s.emailAddressRepository.DeleteEmailAddress(ctx, customerID, emailAddressID); err != nil {
		s.logger.Error("failed to delete customer email address",
			zap.Error(err))
		return err
	}

	return nil
}
