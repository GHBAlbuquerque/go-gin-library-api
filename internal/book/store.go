package book

import (
	"context"
)

type Store interface {
	List(ctx context.Context) ([]Book, error)
	FindById(ctx context.Context, id string) (Book, error)
	Create(ctx context.Context, b Book) (string, error)
	Update(ctx context.Context, b Book) error
	FindByTitle(ctx context.Context, title string) ([]Book, error)
	FindByAuthor(ctx context.Context, author string) ([]Book, error)
}
