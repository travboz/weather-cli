/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// favouritesCmd represents the favourites command
var favouritesCmd = &cobra.Command{
	Use:   "favourites",
	Short: "Fetch weather for all saved favourite locations",
	Long: `Fetch the weather for all of the locations that you have saved.
	
Example usage:
weather favourites
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fetching weather for all saved locations")
	},
}

func init() {
	rootCmd.AddCommand(favouritesCmd)
}
