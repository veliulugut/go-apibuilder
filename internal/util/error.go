package util

import "errors"

var (
	ErrInvalidHashPassword     = errors.New("invalid hash password")
	ErrIncompatibleAlgorithm   = errors.New("incompatible algorithm")
	ErrPasswordNotEmpty        = errors.New("password must not be empty")
	ErrHashAndPasswordNotEmpty = errors.New("password and stored hash cannot be empty")
)
