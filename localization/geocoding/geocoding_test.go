package geocoding

import (
	"context"
	"testing"
)

func TestGet(t *testing.T) {
	client := New(
		"gogistic/1.0 (contato@exemplo.com)",
	)

	address := Address{
		Street:     "Praça da Sé",
		City:       "São Paulo",
		State:      "SP",
		PostalCode: "01001-000",
		Country:    "Brasil",
	}

	location, err := client.Get(
		context.Background(),
		address,
	)

	if err != nil {
		t.Fatal(err)
	}

	if location.Latitude == 0 {
		t.Error("latitude is zero")
	}

	if location.Longitude == 0 {
		t.Error("longitude is zero")
	}

	t.Logf(
		"latitude: %f",
		location.Latitude,
	)

	t.Logf(
		"longitude: %f",
		location.Longitude,
	)

	t.Logf(
		"display name: %s",
		location.DisplayName,
	)
}
