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
// ---------- FindAll ----------
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

// ---------- GetById ----------
func TestService_GetById_Success(t *testing.T) {
	want := book.Book{ID: "1", Title: "Title", Author: "Author", Quantity: 1}
	storeMock := &mockStore{
		FindByIdFunc: func(ctx context.Context, id string) (book.Book, error) {
			if id != "1" {
				t.Fatalf("Expected ID '1', got %q", id)
			}
			return want, nil
		},
	}

	svc := book.NewService(storeMock)
	got, err := svc.GetById(context.Background(), want.ID)
	if err != nil {
		t.Fatalf("GetById returned error: %q", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetById = %#v, want %#v", got, want)
	}
}

func TestService_GetById_Error(t *testing.T) {
	want := book.Book{ID: "1", Title: "Title", Author: "Author", Quantity: 1}
	storeMock := &mockStore{
		FindByIdFunc: func(ctx context.Context, id string) (book.Book, error) {
			if id != "1" {
				t.Fatalf("Expected ID '1', got %q", id)
			}
			return book.Book{}, book.ErrNotFound
		},
	}

	svc := book.NewService(storeMock)
	_, err := svc.GetById(context.Background(), want.ID)
	if err == nil {
		t.Fatalf("GetById didn't return error: %q", err)
	}
	if !errors.Is(err, book.ErrNotFound) {
		t.Fatalf("GetById = %#v, want %#v", book.ErrNotFound.Error(), err.Error())
	}
}

// ---------- Create ----------
func TestService_Create_Success(t *testing.T) {
	req := book.BookRequest{Title: "Title", Author: "Author", Quantity: 1}
	storeMock := &mockStore{
		CreateFunc: func(ctx context.Context, b book.Book) (string, error) {
			if b.Title != req.Title || b.Author != req.Author || b.Quantity != req.Quantity {
				t.Fatalf("Saved book does not correspond to request: %q", b)
			}
			if b.ID == "" {
				t.Fatal("expected service to assign iD before calling the store")
			}

			return b.ID, nil
		},
	}

	svc := book.NewService(storeMock)
	got, err := svc.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Create returned error: %q", err)
	}
	if len(got) < 36 {
		t.Fatalf("Generated ID is not UUID: %q", got)
	}
}

func TestService_Create_Error(t *testing.T) {
	req := book.BookRequest{Title: "Title", Author: "Author", Quantity: 1}
	storeMock := &mockStore{
		CreateFunc: func(ctx context.Context, b book.Book) (string, error) {
			if b.Title != req.Title || b.Author != req.Author || b.Quantity != req.Quantity {
				t.Fatalf("Saved book does not correspond to request: %q", b)
			}
			if b.ID == "" {
				t.Fatal("expected service to assign iD before calling the store")
			}

			return "", errors.New("Error")
		},
	}

	svc := book.NewService(storeMock)
	_, err := svc.Create(context.Background(), req)
	if err == nil {
		t.Fatalf("Create didn't return error: %q", err)
	}
}

// ---------- Checkout ----------

func TestService_Checkout_BookNotExistsError(t *testing.T) {
	want := book.Book{ID: "1", Title: "Title", Author: "Author", Quantity: 1}
	storeMock := &mockStore{
		FindByIdFunc: func(ctx context.Context, id string) (book.Book, error) {
			if id != "1" {
				t.Fatalf("Expected ID '1', got %q", id)
			}
			return book.Book{}, book.ErrNotFound
		},
		UpdateFunc: func(ctx context.Context, b book.Book) error {
			t.Fatal("Update shouldn't have been called")
			return nil
		},
	}

	svc := book.NewService(storeMock)
	_, err := svc.Checkout(context.Background(), want.ID)
	if err == nil {
		t.Fatalf("Checkout didn't return error: %q", err)
	}
	if !errors.Is(err, book.ErrNotFound) {
		t.Fatalf("Checkout = %#v, want %#v", err.Error(), book.ErrNotFound.Error())
	}
}

func TestService_Checkout_BookUnavailableError(t *testing.T) {
	want := book.Book{ID: "1", Title: "Title", Author: "Author", Quantity: 1}
	storeMock := &mockStore{
		FindByIdFunc: func(ctx context.Context, id string) (book.Book, error) {
			if id != "1" {
				t.Fatalf("Expected ID '1', got %q", id)
			}
			return want, nil
		},
		UpdateFunc: func(ctx context.Context, b book.Book) error {
			if b.ID != want.ID || b.Title != want.Title || b.Author != want.Author {
				t.Fatalf("Updated book does not correspond to request: %q and %q", b, want)
			}
			return book.ErrBookUnavailable
		},
	}

	svc := book.NewService(storeMock)
	_, err := svc.Checkout(context.Background(), want.ID)
	if err == nil {
		t.Fatalf("Checkout didn't return error: %q", err)
	}
	if !errors.Is(err, book.ErrBookUnavailable) {
		t.Fatalf("Checkout = %#v, want %#v", err.Error(), book.ErrBookUnavailable.Error())
	}
}

func TestService_Checkout_BookUpdateError(t *testing.T) {
	want := book.Book{ID: "1", Title: "Title", Author: "Author", Quantity: 1}
	storeMock := &mockStore{
		FindByIdFunc: func(ctx context.Context, id string) (book.Book, error) {
			if id != "1" {
				t.Fatalf("Expected ID '1', got %q", id)
			}
			return want, nil
		},
		UpdateFunc: func(ctx context.Context, b book.Book) error {
			if b.ID != want.ID || b.Title != want.Title || b.Author != want.Author {
				t.Fatalf("Updated book does not correspond to request: %q and %q", b, want)
			}
			return errors.New("Error!")
		},
	}

	svc := book.NewService(storeMock)
	_, err := svc.Checkout(context.Background(), want.ID)
	if err == nil {
		t.Fatalf("Checkout didn't return error: %q", err)
	}
}

func TestService_Checkout_Success(t *testing.T) {
	want := book.Book{ID: "1", Title: "Title", Author: "Author", Quantity: 1}
	updated := false
	storeMock := &mockStore{
		FindByIdFunc: func(ctx context.Context, id string) (book.Book, error) {
			if id != "1" {
				t.Fatalf("Expected ID '1', got %q", id)
			}
			return want, nil
		},
		UpdateFunc: func(ctx context.Context, b book.Book) error {
			if b.ID != want.ID || b.Title != want.Title || b.Author != want.Author {
				t.Fatalf("Updated book does not correspond to request: %q and %q", b, want)
			}
			if b.Quantity != want.Quantity-1 {
				t.Fatalf("expected update with quantity %d, got %d", want.Quantity-1, b.Quantity)
			}
			updated = true
			return nil
		},
	}

	svc := book.NewService(storeMock)
	got, err := svc.Checkout(context.Background(), want.ID)
	if err != nil {
		t.Fatalf("Checkout returned error: %q", err)
	}
	if got.Quantity != want.Quantity-1 {
		t.Fatalf("Checkout did not update book quantity correctly. wanted %q, got %q", want.Quantity-1, got.Quantity)
	}
	if updated == false {
		t.Fatal("Update was not called")
	}
}

// ---------- Return ----------

func TestService_Return_BookNotExistsError(t *testing.T) {

}

func TestService_Return_BookUpdateError(t *testing.T) {

}

func TestService_Return_Success(t *testing.T) {

}
