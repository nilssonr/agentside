package skill

import "context"

type Repository interface {
	InsertSkill(ctx context.Context, request *Skill) (*Skill, error)
	GetSkills(ctx context.Context, tenantID string) ([]*Skill, error)
	GetSkill(ctx context.Context, tenantID, skillID string) (*Skill, error)
	UpdateSkill(ctx context.Context, request *Skill) (*Skill, error)
	DeleteSkill(ctx context.Context, tenantID, skillID string) error
}
