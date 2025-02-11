/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display current weather CLI configuration",
	Long: `Show all stored configuration settings, specifically the default city, and temperature units.

Example usage:
weather config show
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
	},
}

func init() {
	ConfigCmd.AddCommand(showCmd)
}
