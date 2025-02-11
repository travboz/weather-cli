package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Data access layer: Calls the weather API

// Client is a wrapper for the weather API.
type Client struct {
	BaseURL string
	ApiKey  string
}

// NewClient initializes an API client.
func NewClient() *Client {
	return &Client{BaseURL: "http://api.weatherapi.com/v1/current.json?", ApiKey: "e3573d692a0848799ba42833252201"}
}

func (c *Client) FetchCurrentWeather(city string, lat, long float64) (*CurrentWeatherResponse, error) {

	var url string

	if city == "" {
		url = fmt.Sprintf("%skey=%s&q=%0.3f,%0.3f&aqi=no", c.BaseURL, c.ApiKey, lat, long)
	} else {
		url = fmt.Sprintf("%skey=%s&q=%s&aqi=no", c.BaseURL, c.ApiKey, city)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	var result CurrentWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil

}
