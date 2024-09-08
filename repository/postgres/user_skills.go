package postgres

import (
	"context"

	"github.com/nilssonr/agentside/repository/postgres/sqlc"
	"github.com/nilssonr/agentside/user"
)

type UserSkillRepository struct {
	DB *sqlc.Queries
}

func NewUserSkillRepository(db *sqlc.Queries) user.SkillRepository {
	return &UserSkillRepository{
		DB: db,
	}
}

// InsertSkill implements user.SkillRepository.
func (usr *UserSkillRepository) UpsertSkill(ctx context.Context, userID, skillID string, level int) (*user.Skill, error) {
	arg := sqlc.UpsertUserSkillParams{
		UserID:     userID,
		SkillID:    skillID,
		SkillLevel: int32(level),
	}

	row, err := usr.DB.UpsertUserSkill(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := &user.Skill{
		ID:    row.ID,
		Name:  row.Name,
		Level: int(row.SkillLevel),
	}

	return result, nil
}

// GetSkills implements user.SkillRepository.
func (usr *UserSkillRepository) GetSkills(ctx context.Context, userID string) ([]*user.Skill, error) {
	rows, err := usr.DB.GetUserSkills(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]*user.Skill, 0, len(rows))
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
	arg := sqlc.GetUserSkillParams{
		UserID: userID,
		ID:     skillID,
	}

	row, err := usr.DB.GetUserSkill(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := &user.Skill{
		ID:    row.ID,
		Name:  row.Name,
		Level: int(row.SkillLevel),
	}

	return result, nil
}

// DeleteSkill implements user.SkillRepository.
func (usr *UserSkillRepository) DeleteSkill(ctx context.Context, userID string, skillID string) error {
	arg := sqlc.DeleteUserSkillParams{
		UserID:  userID,
		SkillID: skillID,
	}

	if err := usr.DB.DeleteUserSkill(ctx, arg); err != nil {
		return err
	}

	return nil
}
