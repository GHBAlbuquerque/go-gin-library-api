# 📚 Tiny Go Library API

A tiny RESTful API built with [Gin](https://github.com/gin-gonic/gin) to manage a small library of books - written for fun and learning.

Features:
- Dynamic store selection (in memory, json file or sql database) and persistence
- External environment variables file with .env
- REST API for library management (CRUD operations)
- Authentication using OAuth 2.0
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

5) Generate a Bearer Token by calling the '/auth/token' endpoint using the following credentials:
    ```
    "grant_type":"client_credentials"
    "client_id":"first_client"
    "client_secret":"first_password" 
    ```

6) Insert token on "Authorization" field of request 

7) Call the endpoints and test out the API 🌼 

* Postman collection inside /misc folder is already pre-configured for correct authentication and use. 

---

## 📖 Available Endpoints

| Method | Endpoint               | Description                                                            | Auth? | 
|--------|------------------------|------------------------------------------------------------------------|-------|
| `POST` | `/auth/token`          | Issues a bearer token when given valid `client_id` and `client_secret` |No auth.|
| `GET`  | `/api/books`           | Returns all books in the library. Can be filtered by `author` or `title` | *(requires `Authorization` header)* |
| `GET`  | `/api/books/:id`       | Returns a specific book by ID | *(requires `Authorization` header)* |
| `POST` | `/api/books`           | Adds a new book to the library | *(requires `Authorization` header)* |
| `PATCH`| `/api/checkout?id=1`   | Checks out (borrows) a book | *(requires `Authorization` header)* |
| `PATCH`| `/api/return?id=1`     | Returns a borrowed book | *(requires `Authorization` header)* |

---

## 🧩 Example Requests

### ➕ Create a book
```bash
curl -X POST http://localhost:8080/books   
    -H "Content-Type: application/json"   
    -H "Authorization: Bearer …"
    -d '{
    "id": "4",
    "title": "Pride and Prejudice",
    "author": "Jane Austen",
    "quantity": 3
  }'
```

### 📗 Get all books
```bash
curl http://localhost:8080/api/books -H "Authorization: Bearer …"
```

Filter by title:
```bash
curl http://localhost:8080/api/books?title=pride -H "Authorization: Bearer …"
```

Filter by author:
```bash
curl http://localhost:8080/api/books?author=fiodor -H "Authorization: Bearer …"
```

### 📘 Get a specific book
```bash
curl http://localhost:8080/api/books/1 -H "Authorization: Bearer …"
```

### 📕 Checkout a book
```bash
curl -X PATCH "http://localhost:8080/api/checkout?id=1" -H "Authorization: Bearer …"
```

### 📙 Return a book
```bash
curl -X PATCH "http://localhost:8080/api/return?id=1" -H "Authorization: Bearer …"
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
- Unit testing

---

> Made with ☕, ❤️, and curiosity - by @GHBAlbuquerque 🌸