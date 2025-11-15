package auth_test

import (
	"example/go-gin-library-api/internal/auth"
	"testing"
)

// ---------- Tests ----------

func TestRepository_Validate_Success(t *testing.T) {
	repo := auth.NewInMemoryClientRepo(map[string]string{"id": "secret"})

	result := repo.Validate("id", "secret")

	if !result {
		t.Fatal("Valid client was deemed invalid.")
	}
}

func TestRepository_Validate_ClientDoesNotExistFailure(t *testing.T) {
	repo := auth.NewInMemoryClientRepo(map[string]string{"id": "secret"})

	result := repo.Validate("not id", "secret")

	if result {
		t.Fatal("Inexistent client was found.")
	}
}

func TestRepository_Validate_BadSecretFailure(t *testing.T) {
	repo := auth.NewInMemoryClientRepo(map[string]string{"id": "secret"})

	result := repo.Validate("id", "bad secret")

	if result {
		t.Fatal("Invalid client was deemed valid.")
	}
}
