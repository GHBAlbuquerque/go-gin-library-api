package book

type Book struct {
	ID       string
	Title    string
	Author   string
	Quantity int
}

type BookRequest struct {
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Quantity int    `json:"quantity" binding:"gte=0"` //greater than or equal to 0
}

type BookResponse struct {
	ID       string `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Quantity int    `json:"quantity" binding:"gte=0"` //greater than or equal to 0
}
