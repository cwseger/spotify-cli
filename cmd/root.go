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
	rootCmd.AddCommand(artistCommands...)
	rootCmd.AddCommand(albumCommands...)
	rootCmd.AddCommand(categoryCommands...)
	rootCmd.AddCommand(commands...)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err, "Failed to execute context")
		os.Exit(1)
	}
}
