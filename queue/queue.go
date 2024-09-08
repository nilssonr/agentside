package queue

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Queue struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	TenantID       string    `json:"tenantId"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}

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
		s.logger.Error("failed to create queue",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetQueues implements Service.
func (s *service) GetQueues(ctx context.Context, tenantID string) ([]*Queue, error) {
	result, err := s.queueRepository.GetQueues(ctx, tenantID)
	if err != nil {
		s.logger.Error("failed to get queues",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetQueue implements Service.
func (s *service) GetQueue(ctx context.Context, tenantID string, queueID string) (*Queue, error) {
	result, err := s.queueRepository.GetQueue(ctx, tenantID, queueID)
	if err != nil {
		s.logger.Error("failed to get queue",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// UpdateQueue implements Service.
func (s *service) UpdateQueue(ctx context.Context, request *Queue) (*Queue, error) {
	result, err := s.queueRepository.UpdateQueue(ctx, request)
	if err != nil {
		s.logger.Error("failed to update queue",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// DeleteQueue implements Service.
func (s *service) DeleteQueue(ctx context.Context, tenantID string, queueID string) error {
	if err := s.queueRepository.DeleteQueue(ctx, tenantID, queueID); err != nil {
		s.logger.Error("failed to delete queue",
			zap.Error(err))
		return err
	}

	return nil
}
