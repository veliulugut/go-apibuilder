package service

import (
	"context"
	"go-apibuilder/db/sqlc"
	"go-apibuilder/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, params sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByID(ctx context.Context, id int64) (sqlc.User, error)
}

type userServiceImplementation struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImplementation{userRepo: userRepo}
}

func (s *userServiceImplementation) CreateUser(ctx context.Context, params sqlc.CreateUserParams) (sqlc.User, error) {
	return s.userRepo.CreateUser(ctx, params)
}

func (s *userServiceImplementation) GetUserByID(ctx context.Context, id int64) (sqlc.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}
