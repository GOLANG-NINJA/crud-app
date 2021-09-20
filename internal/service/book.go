package service

import (
	"context"

	"github.com/GOLANG-NINJA/crud-app/internal/domain"
)

type BooksRepository interface {
	Create(ctx context.Context, book domain.Book) error
	GetByID(ctx context.Context, id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
}

type Books struct {
	repo BooksRepository
}

func NewBooks(repo BooksRepository) *Books {
	return &Books{
		repo: repo,
	}
}

func (b *Books) Create(ctx context.Context, book domain.Book) error {
	return b.repo.Create(ctx, book)
}

func (b *Books) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	return b.repo.GetByID(ctx, id)
}

func (b *Books) GetAll(ctx context.Context) ([]domain.Book, error) {
	return b.repo.GetAll(ctx)
}

func (b *Books) Delete(ctx context.Context, id int64) error {
	return b.repo.Delete(ctx, id)
}
