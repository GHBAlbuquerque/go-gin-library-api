package book

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type BookFilters struct {
	Title, Author string
}

type Service interface {
	FindAll(ctx context.Context, filters BookFilters) ([]Book, error)
	GetById(ctx context.Context, id string) (Book, error)
	Create(ctx context.Context, BookRequest BookRequest) (string, error)
	Checkout(ctx context.Context, id string) (Book, error)
	Return(ctx context.Context, id string) (Book, error)
}

type BookService struct {
	store Store // it's illegal to use pointers to interface types
}

func NewService(store Store) Service {
	s := BookService{
		store: store,
	}

	return &s
}

// FindAll returns a slice of books, filtered or not.
func (s *BookService) FindAll(ctx context.Context, filters BookFilters) ([]Book, error) {

	if title := filters.Title; title != "" {
		books, err := s.store.FindByTitle(ctx, title)
		if err != nil {
			return books, fmt.Errorf("store.FindByTitle: %w", err)
		}

		return books, nil
	}

	if author := filters.Author; author != "" {
		books, err := s.store.FindByAuthor(ctx, author)
		if err != nil {
			return books, fmt.Errorf("store.FindByAuthor: %w", err)
		}

		return books, nil
	}

	out, err := s.store.List(ctx)
	if err != nil {
		return out, fmt.Errorf("store.List: %w", err)
	}

	return out, nil
}

// GetById returns the desired book, found by id.
func (s *BookService) GetById(ctx context.Context, id string) (Book, error) {
	book, err := s.store.FindById(ctx, id)

	if err != nil {
		return Book{}, fmt.Errorf("store.FindById: %w", err)
	}

	return book, nil
}

// Create creates a new book in the store.
func (s *BookService) Create(ctx context.Context, bookRequest BookRequest) (string, error) {
	id := uuid.NewString()
	newBook := Book{id, bookRequest.Title, bookRequest.Author, bookRequest.Quantity}

	out, err := s.store.Create(ctx, newBook)
	if err != nil {
		return "", fmt.Errorf("store.Create: %w", err)
	}

	return out, nil
}

// Checkout retrieves an available book from the library.
func (s *BookService) Checkout(ctx context.Context, id string) (Book, error) {
	book, err := s.store.FindById(ctx, id)

	if err != nil {
		return book, fmt.Errorf("store.FindById: %w", err)
	}

	if book.Quantity == 0 {
		return book, ErrBookUnavailable
	}

	book.Quantity -= 1
	if err := s.store.Update(ctx, book); err != nil {
		return book, fmt.Errorf("store.Update: %w", err)
	}

	return book, nil
}

// Return gives back a book to the library.
func (s *BookService) Return(ctx context.Context, id string) (Book, error) {
	book, err := s.store.FindById(ctx, id)

	if err != nil {
		return book, fmt.Errorf("store.FindById: %w", err)
	}

	book.Quantity += 1
	if err := s.store.Update(ctx, book); err != nil {
		return book, fmt.Errorf("store.Update: %w", err)
	}

	return book, nil
}
