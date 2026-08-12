package address

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Get(ctx context.Context, cep string) (*Address, error) {
	cep = normalizeCEP(cep)

	if len(cep) != 8 {
		return nil, ErrInvalidCEP
	}

	url := fmt.Sprintf("%s/%s/json/", baseURL, cep)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request ViaCEP: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"ViaCEP returned status %d",
			resp.StatusCode,
		)
	}

	var address struct {
		Address
		Erro bool `json:"erro"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return nil, fmt.Errorf(
			"decode ViaCEP response: %w",
			err,
		)
	}

	if address.Erro {
		return nil, ErrCEPNotFound
	}

	return &address.Address, nil
}

func normalizeCEP(cep string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}

		return -1
	}, cep)
}
