/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("forecast called")

		if hourly {
			fmt.Printf("\tand fetching hourly forecast for %s...", forecastCity)
		} else {
			fmt.Printf("\tand fetching %d-day forecast for %s...", dayCount, forecastCity)
		}
	},
}

func init() {
	rootCmd.AddCommand(forecastCmd)
	forecastCmd.Flags().StringVarP(&forecastCity, "city", "c", "", "The city to search for alerts")
	forecastCmd.Flags().IntVarP(&dayCount, "days", "d", 7, "Number of days")
	forecastCmd.Flags().BoolVarP(&hourly, "hourly", "h", false, "Get an hourly forecast instead of daily")
}
