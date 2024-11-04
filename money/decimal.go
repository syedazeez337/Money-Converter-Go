package money

import (
	"fmt"
	"strconv"
	"strings"
)

type Decimal struct {
	subunits  int64
	precision byte
}

// Errors
const (
	ErrInvalidDecimal = Error("Unable to convert the decimal")
	ErrTooLarge       = Error("quantity over 10^12 is too large")
)

func ParseDecimal(value string) (Decimal, error) {
	intPart, fractPart, _ := strings.Cut(value, ".")

	const maxDecimal = 12
	if len(intPart) > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	combinedValue := intPart + fractPart

	subunits, err := strconv.ParseInt(combinedValue, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	precision := byte(len(fractPart))

	return Decimal{
		subunits:  subunits,
		precision: precision,
	}, nil
}
