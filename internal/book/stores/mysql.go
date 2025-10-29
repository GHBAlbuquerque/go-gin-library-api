package stores

import (
	"context"
	"database/sql"
	"errors"
	"example/go-gin-library-api/internal/book"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

// NewMySQL creates a SQLiteStore
func NewMySQL(dsn string) (*MySQL, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	s := &MySQL{
		DB: db,
	}

	return s, nil
}

// (s *SQLite) is my receiver, which can be used to call these functions
// List offers reading for all the current stored books. Returns a slice of books.
func (s *MySQL) List(ctx context.Context) ([]book.Book, error) {
	const q = `SELECT id, title, author, quantity FROM Books
				ORDER BY author;`
	rows, err := s.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out = []book.Book{}
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

func (s *MySQL) FindById(ctx context.Context, id string) (book.Book, error) {
	const q = `SELECT id, title, author, quantity FROM Books where id=?;`
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

func (s *MySQL) Create(ctx context.Context, b book.Book) (string, error) {
	const q = `INSERT INTO Books (id, title, author, quantity)
				VALUES (?, ?, ?, ?);`

	_, err := s.DB.ExecContext(ctx, q, b.ID, b.Title, b.Author, b.Quantity)

	if err != nil {
		return "", err
	}

	return b.ID, nil
}

func (s *MySQL) Update(ctx context.Context, b book.Book) error {
	const q = `UPDATE Books
				SET title=?, author=?, quantity=?
				WHERE id=?;`

	if _, err := s.DB.ExecContext(ctx, q, b.Title, b.Author, b.Quantity, b.ID); err != nil {
		return err
	}

	return nil
}
func (s *MySQL) FindByTitle(ctx context.Context, title string) ([]book.Book, error) { /*TODO*/
	const q = `SELECT * from Books
				WHERE title LIKE ?
				ORDER BY author;`

	like := "%" + title + "%"
	rows, err := s.DB.QueryContext(ctx, q, like)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []book.Book{}
	for rows.Next() {
		var b book.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity); err != nil {
			return nil, err
		}
		out = append(out, b)
	}

	return out, nil
}

func (s *MySQL) FindByAuthor(ctx context.Context, author string) ([]book.Book, error) { /*TODO*/
	const q = `SELECT * from Books
				WHERE author LIKE ?
				ORDER BY author;`

	like := "%" + author + "%"
	rows, err := s.DB.QueryContext(ctx, q, like)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []book.Book{}
	for rows.Next() {
		var b book.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity); err != nil {
			return nil, err
		}
		out = append(out, b)
	}

	return out, nil
}
