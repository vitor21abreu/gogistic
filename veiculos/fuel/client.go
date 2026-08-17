package fuel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Get(
	ctx context.Context,
	fuel FuelType,
	city string,
	state string,
) (FuelPrice, error) {

	params := url.Values{}

	params.Set("mode", "municipality")
	params.Set("fuel", string(fuel))
	params.Set("municipality", city)
	params.Set("state", state)

	requestURL := c.baseURL + "?" + params.Encode()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		requestURL,
		nil,
	)
	if err != nil {
		return FuelPrice{}, fmt.Errorf("create fuel request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return FuelPrice{}, fmt.Errorf("execute fuel request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return FuelPrice{}, fmt.Errorf(
			"fuel provider returned status %d",
			resp.StatusCode,
		)
	}

	var response struct {
		Municipality string  `json:"municipio"`
		Price        float64 `json:"preco_medio"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return FuelPrice{}, fmt.Errorf(
			"decode fuel response: %w",
			err,
		)
	}

	return FuelPrice{
		Fuel:   fuel,
		Price:  response.Price,
		City:   response.Municipality,
		State:  state,
		Source: "ANP",
	}, nil
}
