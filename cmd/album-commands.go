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
			out, err := spotifyClient.GetAlbum(cmd.Context(), strings.Join(args, " "))
			if err != nil {
				fmt.Println("Failed to get album:", err)
			}
			fmt.Println("Name", "|", "Artist", "|", "Popularity")
			fmt.Println(out.Name, "|", out.Artists[0].Name, "|", out.Popularity)
		},
	},
	{
		Use:     "album-tracks",
		Short:   "Get Spotify catalog information about an album’s tracks",
		Example: "spotify-cli album-tracks TODO",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			spotifyClient, err := spotify.NewClient()
			if err != nil {
				fmt.Println("Failed to create new spotify client")
				return
			}

			out, err := spotifyClient.GetAlbumTracks(cmd.Context(), strings.Join(args, " "))
			if err != nil {
				fmt.Println("Failed to get album tracks:", err)
			}

			fmt.Println("Track Number", "|", "Name")
			for _, t := range out.Tracks {
				fmt.Println(t.TrackNumber, "|", t.Name)
			}
		},
	},
}
