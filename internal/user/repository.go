package user

import "context"

type Repository interface {
	InsertUser(ctx context.Context, request *User) (*User, error)
	GetUsers(ctx context.Context, tenantID string) ([]*User, error)
	GetUser(ctx context.Context, tenantID, userID string) (*User, error)
	UpdateUser(ctx context.Context, request *User) (*User, error)
	DeleteUser(ctx context.Context, tenantID, userID string) error
}

type SkillRepository interface {
	UpsertSkill(ctx context.Context, userID, skillID string, level int) (*Skill, error)
	GetSkills(ctx context.Context, userID string) ([]*Skill, error)
	GetSkill(ctx context.Context, userID, skillID string) (*Skill, error)
	DeleteSkill(ctx context.Context, userID, skillID string) error
}
