/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/travboz/weather-cli/cmd/favourite"
)

// favouriteCmd represents the favourite command
var favouriteCmd = &cobra.Command{
	Use:   "favourite",
	Short: "Manage and save favourite locations",
	Long: `Manage your favourite cities around the world by adding to,
	removing from, and listing all them. Can only add by 'city' designation.
	
	Example usage:
	weather favorite add --city "Melbourne"
	weather favorite list
	weather favorite remove "Melbourne"

	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("favourite called")
	},
}

func init() {
	rootCmd.AddCommand(favouriteCmd)

	// Attach subcommands from the favourite package
	favouriteCmd.AddCommand(favourite.AddCmd)
	favouriteCmd.AddCommand(favourite.ListCmd)
	favouriteCmd.AddCommand(favourite.RemoveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// favouriteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// favouriteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
