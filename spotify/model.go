package spotify

// ClientSecrets -
type ClientSecrets struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// GetTokenOutput -
type GetTokenOutput struct {
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

// Followers -
type Followers struct {
	Total int `json:"total"`
}

// Artist -
type Artist struct {
	Name       string    `json:"name"`
	Popularity int       `json:"popularity"`
	Followers  Followers `json:"followers"`
}

// GetArtistInner -
type GetArtistInner struct {
	Artists []Artist `json:"items"`
}

// GetArtistOutput -
type GetArtistOutput struct {
	Inner GetArtistInner `json:"artists"`
}

// GetAlbumOutput -
type GetAlbumOutput struct {
}

// GetAlbumTracksOutput -
type GetAlbumTracksOutput struct {
}

// GetSearchOutput -
type GetSearchOutput struct {
	Inner GetSearchInner `json:"artists"`
}

// GetSearchInner -
type GetSearchInner struct {
	Items []SearchItem `json:"items"`
}

// SearchItem -
type SearchItem struct {
	ID string `json:"id"`
}
