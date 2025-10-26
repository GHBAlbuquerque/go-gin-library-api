package book

import (
	"context"
	"fmt"
)

var (
	ErrNotFound  = fmt.Errorf("not found")
	ErrDuplicate = fmt.Errorf("duplicate")
)

type Store interface {
	List(ctx context.Context) []Book
	FindById(ctx context.Context, id string) (Book, error)
	Create(ctx context.Context, b Book) (string, error)
	Update(ctx context.Context, b Book) error
	FindByTitle(ctx context.Context, title string) []Book
	FindByAuthor(ctx context.Context, author string) []Book
}
