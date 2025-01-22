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
	Long: `Fetch the weather for all of the locations that you
	have saved.
	
	Example usage:
	weather favorites
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fetching weather for all saved locations")
	},
}

func init() {
	rootCmd.AddCommand(favouritesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// favouritesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// favouritesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
