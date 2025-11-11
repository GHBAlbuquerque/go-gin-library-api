package book_test

import (
	"context"
	"errors"
	"example/go-gin-library-api/internal/book"
	"reflect"
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

func TestService_FindAll_TitleSuccess(t *testing.T) {
	want := []book.Book{{ID: "1", Title: "Title", Author: "Author", Quantity: 1}}
	storeMock := &mockStore{
		FindByTitleFunc: func(ctx context.Context, title string) ([]book.Book, error) {
			if title != "Title" {
				t.Fatalf("Expected title Title', got %q", title)
			}
			return want, nil
		},
		// other methods added to avoid calling by mistake.
		ListFunc: func(context.Context) ([]book.Book, error) {
			t.Fatal("List must not run when title filter provided")
			return nil, nil
		},
		FindByAuthorFunc: func(context.Context, string) ([]book.Book, error) {
			t.Fatal("FindByAuthor must not run when title filter provided")
			return nil, nil
		},
	}

	svc := book.NewService(storeMock)
	got, err := svc.FindAll(context.Background(), book.BookFilters{Title: "Title"})
	if err != nil {
		t.Fatalf("FindAll returned error: %q", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("FindAll = %#v, want %#v", got, want)
	}
}

func TestService_FindAll_TitleError(t *testing.T) {
	storeMock := &mockStore{
		FindByTitleFunc: func(ctx context.Context, title string) ([]book.Book, error) {
			if title != "Title" {
				t.Fatalf("Expected title 'Title', got %q", title)
			}
			return nil, errors.New("Error")
		},
		// other methods added to avoid calling by mistake.
		ListFunc: func(context.Context) ([]book.Book, error) {
			t.Fatal("List must not run when title filter provided")
			return nil, nil
		},
		FindByAuthorFunc: func(context.Context, string) ([]book.Book, error) {
			t.Fatal("FindByAuthor must not run when title filter provided")
			return nil, nil
		},
	}

	svc := book.NewService(storeMock)
	_, err := svc.FindAll(context.Background(), book.BookFilters{Title: "Title"})
	if err == nil {
		t.Fatalf("FindAll didn't return error: %q", err)
	}
}

func TestService_FindAll_AuthorSuccess(t *testing.T) {
	want := []book.Book{{ID: "1", Title: "Title", Author: "Author", Quantity: 1}}
	storeMock := &mockStore{
		FindByAuthorFunc: func(ctx context.Context, author string) ([]book.Book, error) {
			if author != "Author" {
				t.Fatalf("Expected author 'Author', got %q", author)
			}
			return want, nil
		},
		// other methods added to avoid calling by mistake.
		ListFunc: func(context.Context) ([]book.Book, error) {
			t.Fatal("List must not run when author filter provided")
			return nil, nil
		},
		FindByTitleFunc: func(context.Context, string) ([]book.Book, error) {
			t.Fatal("FindByTitle must not run when author filter provided")
			return nil, nil
		},
	}

	svc := book.NewService(storeMock)
	got, err := svc.FindAll(context.Background(), book.BookFilters{Author: "Author"})
	if err != nil {
		t.Fatalf("FindAll returned error: %q", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("FindAll = %#v, want %#v", got, want)
	}
}

func TestService_FindAll_AuthorError(t *testing.T) {
	storeMock := &mockStore{
		FindByAuthorFunc: func(ctx context.Context, author string) ([]book.Book, error) {
			if author != "Author" {
				t.Fatalf("Expected author 'Author', got %q", author)
			}
			return nil, errors.New("Error!")
		},
		// other methods added to avoid calling by mistake.
		ListFunc: func(context.Context) ([]book.Book, error) {
			t.Fatal("List must not run when author filter provided")
			return nil, nil
		},
		FindByTitleFunc: func(context.Context, string) ([]book.Book, error) {
			t.Fatal("FindByTitle must not run when author filter provided")
			return nil, nil
		},
	}

	svc := book.NewService(storeMock)
	_, err := svc.FindAll(context.Background(), book.BookFilters{Author: "Author"})
	if err == nil {
		t.Fatalf("FindAll didn't return error: %q", err)
	}
}

func TestService_FindAll_Success(t *testing.T) {
	want := []book.Book{{ID: "1", Title: "Title", Author: "Author", Quantity: 1}}
	storeMock := &mockStore{
		ListFunc: func(context.Context) ([]book.Book, error) {
			return want, nil
		},
		FindByTitleFunc: func(context.Context, string) ([]book.Book, error) {
			t.Fatal("FindByTitle must not run when author filter provided")
			return nil, nil
		},
		FindByAuthorFunc: func(context.Context, string) ([]book.Book, error) {
			t.Fatal("FindByAuthor must not run when title filter provided")
			return nil, nil
		},
	}

	svc := book.NewService(storeMock)
	got, err := svc.FindAll(context.Background(), book.BookFilters{})
	if err != nil {
		t.Fatalf("FindAll returned error: %q", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("FindAll = %#v, want %#v", got, want)
	}
}

func TestService_FindAll_Error(t *testing.T) {
	storeMock := &mockStore{
		ListFunc: func(context.Context) ([]book.Book, error) {
			return nil, errors.New("Error")
		},
		FindByTitleFunc: func(context.Context, string) ([]book.Book, error) {
			t.Fatal("FindByTitle must not run when author filter provided")
			return nil, nil
		},
		FindByAuthorFunc: func(context.Context, string) ([]book.Book, error) {
			t.Fatal("FindByAuthor must not run when title filter provided")
			return nil, nil
		},
	}

	svc := book.NewService(storeMock)
	_, err := svc.FindAll(context.Background(), book.BookFilters{})
	if err == nil {
		t.Fatalf("FindAll didn't return error: %q", err)
	}
}
