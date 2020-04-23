package cmd

import (
	"fmt"

	"github.com/cwseger/spotify-cli/spotify"
	cobra "github.com/spf13/cobra"
)

var commands = []*cobra.Command{
	{
		Use:     "artist",
		Short:   "Get an artist",
		Example: "spotify-cli artist The Black Keys",
		Run: func(cmd *cobra.Command, args []string) {
			spotifyClient, err := spotify.NewClient()
			if err != nil {
				fmt.Println(err, "Failed to create new spotify client")
				return
			}
			out, err := spotifyClient.GetArtist(cmd.Context(), args)
			if err != nil {
				fmt.Println(err, "Failed to get artist")
			}
			fmt.Println(fmt.Sprintf("Name: %s, Follower Count: %d, Popularity: %d", out.Inner.Artists[0].Name, out.Inner.Artists[0].Followers.Total, out.Inner.Artists[0].Popularity))
		},
	},
	{
		Use:     "categories",
		Short:   "Get a list of categories",
		Example: "spotify-cli categories",
		Run: func(cmd *cobra.Command, args []string) {
			spotifyClient, err := spotify.NewClient()
			if err != nil {
				fmt.Println(err, "Failed to create new spotify client")
				return
			}
			out, err := spotifyClient.GetCategoryList(cmd.Context())
			if err != nil {
				fmt.Println(err, "Failed to get category list")
			}
			for _, category := range out.Inner.Items {
				fmt.Println(category.Name)
			}
		},
	},
	{
		Use:     "category-playlist",
		Short:   "Get a list of playlists tagged with the specified category",
		Example: "spotify-cli category-playlist chill",
		Run: func(cmd *cobra.Command, args []string) {
			spotifyClient, err := spotify.NewClient()
			if err != nil {
				fmt.Println("Failed to create new spotify client")
				return
			}
			out, err := spotifyClient.GetCategoryPlaylists(cmd.Context(), args[0])
			if err != nil {
				fmt.Println("Failed to get category")
			}
			for _, playlist := range out.Inner.Items {
				fmt.Println(fmt.Sprintf("%s -- %s", playlist.Name, playlist.URI))
			}
		},
	},
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
		Run: func(cmd *cobra.Command, args []string) {
			spotifyClient, err := spotify.NewClient()
			if err != nil {
				fmt.Println("Failed to create new spotify client")
				return
			}
			out, err := spotifyClient.GetRecommendationsByArtists(cmd.Context(), args...)
			if err != nil {
				fmt.Println("Failed to get category list")
			}
			for _, track := range out.Tracks {
				fmt.Println(track.Album.Name)
			}
		},
	},
}
