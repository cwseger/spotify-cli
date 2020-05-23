package spotify

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	req "github.com/cwseger/spotify-cli/req"

	"github.com/pkg/errors"
)

// Client -
type Client interface {
	GetArtist(ctx context.Context, artist string) (*GetArtistOutput, error)
	GetArtistAlbums(ctx context.Context, artist string) (*GetArtistAlbumOutput, error)
	GetCategoryList(ctx context.Context, limit string) (*GetCategoriesOutput, error)
	GetCategoryPlaylists(ctx context.Context, categoryID string) (*GetCategoryPlaylistsOutput, error)
	GetRecommendationsByArtist(ctx context.Context, artist string) (*GetRecommendationsByArtistOutput, error)
	GetNewReleases(ctx context.Context) (*GetNewReleasesOutput, error)
	GetAlbum(ctx context.Context, album string) (*GetAlbumOutput, error)
	GetAlbumTracks(ctx context.Context, album string) (*GetAlbumTracksOutput, error)
}

var _ Client = &DefaultClient{}

// DefaultClient -
type DefaultClient struct {
	authToken *GetTokenOutput
	requestor req.Requestor
}

// NewClient -
func NewClient() (*DefaultClient, error) {
	getTokenOutput, err := getAccessToken(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get access token")
	}

	return &DefaultClient{
		authToken: getTokenOutput,
		requestor: req.NewRequestor(),
	}, nil
}

// GetArtist -
func (c *DefaultClient) GetArtist(ctx context.Context, artist string) (*GetArtistOutput, error) {
	queryParams := &map[string]string{
		"q":     artist,
		"limit": "1",
		"type":  "artist",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.authToken.AccessToken,
	}
	var output GetArtistOutput
	if err := c.requestor.Get(ctx, &req.GetInput{
		URL:         "https://api.spotify.com/v1/search",
		QueryParams: queryParams,
		Headers:     headers,
		Destination: &output,
	}); err != nil {
		return nil, errors.WithMessage(err, "Failed to get artist")
	}
	return &output, nil
}

// GetArtistAlbums -
func (c *DefaultClient) GetArtistAlbums(ctx context.Context, artist string) (*GetArtistAlbumOutput, error) {
	var getArtistSearchOutput GetArtistSearchOutput
	if err := c.getSpotifyIDForResource(ctx, artist, "artist", &getArtistSearchOutput); err != nil {
		return nil, errors.WithMessage(err, "Failed to get spotify id for artist")
	}
	slugs := &map[string]string{
		"{artistID}": getArtistSearchOutput.Inner.Items[0].ID,
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.authToken.AccessToken,
	}
	var output GetArtistAlbumOutput
	if err := c.requestor.Get(ctx, &req.GetInput{
		URL:         "https://api.spotify.com/v1/artists/{artistID}/albums",
		Slugs:       slugs,
		Headers:     headers,
		Destination: &output,
	}); err != nil {
		return nil, errors.WithMessage(err, "Failed to get artist's albums")
	}
	return &output, nil
}

// GetCategoryList -
func (c *DefaultClient) GetCategoryList(ctx context.Context, limit string) (*GetCategoriesOutput, error) {
	queryParams := &map[string]string{
		"limit": limit,
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.authToken.AccessToken,
	}
	var output GetCategoriesOutput
	if err := c.requestor.Get(ctx, &req.GetInput{
		URL:         "https://api.spotify.com/v1/browse/categories",
		QueryParams: queryParams,
		Headers:     headers,
		Destination: &output,
	}); err != nil {
		return nil, errors.WithMessage(err, "Failed to get category list")
	}
	return &output, nil
}

// GetCategoryPlaylists -
func (c *DefaultClient) GetCategoryPlaylists(ctx context.Context, categoryID string) (*GetCategoryPlaylistsOutput, error) {
	slugs := &map[string]string{
		"{categoryID}": categoryID,
	}
	queryParams := &map[string]string{
		"limit": "5",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.authToken.AccessToken,
	}
	var output GetCategoryPlaylistsOutput
	if err := c.requestor.Get(ctx, &req.GetInput{
		URL:         "https://api.spotify.com/v1/browse/categories/{categoryID}/playlists",
		Slugs:       slugs,
		QueryParams: queryParams,
		Headers:     headers,
		Destination: &output,
	}); err != nil {
		return nil, errors.WithMessage(err, "Failed to get category's playlists")
	}
	return &output, nil
}

// GetRecommendationsByArtist -
func (c *DefaultClient) GetRecommendationsByArtist(ctx context.Context, artist string) (*GetRecommendationsByArtistOutput, error) {
	var getArtistSearchOutput GetArtistSearchOutput
	if err := c.getSpotifyIDForResource(ctx, artist, "artist", &getArtistSearchOutput); err != nil {
		return nil, errors.WithMessage(err, "Failed to get spotify id for artist")
	}
	queryParams := &map[string]string{
		"seed_artists": getArtistSearchOutput.Inner.Items[0].ID,
		"limit":        "10",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.authToken.AccessToken,
	}
	var output GetRecommendationsByArtistOutput
	if err := c.requestor.Get(ctx, &req.GetInput{
		URL:         "https://api.spotify.com/v1/recommendations",
		QueryParams: queryParams,
		Headers:     headers,
		Destination: &output,
	}); err != nil {
		return nil, errors.WithMessage(err, "Failed to get recommendations by artist")
	}
	return &output, nil
}

// GetNewReleases -
func (c *DefaultClient) GetNewReleases(ctx context.Context) (*GetNewReleasesOutput, error) {
	queryParams := &map[string]string{
		"limit": "50",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.authToken.AccessToken,
	}
	var output GetNewReleasesOutput
	if err := c.requestor.Get(ctx, &req.GetInput{
		URL:         "https://api.spotify.com/v1/browse/new-releases",
		QueryParams: queryParams,
		Headers:     headers,
		Destination: &output,
	}); err != nil {
		return nil, errors.WithMessage(err, "Failed to get new releases")
	}

	return &output, nil
}

// GetAlbum -
func (c *DefaultClient) GetAlbum(ctx context.Context, album string) (*GetAlbumOutput, error) {
	var getAlbumSearchOutput GetAlbumSearchOutput
	if err := c.getSpotifyIDForResource(ctx, album, "album", &getAlbumSearchOutput); err != nil {
		return nil, errors.WithMessage(err, "Failed to get spotify id for album")
	}
	queryParams := &map[string]string{
		"q":     album,
		"limit": "1",
		"type":  "album",
	}
	slugs := &map[string]string{
		"{albumID}": getAlbumSearchOutput.Inner.Items[0].ID,
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.authToken.AccessToken,
	}

	var output GetAlbumOutput
	if err := c.requestor.Get(ctx, &req.GetInput{
		URL:         "https://api.spotify.com/v1/albums/{albumID}",
		Slugs:       slugs,
		QueryParams: queryParams,
		Headers:     headers,
		Destination: &output,
	}); err != nil {
		return nil, errors.WithMessage(err, "Failed to get album")
	}
	return &output, nil
}

// GetAlbumTracks -
func (c *DefaultClient) GetAlbumTracks(ctx context.Context, album string) (*GetAlbumTracksOutput, error) {
	var getAlbumSearchOutput GetAlbumSearchOutput
	if err := c.getSpotifyIDForResource(ctx, album, "album", &getAlbumSearchOutput); err != nil {
		return nil, errors.WithMessage(err, "Failed to get spotify id for album")
	}
	queryParams := &map[string]string{
		"limit": "50",
	}
	slugs := &map[string]string{
		"{albumID}": getAlbumSearchOutput.Inner.Items[0].ID,
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.authToken.AccessToken,
	}

	var output GetAlbumTracksOutput
	if err := c.requestor.Get(ctx, &req.GetInput{
		URL:         "https://api.spotify.com/v1/albums/{albumID}/tracks",
		Slugs:       slugs,
		QueryParams: queryParams,
		Headers:     headers,
		Destination: &output,
	}); err != nil {
		return nil, errors.WithMessage(err, "Failed to get album tracks")
	}
	return &output, nil
}

func (c *DefaultClient) getSpotifyIDForResource(ctx context.Context, resource string, resourceType string, output interface{}) error {
	queryParams := &map[string]string{
		"q":    resource,
		"type": resourceType,
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.authToken.AccessToken,
	}
	if err := c.requestor.Get(ctx, &req.GetInput{
		URL:         "https://api.spotify.com/v1/search",
		QueryParams: queryParams,
		Headers:     headers,
		Destination: output,
	}); err != nil {
		return errors.WithMessage(err, "Failed to search for spotify id")
	}
	return nil
}

func getAccessToken(clientID, clientSecret string) (*GetTokenOutput, error) {
	data := []byte(fmt.Sprintf("%s:%s", clientID, clientSecret))
	str := base64.StdEncoding.EncodeToString(data)

	headers := &map[string]string{
		"Authorization": "Basic " + str,
		"Content-Type":  "application/x-www-form-urlencoded",
	}
	body := &map[string]string{
		"grant_type": "client_credentials",
	}
	var output GetTokenOutput
	if err := req.NewRequestor().Post(context.Background(), &req.PostInput{
		URL:         "https://accounts.spotify.com/api/token",
		Headers:     headers,
		Body:        body,
		Destination: &output,
	}); err != nil {
		return nil, errors.WithMessage(err, "Failed to get access token")
	}
	return &output, nil
}
