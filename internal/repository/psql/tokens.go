package psql

import (
	"context"
	"database/sql"

	"github.com/GOLANG-NINJA/crud-app/internal/domain"
)

type Tokens struct {
	db *sql.DB
}

func NewTokens(db *sql.DB) *Tokens {
	return &Tokens{db}
}

func (r *Tokens) Create(ctx context.Context, token domain.RefreshToken) error {
	_, err := r.db.Exec("INSERT INTO refresh_tokens (user_id, token, expires_at) values ($1, $2, $3)",
		token.UserID, token.Token, token.ExpiresAt)

	return err
}

func (r *Tokens) Get(ctx context.Context, token string) (domain.RefreshToken, error) {
	var t domain.RefreshToken
	err := r.db.QueryRow("SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token=$1", token).
		Scan(&t.ID, &t.UserID, &t.Token, &t.ExpiresAt)

	return t, err
}

func (r *Tokens) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec("DELETE FROM refresh_tokens WHERE id=$3", id)

	return err
}
