package stores

import (
	"context"
	"example/go-gin-library-api/internal/book"
	"strings"
	"sync"
)

// Fastest way to get going, stores items on a map. Thread-safe with a RWMutex.
type Memory struct {
	mu    sync.RWMutex         // embedding a value of type sync.RWMutex, not a pointer. That type is ready to use in its zero value.
	items map[string]book.Book // key is the book id
}

// NewMemory creates a MemoryStore
func NewMemory(seed []book.Book) *Memory {
	m := &Memory{items: make(map[string]book.Book, len(seed))}

	for _, b := range seed {
		m.items[b.ID] = b
	}

	return m
}

// (m *Memory) is my receiver, which can be used to call these functions

// List offers thread-safe reading for all the current in-memory stored books. Returns a slice of books.
func (m *Memory) List(ctx context.Context) []book.Book {
	m.mu.RLock()         //locks for writes, but keep reads going
	defer m.mu.RUnlock() // defer runs the code when the concurrent function return (no matter the result)

	out := make([]book.Book, 0, len(m.items))
	for _, b := range m.items {
		out = append(out, b)
	}

	return out
}

// Get offers thread-safe read of a book by its id.
func (m *Memory) Get(ctx context.Context, id string) (book.Book, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	b, ok := m.items[id]
	if !ok {
		return book.Book{}, book.ErrNotFound
	}

	return b, nil
}

// Create offers thread-safe writing in memory.
func (m *Memory) Create(ctx context.Context, b book.Book) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.items[b.ID]; exists {
		return b.ID, book.ErrDuplicate
	}

	for _, current := range m.items {
		if strings.EqualFold(current.Title, b.Title) && strings.EqualFold(current.Author, b.Author) {
			return b.ID, book.ErrDuplicate
		}
	}

	m.items[b.ID] = b
	return b.ID, nil
}

// Update offers thread-safe writing in memory for an existing book (found by ID).
func (m *Memory) Update(ctx context.Context, b book.Book) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.items[b.ID]; !exists {
		return book.ErrNotFound
	}

	m.items[b.ID] = b
	return nil
}

// FindByTitle returns a slice of books found by a title
func (m *Memory) FindByTitle(ctx context.Context, title string) []book.Book {
	m.mu.RLock()
	defer m.mu.RUnlock()

	out := make([]book.Book, 0)

	for _, b := range m.items {
		if strings.EqualFold(b.Title, title) {
			out = append(out, b)
		}
	}

	return out
}

// FindByAuthor returns a slice of books found by an author
func (m *Memory) FindByAuthor(ctx context.Context, author string) []book.Book {
	m.mu.RLock()
	defer m.mu.RUnlock()

	out := make([]book.Book, 0)

	for _, b := range m.items {
		if strings.EqualFold(b.Author, author) {
			out = append(out, b)
		}
	}

	return out
}
