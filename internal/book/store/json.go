package stores

import (
	"context"
	"example/go-gin-library-api/internal/book"
	"sync"
)

type Json struct {
	mu   sync.Mutex
	path string
	data map[string]book.Book
}

// (m *Json) is my receiver, which can be used to call these functions
func (m *Json) List(ctx context.Context) ([]book.Book, error)           {}
func (m *Json) Get(ctx context.Context, id string) (book.Book, error)   {}
func (m *Json) Create(ctx context.Context, b book.Book) (string, error) {}
func (m *Json) Update(ctx context.Context, b book.Book) error           {}
func (m *Json) FindByTitle(ctx context.Context) ([]book.Book, error)    {}
func (m *Json) FindByAuthor(ctx context.Context) ([]book.Book, error)   {}
