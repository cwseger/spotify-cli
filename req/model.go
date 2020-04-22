package req

// GetInput -
type GetInput struct {
	URL         string
	Slugs       *map[string]string
	QueryParams *map[string]string
	Headers     *map[string]string
	Destination interface{}
}

// PostInput -
type PostInput struct {
	URL         string
	Slugs       *map[string]string
	Headers     *map[string]string
	Body        *map[string]string
	Destination interface{}
}
