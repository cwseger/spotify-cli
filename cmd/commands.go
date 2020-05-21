package cmd

import (
	"fmt"
	"strings"

	"github.com/cwseger/spotify-cli/spotify"
	cobra "github.com/spf13/cobra"
)

var commands = []*cobra.Command{
	{
		Use:     "new-releases",
		Short:   "Get a list of new album releases featured in Spotify",
		Example: "spotify-cli new-releases",
		Run: func(cmd *cobra.Command, args []string) {
			spotifyClient, err := spotify.NewClient()
			if err != nil {
				fmt.Println("Failed to create new spotify client")
			}
			out, err := spotifyClient.GetNewReleases(cmd.Context())
			if err != nil {
				fmt.Println("Failed to get new releases")
			}

			for _, album := range out.Inner.Items {
				fmt.Println(album.Name)
			}
		},
	},
	{
		Use:   "recommendations",
		Short: "Get recommended tracks based on the provided artist",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			spotifyClient, err := spotify.NewClient()
			if err != nil {
				fmt.Println("Failed to create new spotify client", err)
				return
			}
			out, err := spotifyClient.GetRecommendationsByArtist(cmd.Context(), strings.Join(args, ""))
			if err != nil {
				fmt.Println("Failed to get recommendations by artist", err)
			}
			for _, track := range out.Tracks {
				fmt.Println(track.Album.Name)
			}
		},
	},
}
