package geocoding

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://nominatim.openstreetmap.org/search"

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
	Country    string
}

type Location struct {
	Latitude    float64
	Longitude   float64
	DisplayName string
}

type Client struct {
	httpClient *http.Client
	userAgent  string
}

func New(userAgent string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		userAgent: userAgent,
	}
}

func (c *Client) Get(
	ctx context.Context,
	address Address,
) (*Location, error) {
	params := url.Values{}

	params.Set("street", address.Street)
	params.Set("city", address.City)
	params.Set("state", address.State)
	params.Set("postalcode", address.PostalCode)
	params.Set("country", address.Country)

	params.Set("format", "jsonv2")
	params.Set("limit", "1")
	params.Set("countrycodes", "br")

	requestURL := fmt.Sprintf(
		"%s?%s",
		baseURL,
		params.Encode(),
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		requestURL,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create request: %w",
			err,
		)
	}

	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(
			"request geocoding: %w",
			err,
		)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"geocoding returned status %d",
			resp.StatusCode,
		)
	}

	var results []struct {
		Latitude    string `json:"lat"`
		Longitude   string `json:"lon"`
		DisplayName string `json:"display_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf(
			"decode geocoding response: %w",
			err,
		)
	}

	if len(results) == 0 {
		return nil, ErrAddressNotFound
	}

	latitude, err := strconv.ParseFloat(
		results[0].Latitude,
		64,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"parse latitude: %w",
			err,
		)
	}

	longitude, err := strconv.ParseFloat(
		results[0].Longitude,
		64,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"parse longitude: %w",
			err,
		)
	}

	return &Location{
		Latitude:    latitude,
		Longitude:   longitude,
		DisplayName: results[0].DisplayName,
	}, nil
}
