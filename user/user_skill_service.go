package user

import (
	"context"

	"go.uber.org/zap"
)

type SkillService interface {
	UpsertSkill(ctx context.Context, userID, skillID string, level int) (*Skill, error)
	GetSkills(ctx context.Context, userID string) ([]*Skill, error)
	GetSkill(ctx context.Context, userID, skillID string) (*Skill, error)
	DeleteSkill(ctx context.Context, userID, skillID string) error
}

type skillService struct {
	skillRepository SkillRepository
	logger          *zap.Logger
}

func NewSkillService(sr SkillRepository, l *zap.Logger) SkillService {
	return &skillService{
		skillRepository: sr,
		logger:          l,
	}
}

// CreateSkill implements SkillService.
func (s *skillService) UpsertSkill(ctx context.Context, userID string, skillID string, level int) (*Skill, error) {
	result, err := s.skillRepository.UpsertSkill(ctx, userID, skillID, level)
	if err != nil {
		s.logger.Error("failed to upsert skill",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetSkills implements SkillService.
func (s *skillService) GetSkills(ctx context.Context, userID string) ([]*Skill, error) {
	result, err := s.skillRepository.GetSkills(ctx, userID)
	if err != nil {
		s.logger.Error("failed to get skills",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetSkill implements SkillService.
func (s *skillService) GetSkill(ctx context.Context, userID string, skillID string) (*Skill, error) {
	result, err := s.skillRepository.GetSkill(ctx, userID, skillID)
	if err != nil {
		s.logger.Error("failed to get skill",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// DeleteSkill implements SkillService.
func (s *skillService) DeleteSkill(ctx context.Context, userID string, skillID string) error {
	if err := s.skillRepository.DeleteSkill(ctx, userID, skillID); err != nil {
		s.logger.Error("failed to delete skill",
			zap.Error(err))
		return err
	}

	return nil
}
