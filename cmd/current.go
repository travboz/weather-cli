/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var city string
var lat, lon float64

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get the current weather in a city",
	Long: `Get the current weather for a city you input by either searching
	by city name or latitude/longitude.

Also, specify the preferred units as well (metric as default).
	
Example usage:
weather current --city "Sydney"
weather current --lat -33.87 --lon 151.21
weather current -c "New York" -u imperial
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
			fmt.Printf("Fetching weather for city: %s\n", city)
		} else {
			fmt.Printf("Fetching weather for coordinates: %f, %f\n", lat, lon)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)

	currentCmd.Flags().StringVarP(&city, "city", "c", "", "The city to search for alerts")
	currentCmd.Flags().Float64Var(&lat, "lat", 0.0, "Latitude")
	currentCmd.Flags().Float64Var(&lon, "lon", 0.0, "Longitude")
}
