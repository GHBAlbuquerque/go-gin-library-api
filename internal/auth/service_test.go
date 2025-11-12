package auth_test

import (
	"errors"
	"example/go-gin-library-api/internal/auth"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ---------- Tests ----------

func TestService_IssueToken_Success(t *testing.T) {
	svc := auth.NewService("XXXXX", "ISSUER", "AUDIENCE")
	got, err := svc.IssueToken("clientId", 60*time.Second)
	if err != nil {
		t.Fatalf("IssueToken returned error: %q", err)
	}
	if got.AccessToken == "" || got.ExpiresIn == 0 {
		t.Fatalf("IssueToken did not return correct TokenRes. \nGot %q", got)
	}
}

func TestService_IssueToken_Error(t *testing.T) {
	// Currently untestable, needs interface before jwt
}

func TestService_ParseAndValidate_Success(t *testing.T) {
	svc := auth.NewService("XXXXX", "ISSUER", "AUDIENCE")
	got, err := svc.ParseAndValidate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaWQiOiJmaXJzdF9jbGllbnQiLCJpc3MiOiJJU1NVRVIiLCJhdWQiOlsiQVVESUVOQ0UiXSwiZXhwIjoxNzYyOTEzMjQ4LCJpYXQiOjE3NjI5MDk2NDh9.kcjIpiZPSIBNhL1n8lTybVrcEdR0v2Tta7lAeQUSHSU")
	if err != nil {
		t.Fatalf("ParseAndValidate returned error: %q", err)
	}
	if got.ClientID == "" {
		t.Fatalf("ParseAndValidate did not return correct TokenRes. \nGot %q", got)
	}
}

func TestService_ParseAndValidate_ParseWithClaimError(t *testing.T) {
	svc := auth.NewService("XXXXX", "ISSUER", "AUDIENCE")
	_, err := svc.ParseAndValidate("x.xxxx2NDh9.kcjIpiZPSIBNhL1n8lTybVrcEdR0v2Tta7lAeQUSHSU")

	if err == nil {
		t.Fatal("ParseAndValidate did not return error")
	}
	if errors.Is(err, jwt.ErrTokenInvalidClaims) {
		t.Fatal("ParseAndValidate failed on validation and not parsing")
	}
}

func TestService_ParseAndValidate_ValidateError(t *testing.T) {
	svc := auth.NewService("XXXXX", "ISSUER", "AUDIENCE")
	_, err := svc.ParseAndValidate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaWQiOiJmaXJzdF9jbGllbnQiLCJpc3MiOiJJU1NVRVIyIiwiYXVkIjpbIkFVRElFTkNFIl0sImV4cCI6MTc2MjkxMzY3MSwiaWF0IjoxNzYyOTEwMDcxfQ.E5aTrvjK2CM4Nwt2oR-bFVGiNCbLuPq-bdF3Q72PeeM")

	if err == nil {
		t.Fatal("ParseAndValidate did not return error")
	}
	if !errors.Is(err, jwt.ErrTokenInvalidClaims) {
		t.Fatalf("ParseAndValidate did not fail on parsing: %q", err.Error())
	}
}
