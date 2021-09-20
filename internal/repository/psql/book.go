package psql

import (
	"context"
	"database/sql"
	"web-app/internal/domain"
)

type Books struct {
	db *sql.DB
}

func NewBooks(db *sql.DB) *Books {
	return &Books{db}
}

func (b *Books) Create(ctx context.Context, book domain.Book) error {
	_, err := b.db.Exec("INSERT INTO books (title, author, publish_date, rating) values ($1, $2, $3, $4)",
		book.Title, book.Author, book.PublishDate, book.Rating)

	return err
}

func (b *Books) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	var book domain.Book
	err := b.db.QueryRow("SELECT id, title, author, publish_date, rating FROM books WHERE id=$1", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Rating)

	return book, err
}

func (b *Books) GetAll(ctx context.Context) ([]domain.Book, error) {
	return nil, nil
}

func (b *Books) Delete(ctx context.Context, id int64) error {
	return nil
}
