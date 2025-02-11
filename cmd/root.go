/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/travboz/weather-cli/cmd/config"
	"github.com/travboz/weather-cli/cmd/favourite"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "weather [command] [arguments] [flags]",
	Short: "Get the weather and related info for a location",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	/*
		net is added as a palette
		the package `net` gets added and all that is within net is kept private.
	*/
	rootCmd.AddCommand(config.ConfigCmd)

	// now we add the info palette of commands
	rootCmd.AddCommand(favourite.FavouriteCmd)

}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommandPalettes()
}
