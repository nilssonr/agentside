package skill

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Skill struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	TenantID       string    `json:"tenantId"`
	LastModifiedBy string    `json:"lastModifiedBy"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}

type Service interface {
	CreateSkill(ctx context.Context, request *Skill) (*Skill, error)
	GetSkills(ctx context.Context, tenantID string) ([]*Skill, error)
	GetSkill(ctx context.Context, tenantID, skillID string) (*Skill, error)
	UpdateSkill(ctx context.Context, request *Skill) (*Skill, error)
	DeleteSkill(ctx context.Context, tenantID, skillID string) error
}

type service struct {
	skillRepository Repository
	logger          *zap.Logger
}

func NewService(sr Repository, l *zap.Logger) Service {
	return &service{
		skillRepository: sr,
		logger:          l,
	}
}

// CreateSkill implements Service.
func (s *service) CreateSkill(ctx context.Context, request *Skill) (*Skill, error) {
	result, err := s.skillRepository.InsertSkill(ctx, request)
	if err != nil {
		s.logger.Error("failed to create skill",
			zap.Error(err),
			zap.Any("request", request))
		return nil, err
	}

	return result, nil
}

// GetSkills implements Service.
func (s *service) GetSkills(ctx context.Context, tenantID string) ([]*Skill, error) {
	result, err := s.skillRepository.GetSkills(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetSkill implements Service.
func (s *service) GetSkill(ctx context.Context, tenantID string, skillID string) (*Skill, error) {
	result, err := s.skillRepository.GetSkill(ctx, tenantID, skillID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateSkill implements Service.
func (s *service) UpdateSkill(ctx context.Context, request *Skill) (*Skill, error) {
	result, err := s.skillRepository.UpdateSkill(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteSkill implements Service.
func (s *service) DeleteSkill(ctx context.Context, tenantID string, skillID string) error {
	if err := s.skillRepository.DeleteSkill(ctx, tenantID, skillID); err != nil {
		return nil
	}

	return nil
}
