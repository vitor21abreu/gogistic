package address

import "context"

type Cep struct {
	Sender    string
	Recipient string
}

func (c *Cep) GetAddress() ([]Address, error) {
	client := New()

	recipient, err := client.Get(
		context.Background(),
		c.Recipient,
	)
	if err != nil {
		return nil, err
	}

	sender, err := client.Get(
		context.Background(),
		c.Sender,
	)
	if err != nil {
		return nil, err
	}

	return []Address{
		*recipient,
		*sender,
	}, nil
}
