package fuel

import "context"

// PriceProvider defines a source of fuel prices.
type PriceProvider interface {
	GetPrice(
		ctx context.Context,
		fuel FuelType,
		city string,
		state string,
	) (FuelPrice, error)
}
