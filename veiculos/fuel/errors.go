package fuel

import "errors"

var (
	ErrInvalidFuel   = errors.New("invalid fuel type")
	ErrPriceNotFound = errors.New("fuel price not found")
)
