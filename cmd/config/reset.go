/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset configuration settings",
	Long: `Restore all configuration settings to their default values.
	
Example usage:
weather config reset
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("weather config reset called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
