package postgres

import (
	"context"

	"github.com/nilssonr/agentside/repository/postgres/sqlc"
	"github.com/nilssonr/agentside/skill"
)

type SkillRepository struct {
	DB *sqlc.Queries
}

func NewSkillRepository(db *sqlc.Queries) skill.Repository {
	return &SkillRepository{
		DB: db,
	}
}

// InsertSkill implements skill.Repository.
func (sr *SkillRepository) InsertSkill(ctx context.Context, request *skill.Skill) (*skill.Skill, error) {
	arg := sqlc.CreateSkillParams{
		Name:           request.Name,
		TenantID:       request.TenantID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	s, err := sr.DB.CreateSkill(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result skill.Skill
	result.ID = s.ID
	result.Name = s.Name
	result.TenantID = s.TenantID
	result.LastModifiedAt = s.LastModifiedAt.Time
	result.LastModifiedBy = s.LastModifiedBy

	return &result, nil
}

// GetSkills implements skill.Repository.
func (sr *SkillRepository) GetSkills(ctx context.Context, tenantID string) ([]*skill.Skill, error) {
	s, err := sr.DB.GetSkills(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	result := []*skill.Skill{}
	for _, v := range s {
		result = append(result, &skill.Skill{
			ID:             v.ID,
			Name:           v.Name,
			TenantID:       v.TenantID,
			LastModifiedBy: v.LastModifiedBy,
			LastModifiedAt: v.LastModifiedAt.Time,
		})
	}

	return result, nil
}

// GetSkill implements skill.Repository.
func (sr *SkillRepository) GetSkill(ctx context.Context, tenantID string, skillID string) (*skill.Skill, error) {
	arg := sqlc.GetSkillParams{
		ID:       skillID,
		TenantID: tenantID,
	}

	s, err := sr.DB.GetSkill(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result skill.Skill
	result.ID = s.ID
	result.Name = s.Name
	result.TenantID = s.TenantID
	result.LastModifiedAt = s.LastModifiedAt.Time
	result.LastModifiedBy = s.LastModifiedBy

	return &result, nil
}

// UpdateSkill implements skill.Repository.
func (sr *SkillRepository) UpdateSkill(ctx context.Context, request *skill.Skill) (*skill.Skill, error) {
	arg := sqlc.UpdateSkillParams{
		ID:             request.ID,
		Name:           request.Name,
		TenantID:       request.TenantID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	s, err := sr.DB.UpdateSkill(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result skill.Skill
	result.ID = s.ID
	result.Name = s.Name
	result.TenantID = s.TenantID
	result.LastModifiedAt = s.LastModifiedAt.Time
	result.LastModifiedBy = s.LastModifiedBy

	return &result, nil
}

// DeleteSkill implements skill.Repository.
func (sr *SkillRepository) DeleteSkill(ctx context.Context, tenantID string, skillID string) error {
	arg := sqlc.DeleteSkillParams{
		ID:       skillID,
		TenantID: tenantID,
	}

	if err := sr.DB.DeleteSkill(ctx, arg); err != nil {
		return err
	}

	return nil
}
