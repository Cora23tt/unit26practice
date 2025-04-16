package application

import "errors"

var (
	ErrInvalidParams = errors.New("invalid parameters")
	ErrNotFound      = errors.New("not found")
	ErrInternal      = errors.New("internal error")
)
