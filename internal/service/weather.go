package service

import (
	"errors"
	"fmt"

	"github.com/travboz/weather-cli/internal/api"
)

// Business logic:  Processes weather data

// Service handles weather logic.
type Service struct {
	apiClient *api.Client
}

// NewService initializes a weather service.
func NewService(apiClient *api.Client) *Service {
	return &Service{apiClient: apiClient}
}

func (s *Service) GetCurrentWeather(city string, lat, long float64) error {
	if city == "" && (lat == 0.0 || long == 0.0) {
		return errors.New("you must provide either a city or both latitude and longitude")
	}

	weatherData, err := s.apiClient.FetchCurrentWeather(city, lat, long)
	if err != nil {
		return fmt.Errorf("failed to fetch current weather data from service: %w", err)
	}

	// print response
	fmt.Printf(
		"Weather Report:\n"+
			"Location: %s, %s, %s (Lat: %.2f, Lon: %.2f)\n"+
			"Last Updated: %s\n"+
			"Temperature: %.1f°C (%.1f°F)\n"+
			"Condition: %s\n"+
			"Wind: %.1f mph (%.1f kph)\n"+
			"Humidity: %d%% | Cloud Cover: %d%%\n"+
			"Feels Like: %.1f°C (%.1f°F)\n"+
			"Visibility: %.1f km | UV Index: %.1f\n",
		weatherData.Location.Name, weatherData.Location.Region, weatherData.Location.Country, weatherData.Location.Lat, weatherData.Location.Lon,
		weatherData.Current.LastUpdated,
		weatherData.Current.TempC, weatherData.Current.TempF,
		weatherData.Current.Condition.Text,
		weatherData.Current.WindMph, weatherData.Current.WindKph,
		weatherData.Current.Humidity, weatherData.Current.Cloud,
		weatherData.Current.FeelslikeC, weatherData.Current.FeelslikeF,
		weatherData.Current.VisKm, weatherData.Current.Uv,
	)

	return nil
}
