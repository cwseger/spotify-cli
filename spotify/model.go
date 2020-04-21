package spotify

// ClientSecrets -
type ClientSecrets struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// TokenResponse -
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

// Category -
type Category struct {
	Href string `json:"href"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetCategoriesInner -
type GetCategoriesInner struct {
	Items    []Category `json:"items"`
	Limit    int        `json:"limit"`
	Next     string     `json:"next"`
	Offset   int        `json:"offset"`
	Previous *string    `json:"previous"`
	Total    int        `json:"total"`
}

// GetCategoriesOutput -
type GetCategoriesOutput struct {
	Inner GetCategoriesInner `json:"categories"`
}

// Playlist -
type Playlist struct {
	Name          string `json:"name"`
	URI           string `json:"uri"`
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
}

// GetCategoryPlaylistsInner -
type GetCategoryPlaylistsInner struct {
	Items []Playlist `json:"items"`
}

// GetCategoryPlaylistsOutput -
type GetCategoryPlaylistsOutput struct {
	Inner GetCategoryPlaylistsInner `json:"playlists"`
}

// Album -
type Album struct {
	Name string `json:"name"`
}

// Track -
type Track struct {
	Album Album `json:"album"`
}

// GetRecommendationsByArtistInner -
type GetRecommendationsByArtistInner struct {
	Tracks []Track `json:"tracks"`
}

// GetRecommendationsByArtistOutput -
type GetRecommendationsByArtistOutput struct {
	Tracks []Track `json:"tracks"`
}

// GetNewReleasesInner -
type GetNewReleasesInner struct {
	Items []Album `json:"items"`
}

// GetNewReleasesOutput -
type GetNewReleasesOutput struct {
	Inner GetNewReleasesInner `json:"albums"`
}
