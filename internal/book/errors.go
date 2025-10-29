package book

import "fmt"

var (
	ErrInvalidFilter   = fmt.Errorf("Can't filter by author and title simultaneously")
	ErrNotFound        = fmt.Errorf("Book not found")
	ErrDuplicate       = fmt.Errorf("Book already exists")
	ErrBookUnavailable = fmt.Errorf("Book is not available for checkout")
)
