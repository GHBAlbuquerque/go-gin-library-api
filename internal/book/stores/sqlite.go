package stores

import (
	"context"
	"database/sql"
	"example/go-gin-library-api/internal/book"
)

type SQLite struct {
	DB    *sql.DB
	items map[string]book.Book
}

// (m *Memory) is my receiver, which can be used to call these functions
func (s *SQLite) List(ctx context.Context) []book.Book { /*TODO*/ return []book.Book{} }
func (s *SQLite) FindById(ctx context.Context, id string) (book.Book, error) { /*TODO*/
	return book.Book{}, nil
}
func (s *SQLite) Create(ctx context.Context, b book.Book) (string, error) { /*TODO*/ return "ok", nil }

func (s *SQLite) Update(ctx context.Context, b book.Book) error { /*TODO*/
	return nil
}
func (s *SQLite) FindByTitle(ctx context.Context) []book.Book  { /*TODO*/ return []book.Book{} }
func (s *SQLite) FindByAuthor(ctx context.Context) []book.Book { /*TODO*/ return []book.Book{} }
