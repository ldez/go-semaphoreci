package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://semaphoreci.com/api/v1/"
)

// Client API v1 client
type Client struct {
	BaseURL   *url.URL
	UserAgent string
	client    *http.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Branch   *BranchesService
	Builds   *BuildsService
	Projects *ProjectsService
	Servers  *ServersService
	Deploys  *DeploysService
	WebHooks *WebHooksService
}

// NewClient create a new API v1 client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{BaseURL: baseURL, client: httpClient}

	c.common.client = c
	c.Branch = (*BranchesService)(&c.common)
	c.Builds = (*BuildsService)(&c.common)
	c.Projects = (*ProjectsService)(&c.common)
	c.Servers = (*ServersService)(&c.common)
	c.Deploys = (*DeploysService)(&c.common)
	c.WebHooks = (*WebHooksService)(&c.common)

	return c
}

// NewRequest creates a request
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(method, u.String(), body)
}

// Do execute a request
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer safeClose(resp.Body)

	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

	if v != nil {
		raw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, &ErrorResponse{
				Response: resp,
				Message:  fmt.Sprintf("faild to read body: %v", err),
			}
		}

		err = json.Unmarshal(raw, v)
		if err != nil {
			return nil, &ErrorResponse{
				Response: resp,
				Message:  fmt.Sprintf("unmarshaling error: %v: %s", err, string(raw)),
			}
		}
	}

	return newResponse(resp), nil
}

func checkResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: resp}
	data, err := ioutil.ReadAll(resp.Body)
	if err == nil && data != nil {
		errJSON := json.Unmarshal(data, errorResponse)
		if errJSON != nil {
			errorResponse.Message = fmt.Sprintf("unmarshaling error: %v: %s", errJSON.Error(), string(data))
		}
	}
	return errorResponse
}

func newResponse(resp *http.Response) *Response {
	response := &Response{Response: resp}

	err := response.populatePageValues()
	if err != nil {
		log.Println(err)
	}

	return response
}

func safeClose(elt io.Closer) {
	err := elt.Close()
	if err != nil {
		log.Println(err)
	}
}
