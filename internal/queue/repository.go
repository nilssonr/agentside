package queue

import "context"

type Repository interface {
	InsertQueue(ctx context.Context, request *Queue) (*Queue, error)
	GetQueues(ctx context.Context, tenantID string) ([]*Queue, error)
	GetQueue(ctx context.Context, tenantID, queueID string) (*Queue, error)
	UpdateQueue(ctx context.Context, request *Queue) (*Queue, error)
	DeleteQueue(ctx context.Context, tenantID, queueID string) error
}
