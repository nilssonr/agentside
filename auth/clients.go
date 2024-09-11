package auth

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Client struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Secret         string    `json:"secret"`
	TenantID       string    `json:"tenantId"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}

type ClientService interface {
	CreateClient(ctx context.Context, request *Client) (*Client, error)
	GetClients(ctx context.Context, tenantID string) ([]*Client, error)
	GetClient(ctx context.Context, tenantID, clientID string) (*Client, error)
	UpdateClient(ctx context.Context, request *Client) (*Client, error)
	DeleteClient(ctx context.Context, tenantID, clientID string) error
}

type clientService struct {
	clientRepository ClientRepository
	logger           *zap.Logger
}

func NewClientService(r ClientRepository, l *zap.Logger) ClientService {
	return &clientService{
		clientRepository: r,
		logger:           l,
	}
}

// CreateClient implements ClientService.
func (s *clientService) CreateClient(ctx context.Context, request *Client) (*Client, error) {
	secret, err := generateClientSecret(24)
	if err != nil {
		s.logger.Error("failed to create auth client", zap.Error(err))
		return nil, err
	}

	request.Secret = secret

	result, err := s.clientRepository.InsertClient(ctx, request)
	if err != nil {
		s.logger.Error("failed to create auth client",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetClients implements ClientService.
func (s *clientService) GetClients(ctx context.Context, tenantID string) ([]*Client, error) {
	result, err := s.clientRepository.GetClients(ctx, tenantID)
	if err != nil {
		s.logger.Error("failed to get auth clients",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetClient implements ClientService.
func (s *clientService) GetClient(ctx context.Context, tenantID, clientID string) (*Client, error) {
	result, err := s.clientRepository.GetClient(ctx, tenantID, clientID)
	if err != nil {
		s.logger.Error("failed to get auth client",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// UpdateClient implements ClientService.
func (s *clientService) UpdateClient(ctx context.Context, request *Client) (*Client, error) {
	result, err := s.clientRepository.UpdateClient(ctx, request)
	if err != nil {
		s.logger.Error("failed to update auth client",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// DeleteClient implements ClientService.
func (s *clientService) DeleteClient(ctx context.Context, tenantID, clientID string) error {
	if err := s.clientRepository.DeleteClient(ctx, tenantID, clientID); err != nil {
		s.logger.Error("failed to delete auth client",
			zap.Error(err))
		return err
	}

	return nil
}
