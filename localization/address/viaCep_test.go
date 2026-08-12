package address

import (
	"context"
	"errors"
	"testing"
)

func TestGet(t *testing.T) {
	client := New()

	address, err := client.Get(
		context.Background(),
		"01001-000",
	)

	if err != nil {
		t.Fatal(err)
	}

	if address.CEP != "01001-000" {
		t.Errorf(
			"unexpected CEP: %s",
			address.CEP,
		)
	}

	if address.UF != "SP" {
		t.Errorf(
			"unexpected UF: %s",
			address.UF,
		)
	}
}

func TestInvalidCEP(t *testing.T) {
	client := New()

	_, err := client.Get(
		context.Background(),
		"123",
	)

	if !errors.Is(err, ErrInvalidCEP) {
		t.Errorf(
			"expected ErrInvalidCEP, got %v",
			err,
		)
	}
}

func TestGetAddress(t *testing.T) {
	cep := &Cep{
		Sender:    "01001-000",
		Recipient: "20040-020",
	}

	addresses, err := cep.GetAddress()
	if err != nil {
		t.Fatal(err)
	}

	if len(addresses) != 2 {
		t.Fatalf(
			"expected 2 addresses, got %d",
			len(addresses),
		)
	}
}
