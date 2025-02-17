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
func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL, ApiKey: "e3573d692a0848799ba42833252201"}
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

func (c *Client) FetchForecastForDays(city string, lat, long float64, days int) (*ForecastWeatherResponseDay, error) {
	var url string

	if city == "" {
		url = fmt.Sprintf("%skey=%s&q=%0.3f,%0.3f&days=%d&aqi=no&alerts=no", c.BaseURL, c.ApiKey, lat, long, days)
	} else {
		url = fmt.Sprintf("%skey=%s&q=%s&days=%d&aqi=no&alerts=no", c.BaseURL, c.ApiKey, city, days)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	var result ForecastWeatherResponseDay
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FetchForecastHourlyForNextDay(city string, lat, long float64) (*ForecastWeatherResponseHourly, error) {
	var url string

	if city == "" {
		url = fmt.Sprintf("%skey=%s&q=%0.3f,%0.3f&days=1&aqi=no&alerts=no", c.BaseURL, c.ApiKey, lat, long)
	} else {
		url = fmt.Sprintf("%skey=%s&q=%s&days=1&aqi=no&alerts=no", c.BaseURL, c.ApiKey, city)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	var result ForecastWeatherResponseHourly
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
