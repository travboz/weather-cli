/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package favourite

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all favourite locations",
	Long: `Display all saved favourite locations along with their corresponding city names, latitude, and longitude if available.
	
	Example usage:
	weather favourite list
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("favourites list called")
	},
}

func init() {
	FavouriteCmd.AddCommand(listCmd)
}
