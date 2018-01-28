package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

const (
	apiURL = "https://semaphoreci.com/api/v1"
)

// Client API v1 client
type Client struct {
	URLApi    string
	AuthToken string
}

// NewClient create a new API v1 client
func NewClient(authToken string) *Client {
	return &Client{
		URLApi:    apiURL,
		AuthToken: authToken,
	}
}

func (c Client) newRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	u := fmt.Sprintf("%s/%s?auth_token=%s", c.URLApi, urlStr, c.AuthToken)
	return http.NewRequest(method, u, body)
}

func do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = checkResponse(resp)
	if err != nil {
		return nil, errors.Cause(err)
	}

	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		defer safeClose(resp.Body)

		err = decoder.Decode(v)
		if err != nil {
			return nil, err
		}
	}

	return newResponse(resp), nil
}

func checkResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 > c || c > 299 {
		errorResponse := &ErrorResponse{Response: resp}
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil && data != nil {
			errJSON := json.Unmarshal(data, errorResponse)
			if errJSON != nil {
				log.Println(errJSON)
			}
		}
		defer safeClose(resp.Body)
		return errorResponse
	}

	return nil
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
