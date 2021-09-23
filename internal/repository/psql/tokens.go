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
	tx, err := r.db.Begin()
	if err != nil {
		return domain.RefreshToken{}, err
	}
	defer tx.Rollback()

	var t domain.RefreshToken
	err := tx.QueryRow("SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token=$1", token).
		Scan(&t.ID, &t.UserID, &t.Token, &t.ExpiresAt)

	return t, err
}

func (r *Tokens) Update(ctx context.Context, id int, token domain.RefreshToken) error {
	_, err := r.db.Exec("UPDATE refresh_tokens SET token=$1, expires_at=$2 WHERE id=$3", token.Token, token.ExpiresAt, id)

	return err
}
