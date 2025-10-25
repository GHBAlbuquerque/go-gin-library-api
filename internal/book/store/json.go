package stores

import (
	"context"
	"encoding/json"
	"errors"
	"example/go-gin-library-api/internal/book"
	"os"
	"path/filepath"
	"sync"
)

type JSON struct {
	mu   sync.Mutex // Allows one goroutine at a time to access a critical section.
	path string
	data map[string]book.Book
}

// (m *JSON) is my receiver, which can be used to call these functions

func NewJson(path string, seed []book.Book) (*JSON, error) {
	j := &JSON{
		path: path,
		data: map[string]book.Book{},
	}

	// ensure dir on desired path exists
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, err
		}
	}

	// load the file from dir
	_, err := os.Stat(path)

	// use seed to create the file if it doesn't exist
	if errors.Is(err, os.ErrNotExist) {
		for _, b := range seed {
			j.data[b.ID] = b
		}

		if err := j.persist(); err != nil {
			return nil, err
		}

		return j, nil
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(bytes) > 0 {
		err := json.Unmarshal(bytes, &j.data)
		if err != nil {
			return nil, err
		}
	}

	return j, nil
}

func (m *JSON) List(ctx context.Context) ([]book.Book, error)           { /*TODO*/ }
func (m *JSON) Get(ctx context.Context, id string) (book.Book, error)   { /*TODO*/ }
func (m *JSON) Create(ctx context.Context, b book.Book) (string, error) { /*TODO*/ }
func (m *JSON) Update(ctx context.Context, b book.Book) error           { /*TODO*/ }
func (m *JSON) FindByTitle(ctx context.Context) ([]book.Book, error)    { /*TODO*/ }
func (m *JSON) FindByAuthor(ctx context.Context) ([]book.Book, error)   { /*TODO*/ }
