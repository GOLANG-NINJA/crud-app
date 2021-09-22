package service

import (
	"context"
	"time"

	"github.com/GOLANG-NINJA/crud-app/internal/domain"
)

// PasswordHasher provides hashing logic to securely store passwords.
type PasswordHasher interface {
	Hash(password string) (string, error)
}

type UsersRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByCredentials(ctx context.Context, email, password string) (domain.User, error)
}

type Users struct {
	repo   UsersRepository
	hasher PasswordHasher
}

func NewUsers(repo UsersRepository, hasher PasswordHasher) *Users {
	return &Users{
		repo:   repo,
		hasher: hasher,
	}
}

func (u *Users) SignUp(ctx context.Context, inp domain.SignUpInput) error {
	password, err := u.hasher.Hash(inp.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:         inp.Name,
		Email:        inp.Email,
		Password:     password,
		RegisteredAt: time.Now(),
	}

	return u.repo.Create(ctx, user)
}

func (u *Users) SignIn(ctx context.Context, email, password string) (string, error) {
	return "", nil
}
