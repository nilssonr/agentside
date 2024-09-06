package customer

import (
	"context"
	"time"
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
}

func NewPhoneNumberService(r PhoneNumberRepository) PhoneNumberService {
	return &phoneNumberService{
		phoneNumberRepository: r,
	}
}

// CreatePhoneNumber implements PhoneNumberService.
func (p *phoneNumberService) CreatePhoneNumber(ctx context.Context, request *PhoneNumber) (*PhoneNumber, error) {
	result, err := p.phoneNumberRepository.InsertPhoneNumber(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetPhoneNumbers implements PhoneNumberService.
func (p *phoneNumberService) GetPhoneNumbers(ctx context.Context, customerID string) ([]*PhoneNumber, error) {
	result, err := p.phoneNumberRepository.GetPhoneNumbers(ctx, customerID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetPhoneNumber implements PhoneNumberService.
func (p *phoneNumberService) GetPhoneNumber(ctx context.Context, customerID string, phoneNumberID string) (*PhoneNumber, error) {
	result, err := p.phoneNumberRepository.GetPhoneNumber(ctx, customerID, phoneNumberID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdatePhoneNumber implements PhoneNumberService.
func (p *phoneNumberService) UpdatePhoneNumber(ctx context.Context, request *PhoneNumber) (*PhoneNumber, error) {
	result, err := p.phoneNumberRepository.UpdatePhoneNumber(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeletePhoneNumber implements PhoneNumberService.
func (p *phoneNumberService) DeletePhoneNumber(ctx context.Context, customerID string, phoneNumberID string) error {
	if err := p.phoneNumberRepository.DeletePhoneNumber(ctx, customerID, phoneNumberID); err != nil {
		return err
	}

	return nil
}
