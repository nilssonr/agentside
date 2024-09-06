package postgres

import (
	"context"

	"github.com/nilssonr/agentside/skill"
)

type SkillRepository struct {
	db *Queries
}

func NewSkillRepository(db *Queries) skill.Repository {
	return &SkillRepository{
		db: db,
	}
}

// InsertSkill implements skill.Repository.
func (sr *SkillRepository) InsertSkill(ctx context.Context, request *skill.Skill) (*skill.Skill, error) {
	arg := CreateSkillParams{
		Name:           request.Name,
		TenantID:       request.TenantID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	s, err := sr.db.CreateSkill(ctx, arg)
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
	s, err := sr.db.GetSkills(ctx, tenantID)
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
	arg := GetSkillParams{
		ID:       skillID,
		TenantID: tenantID,
	}

	s, err := sr.db.GetSkill(ctx, arg)
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
	arg := UpdateSkillParams{
		ID:             request.ID,
		Name:           request.Name,
		TenantID:       request.TenantID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	s, err := sr.db.UpdateSkill(ctx, arg)
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
	arg := DeleteSkillParams{
		ID:       skillID,
		TenantID: tenantID,
	}

	if err := sr.db.DeleteSkill(ctx, arg); err != nil {
		return err
	}

	return nil
}
