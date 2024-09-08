package user

import (
	"context"

	"go.uber.org/zap"
)

type Service interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUsers(ctx context.Context, tenantID string) ([]*User, error)
	GetUser(ctx context.Context, tenantID, userID string) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	DeleteUser(ctx context.Context, tenantID, userID string) error
}

type service struct {
	userRepository Repository
	logger         *zap.Logger
}

func NewService(ur Repository, l *zap.Logger) Service {
	return &service{
		userRepository: ur,
		logger:         l,
	}
}

// CreateUser implements Service.
func (s *service) CreateUser(ctx context.Context, u *User) (*User, error) {
	u, err := s.userRepository.InsertUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUsers implements Service.
func (s *service) GetUsers(ctx context.Context, tenantID string) ([]*User, error) {
	u, err := s.userRepository.GetUsers(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUser implements Service.
func (s *service) GetUser(ctx context.Context, tenantID string, userID string) (*User, error) {
	u, err := s.userRepository.GetUser(ctx, tenantID, userID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// UpdateUser implements Service.
func (s *service) UpdateUser(ctx context.Context, u *User) (*User, error) {
	u, err := s.userRepository.UpdateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// DeleteUser implements Service.
func (s *service) DeleteUser(ctx context.Context, tenantID string, userID string) error {
	if err := s.userRepository.DeleteUser(ctx, tenantID, userID); err != nil {
		return err
	}

	return nil
}
