package stores

import (
	"context"
	"example/go-gin-library-api/internal/book"
)

type Store interface {
	List(ctx context.Context) ([]book.Book, error)
	Get(ctx context.Context, id string) (book.Book, error)
	Create(ctx context.Context, b book.Book) (string, error)
	Update(ctx context.Context, b book.Book) error
	FindByTitleAuthor(ctx context.Context)
}
