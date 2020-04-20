package http_util

// GetInput -
type GetInput struct {
	URL         string
	Slugs       *map[string]string
	QueryParams *map[string]string
	Headers     *map[string]string
	Destination interface{}
}
