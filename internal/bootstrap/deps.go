package bootstrap

import (
	"example/go-gin-library-api/internal/auth"
	"example/go-gin-library-api/internal/book"
)

// BuildDeps wires together the auth and book handlers by pulling stores and
// clients from the environment and instantiating the required services.
func BuildDeps(authConfig AuthConfig) (*auth.Handler, *book.Handler, error) {
	// Create stores and repositories
	store, err := newStoreFromEnv()
	if err != nil {
		return nil, nil, err
	}

	clientRepo, err := newClientRepoFromEnv()
	if err != nil {
		return nil, nil, err
	}

	// Create services
	authSvc := auth.NewService(authConfig.JWTSecret, authConfig.Issuer, authConfig.Audience)
	bookSvc := book.NewService(store)

	// Create handlers
	authHandler := auth.NewHandler(clientRepo, authSvc)
	bookHandler := book.NewHandler(bookSvc)

	return authHandler, bookHandler, nil
}
