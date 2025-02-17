/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/travboz/weather-cli/internal/api"
	"github.com/travboz/weather-cli/internal/service"
)

var (
	forecastCity string
	dayCount     int
	hourly       bool
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get the weather forecast for a location",
	Long: `Retrieve detailed weather forecasts for a specified location, including temperature, precipitation, and other key weather conditions for a given number of days.
	
Example usage:
weather forecast --city "Sydney" --days 5
weather forecast --city "Sydney" --hourly
weather forecast -c "Tokyo" -d 3

	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Enforce mutual dependency between lat/lon
		if (cmd.Flags().Changed("lat") || cmd.Flags().Changed("lon")) &&
			!(cmd.Flags().Changed("lat") && cmd.Flags().Changed("lon")) {
			return errors.New("both --lat and --lon must be provided together")
		}

		// Ensure either lat/lon or city is provided
		if !cmd.Flags().Changed("city") && !(cmd.Flags().Changed("lat") && cmd.Flags().Changed("lon")) {
			return errors.New("you must provide either --city or both --lat and --lon")
		}

		// Execute the command logic
		if cmd.Flags().Changed("city") {
			fmt.Printf("Fetching forecast for city: %s\n", forecastCity)
		} else {
			fmt.Printf("Fetching forecast for coordinates: %f, %f\n", lat, lon)
		}

		if dayCount < 1 || dayCount > 10 {
			return errors.New("number of days must be between 1 and 10 (inclusive)")
		}

		// Initialize API client and service
		client := api.NewClient("http://api.weatherapi.com/v1/forecast.json?")
		service := service.NewService(client)

		if hourly {
			return service.GetHourlyForecast(forecastCity, lat, lon)
		} else {
			fmt.Printf("Fetching daily forecast for the next %d days\n\n", dayCount)
			return service.GetDailyForecast(forecastCity, lat, lon, dayCount)
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(forecastCmd)
	forecastCmd.Flags().StringVarP(&forecastCity, "city", "c", "", "The city to search for alerts")
	forecastCmd.Flags().IntVarP(&dayCount, "days", "d", 1, "Number of days (max of 10)")
	forecastCmd.Flags().BoolVarP(&hourly, "hourly", "", false, "Get an hourly forecast for the next 24 hours")
}
