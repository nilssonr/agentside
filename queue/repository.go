package queue

import "context"

type Repository interface {
	InsertQueue(ctx context.Context, request *Queue) (*Queue, error)
	GetQueues(ctx context.Context, tenantID string) ([]*Queue, error)
	GetQueue(ctx context.Context, tenantID, queueID string) (*Queue, error)
	UpdateQueue(ctx context.Context, request *Queue) (*Queue, error)
	DeleteQueue(ctx context.Context, tenantID, queueID string) error
}

type SkillRepository interface {
	UpsertSkill(ctx context.Context, queueID string, request *Skill) (*Skill, error)
	GetSkills(ctx context.Context, queueID string) ([]*Skill, error)
	GetSkill(ctx context.Context, queueID, skillID string) (*Skill, error)
	DeleteSkill(ctx context.Context, queueID, skillID string) error
}
