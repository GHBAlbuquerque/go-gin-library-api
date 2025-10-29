# 📚 Tiny Go Library API

A tiny RESTful API built with [Gin](https://github.com/gin-gonic/gin) to manage a small library of books - written for fun and learning.

Features:
- Dynamic store selection (in memory, json file or sql database) and persistence
- External environment variables file with .env
- REST API for library management (CRUD operations)
- Local infrastructure with docker compose


## 🚀 How to run

1) Select the desired store by changing the BOOK_STORE on .env file
  
    `STORE OPTIONS: memory, json, mysql`

2) (optional) If choosing mysql, run these commands to create the necessary infrastructure (mysql database)
    ```bash
    cd go-gin-library-infra 
    docker compose up
    ```
    to delete the volumes and reestart the database, run:
    ```bash
    docker compose down -v
    ```

3) To start the server, run the following command `(Requires Go 1.22+)`
    ```bash
    go run ./cmd/booksrv
    ```
4) The server will start at:
    ```
    http://localhost:8080
    ```
5) Call the endpoints and test out the API 🌼 

---

## 📖 Available Endpoints

| Method | Endpoint           | Description |
|--------|--------------------|--------------|
| `GET`  | `/books`           | Returns all books in the library. Can be filtered by `author` or `title` |
| `GET`  | `/books/:id`       | Returns a specific book by ID |
| `POST` | `/books`           | Adds a new book to the library |
| `PATCH`| `/checkout?id=1`   | Checks out (borrows) a book |
| `PATCH`| `/return?id=1`     | Returns a borrowed book |

---

## 🧩 Example Requests

### ➕ Create a book
```bash
curl -X POST http://localhost:8080/books   -H "Content-Type: application/json"   -d '{
    "id": "4",
    "title": "Pride and Prejudice",
    "author": "Jane Austen",
    "quantity": 3
  }'
```

### 📗 Get all books
```bash
curl http://localhost:8080/books
```

Filter by title:
```bash
curl http://localhost:8080/books?title=pride
```

Filter by author:
```bash
curl http://localhost:8080/books?author=fiodor
```


### 📘 Get a specific book
```bash
curl http://localhost:8080/books/1
```

### 📕 Checkout a book
```bash
curl -X PATCH "http://localhost:8080/checkout?id=1"
```

### 📙 Return a book
```bash
curl -X PATCH "http://localhost:8080/return?id=1"
```

---

## 🛠️ Tech Stack
- **Go 1.22+**
- **Gin** web framework
- **MySQL 8.4** database

---

## 🌼 Future ideas
- Add validation and better error handling
- Pagination
- Authentication

---

> Made with ☕, ❤️, and curiosity - by @GHBAlbuquerque 🌸
