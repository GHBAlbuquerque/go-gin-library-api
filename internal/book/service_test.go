package book_test

import (
	"context"
	"errors"
	"example/go-gin-library-api/internal/book"
	"testing"
)

// ---------- Mock with function fields ----------

type mockStore struct {
	ListFunc         func(ctx context.Context) ([]book.Book, error)
	FindByIdFunc     func(ctx context.Context, id string) (book.Book, error)
	CreateFunc       func(ctx context.Context, b book.Book) (string, error)
	UpdateFunc       func(ctx context.Context, b book.Book) error
	FindByTitleFunc  func(ctx context.Context, title string) ([]book.Book, error)
	FindByAuthorFunc func(ctx context.Context, author string) ([]book.Book, error)
}

func (s *mockStore) List(ctx context.Context) ([]book.Book, error) {
	if s.ListFunc == nil {
		panic("List called but ListFunc is nil")
	}
	return s.ListFunc(ctx)
}

func (s *mockStore) FindById(ctx context.Context, id string) (book.Book, error) {
	if s.FindByIdFunc == nil {
		panic("FindById called but FindByIdFunc is nil")
	}
	return s.FindByIdFunc(ctx, id)
}

func (s *mockStore) Create(ctx context.Context, b book.Book) (string, error) {
	if s.CreateFunc == nil {
		panic("Create called but CreateFunc is nil")
	}
	return s.CreateFunc(ctx, b)
}

func (s *mockStore) Update(ctx context.Context, b book.Book) error {
	if s.UpdateFunc == nil {
		panic("Update called but UpdateFunc is nil")
	}
	return s.UpdateFunc(ctx, b)
}

func (s *mockStore) FindByTitle(ctx context.Context, title string) ([]book.Book, error) {
	if s.FindByTitleFunc == nil {
		panic("FindByTitle called but FindByTitleFunc is nil")
	}
	return s.FindByTitleFunc(ctx, title)
}

func (s *mockStore) FindByAuthor(ctx context.Context, author string) ([]book.Book, error) {
	if s.FindByAuthorFunc == nil {
		panic("FindByAuthor called but FindByAuthorFunc is nil")
	}
	return s.FindByAuthorFunc(ctx, author)
}

var _ book.Store = (*mockStore)(nil)

// ---------- Tests ----------

func TestService_FindAll(t *testing.T) {
	tests := []struct {
		name string
		str  *mockStore
		pass bool
	}{
		{
			name: "Returns error if findByTitle returns error",
			str: &mockStore{FindByTitleFunc: func(ctx context.Context, title string) ([]book.Book, error) {
				return nil, errors.New("something went wrong")
			},
			},
			pass: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			s := book.NewService(testCase.str)

			_, err := s.FindAll(context.Background(), book.BookFilters{Title: "Tit"})

			if err != nil && testCase.pass {
				t.Fatalf("pass=%t; error=%s", testCase.pass, err.Error())
			}
		})
	}
}
