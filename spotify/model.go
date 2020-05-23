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

// GetArtistAlbumOutput -
type GetArtistAlbumOutput struct {
	Albums []Album `json:"items"`
}

// GetArtistOutput -
type GetArtistOutput struct {
	Inner GetArtistInner `json:"artists"`
}

// GetAlbumTrack -
type GetAlbumTrack struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	TrackNumber int    `json:"track_number"`
	DurationMS  int    `json:"duration_ms"`
	PreviewURL  string `json:"preview_url"`
}

// GetAlbumTracksOutput -
type GetAlbumTracksOutput struct {
	Tracks []GetAlbumTrack `json:"items"`
	Total  int             `json:"total"`
}

// GetAlbumOutput -
type GetAlbumOutput struct {
	Name       string   `json:"name"`
	Popularity int      `json:"popularity"`
	Artists    []Artist `json:"artists"`
}

// GetAlbumInner -
type GetAlbumInner struct {
}

// GetArtistSearchOutput -
type GetArtistSearchOutput struct {
	Inner GetSearchInner `json:"artists"`
}

// GetAlbumSearchOutput -
type GetAlbumSearchOutput struct {
	Inner GetSearchInner `json:"albums"`
}

// GetSearchInner -
type GetSearchInner struct {
	Items []SearchItem `json:"items"`
}

// SearchItem -
type SearchItem struct {
	ID string `json:"id"`
}
