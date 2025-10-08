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
	List(ctx context.Context) ([]Book, error)
	Get(ctx context.Context, id string) (Book, error)
	Create(ctx context.Context, b Book) (string, error)
	Update(ctx context.Context, b Book) error
	FindByTitleAuthor(ctx context.Context)
}
