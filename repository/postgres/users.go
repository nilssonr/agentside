package postgres

import (
	"context"
	"time"

	"github.com/nilssonr/agentside/repository/postgres/sqlc"
	"github.com/nilssonr/agentside/user"
)

type UserRepository struct {
	DB *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) user.Repository {
	return &UserRepository{
		DB: db,
	}
}

// CreateUser implements user.Repository.
func (ur *UserRepository) InsertUser(ctx context.Context, u *user.User) (*user.User, error) {
	arg := sqlc.InsertUserParams{
		FirstName:      u.Firstname,
		LastName:       u.Lastname,
		EmailAddress:   u.EmailAddress,
		TenantID:       u.TenantID,
		LastModifiedAt: mustCreateTime(u.LastModifiedAt),
		LastModifiedBy: u.LastModifiedBy,
	}

	row, err := ur.DB.InsertUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := &user.User{
		ID:             row.ID,
		Firstname:      row.FirstName,
		Lastname:       row.LastName,
		EmailAddress:   row.EmailAddress,
		TenantID:       row.TenantID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// GetUsers implements user.Repository.
func (ur *UserRepository) GetUsers(ctx context.Context, tenantID string) ([]*user.User, error) {
	rows, err := ur.DB.GetUsers(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	result := make([]*user.User, 0, len(rows))
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
	arg := sqlc.GetUserParams{
		TenantID: tenantID,
		ID:       userID,
	}

	row, err := ur.DB.GetUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := &user.User{
		ID:             row.ID,
		Firstname:      row.FirstName,
		Lastname:       row.LastName,
		EmailAddress:   row.EmailAddress,
		TenantID:       row.TenantID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// GetUser implements user.Repository.
func (ur *UserRepository) GetUserByEmailAddress(ctx context.Context, emailAddress string) (*user.User, error) {
	row, err := ur.DB.GetUserByEmailAddress(ctx, emailAddress)
	if err != nil {
		return nil, err
	}

	result := &user.User{
		ID:             row.ID,
		Firstname:      row.FirstName,
		Lastname:       row.LastName,
		EmailAddress:   row.EmailAddress,
		TenantID:       row.TenantID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// UpdateUser implements user.Repository.
func (ur *UserRepository) UpdateUser(ctx context.Context, u *user.User) (*user.User, error) {
	arg := sqlc.UpdateUserParams{
		ID:             u.ID,
		FirstName:      u.Firstname,
		LastName:       u.Lastname,
		LastModifiedAt: mustCreateTime(u.LastModifiedAt),
		LastModifiedBy: u.LastModifiedBy,
		TenantID:       u.TenantID,
	}

	row, err := ur.DB.UpdateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := &user.User{
		ID:             row.ID,
		Firstname:      row.FirstName,
		Lastname:       row.LastName,
		EmailAddress:   row.EmailAddress,
		TenantID:       row.TenantID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// DeleteUser implements user.Repository.
func (ur *UserRepository) DeleteUser(ctx context.Context, tenantID string, userID string) error {
	arg := sqlc.DeleteUserParams{
		TenantID:  tenantID,
		ID:        userID,
		DeletedAt: mustCreateTime(time.Now()),
	}

	if err := ur.DB.DeleteUser(ctx, arg); err != nil {
		return err
	}

	return nil
}
