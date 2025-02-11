/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package favourite

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cityRemove string

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a location from favourites",
	Long: `Delete a saved favourite location by specifying its name or coordinates, ensuring it no longer appears in your favourites list.
	
	Example usage:
	weather favourite remove -c "Melbourne"
	weather favourite remove --city "London"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("favourites remove called")
	},
}

func init() {
	removeCmd.Flags().StringVarP(&cityRemove, "city", "c", "", "The city to search for alerts")
	FavouriteCmd.AddCommand(removeCmd)
}
