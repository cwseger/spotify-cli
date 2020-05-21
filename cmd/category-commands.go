package cmd

import (
	"errors"
	"fmt"

	"github.com/cwseger/spotify-cli/spotify"
	cobra "github.com/spf13/cobra"
)

var categoryCommands = []*cobra.Command{
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
}
