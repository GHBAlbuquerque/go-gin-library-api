package stores

import (
	"context"
	"encoding/json"
	"errors"
	"example/go-gin-library-api/internal/book"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type JSON struct {
	mu   sync.Mutex // Allows one go routine at a time to access a critical section.
	path string
	data map[string]book.Book
}

// (j *JSON) is my receiver, which can be used to call these functions
// NewJSON creates a JSONstore
func NewJSON(path string, seed []book.Book) (*JSON, error) {
	j := &JSON{
		path: path,
		data: map[string]book.Book{},
	}

	if err := ensureDir(path); err != nil {
		return nil, fmt.Errorf("ensureDir: %w", err)
	}

	created, err := ensureFileWithSeed(path, seed, j)
	if err != nil {
		return nil, fmt.Errorf("ensureFileWithSeed: %w", err)
	}

	if !created {
		if err := loadFromFile(path, j); err != nil {
			return nil, fmt.Errorf("loadFromFile: %w", err)
		}
	}

	return j, nil
}

// ensureDir ensures the directory exists on the desired path
func ensureDir(path string) error {
	// ensure dir on desired path exists
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	return nil
}

// ensureFileWithSeed creates the file with the seed if necessary, and returns a bool (created) to confirm creation
func ensureFileWithSeed(path string, seed []book.Book, j *JSON) (bool, error) {
	// load the file from dir
	_, err := os.Stat(path)

	// use seed to create the file if it doesn't exist
	if errors.Is(err, os.ErrNotExist) {
		for _, b := range seed {
			j.data[b.ID] = b
		}

		if err := j.persist(); err != nil {
			return false, err
		}

		return true, nil
	}

	if err != nil {
		return false, err
	}

	// no file created, data is not yest inside j.data
	return false, nil
}

// persist writes the file atomically if it doesn't exist on the path
func (j *JSON) persist() error {
	tmp := j.path + ".tmp"

	bytes, err := json.MarshalIndent(j.data, "", " ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(tmp, bytes, 0o644); err != nil {
		return err
	}

	return os.Rename(tmp, j.path)
}

// loadFromFile reads the Json file from path and unmarshalls the content to the data map on JSON struct
func loadFromFile(path string, j *JSON) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if len(bytes) > 0 {
		err := json.Unmarshal(bytes, &j.data)
		if err != nil {
			return err
		}
	}

	return nil
}

// List offers thread-safe reading for all the current in-memory stored books. Returns a slice of books.
func (j *JSON) List(ctx context.Context) []book.Book {
	j.mu.Lock()
	defer j.mu.Unlock()

	out := make([]book.Book, 0, len(j.data))
	for _, b := range j.data {
		out = append(out, b)
	}

	return out
}

func (j *JSON) Get(ctx context.Context, id string) (book.Book, error) {
	j.mu.Lock()
	defer j.mu.Unlock()

	b, exists := j.data[id]

	if !exists {
		return book.Book{}, book.ErrNotFound
	}

	return b, nil
}

func (j *JSON) Create(ctx context.Context, b book.Book) (string, error) {
	j.mu.Lock()
	defer j.mu.Unlock()

	if _, exists := j.data[b.ID]; exists {
		return b.ID, book.ErrDuplicate
	}

	for _, current := range j.data {
		if strings.EqualFold(current.Title, b.Title) && strings.EqualFold(current.Author, b.Author) {
			return b.ID, book.ErrDuplicate
		}
	}

	j.data[b.ID] = b
	return b.ID, j.persist()
}

func (j *JSON) Update(ctx context.Context, b book.Book) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	_, exists := j.data[b.ID]
	if !exists {
		return book.ErrNotFound
	}

	j.data[b.ID] = b
	return j.persist()
}

func (j *JSON) FindByTitle(ctx context.Context, title string) []book.Book {
	j.mu.Lock()
	defer j.mu.Unlock()

	var out = make([]book.Book, 0, len(j.data))

	for _, current := range j.data {
		if strings.EqualFold(current.Title, title) {
			out = append(out, current)
		}
	}

	return out
}

func (j *JSON) FindByAuthor(ctx context.Context, author string) []book.Book {
	j.mu.Lock()
	defer j.mu.Unlock()

	var out = make([]book.Book, 0, len(j.data))

	for _, current := range j.data {
		if strings.EqualFold(current.Author, author) {
			out = append(out, current)
		}
	}

	return out
}
