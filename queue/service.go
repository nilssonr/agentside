package queue

import (
	"context"

	"go.uber.org/zap"
)

type Service interface {
	CreateQueue(ctx context.Context, request *Queue) (*Queue, error)
	GetQueues(ctx context.Context, tenantID string) ([]*Queue, error)
	GetQueue(ctx context.Context, tenantID, queueID string) (*Queue, error)
	UpdateQueue(ctx context.Context, request *Queue) (*Queue, error)
	DeleteQueue(ctx context.Context, tenantID, queueID string) error
}

type service struct {
	queueRepository Repository
	logger          *zap.Logger
}

func NewService(qr Repository, l *zap.Logger) Service {
	return &service{
		queueRepository: qr,
		logger:          l,
	}
}

// CreateQueue implements Service.
func (s *service) CreateQueue(ctx context.Context, request *Queue) (*Queue, error) {
	result, err := s.queueRepository.InsertQueue(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetQueues implements Service.
func (s *service) GetQueues(ctx context.Context, tenantID string) ([]*Queue, error) {
	result, err := s.queueRepository.GetQueues(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetQueue implements Service.
func (s *service) GetQueue(ctx context.Context, tenantID string, queueID string) (*Queue, error) {
	result, err := s.queueRepository.GetQueue(ctx, tenantID, queueID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateQueue implements Service.
func (s *service) UpdateQueue(ctx context.Context, request *Queue) (*Queue, error) {
	result, err := s.queueRepository.UpdateQueue(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteQueue implements Service.
func (s *service) DeleteQueue(ctx context.Context, tenantID string, queueID string) error {
	if err := s.queueRepository.DeleteQueue(ctx, tenantID, queueID); err != nil {
		return err
	}

	return nil
}
