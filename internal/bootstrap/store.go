package bootstrap

import (
	"example/go-gin-library-api/internal/book"
	"example/go-gin-library-api/internal/book/stores"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var books = []book.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func NewStoreFromEnv() (book.Store, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("godotenv.Load: %w", err)
	}

	env, ok := os.LookupEnv("BOOK_STORE")
	if !ok {
		return nil, fmt.Errorf("godotenv.Load: variable BOOK_STORE not found")
	}

	log.Printf("Starting application with store %q", env)

	switch strings.ToLower(env) {
	case "mysql":
		dsn := os.Getenv("BOOK_MYSQL_DSN")
		return stores.NewMySQL(dsn)
	case "json":
		path := os.Getenv("BOOK_JSON_PATH")
		return stores.NewJSON(path, books)
	case "memory":
		return stores.NewMemory(books)
	default:
		return nil, fmt.Errorf("unknown BOOK_STORE %q", env)
	}
}
