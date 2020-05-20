package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cwseger/spotify-cli/spotify"
	cobra "github.com/spf13/cobra"
)

var commands = []*cobra.Command{
	{
		Use:     "artist",
		Short:   "Get an artist",
		Example: "spotify-cli artist The Black Keys",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Must provide an artist as an argument")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			spotifyClient, err := spotify.NewClient()
			if err != nil {
				fmt.Println(err, "Failed to create new spotify client")
				return
			}

			out, err := spotifyClient.GetArtist(cmd.Context(), strings.Join(args, ""))
			if err != nil {
				fmt.Println(err, "Failed to get artist")
			}
			fmt.Println(fmt.Sprintf("Name: %s, Follower Count: %d, Popularity: %d", out.Inner.Artists[0].Name, out.Inner.Artists[0].Followers.Total, out.Inner.Artists[0].Popularity))
		},
	},
	{
		Use:     "categories",
		Short:   "Get a list of categories",
		Example: "spotify-cli categories 2",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				return errors.New("Too many args")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			spotifyClient, err := spotify.NewClient()
			if err != nil {
				fmt.Println(err, "Failed to create new spotify client")
				return
			}
			if len(args) < 1 {
				args = []string{"50"}
			}
			out, err := spotifyClient.GetCategoryList(cmd.Context(), args[0])
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
		Args:    cobra.ExactArgs(1),
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
	{
		Use:     "album",
		Short:   "Get Spotify catalog information for a single album",
		Args:    cobra.ExactArgs(1),
		Example: "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			// spotifyClient, err := spotify.NewClient()
			// if err != nil {
			// 	fmt.Println("Failed to create new spotify client")
			// 	return
			// }
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
