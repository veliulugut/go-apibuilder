package repository

import (
	"context"
	"go-apibuilder/db/sqlc"
)

var _ UserRepository = (*DBUserRepository)(nil)

type UserRepository interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) error
	DeleteUser(ctx context.Context, id int64) error
	GetUserByID(ctx context.Context, id int64) (sqlc.User, error)
	GetUserByEmail(ctx context.Context, email string) (sqlc.User, error)
	ListUsers(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error)
	UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error)
}

func NewDBUserRepository(querier sqlc.Querier) UserRepository {
	return &DBUserRepository{db: querier}
}

type DBUserRepository struct {
	db sqlc.Querier
}

// CreateUser implements UserRepository.
func (d *DBUserRepository) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) error {
	panic("unimplemented")
}

// DeleteUser implements UserRepository.
func (d *DBUserRepository) DeleteUser(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetUserByEmail implements UserRepository.
func (d *DBUserRepository) GetUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	panic("unimplemented")
}

// GetUserByID implements UserRepository.
func (d *DBUserRepository) GetUserByID(ctx context.Context, id int64) (sqlc.User, error) {
	panic("unimplemented")
}

// ListUsers implements UserRepository.
func (d *DBUserRepository) ListUsers(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserRepository.
func (d *DBUserRepository) UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	panic("unimplemented")
}
