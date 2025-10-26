package stores

import (
	"context"
	"database/sql"
	"errors"
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
func (s *SQLite) List(ctx context.Context) ([]book.Book, error) {
	const q = `SELECT id, title, author, quantity FROM books`
	rows, err := s.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []book.Book
	for rows.Next() {
		var b book.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity); err != nil {
			return nil, err
		}
		out = append(out, b)
	}

	if err := rows.Err(); err != nil { // checks if rows.Next() stops returning true either because you reached EOF or because an iteration error occurred
		return nil, err
	}

	return out, nil
}

func (s *SQLite) FindById(ctx context.Context, id string) (book.Book, error) {
	const q = `SELECT id, title, author, quantity FROM books where id=?`
	row := s.DB.QueryRowContext(ctx, q, id) // QueryRowContext runs a query supposed to bring only one result

	var b book.Book
	if err := row.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return book.Book{}, book.ErrNotFound
		}

		return book.Book{}, err
	}

	return b, nil
}
func (s *SQLite) Create(ctx context.Context, b book.Book) (string, error) { /*TODO*/ return "ok", nil }

func (s *SQLite) Update(ctx context.Context, b book.Book) error { /*TODO*/
	return nil
}
func (s *SQLite) FindByTitle(ctx context.Context, title string) ([]book.Book, error) { /*TODO*/
	return []book.Book{}, nil
}
func (s *SQLite) FindByAuthor(ctx context.Context, author string) ([]book.Book, error) { /*TODO*/
	return []book.Book{}, nil
}
