package cmd

import (
	"fmt"
	"os"

	cobra "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spotify-cli",
	Short: "This CLI allows you to interact with Spotify via the command line",
}

// Execute -
func Execute() {
	rootCmd.AddCommand(categoryListCmd)
	rootCmd.AddCommand(categoryPlaylistsCmd)
	rootCmd.AddCommand(recommendationsByArtistsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Failed to execute context")
		os.Exit(1)
	}
}
