package v2

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Project An API project representation.
type Project struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	HTMLURL    string     `json:"html_url"`
	UsersURL   string     `json:"users_url"`
	SecretsURL string     `json:"secrets_url"`
	UpdatedAt  *time.Time `json:"updated_at"`
	CreatedAt  *time.Time `json:"created_at"`
}

// Organization An API organization representation.
type Organization struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	URL         string     `json:"url"`
	ProjectsURL string     `json:"projects_url"`
	SecretsURL  string     `json:"secrets_url"`
	UsersURL    string     `json:"users_url"`
	TeamsURL    string     `json:"teams_url"`
	Username    string     `json:"username"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// ConfigFile An API configuration file representation.
type ConfigFile struct {
	ID        string     `json:"id"`
	Path      string     `json:"path"`
	URL       string     `json:"url"`
	Content   string     `json:"content"`
	Shared    bool       `json:"shared"`
	Encrypted bool       `json:"encrypted"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// EnvVar An API environment variables representation.
type EnvVar struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	URL       string     `json:"url"`
	Content   string     `json:"content"`
	Shared    bool       `json:"shared"`
	Encrypted bool       `json:"encrypted"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// Secret An API secret representation.
type Secret struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	URL            string     `json:"url"`
	ProjectsURL    string     `json:"projects_url"`
	TeamsURL       string     `json:"teams_url"`
	EnvVarsURL     string     `json:"env_vars_url"`
	ConfigFilesURL string     `json:"config_files_url"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

// Team An API team representation.
type Team struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	URL         string     `json:"url"`
	Permission  string     `json:"permission"`
	UsersURL    string     `json:"users_url"`
	ProjectsURL string     `json:"projects_url"`
	SecretsURL  string     `json:"secrets_url"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// User An API user representation.
type User struct {
	UID       string     `json:"uid"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type service struct {
	client *Client
}

// ErrorResponse An API error representation.
type ErrorResponse struct {
	Response         *http.Response    `json:"-"`
	Message          string            `json:"message,omitempty"`
	DocumentationURL string            `json:"documentation_url,omitempty"`
	Errors           []ErrorValidation `json:"errors,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

// ErrorValidation An API validation error representation.
type ErrorValidation struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

// Response API response.
type Response struct {
	*http.Response

	PerPage   int
	Total     int
	NextPage  int
	PrevPage  int
	FirstPage int
	LastPage  int

	Rate
}

func (r *Response) populatePageValues() {
	if raw, ok := r.Response.Header["Total"]; ok {
		r.Total, _ = strconv.Atoi(raw[0])
	}

	if raw, ok := r.Response.Header["Per-Page"]; ok {
		r.PerPage, _ = strconv.Atoi(raw[0])
	}

	if links, ok := r.Response.Header["Link"]; ok && len(links) > 0 {
		for _, link := range strings.Split(links[0], ",") {
			r.loadPaginatedLink(link)
		}
	}
}

func (r *Response) loadPaginatedLink(link string) {
	segments := strings.Split(strings.TrimSpace(link), ";")

	// link must at least have href and rel
	if len(segments) < 2 {
		return
	}

	// ensure href is properly formatted
	if !strings.HasPrefix(segments[0], "<") || !strings.HasSuffix(segments[0], ">") {
		return
	}

	// try to pull out page parameter
	u, err := url.Parse(segments[0][1 : len(segments[0])-1])
	if err != nil {
		return
	}
	page := u.Query().Get("page")
	if page == "" {
		return
	}

	for _, segment := range segments[1:] {
		switch strings.TrimSpace(segment) {
		case `rel="next"`:
			r.NextPage, _ = strconv.Atoi(page)
		case `rel="prev"`:
			r.PrevPage, _ = strconv.Atoi(page)
		case `rel="first"`:
			r.FirstPage, _ = strconv.Atoi(page)
		case `rel="last"`:
			r.LastPage, _ = strconv.Atoi(page)
		}
	}
}

// Rate a rate limit representation.
type Rate struct {
	Limit     int       `json:"limit"`
	Remaining int       `json:"remaining"`
	Reset     time.Time `json:"reset"`
}

// RateLimitError An API rate limit error representation.
type RateLimitError struct {
	Rate     Rate
	Response *http.Response
	Message  string `json:"message"`
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Rate.Reset)
}
