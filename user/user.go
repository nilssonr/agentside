package user

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type User struct {
	ID             string    `json:"id"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	EmailAddress   string    `json:"emailAddress"`
	Password       string    `json:"-"`
	TenantID       string    `json:"tenantId"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}

type Service interface {
	CreateUser(ctx context.Context, request *User) (*User, error)
	GetUsers(ctx context.Context, tenantID string) ([]*User, error)
	GetUser(ctx context.Context, tenantID, userID string) (*User, error)
	GetUserByEmailAddress(ctx context.Context, tenantID, emailAddress string) (*User, error)
	UpdateUser(ctx context.Context, request *User) (*User, error)
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
		s.logger.Error("failed to create user",
			zap.Error(err))
		return nil, err
	}

	return u, nil
}

// GetUsers implements Service.
func (s *service) GetUsers(ctx context.Context, tenantID string) ([]*User, error) {
	u, err := s.userRepository.GetUsers(ctx, tenantID)
	if err != nil {
		s.logger.Error("failed to get users",
			zap.Error(err))
		return nil, err
	}

	return u, nil
}

// GetUser implements Service.
func (s *service) GetUser(ctx context.Context, tenantID string, userID string) (*User, error) {
	u, err := s.userRepository.GetUser(ctx, tenantID, userID)
	if err != nil {
		s.logger.Error("failed to get user",
			zap.Error(err))
		return nil, err
	}

	return u, nil
}

func (s *service) GetUserByEmailAddress(ctx context.Context, tenantID, emailAddress string) (*User, error) {
	result, err := s.userRepository.GetUserByEmailAddress(ctx, tenantID, emailAddress)
	if err != nil {
		s.logger.Error("failed to get user by email address",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// UpdateUser implements Service.
func (s *service) UpdateUser(ctx context.Context, u *User) (*User, error) {
	u, err := s.userRepository.UpdateUser(ctx, u)
	if err != nil {
		s.logger.Error("failed to update user",
			zap.Error(err))
		return nil, err
	}

	return u, nil
}

// DeleteUser implements Service.
func (s *service) DeleteUser(ctx context.Context, tenantID string, userID string) error {
	if err := s.userRepository.DeleteUser(ctx, tenantID, userID); err != nil {
		s.logger.Error("failed to delete user",
			zap.Error(err))
		return err
	}

	return nil
}
