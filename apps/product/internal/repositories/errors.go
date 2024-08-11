package repositories

import "errors"

var (
	ErrNotFound      = errors.New("product not found")
	ErrAlreadyExists = errors.New("product already exists")
	ErrTitleRequired = errors.New("product title is required")
)
