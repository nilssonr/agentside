package queue

import (
	"context"

	"go.uber.org/zap"
)

type Skill struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Choice int    `json:"choice"`
}

type SkillService interface {
	UpsertSkill(ctx context.Context, queueID string, request *Skill) (*Skill, error)
	GetSkills(ctx context.Context, queueID string) ([]*Skill, error)
	GetSkill(ctx context.Context, queueID, skillID string) (*Skill, error)
	DeleteSkill(ctx context.Context, queueID, skillID string) error
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

// UpsertSkill implements SkillService.
func (s *skillService) UpsertSkill(ctx context.Context, queueID string, request *Skill) (*Skill, error) {
	result, err := s.skillRepository.UpsertSkill(ctx, queueID, request)
	if err != nil {
		s.logger.Error("failed to upsert queue skill",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetSkills implements SkillService.
func (s *skillService) GetSkills(ctx context.Context, queueID string) ([]*Skill, error) {
	result, err := s.skillRepository.GetSkills(ctx, queueID)
	if err != nil {
		s.logger.Error("failed to get queue skills",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetSkill implements SkillService.
func (s *skillService) GetSkill(ctx context.Context, queueID, skillID string) (*Skill, error) {
	result, err := s.skillRepository.GetSkill(ctx, queueID, skillID)
	if err != nil {
		s.logger.Error("failed to get queue skill",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// DeleteSkill implements SkillService.
func (s *skillService) DeleteSkill(ctx context.Context, queueID, skillID string) error {
	if err := s.skillRepository.DeleteSkill(ctx, queueID, skillID); err != nil {
		s.logger.Error("failed to delete queue skill",
			zap.Error(err))
		return err
	}

	return nil
}
