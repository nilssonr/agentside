package postgres

import (
	"context"
	"fmt"

	"github.com/nilssonr/agentside/queue"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

type QueueSkillRepository struct {
	DB *sqlc.Queries
}

func NewQueueSkillRepository(db *sqlc.Queries) queue.SkillRepository {
	return &QueueSkillRepository{
		DB: db,
	}
}

// UpsertSkill implements queue.SkillRepository.
func (r *QueueSkillRepository) UpsertSkill(ctx context.Context, queueID string, request *queue.Skill) (*queue.Skill, error) {
	arg := sqlc.UpsertQueueSkillParams{
		QueueID: queueID,
		SkillID: request.ID,
		Level:   int32(request.Level),
		Choice:  int32(request.Choice),
	}

	row, err := r.DB.UpsertQueueSkill(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &queue.Skill{
		ID:     row.ID,
		Name:   row.Name,
		Level:  int(row.Level),
		Choice: int(row.Choice),
	}

	return result, nil
}

// GetSkills implements queue.SkillRepository.
func (r *QueueSkillRepository) GetSkills(ctx context.Context, queueID string) ([]*queue.Skill, error) {
	rows, err := r.DB.GetQueueSkills(ctx, queueID)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := make([]*queue.Skill, 0, len(rows))
	for _, v := range rows {
		result = append(result, &queue.Skill{
			ID:     v.ID,
			Name:   v.Name,
			Level:  int(v.Level),
			Choice: int(v.Choice),
		})
	}

	return result, nil
}

// GetSkill implements queue.SkillRepository.
func (r *QueueSkillRepository) GetSkill(ctx context.Context, queueID, skillID string) (*queue.Skill, error) {
	arg := sqlc.GetQueueSkillParams{
		QueueID: queueID,
		ID:      skillID,
	}

	row, err := r.DB.GetQueueSkill(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &queue.Skill{
		ID:     row.ID,
		Name:   row.Name,
		Level:  int(row.Level),
		Choice: int(row.Choice),
	}

	return result, nil
}

// DeleteSkill implements queue.SkillRepository.
func (r *QueueSkillRepository) DeleteSkill(ctx context.Context, queueID, skillID string) error {
	arg := sqlc.DeleteQueueSkillParams{
		QueueID: queueID,
		SkillID: skillID,
	}

	if err := r.DB.DeleteQueueSkill(ctx, arg); err != nil {
		return fmt.Errorf("repository: %w", err)
	}

	return nil
}
