package errs

import "errors"

var (
	ErrNoOrder = errors.New("order with that ID not found")

	ErrUnmarshal = errors.New("failed to unmarshal data")

	ErrNatsMetadata    = errors.New("failed to get metadata")
	ErrNatsInvalidData = errors.New("no order_uid in data")
)
