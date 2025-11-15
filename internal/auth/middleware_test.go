package auth_test

import (
	"errors"
	"example/go-gin-library-api/internal/auth"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

// ---------- Mock with function fields ----------

type mockRepo struct {
	ValidateFunc func(clientID, clientSecret string) bool
}

func (r *mockRepo) Validate(clientID, clientSecret string) bool {
	if r.ValidateFunc == nil {
		panic("Validate called but ValidateFunc is nil")
	}
	return r.ValidateFunc(clientID, clientSecret)
}

type mockSvc struct {
	IssueTokenFunc       func(clientID string, ttl time.Duration) (auth.TokenRes, error)
	ParseAndValidateFunc func(token string) (*auth.Claims, error)
}

func (m *mockSvc) IssueToken(clientID string, ttl time.Duration) (auth.TokenRes, error) {
	if m.IssueTokenFunc == nil {
		panic("IssueToken called but IssueTokenFunc is nil")
	}
	return m.IssueTokenFunc(clientID, ttl)
}

func (m *mockSvc) ParseAndValidate(token string) (*auth.Claims, error) {
	if m.ParseAndValidateFunc == nil {
		panic("ParseAndValidate called but ParseAndValidateFunc is nil")
	}
	return m.ParseAndValidateFunc(token)
}

var _ auth.ClientRepository = (*mockRepo)(nil)
var _ auth.Service = (*mockSvc)(nil)

// ---------- Tests ----------

func TestMiddleware_RequireAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		header       string
		mockSvc      *mockSvc
		wantedStatus int
		wantNext     bool
	}{
		{name: "ok",
			header: "Bearer good",
			mockSvc: &mockSvc{
				ParseAndValidateFunc: func(token string) (*auth.Claims, error) {
					return &auth.Claims{ClientID: "c"}, nil
				}},
			wantedStatus: http.StatusOK,
			wantNext:     true},
		{name: "missing bearer",
			header:       "",
			mockSvc:      nil,
			wantedStatus: http.StatusUnauthorized,
			wantNext:     false},
		{name: "invalid token",
			header: "Bearer bad",
			mockSvc: &mockSvc{
				ParseAndValidateFunc: func(token string) (*auth.Claims, error) {
					return nil, errors.New("bad")
				}},
			wantedStatus: http.StatusUnauthorized,
			wantNext:     false},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			h := auth.NewHandler(&mockRepo{}, testCase.mockSvc)
			r := gin.New()

			// create a new protected path for testing
			protected := r.Group("/protected")
			protected.Use(h.RequireAuth())
			protected.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

			// call the protected path
			req := httptest.NewRequest(http.MethodGet, "/protected/ping", nil)
			req.Header.Set("Authorization", testCase.header)
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			if rec.Code != testCase.wantedStatus {
				t.Fatalf("status=%d; wanted=%d body=%s", rec.Code, testCase.wantedStatus, rec.Body.String())
			}
		})
	}
}
