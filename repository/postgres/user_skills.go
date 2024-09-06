package postgres

import (
	"context"

	"github.com/nilssonr/agentside/user"
)

type UserSkillRepository struct {
	db *Queries
}

func NewUserSkillRepository(db *Queries) user.SkillRepository {
	return &UserSkillRepository{
		db: db,
	}
}

// InsertSkill implements user.SkillRepository.
func (usr *UserSkillRepository) UpsertSkill(ctx context.Context, userID, skillID string, level int) (*user.Skill, error) {
	arg := UpsertUserSkillParams{
		UserID:     userID,
		SkillID:    skillID,
		SkillLevel: int32(level),
	}

	row, err := usr.db.UpsertUserSkill(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result user.Skill
	result.ID = row.ID
	result.Name = row.Name
	result.Level = int(row.SkillLevel)

	return &result, nil
}

// GetSkills implements user.SkillRepository.
func (usr *UserSkillRepository) GetSkills(ctx context.Context, userID string) ([]*user.Skill, error) {
	rows, err := usr.db.GetUserSkills(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := []*user.Skill{}
	for _, v := range rows {
		result = append(result, &user.Skill{
			ID:    v.ID,
			Name:  v.Name,
			Level: int(v.SkillLevel),
		})
	}

	return result, nil
}

// GetSkill implements user.SkillRepository.
func (usr *UserSkillRepository) GetSkill(ctx context.Context, userID string, skillID string) (*user.Skill, error) {
	arg := GetUserSkillParams{
		UserID: userID,
		ID:     skillID,
	}

	row, err := usr.db.GetUserSkill(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result user.Skill
	result.ID = row.ID
	result.Name = row.Name
	result.Level = int(row.SkillLevel)

	return &result, nil
}

// DeleteSkill implements user.SkillRepository.
func (usr *UserSkillRepository) DeleteSkill(ctx context.Context, userID string, skillID string) error {
	arg := DeleteUserSkillParams{
		UserID:  userID,
		SkillID: skillID,
	}

	if err := usr.db.DeleteUserSkill(ctx, arg); err != nil {
		return err
	}

	return nil
}
