package domain

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrEmailAlreadyTaken = errors.New("email already taken")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken      = errors.New("invalid token")
)
