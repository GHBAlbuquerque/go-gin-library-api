# 📚 Tiny Go Library API

A tiny RESTful API built with [Gin](https://github.com/gin-gonic/gin) to manage a small library of books — written just for fun and learning 🐣✨

---

## 🚀 How to run

```bash
go run main.go
```

The server will start at:
```
http://localhost:8080
```

---

## 📖 Available Endpoints

| Method | Endpoint           | Description |
|--------|--------------------|--------------|
| `GET`  | `/books`           | Returns all books in the library |
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
- Simple in-memory slice (no database yet)

---

## 🌼 Future ideas
- Add persistence (SQLite or JSON file)
- Add book search by title or author
- Add validation and better error handling
- Split into routes, services, and models ✨

---

> Made with ☕, ❤️, and curiosity — by [your name here] 🌸
