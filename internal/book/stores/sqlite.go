package stores

import (
	"context"
	"database/sql"
	"example/go-gin-library-api/internal/book"
)

type SQLite struct {
	DB *sql.DB
}

// NewSQLite creates a SQLiteStore
func NewSQLite() (*SQLite, error) {
	s := &SQLite{}

	return s, nil
}

// (s *SQLite) is my receiver, which can be used to call these functions
// List offers reading for all the current stored books. Returns a slice of books.
func (s *SQLite) List(ctx context.Context) []book.Book {
	return []book.Book{}
}

func (s *SQLite) FindById(ctx context.Context, id string) (book.Book, error) { /*TODO*/
	return book.Book{}, nil
}
func (s *SQLite) Create(ctx context.Context, b book.Book) (string, error) { /*TODO*/ return "ok", nil }

func (s *SQLite) Update(ctx context.Context, b book.Book) error { /*TODO*/
	return nil
}
func (s *SQLite) FindByTitle(ctx context.Context, title string) []book.Book { /*TODO*/
	return []book.Book{}
}
func (s *SQLite) FindByAuthor(ctx context.Context, author string) []book.Book { /*TODO*/
	return []book.Book{}
}
