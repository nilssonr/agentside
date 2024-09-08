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
func (e *emailAddressService) CreateEmailAddress(ctx context.Context, request *EmailAddress) (*EmailAddress, error) {
	result, err := e.emailAddressRepository.InsertEmailAddress(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetEmailAddresses implements EmailAddressService.
func (e *emailAddressService) GetEmailAddresses(ctx context.Context, customerID string) ([]*EmailAddress, error) {
	result, err := e.emailAddressRepository.GetEmailAddresses(ctx, customerID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetEmailAddress implements EmailAddressService.
func (e *emailAddressService) GetEmailAddress(ctx context.Context, customerID, emailAddressID string) (*EmailAddress, error) {
	result, err := e.emailAddressRepository.GetEmailAddress(ctx, customerID, emailAddressID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateEmailAddress implements EmailAddressService.
func (e *emailAddressService) UpdateEmailAddress(ctx context.Context, request *EmailAddress) (*EmailAddress, error) {
	result, err := e.emailAddressRepository.UpdateEmailAddress(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteEmailAddress implements EmailAddressService.
func (e *emailAddressService) DeleteEmailAddress(ctx context.Context, customerID, emailAddressID string) error {
	if err := e.emailAddressRepository.DeleteEmailAddress(ctx, customerID, emailAddressID); err != nil {
		return err
	}

	return nil
}
