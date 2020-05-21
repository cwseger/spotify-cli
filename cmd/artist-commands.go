package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cwseger/spotify-cli/spotify"
	cobra "github.com/spf13/cobra"
)

var artistCommands = []*cobra.Command{
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
			fmt.Println("Name:", out.Inner.Artists[0].Name)
			fmt.Println("Popularity:", out.Inner.Artists[0].Popularity)
			fmt.Println("Followers:", out.Inner.Artists[0].Followers.Total)
		},
	},
	{
		Use:     "artist-albums",
		Short:   "Get an artist's albums",
		Example: "spotify-cli artist-albums The Black Keys",
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

			out, err := spotifyClient.GetArtistAlbums(cmd.Context(), strings.Join(args, ""))
			if err != nil {
				fmt.Println(err, "Failed to get artist")
			}
			for i := range out.Albums {
				fmt.Println("Name:", out.Albums[i].Name)
			}
		},
	},
}
