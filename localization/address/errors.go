package address

import "errors"

var (
	ErrInvalidCEP  = errors.New("invalid CEP")
	ErrCEPNotFound = errors.New("CEP not found")
)
