/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package favourite

import (
	"github.com/spf13/cobra"
)

// favouriteCmd represents the favourite command
var FavouriteCmd = &cobra.Command{
	Use:   "favourite",
	Short: "Manage and save favourite locations",
	Long: `Manage your favourite cities around the world by adding, removing or listing your favourites.
	
Example usage:
weather favourite add --city "Melbourne"
weather favourite list
weather favourite remove "Melbourne"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
