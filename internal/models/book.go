package models

type Book struct {
	ID       string `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Quantity int    `json:"quantity" binding:"gte=0"` //greater than or equal to 0
}
