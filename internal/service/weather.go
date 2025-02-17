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

func (s *Service) GetDailyForecast(city string, lat, long float64, days int) error {
	if city == "" && (lat == 0.0 || long == 0.0) {
		return errors.New("you must provide either a city or both latitude and longitude")
	}

	forecastDays, err := s.apiClient.FetchForecastForDays(city, lat, long, days)
	if err != nil {
		return fmt.Errorf("failed to fetch current weather data from service: %w", err)
	}

	fmt.Printf("Daily weather forecast:\n")
	for _, day := range forecastDays.Forecast.Forecastday {
		fmt.Printf(
			"Date: %s\n"+
				"Temperature: Max %.1f°C (%.1f°F) | Min %.1f°C (%.1f°F)\n"+
				"Chance of Rain: %d%% | Will it rain: %t\n"+
				"Chance of Snow: %d%% | Will it snow: %t\n"+
				"Wind: %.1f mph (%.1f kph)\n"+
				"Condition: %s\n\n",
			day.Date,
			day.Day.MaxtempC, day.Day.MaxtempF,
			day.Day.MintempC, day.Day.MintempF,
			day.Day.DailyChanceOfRain, day.Day.DailyWillItRain == 1,
			day.Day.DailyChanceOfSnow, day.Day.DailyWillItSnow == 1,
			day.Day.MaxwindMph, day.Day.MaxwindKph,
			day.Day.Condition.Text,
		)

	}

	return nil
}

func (s *Service) GetHourlyForecast(city string, lat, long float64) error {
	if city == "" && (lat == 0.0 || long == 0.0) {
		return errors.New("you must provide either a city or both latitude and longitude")
	}

	hourlyForecast, err := s.apiClient.FetchForecastHourlyForNextDay(city, lat, long)
	if err != nil {
		return fmt.Errorf("failed to fetch current weather data from service: %w", err)
	}

	fmt.Println("Hourly Forecast:")
	for _, hour := range hourlyForecast.Forecast.Forecastday[0].Hour {
		fmt.Printf(
			"Time: %s\n"+
				"Temperature: %.1f°C (%.1f°F)\n"+
				"Condition: %s\n"+
				"Wind: %.1f mph (%.1f kph) from %s\n"+
				"Chance of Rain: %d%% | Will it rain: %t\n"+
				"Chance of Snow: %d%% | Will it snow: %t\n"+
				"Humidity: %d%% | Cloud Cover: %d%%\n\n",
			hour.Time,
			hour.TempC, hour.TempF,
			hour.Condition.Text,
			hour.WindMph, hour.WindKph, hour.WindDir,
			hour.ChanceOfRain, hour.WillItRain == 1,
			hour.ChanceOfSnow, hour.WillItSnow == 1,
			hour.Humidity, hour.Cloud,
		)
	}

	return nil
}
