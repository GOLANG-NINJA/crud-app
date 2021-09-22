package service

import (
	"context"

	"github.com/GOLANG-NINJA/crud-app/internal/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByCredentials(ctx context.Context, email, password string) (domain.User, error)
}

type Users struct {
	repo UsersRepository
}

func NewUsers(repo UsersRepository) *Users {
	return &Users{
		repo: repo,
	}
}

func SignUp(ctx context.Context, inp domain.SignUpInput) error {
	return nil
}

func SignIn(ctx context.Context, email, password string) (string, error) {
	return "", nil
}
