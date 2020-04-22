package spotify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"strings"

	httputil "github.com/cwseger/spotify-cli/http_util"

	"github.com/pkg/errors"
)

// Client -
type Client interface {
	GetArtist(ctx context.Context, artist ...string)
	GetCategoryList(ctx context.Context) (*GetCategoriesOutput, error)
	GetCategoryPlaylists(ctx context.Context, categoryID string) (*GetCategoryPlaylistsOutput, error)
	GetRecommendationsByArtists(ctx context.Context, artists string) (*GetRecommendationsByArtistOutput, error)
	GetNewReleases(ctx context.Context) (*GetNewReleasesOutput, error)
	GetToken() (*TokenResponse, error)
}

// DefaultClient -
type DefaultClient struct {
	clientID     string
	clientSecret string
	requestor    httputil.Requestor
}

// NewClient -
func NewClient() (*DefaultClient, error) {
	clientSecretsFile, err := os.Open("client-secrets.json")
	if err != nil {
		fmt.Println("Failed to open secrets file")
		return nil, err
	}
	defer clientSecretsFile.Close()
	var secrets ClientSecrets
	bytesValue, _ := ioutil.ReadAll(clientSecretsFile)
	err = json.Unmarshal(bytesValue, &secrets)
	if err != nil {
		fmt.Println("Failed to unmarshal secrets")
		return nil, err
	}
	return &DefaultClient{
		clientID:     secrets.ClientID,
		clientSecret: secrets.ClientSecret,
		requestor:    httputil.NewRequestor(),
	}, nil
}

// GetArtist -
func (c *DefaultClient) GetArtist(ctx context.Context, artist string) (*GetArtistOutput, error) {
	tokenResponse, err := c.getAccessToken()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get access token")
	}

	queryParams := &map[string]string{
		"q":     artist,
		"limit": "1",
		"type":  "artist",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + tokenResponse.AccessToken,
	}
	var output GetArtistOutput
	if err := c.requestor.Get(ctx, &httputil.GetInput{
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
	tokenResponse, err := c.getAccessToken()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get access token")
	}

	queryParams := &map[string]string{
		"limit": "50",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + tokenResponse.AccessToken,
	}
	var output GetCategoriesOutput
	if err := c.requestor.Get(ctx, &httputil.GetInput{
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
	tokenResponse, err := c.getAccessToken()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get access token")
	}

	slugs := &map[string]string{
		"{categoryID}": categoryID,
	}
	queryParams := &map[string]string{
		"limit": "5",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + tokenResponse.AccessToken,
	}
	var output GetCategoryPlaylistsOutput
	if err := c.requestor.Get(ctx, &httputil.GetInput{
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
	tokenResponse, err := c.getAccessToken()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get access token")
	}

	queryParams := &map[string]string{
		"seed_artists": artists[0],
		"limit":        "3",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + tokenResponse.AccessToken,
	}
	var output GetRecommendationsByArtistOutput
	if err := c.requestor.Get(ctx, &httputil.GetInput{
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
	tokenResponse, err := c.getAccessToken()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get access token")
	}

	queryParams := &map[string]string{
		"limit": "50",
	}
	headers := &map[string]string{
		"Authorization": "Bearer " + tokenResponse.AccessToken,
	}
	var output GetNewReleasesOutput
	if err := c.requestor.Get(ctx, &httputil.GetInput{
		URL:         "https://api.spotify.com/v1/browse/new-releases",
		QueryParams: queryParams,
		Headers:     headers,
		Destination: &output,
	}); err != nil {
		return nil, errors.WithMessage(err, "Failed to get new releases")
	}

	return &output, nil
}

// GetToken -
func (c *DefaultClient) GetToken() (*TokenResponse, error) {
	return c.getAccessToken()
}

func (c *DefaultClient) getAccessToken() (*TokenResponse, error) {
	// fmt.Println("Getting access token")
	data := []byte(fmt.Sprintf("%s:%s", c.clientID, c.clientSecret))
	str := base64.StdEncoding.EncodeToString(data)
	// fmt.Println(fmt.Sprintf("Encoded data: %+v", str))

	// header := &map[string]string{
	// 	"Authorization": "Basic " + str,
	// 	"Content-Type":  "application/x-www-form-urlencoded",
	// }
	// body := &map[string]string{
	// 	"grant_type": "client_credentials",
	// }
	bodyValues := url.Values{}
	bodyValues.Set("grant_type", "client_credentials")
	tokenReq, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(bodyValues.Encode()))
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to build token request. Err: %+v", err))
		return nil, err
	}
	tokenReq.Header.Add("Authorization", "Basic "+str)
	tokenReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	httpClient := http.Client{}
	tokenResp, err := httpClient.Do(tokenReq)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to do token request. Err: %+v", err))
		return nil, err
	}
	defer tokenResp.Body.Close()
	tokenBody, err := ioutil.ReadAll(tokenResp.Body)
	var resp TokenResponse
	err = json.Unmarshal(tokenBody, &resp)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to unmarshal token response. Err: %+v", err))
	}
	// fmt.Println(fmt.Sprintf("Token resp body: (%+v) %+v (%+v)", tokenResp.Status, string(tokenBody), resp))
	return &resp, nil
}
