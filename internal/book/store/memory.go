package stores

import (
	"context"
	"example/go-gin-library-api/internal/book"
	"sync"
)

type Memory struct {
	mu    sync.RWMutex
	items map[string]book.Book
}

// (m *Memory) is my receiver, which can be used to call these functions
func (m *Memory) List(ctx context.Context) ([]book.Book, error)           { /*TODO*/ }
func (m *Memory) Get(ctx context.Context, id string) (book.Book, error)   { /*TODO*/ }
func (m *Memory) Create(ctx context.Context, b book.Book) (string, error) { /*TODO*/ }
func (m *Memory) Update(ctx context.Context, b book.Book) error           { /*TODO*/ }
func (m *Memory) FindByTitle(ctx context.Context) ([]book.Book, error)    { /*TODO*/ }
func (m *Memory) FindByAuthor(ctx context.Context) ([]book.Book, error)   { /*TODO*/ }
