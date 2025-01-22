package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	unit        string
	defaultCity string
)

// setCmd represents the set command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Long: `Update a specific configuration setting, such as the default city or temperature units. 
	
Example usage:
weather config set --unit "metric" --default-city "Brisbane"
weather config set --default-city "Sydney"
weather config set --unit "imperial"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if unit != "" {
			fmt.Printf("Setting unit to: %s\n", unit)
		}
		if defaultCity != "" {
			fmt.Printf("Setting default city to: %s\n", defaultCity)
		}
	},
}

func init() {
	SetCmd.Flags().StringVar(&unit, "unit", "", "Set temperature unit (metric/imperial)")
	SetCmd.Flags().StringVar(&defaultCity, "default-city", "", "Set the default city for weather queries")
}
