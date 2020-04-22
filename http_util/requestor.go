package http_util

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// Requestor -
type Requestor interface {
	Get(ctx context.Context, input *GetInput) error
	// Post() error
}

// DefaultRequestor -
type DefaultRequestor struct {
	httpClient http.Client
}

// NewRequestor -
func NewRequestor() *DefaultRequestor {
	return &DefaultRequestor{
		httpClient: http.Client{},
	}
}

var _ Requestor = &DefaultRequestor{}

// Get -
func (r *DefaultRequestor) Get(ctx context.Context, input *GetInput) error {
	url, err := r.replaceSlugsWithValues(input.URL, input.Slugs)
	if err != nil {
		return errors.WithMessage(err, "Failed to replace slugs with values")
	}
	url, err = r.addQueryParamsToURL(url, input.QueryParams)
	if err != nil {
		return errors.WithMessage(err, "Failed to add query params to url")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return errors.WithMessage(err, "Failed to build new request with context")
	}
	r.addHeadersToRequest(input.Headers, req)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return errors.WithMessage(err, "Failed to execute request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithMessage(err, "Failed to read response body")
	}

	if err := json.Unmarshal(body, input.Destination); err != nil {
		return errors.WithMessage(err, "Failed to unmarshal response body")
	}
	return nil
}

func (r *DefaultRequestor) replaceSlugsWithValues(rawURL string, slugs *map[string]string) (string, error) {
	if slugs == nil {
		return rawURL, nil
	}
	for k, v := range *slugs {
		rawURL = strings.ReplaceAll(rawURL, k, v)
	}
	return rawURL, nil
}

func (r *DefaultRequestor) addQueryParamsToURL(rawURL string, queryParams *map[string]string) (string, error) {
	if queryParams == nil {
		return rawURL, nil
	}
	if len(*queryParams) == 0 {
		return rawURL, nil
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return "", errors.WithMessage(err, "Failed to parse request uri")
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", errors.WithMessage(err, "Failed to parse raw query")
	}

	for k, v := range *queryParams {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (r *DefaultRequestor) addHeadersToRequest(headers *map[string]string, req *http.Request) {
	if headers == nil {
		return
	}
	for k, v := range *headers {
		req.Header.Add(k, v)
	}
}

// // Post -
// func (r *DefaultRequestor) Post(ctx context.Context, url string, slugs *map[string]string, headers *map[string]string, body *map[string]string, destination interface{}) error {
// 	url, err := r.replaceSlugsWithValues(url, slugs)
// 	if err != nil {
// 		return errors.WithMessage(err, "Failed to replace slugs with values")
// 	}

// 	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
// 	if err != nil {
// 		return errors.WithMessage(err, "Failed to build new request with context")
// 	}
// 	r.addHeadersToRequest(headers, req)

// 	resp, err := r.httpClient.Do(req)
// 	if err != nil {
// 		return errors.WithMessage(err, "Failed to execute request")
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return errors.WithMessage(err, "Failed to read response body")
// 	}

// 	if err := json.Unmarshal(body, destination); err != nil {
// 		return errors.WithMessage(err, "Failed to unmarshal response body")
// 	}
// 	return nil
// }
