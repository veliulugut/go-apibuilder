package repository

import (
	"context"
	"go-apibuilder/db/sqlc"
	"go-apibuilder/internal/util"

	"github.com/jackc/pgx/v5/pgtype"
)

var _ UserRepository = (*DBUserRepository)(nil)

type UserRepository interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
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
func (d *DBUserRepository) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	hashedPassword, err := util.HashPassword(arg.HashedPassword)
	if err != nil {
		return sqlc.User{}, err
	}

	arg.HashedPassword = hashedPassword

	return d.db.CreateUser(ctx, arg)
}

// DeleteUser implements UserRepository.
func (d *DBUserRepository) DeleteUser(ctx context.Context, id int64) error {
	return d.db.DeleteUser(ctx, id)
}

// GetUserByEmail implements UserRepository.
func (d *DBUserRepository) GetUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	return d.db.GetUserByEmail(ctx, email)
}

// GetUserByID implements UserRepository.
func (d *DBUserRepository) GetUserByID(ctx context.Context, id int64) (sqlc.User, error) {
	return d.db.GetUserByID(ctx, id)
}

// ListUsers implements UserRepository.
func (d *DBUserRepository) ListUsers(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error) {
	return d.db.ListUsers(ctx, arg)
}

// UpdateUser implements UserRepository.
func (d *DBUserRepository) UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	if arg.HashedPassword.Valid && arg.HashedPassword.String != "" {
		newHashedPassword, err := util.HashPassword(arg.HashedPassword.String)
		if err != nil {
			return sqlc.User{}, err
		}

		arg.HashedPassword = pgtype.Text{String: newHashedPassword, Valid: true}
	} else {
		if !arg.HashedPassword.Valid {
		} else if arg.HashedPassword.String == "" {
			arg.HashedPassword.Valid = false
		}
	}

	return d.db.UpdateUser(ctx, arg)
}
