package errs

import "errors"

var (
	ErrNoOrder = errors.New("order with that ID not found")
)
