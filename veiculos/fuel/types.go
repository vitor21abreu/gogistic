package fuel

import "time"

// FuelPrice represents a fuel price returned by a provider.
type FuelPrice struct {
	Fuel          FuelType
	Price         float64
	City          string
	State         string
	ReferenceDate time.Time
	Source        string
}
