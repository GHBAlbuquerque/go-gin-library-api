package book_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"

	"example/go-gin-library-api/internal/book"
	"testing"

	"github.com/gin-gonic/gin"
)

// ---------- Mock with function fields ----------

type mockSvc struct {
	FindAllFunc  func(ctx context.Context, filters book.BookFilters) ([]book.Book, error)
	GetByIdFunc  func(ctx context.Context, id string) (book.Book, error)
	CreateFunc   func(ctx context.Context, BookRequest book.BookRequest) (string, error)
	CheckoutFunc func(ctx context.Context, id string) (book.Book, error)
	ReturnFunc   func(ctx context.Context, id string) (book.Book, error)
}

func (m *mockSvc) FindAll(ctx context.Context, filters book.BookFilters) ([]book.Book, error) {
	if m.FindAllFunc == nil {
		panic("FindAll called but FindAllFunc is nil")
	}
	return m.FindAllFunc(ctx, filters)
}
func (m *mockSvc) GetById(ctx context.Context, id string) (book.Book, error) {
	if m.GetByIdFunc == nil {
		panic("GetById called but GetByIdFunc is nil")
	}
	return m.GetByIdFunc(ctx, id)
}
func (m *mockSvc) Create(ctx context.Context, req book.BookRequest) (string, error) {
	if m.CreateFunc == nil {
		panic("Create called but CreateFunc is nil")
	}
	return m.CreateFunc(ctx, req)
}
func (m *mockSvc) Checkout(ctx context.Context, id string) (book.Book, error) {
	if m.CheckoutFunc == nil {
		panic("Checkout called but CheckoutFunc is nil")
	}
	return m.CheckoutFunc(ctx, id)
}
func (m *mockSvc) Return(ctx context.Context, id string) (book.Book, error) {
	if m.ReturnFunc == nil {
		panic("Return called but ReturnFunc is nil")
	}
	return m.ReturnFunc(ctx, id)
}

var _ book.Service = (*mockSvc)(nil)

// ---------- Tests ----------

func TestHandler_FindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		svc          *mockSvc
		path         string
		wantedStatus int
	}{
		{
			name:         "400 when both title and author are provided",
			svc:          &mockSvc{ /* should not be called; keep nil to panic if it is */ },
			path:         "/books?title=aaa&author=bbb",
			wantedStatus: http.StatusBadRequest,
		},
		{
			name: "500 when service returns error",
			svc: &mockSvc{FindAllFunc: func(ctx context.Context, filters book.BookFilters) ([]book.Book, error) {
				return nil, errors.New("something went wrong")
			}},
			path:         "/books",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "200 on success (no filters)",
			svc: &mockSvc{FindAllFunc: func(ctx context.Context, filters book.BookFilters) ([]book.Book, error) {
				return []book.Book{{ID: "1", Title: "Title", Author: "Author", Quantity: 2}}, nil
			}},
			path:         "/books",
			wantedStatus: http.StatusOK,
		},
		{
			name: "200 on success (title filter only)",
			svc: &mockSvc{
				FindAllFunc: func(ctx context.Context, f book.BookFilters) ([]book.Book, error) {
					return []book.Book{{ID: "2", Title: "Title", Author: "Author", Quantity: 2}}, nil
				},
			},
			path:         "/books?title=Tit",
			wantedStatus: http.StatusOK,
		},
		{
			name: "200 on success (author filter only)",
			svc: &mockSvc{
				FindAllFunc: func(ctx context.Context, f book.BookFilters) ([]book.Book, error) {
					return []book.Book{{ID: "3", Title: "Title", Author: "Author", Quantity: 2}}, nil
				},
			},
			path:         "/books?author=Auth",
			wantedStatus: http.StatusOK,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			h := book.NewHandler(testCase.svc)
			r := gin.New()
			r.GET("/books", h.FindAll)

			req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			if rec.Code != testCase.wantedStatus {
				t.Fatalf("status=%d; wanted=%d body=%s", rec.Code, testCase.wantedStatus, rec.Body.String())
			}
		})
	}
}

func TestHandler_GetById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		svc          *mockSvc
		path         string
		wantedStatus int
	}{
		{
			name: "404 when book is not found",
			svc: &mockSvc{GetByIdFunc: func(ctx context.Context, id string) (book.Book, error) {
				return book.Book{}, errors.New("not found")
			}},
			path:         "/books/1",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "200 when book is found",
			svc: &mockSvc{GetByIdFunc: func(ctx context.Context, id string) (book.Book, error) {
				return book.Book{ID: "1", Title: "Title", Author: "Author", Quantity: 2}, nil
			}},
			path:         "/books/1",
			wantedStatus: http.StatusOK,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			h := book.NewHandler(testCase.svc)
			r := gin.New()
			r.GET("/books/1", h.GetById)

			req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			if rec.Code != testCase.wantedStatus {
				t.Fatalf("status=%d; wanted=%d body=%s", rec.Code, testCase.wantedStatus, rec.Body.String())
			}
		})
	}

}

/*
func TestHander_Create(t *testing.T) {
	tests := []struct {
		name         string
		svc          *mockSvc
		path         string
		wantedStatus int
		body         *bytes.Reader
	}{
		{
			name:         "400 when JSON can't be binded",
			svc:          &mockSvc{  },
			path:         "/books",
			wantedStatus: http.StatusBadRequest,
			body:         bytes.NewReader([]byte(`{"title":`)),
		},
		{
			name: "204 when book there is a conflict",
			svc: &mockSvc{CreateFunc: func(ctx context.Context, BookRequest book.BookRequest) (string, error) {
				//TODO
			}},
			path:         "/books",
			wantedStatus: http.StatusConflict,
			body:         bodyJSON(t, map[string]any{"title": "Clean Code", "author": "Robert C. Martin"}),
		},
		{
			name: "201 when book is created",
			svc: &mockSvc{CreateFunc: func(ctx context.Context, BookRequest book.BookRequest) (string, error) {
				//TODO
			}},
			path:         "/books",
			wantedStatus: http.StatusCreated,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			h := book.NewHandler(testCase.svc)
			r := gin.New()
			r.GET("/books/1", h.GetById)

			req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			if rec.Code != testCase.wantedStatus {
				t.Fatalf("status=%d; wanted=%d body=%s", rec.Code, testCase.wantedStatus, rec.Body.String())
			}
		})
	}
}

func TestHandler_Checkout(t *testing.T) {
	tests := []struct {
		name         string
		svc          *mockSvc
		path         string
		wantedStatus int
	}{}
}

func TestHandler_Return(t *testing.T) {
	tests := []struct {
		name         string
		svc          *mockSvc
		path         string
		wantedStatus int
	}{}
}*/
