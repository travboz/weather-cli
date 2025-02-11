/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
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
	ConfigCmd.AddCommand(resetCmd)

}
