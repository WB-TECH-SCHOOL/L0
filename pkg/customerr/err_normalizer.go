package customerr

import (
	"fmt"
	"strings"
)

const (
	ScanErr  = "scan error: %v"
	QueryErr = "query error: %v"
	ExecErr  = "execution error: %v"
	RowsErr  = "rows error: %v"
)

type ErrorPair struct {
	Message string
	Err     error
}

func ErrNormalizer(errPairs ...ErrorPair) error {
	var errStrings []string
	for _, pair := range errPairs {
		if pair.Err != nil {
			errStrings = append(errStrings, fmt.Sprintf(pair.Message, pair.Err))
		}
	}
	return fmt.Errorf(strings.Join(errStrings, ", "))
}
