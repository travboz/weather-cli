/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/travboz/weather-cli/cmd/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage weather CLI configuration settings",
	Long: `Configure and manage settings for the weather CLI, such as default location and the preferred units (Celsius/Fahrenheit),

weather config set --unit metric --default-city "Brisbane"
weather config show
weather config reset
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(config.SetCmd)
	configCmd.AddCommand(config.ShowCmd)
	configCmd.AddCommand(config.ResetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
