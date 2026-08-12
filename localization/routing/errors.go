package routing

import "errors"

var (
	ErrInvalidPoint  = errors.New("invalid point")
	ErrRouteNotFound = errors.New("route not found")
)
