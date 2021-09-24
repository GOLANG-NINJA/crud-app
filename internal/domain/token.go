package domain

import "time"

type RefreshSession struct {
	ID        int64
	UserID    int64
	Token     string
	ExpiresAt time.Time
}
