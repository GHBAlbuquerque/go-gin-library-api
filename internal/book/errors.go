package book

import "fmt"

var (
	ErrInvalidFilter   = fmt.Errorf("can't filter by author and title simultaneously")
	ErrNotFound        = fmt.Errorf("book not found")
	ErrDuplicate       = fmt.Errorf("book already exists")
	ErrBookUnavailable = fmt.Errorf("book is not available for checkout")
)
