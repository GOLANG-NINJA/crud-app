package service

import (
	"context"
	"time"

	"github.com/GOLANG-NINJA/crud-app/internal/domain"
)

type BooksRepository interface {
	Create(ctx context.Context, book domain.Book) error
	GetByID(ctx context.Context, id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp domain.UpdateBookInput) error
}

type Books struct {
	repo BooksRepository
}

func NewBooks(repo BooksRepository) *Books {
	return &Books{
		repo: repo,
	}
}

func (s *Books) Create(ctx context.Context, book domain.Book) error {
	if book.PublishDate.IsZero() {
		book.PublishDate = time.Now()
	}

	return s.repo.Create(ctx, book)
}

func (s *Books) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Books) GetAll(ctx context.Context) ([]domain.Book, error) {
	return s.repo.GetAll(ctx)
}

func (s *Books) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *Books) Update(ctx context.Context, id int64, inp domain.UpdateBookInput) error {
	return s.repo.Update(ctx, id, inp)
}
