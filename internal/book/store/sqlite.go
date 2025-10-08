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
func (m *SQLite) List(ctx context.Context) ([]book.Book, error)           { /*TODO*/ }
func (m *SQLite) Get(ctx context.Context, id string) (book.Book, error)   { /*TODO*/ }
func (m *SQLite) Create(ctx context.Context, b book.Book) (string, error) { /*TODO*/ }
func (m *SQLite) Update(ctx context.Context, b book.Book) error           { /*TODO*/ }
func (m *SQLite) FindByTitle(ctx context.Context) ([]book.Book, error)    { /*TODO*/ }
func (m *SQLite) FindByAuthor(ctx context.Context) ([]book.Book, error)   { /*TODO*/ }

/*TODO*/
