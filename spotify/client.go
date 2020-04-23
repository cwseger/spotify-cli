package spotify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"strings"

	req "github.com/cwseger/spotify-cli/req"

	"github.com/pkg/errors"
)

// Client -
type Client interface {
	GetArtist(ctx context.Context, artist []string)
	GetCategoryList(ctx context.Context) (*GetCategoriesOutput, error)
	GetCategoryPlaylists(ctx context.Context, categoryID string) (*GetCategoryPlaylistsOutput, error)
	GetRecommendationsByArtists(ctx context.Context, artists string) (*GetRecommendationsByArtistOutput, error)
	GetNewReleases(ctx context.Context) (*GetNewReleasesOutput, error)
	GetToken(ctx context.Context) (*GetTokenOutput, error)
}

// DefaultClient -
type DefaultClient struct {
	token     *GetTokenOutput
	requestor req.Requestor
}

// NewClient -
func NewClient() (*DefaultClient, error) {
	clientSecretsFile, err := os.Open("client-secrets.json")
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to open secrets file")
	}
	defer clientSecretsFile.Close()

	var secrets ClientSecrets
	bytesValue, _ := ioutil.ReadAll(clientSecretsFile)
	err = json.Unmarshal(bytesValue, &secrets)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to unmarshal client secrets file")
	}

	getTokenOutput, err := getAccessToken(secrets.ClientID, secrets.ClientSecret)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get access token")
	}

	return &DefaultClient{
		token:     getTokenOutput,
		requestor: req.NewRequestor(),
	}, nil
}

// GetArtist -
func (c *DefaultClient) GetArtist(ctx context.Context, artist []string) (*GetArtistOutput, error) {
	queryParams := &map[string]string{
		"q":     strings.Join(artist, ""),
		"limit": "1",
		"type":  "artist",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.token.AccessToken,
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

// GetCategoryList -
func (c *DefaultClient) GetCategoryList(ctx context.Context) (*GetCategoriesOutput, error) {
	queryParams := &map[string]string{
		"limit": "50",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.token.AccessToken,
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
		"Authorization": "Bearer " + c.token.AccessToken,
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

// GetRecommendationsByArtists -
func (c *DefaultClient) GetRecommendationsByArtists(ctx context.Context, artists ...string) (*GetRecommendationsByArtistOutput, error) {
	queryParams := &map[string]string{
		"seed_artists": artists[0],
		"limit":        "3",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + c.token.AccessToken,
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
		"Authorization": "Bearer " + c.token.AccessToken,
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
