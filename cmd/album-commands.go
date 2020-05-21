package cmd

import (
	"fmt"
	"strings"

	"github.com/cwseger/spotify-cli/spotify"
	cobra "github.com/spf13/cobra"
)

var albumCommands = []*cobra.Command{
	{
		Use:     "album",
		Short:   "Get Spotify catalog information for a single album",
		Args:    cobra.MinimumNArgs(1),
		Example: "spotify-cli Control",
		Run: func(cmd *cobra.Command, args []string) {
			spotifyClient, err := spotify.NewClient()
			if err != nil {
				fmt.Println("Failed to create new spotify client")
				return
			}
			out, err := spotifyClient.GetAlbum(cmd.Context(), strings.Join(args, ""))
			if err != nil {
				fmt.Println("Failed to get album:", err)
			}
			fmt.Println(*out)
		},
	},
	{
		Use:     "album-tracks",
		Short:   "Get Spotify catalog information about an album’s tracks",
		Example: "spotify-cli album-tracks TODO",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// spotifyClient, err := spotify.NewClient()
			// if err != nil {
			// 	fmt.Println("Failed to create new spotify client")
			// 	return
			// }
		},
	},
}
