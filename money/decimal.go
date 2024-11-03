package money

import (
	"strings"
)

type Decimal struct {
	subunits int64
	precision byte
}

// Errors
const (
	ErrInvalidDecimal = Error("Unable to convert the decimal")
	ErrTooLarge       = Error("quantity over 10^12 is too large")
)

func ParseDecimal(value string) (Decimal, error) {
	intPart, fractPart, _ := strings.Cut(value, ".")
}
