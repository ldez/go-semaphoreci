package v1

import (
	"net/http"
)

// TokenTransport HTTP transport for API authentication
type TokenTransport struct {
	Token string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip executes a single HTTP transaction
func (t *TokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	enrichedReq := &http.Request{}
	*enrichedReq = *req

	enrichedReq.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		enrichedReq.Header[k] = append([]string(nil), s...)
	}

	if t.Token != "" {
		query := enrichedReq.URL.Query()
		query.Add("auth_token", t.Token)
		enrichedReq.URL.RawQuery = query.Encode()
	}

	return t.transport().RoundTrip(enrichedReq)
}

// Client Creates a new HTTP client
func (t *TokenTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *TokenTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
