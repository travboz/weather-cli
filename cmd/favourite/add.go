/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package favourite

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cityAdd string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a location to favourites",
	Long: `Save a location as a favourite so you can quickly retrieve weather information without entering the coordinates or city name every time.

	Example usage:
	weather favorite add --city "Melbourne"
	weather favorite add -c "London"

	Note, can only add using -c or --city flags.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("favourites added called")
	},
}

func init() {
	addCmd.Flags().StringVarP(&cityAdd, "city", "c", "", "The city to search for alerts")

}
