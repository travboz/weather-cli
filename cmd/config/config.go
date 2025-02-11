/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage weather CLI configuration settings",
	Long: `Configure and manage settings for the weather CLI, such as default location and the preferred units (Celsius/Fahrenheit),

weather config set --unit metric --default-city "Brisbane"
weather config show
weather config reset
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
