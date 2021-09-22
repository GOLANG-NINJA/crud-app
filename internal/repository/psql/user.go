package psql

import (
	"context"
	"database/sql"

	"github.com/GOLANG-NINJA/crud-app/internal/domain"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (u *Users) Create(ctx context.Context, user domain.User) error {
	_, err := u.db.Exec("INSERT INTO users (name, email, password, registered_at) values ($1, $2, $3, $4)",
		user.Name, user.Email, user.Password, user.RegisteredAt)

	return err
}

func (u *Users) GetByCredentials(ctx context.Context, email, password string) (domain.User, error) {
	return domain.User{}, nil
}
