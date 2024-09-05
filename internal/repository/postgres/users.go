package postgres

import (
	"context"
	"time"

	"github.com/nilssonr/agentside/internal/user"
)

type UserRepository struct {
	db *Queries
}

func NewUserRepository(db *Queries) user.Repository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser implements user.Repository.
func (ur *UserRepository) InsertUser(ctx context.Context, u *user.User) (*user.User, error) {
	arg := CreateUserParams{
		FirstName:      u.Firstname,
		LastName:       u.Lastname,
		EmailAddress:   u.EmailAddress,
		TenantID:       u.TenantID,
		LastModifiedAt: mustCreateTime(u.LastModifiedAt),
		LastModifiedBy: u.LastModifiedBy,
	}
	row, err := ur.db.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result user.User
	result.ID = row.ID
	result.Firstname = row.FirstName
	result.Lastname = row.LastName
	result.EmailAddress = row.EmailAddress
	result.TenantID = row.TenantID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// GetUsers implements user.Repository.
func (ur *UserRepository) GetUsers(ctx context.Context, tenantID string) ([]*user.User, error) {
	rows, err := ur.db.GetUsers(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	result := []*user.User{}
	for _, v := range rows {
		result = append(result, &user.User{
			ID:             v.ID,
			Firstname:      v.FirstName,
			Lastname:       v.LastName,
			EmailAddress:   v.EmailAddress,
			TenantID:       v.TenantID,
			LastModifiedAt: v.LastModifiedAt.Time,
			LastModifiedBy: v.LastModifiedBy,
		})
	}

	return result, nil
}

// GetUser implements user.Repository.
func (ur *UserRepository) GetUser(ctx context.Context, tenantID string, userID string) (*user.User, error) {
	arg := GetUserParams{
		TenantID: tenantID,
		ID:       userID,
	}
	row, err := ur.db.GetUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result user.User
	result.ID = row.ID
	result.Firstname = row.FirstName
	result.Lastname = row.LastName
	result.EmailAddress = row.EmailAddress
	result.TenantID = row.TenantID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// GetUser implements user.Repository.
func (ur *UserRepository) GetUserByEmailAddress(ctx context.Context, emailAddress string) (*user.User, error) {
	row, err := ur.db.GetUserByEmailAddress(ctx, emailAddress)
	if err != nil {
		return nil, err
	}

	var result user.User
	result.ID = row.ID
	result.Firstname = row.FirstName
	result.Lastname = row.LastName
	result.EmailAddress = row.EmailAddress
	result.TenantID = row.TenantID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// UpdateUser implements user.Repository.
func (ur *UserRepository) UpdateUser(ctx context.Context, u *user.User) (*user.User, error) {
	arg := UpdateUserParams{
		ID:             u.ID,
		FirstName:      u.Firstname,
		LastName:       u.Lastname,
		LastModifiedAt: mustCreateTime(u.LastModifiedAt),
		LastModifiedBy: u.LastModifiedBy,
		TenantID:       u.TenantID,
	}
	row, err := ur.db.UpdateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result user.User
	result.ID = row.ID
	result.Firstname = row.FirstName
	result.Lastname = row.LastName
	result.EmailAddress = row.EmailAddress
	result.TenantID = row.TenantID
	result.LastModifiedAt = row.LastModifiedAt.Time
	result.LastModifiedBy = row.LastModifiedBy

	return &result, nil
}

// DeleteUser implements user.Repository.
func (ur *UserRepository) DeleteUser(ctx context.Context, tenantID string, userID string) error {
	arg := DeleteUserParams{
		TenantID:  tenantID,
		ID:        userID,
		DeletedAt: mustCreateTime(time.Now()),
	}
	if err := ur.db.DeleteUser(ctx, arg); err != nil {
		return err
	}

	return nil
}
