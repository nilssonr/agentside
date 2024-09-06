package postgres

import (
	"context"
	"time"

	"github.com/nilssonr/agentside/queue"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

type QueueRepository struct {
	DB *sqlc.Queries
}

func NewQueueRepository(db *sqlc.Queries) queue.Repository {
	return &QueueRepository{
		DB: db,
	}
}

// InsertQueue implements queue.Repository.
func (q *QueueRepository) InsertQueue(ctx context.Context, request *queue.Queue) (*queue.Queue, error) {
	arg := sqlc.InsertQueueParams{
		Name:           request.Name,
		TenantID:       request.TenantID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := q.DB.InsertQueue(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result queue.Queue
	result.ID = row.ID
	result.Name = row.Name
	result.TenantID = row.TenantID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// GetQueues implements queue.Repository.
func (q *QueueRepository) GetQueues(ctx context.Context, tenantID string) ([]*queue.Queue, error) {
	rows, err := q.DB.GetQueues(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	result := []*queue.Queue{}
	for _, v := range rows {
		result = append(result, &queue.Queue{
			ID:             v.ID,
			Name:           v.Name,
			TenantID:       v.TenantID,
			LastModifiedAt: v.LastModifiedAt.Time,
			LastModifiedBy: v.LastModifiedBy,
		})
	}

	return result, nil
}

// GetQueue implements queue.Repository.
func (q *QueueRepository) GetQueue(ctx context.Context, tenantID string, queueID string) (*queue.Queue, error) {
	arg := sqlc.GetQueueParams{
		TenantID: tenantID,
		ID:       queueID,
	}

	row, err := q.DB.GetQueue(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result queue.Queue
	result.ID = row.ID
	result.Name = row.Name
	result.TenantID = row.TenantID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// UpdateQueue implements queue.Repository.
func (q *QueueRepository) UpdateQueue(ctx context.Context, request *queue.Queue) (*queue.Queue, error) {
	arg := sqlc.UpdateQueueParams{
		ID:             request.ID,
		Name:           request.Name,
		TenantID:       request.TenantID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := q.DB.UpdateQueue(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result queue.Queue
	result.ID = row.ID
	result.Name = row.Name
	result.TenantID = row.TenantID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// DeleteQueue implements queue.Repository.
func (q *QueueRepository) DeleteQueue(ctx context.Context, tenantID string, queueID string) error {
	arg := sqlc.DeleteQueueParams{
		TenantID:  tenantID,
		ID:        queueID,
		DeletedAt: mustCreateTime(time.Now()),
	}

	if err := q.DB.DeleteQueue(ctx, arg); err != nil {
		return err
	}

	return nil
}
