package cmd

import (
	"fmt"

	"github.com/cwseger/spotify-cli/spotify"
	cobra "github.com/spf13/cobra"
)

var categoryListCmd = &cobra.Command{
	Use:   "categories",
	Short: "Get a list of categories",
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
}

var categoryPlaylistsCmd = &cobra.Command{
	Use:   "category-playlist",
	Short: "Get a list of playlists tagged with the specified category",
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
}

var recommendationsByArtistsCmd = &cobra.Command{
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
}
