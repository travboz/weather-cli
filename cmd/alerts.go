/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// alertsCmd represents the alerts command
var alertsCmd = &cobra.Command{
	Use:   "alerts",
	Short: "Displays severe weather warnings if available",
	Long: `Show any future alerts that could occur at a location in the next day.

Example usage:
weather alerts --city "Sydney"
weather alerts -c "London"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		city := cmd.Flags().Lookup("city").Value.String()
		fmt.Println("Searching for alerts in", city)

	},
}

func init() {
	alertsCmd.Flags().StringP("city", "c", "", "The city to search for alerts")
	_ = alertsCmd.MarkFlagRequired("city")
	rootCmd.AddCommand(alertsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alertsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alertsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
