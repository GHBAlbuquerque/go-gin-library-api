package bootstrap

import (
	"example/go-gin-library-api/internal/auth"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AuthConfig struct {
	JWTSecret string
	Issuer    string
	Audience  string
	Addr      string
}

// LoadAuthConfigFromEnv reads the auth settings from .env variables, ensuring
// required values exist before returning a populated AuthConfig.
func LoadAuthConfigFromEnv() (AuthConfig, error) {
	if err := godotenv.Load(".env"); err != nil {
		return AuthConfig{}, fmt.Errorf("godotenv.Load: %w", err)
	}

	log.Print("Loading Auth configs from env")

	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return AuthConfig{}, fmt.Errorf("godotenv.Load: variable JWT_SECRET not found")
	}

	issuer, ok := os.LookupEnv("ISSUER")
	if !ok {
		return AuthConfig{}, fmt.Errorf("godotenv.Load: variable ISSUER not found")
	}

	audience, ok := os.LookupEnv("AUDIENCE")
	if !ok {
		return AuthConfig{}, fmt.Errorf("godotenv.Load: variable AUDIENCE not found")
	}

	address, ok := os.LookupEnv("ADDR")
	if !ok {
		return AuthConfig{}, fmt.Errorf("godotenv.Load: variable ADDR not found")
	}

	return AuthConfig{
		JWTSecret: secret,
		Issuer:    issuer,
		Audience:  audience,
		Addr:      address,
	}, nil
}

// newClientRepoFromEnv builds an in-memory client repository using credentials
// sourced from the CLIENT_ID and CLIENT_SECRET environment variables.
func newClientRepoFromEnv() (auth.ClientRepository, error) {
	clientID, ok := os.LookupEnv("CLIENT_ID")
	if !ok {
		return nil, fmt.Errorf("godotenv.Load: variable CLIENT_ID not found")
	}

	clientSecret, ok := os.LookupEnv("CLIENT_SECRET")
	if !ok {
		return nil, fmt.Errorf("godotenv.Load: variable CLIENT_SECRET not found")
	}

	return auth.NewInMemoryClientRepo(map[string]string{clientID: clientSecret}), nil
}
