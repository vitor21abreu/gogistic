package routing

import (
	"context"
	"testing"
)

func TestCalculate(t *testing.T) {
	client := New()

	origin := Point{
		Latitude:  -12.9714,
		Longitude: -38.5014,
	}

	destination := Point{
		Latitude:  -12.9777,
		Longitude: -38.5016,
	}

	route, err := client.Calculate(
		context.Background(),
		origin,
		destination,
	)

	if err != nil {
		t.Fatal(err)
	}

	if route.Distance <= 0 {
		t.Errorf(
			"expected distance > 0, got %f",
			route.Distance,
		)
	}

	if route.Duration <= 0 {
		t.Errorf(
			"expected duration > 0, got %f",
			route.Duration,
		)
	}

	t.Logf(
		"distance: %.2f km",
		route.Distance/1000,
	)

	t.Logf(
		"duration: %.2f minutes",
		route.Duration/60,
	)
}
