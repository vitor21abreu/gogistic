package routing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://router.project-osrm.org"

type Point struct {
	Latitude  float64
	Longitude float64
}

type Route struct {
	Distance float64
	Duration float64
}

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

func (c *Client) Calculate(
	ctx context.Context,
	origin Point,
	destination Point,
) (*Route, error) {

	if origin.Latitude < -90 || origin.Latitude > 90 {
		return nil, ErrInvalidPoint
	}

	if destination.Latitude < -90 || destination.Latitude > 90 {
		return nil, ErrInvalidPoint
	}

	url := fmt.Sprintf(
		"%s/route/v1/driving/%f,%f;%f,%f?overview=false",
		baseURL,
		origin.Longitude,
		origin.Latitude,
		destination.Longitude,
		destination.Latitude,
	)

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
		return nil, fmt.Errorf("request OSRM: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"OSRM returned status %d",
			resp.StatusCode,
		)
	}

	var response struct {
		Code   string `json:"code"`
		Routes []struct {
			Distance float64 `json:"distance"`
			Duration float64 `json:"duration"`
		} `json:"routes"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf(
			"decode OSRM response: %w",
			err,
		)
	}

	if response.Code != "Ok" || len(response.Routes) == 0 {
		return nil, ErrRouteNotFound
	}

	return &Route{
		Distance: response.Routes[0].Distance,
		Duration: response.Routes[0].Duration,
	}, nil
}
