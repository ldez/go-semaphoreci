/*
Package v2 Access to the API v2 of SemaphoreCI.

Reference: http://semaphoreci.com/docs/api-v2-overview.html
*/
package v2

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const defaultBaseURL = "https://api.semaphoreci.com/v2/"

const (
	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "X-RateLimit-Remaining"
	headerRateReset     = "X-RateLimit-Reset"
)

// Client API v2 client.
type Client struct {
	BaseURL   *url.URL
	UserAgent string
	client    *http.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	ConfigFiles   *ConfigFilesService
	EnvVars       *EnvVarsService
	Organizations *OrganizationsService
	Projects      *ProjectsService
	Secrets       *SecretsService
	Teams         *TeamsService
	Users         *UsersService
}

// NewClient creates a new API v2 client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{BaseURL: baseURL, client: httpClient}

	c.common.client = c
	c.ConfigFiles = (*ConfigFilesService)(&c.common)
	c.EnvVars = (*EnvVarsService)(&c.common)
	c.Organizations = (*OrganizationsService)(&c.common)
	c.Projects = (*ProjectsService)(&c.common)
	c.Secrets = (*SecretsService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	c.Users = (*UsersService)(&c.common)

	return c
}

// NewRequest creates a request.
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do execute a request.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer safeClose(resp.Body)

	response := newResponse(resp)

	err = checkResponse(resp)
	if err != nil {
		return response, err
	}

	if v == nil {
		return response, nil
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, &ErrorResponse{
			Response: resp,
			Message:  fmt.Sprintf("failed to read body: %v", err),
		}
	}

	err = json.Unmarshal(raw, v)
	if err != nil {
		return response, &ErrorResponse{
			Response: resp,
			Message:  fmt.Sprintf("unmarshaling error: %v: %s", err, string(raw)),
		}
	}

	return response, nil
}

func checkResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: resp}
	data, err := io.ReadAll(resp.Body)
	if err == nil && data != nil {
		errJSON := json.Unmarshal(data, errorResponse)
		if errJSON != nil {
			return &ErrorResponse{
				Response: resp,
				Message:  fmt.Sprintf("unmarshaling error: %v: %s", errJSON.Error(), string(data)),
			}
		}
	}
	defer safeClose(resp.Body)

	if resp.StatusCode == http.StatusForbidden && resp.Header.Get(headerRateRemaining) == "0" {
		return &RateLimitError{
			Rate:     parseRate(resp),
			Response: errorResponse.Response,
			Message:  errorResponse.Message,
		}
	}

	return errorResponse
}

func newResponse(resp *http.Response) *Response {
	response := &Response{Response: resp}
	response.populatePageValues()
	response.Rate = parseRate(resp)
	return response
}

func parseRate(r *http.Response) Rate {
	var rate Rate
	if limit := r.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}
	if reset := r.Header.Get(headerRateReset); reset != "" {
		rate.Reset, _ = time.Parse(time.RFC3339, reset)
	}
	return rate
}

func safeClose(elt io.Closer) {
	err := elt.Close()
	if err != nil {
		log.Println(err)
	}
}
