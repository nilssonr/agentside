package user

import (
	"context"
	"log"
)

type SkillService interface {
	UpsertSkill(ctx context.Context, userID, skillID string, level int) (*Skill, error)
	GetSkills(ctx context.Context, userID string) ([]*Skill, error)
	GetSkill(ctx context.Context, userID, skillID string) (*Skill, error)
	DeleteSkill(ctx context.Context, userID, skillID string) error
}

type skillService struct {
	skillRepository SkillRepository
}

func NewSkillService(sr SkillRepository) SkillService {
	return &skillService{
		skillRepository: sr,
	}
}

// CreateSkill implements SkillService.
func (s *skillService) UpsertSkill(ctx context.Context, userID string, skillID string, level int) (*Skill, error) {
	result, err := s.skillRepository.UpsertSkill(ctx, userID, skillID, level)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

// GetSkills implements SkillService.
func (s *skillService) GetSkills(ctx context.Context, userID string) ([]*Skill, error) {
	result, err := s.skillRepository.GetSkills(ctx, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetSkill implements SkillService.
func (s *skillService) GetSkill(ctx context.Context, userID string, skillID string) (*Skill, error) {
	result, err := s.skillRepository.GetSkill(ctx, userID, skillID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteSkill implements SkillService.
func (s *skillService) DeleteSkill(ctx context.Context, userID string, skillID string) error {
	if err := s.skillRepository.DeleteSkill(ctx, userID, skillID); err != nil {
		return err
	}

	return nil
}
