package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Branch a branch.
type Branch struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BranchURL string `json:"branch_url"`
}

// BranchStatus a branch status.
type BranchStatus struct {
	BranchName       string     `json:"branch_name"`
	BranchURL        string     `json:"branch_url"`
	BranchStatusURL  string     `json:"branch_status_url"`
	BranchHistoryURL string     `json:"branch_history_url"`
	ProjectName      string     `json:"project_name"`
	BuildURL         string     `json:"build_url"`
	BuildInfoURL     string     `json:"build_info_url"`
	BuildNumber      int        `json:"build_number,omitempty"`
	Result           string     `json:"result"`
	StartedAt        *time.Time `json:"started_at"`
	FinishedAt       *time.Time `json:"finished_at"`
	Commit           Commit     `json:"commit"`
}

// BranchHistoryOptions the branch history pagination options.
type BranchHistoryOptions struct {
	Page int
}

// BranchHistory a branch history.
type BranchHistory struct {
	BranchName       string  `json:"branch_name"`
	BranchURL        string  `json:"branch_url"`
	BranchStatusURL  string  `json:"branch_status_url"`
	BranchHistoryURL string  `json:"branch_history_url"`
	ProjectName      string  `json:"project_name"`
	Builds           []Build `json:"builds"`
}

// Build a build.
type Build struct {
	BuildURL     string     `json:"build_url"`
	BuildInfoURL string     `json:"build_info_url"`
	BuildNumber  int        `json:"build_number"`
	Result       string     `json:"result"`
	StartedAt    *time.Time `json:"started_at"`
	FinishedAt   *time.Time `json:"finished_at"`
	Commit       Commit     `json:"commit"`
}

// BuildInformation information about the build.
type BuildInformation struct {
	Commits      []Commit   `json:"commits"`
	ProjectName  string     `json:"project_name"`
	BranchName   string     `json:"branch_name"`
	Number       int        `json:"number,omitempty"`
	Result       string     `json:"result"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	StartedAt    *time.Time `json:"started_at"`
	FinishedAt   *time.Time `json:"finished_at"`
	HTMLURL      string     `json:"html_url"`
	BuildLogURL  string     `json:"build_log_url"`
	BuildInfoURL string     `json:"build_info_url"`
}

// Commit a commit.
type Commit struct {
	ID          string     `json:"id"`
	URL         string     `json:"url"`
	AuthorName  string     `json:"author_name"`
	AuthorEmail string     `json:"author_email"`
	Message     string     `json:"message"`
	Timestamp   *time.Time `json:"timestamp" `
}

// BuildLog a build log.
type BuildLog struct {
	Threads      []Thread `json:"threads"`
	BuildInfoURL string   `json:"build_info_url"`
}

// DeployLog a deploy log.
type DeployLog struct {
	Threads       []Thread `json:"threads"`
	DeployInfoURL string   `json:"deploy_info_url"`
}

// Thread a thread.
type Thread struct {
	Number   int       `json:"number"`
	Commands []Command `json:"commands"`
}

// Command a command.
type Command struct {
	Name       string     `json:"name"`
	Result     string     `json:"result"`
	Output     string     `json:"output"`
	StartTime  *time.Time `json:"start_time"`
	FinishTime *time.Time `json:"finish_time"`
	Duration   string     `json:"duration"`
}

// Project a project.
type Project struct {
	ID        int             `json:"id"`
	HashID    string          `json:"hash_id"`
	Name      string          `json:"name"`
	Owner     string          `json:"owner"`
	HTMLURL   string          `json:"html_url"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
	Branches  []ProjectBranch `json:"branches"`
	Servers   []ServerStatus  `json:"servers"`
}

// ProjectBranch a branch in project.
type ProjectBranch struct {
	Name         string     `json:"branch_name"`
	URL          string     `json:"branch_url"`
	StatusURL    string     `json:"branch_status_url"`
	HistoryURL   string     `json:"branch_history_url"`
	ProjectName  string     `json:"project_name"`
	BuildURL     string     `json:"build_url"`
	BuildInfoURL string     `json:"build_info_url"`
	BuildNumber  int        `json:"build_number"`
	Result       string     `json:"result"`
	StartedAt    *time.Time `json:"started_at"`
	FinishedAt   *time.Time `json:"finished_at"`
}

// Server a server.
type Server struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ServerURL string `json:"server_url"`
}

// ServerStatus a server.
type ServerStatus struct {
	ServerName       string     `json:"server_name"`
	ServerURL        string     `json:"server_url"`
	ServerStatusURL  string     `json:"server_status_url"`
	ServerHistoryURL string     `json:"server_history_url"`
	DeploymentMethod string     `json:"deployment_method"`
	Strategy         string     `json:"strategy"`
	BranchName       string     `json:"branch_name"`
	ProjectName      string     `json:"project_name"`
	Number           int        `json:"number"`
	Result           string     `json:"result"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	StartedAt        *time.Time `json:"started_at"`
	FinishedAt       *time.Time `json:"finished_at"`
	HTMLURL          string     `json:"html_url"`
	DeployURL        string     `json:"deploy_url"`
	DeployLogURL     string     `json:"deploy_log_url"`
	BuildURL         string     `json:"build_url"`
	BuildHTMLURL     string     `json:"build_html_url"`
	Commit           *Commit    `json:"commit"`
}

// ServerDeploy a server deploy.
type ServerDeploy struct {
	ServerName       string   `json:"server_name"`
	ServerURL        string   `json:"server_url"`
	ServerStatusURL  string   `json:"server_status_url"`
	ServerHistoryURL string   `json:"server_history_url"`
	DeploymentMethod string   `json:"deployment_method"`
	Strategy         string   `json:"strategy"`
	BranchName       string   `json:"branch_name"`
	ProjectName      string   `json:"project_name"`
	Deploys          []Deploy `json:"deploys"`
}

// ServerHistoryOptions the server history pagination options.
type ServerHistoryOptions struct {
	Page int
}

// Deploy a deploy.
type Deploy struct {
	ProjectName  string     `json:"project_name"`
	ServerName   string     `json:"server_name"`
	Number       int        `json:"number"`
	Result       string     `json:"result"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	StartedAt    *time.Time `json:"started_at"`
	FinishedAt   *time.Time `json:"finished_at"`
	HTMLURL      string     `json:"html_url"`
	DeployURL    string     `json:"deploy_url"`
	DeployLogURL string     `json:"deploy_log_url"`
	BuildURL     string     `json:"build_url"`
	BuildHTMLURL string     `json:"build_html_url"`
	Commit       *Commit    `json:"commit"`
}

// WebHook a WebHook.
type WebHook struct {
	ID       int    `json:"id,omitempty"`
	URL      string `json:"url"`
	HookType string `json:"hook_type"`
}

// Pagination Pagination headers.
type Pagination struct {
	TotalEntries int  `json:"total_entries,omitempty"`
	TotalPages   int  `json:"total_pages,omitempty"`
	PerPage      int  `json:"per_page,omitempty"`
	CurrentPage  int  `json:"current_page,omitempty"`
	FirstPage    bool `json:"first_page"`
	LastPage     bool `json:"last_page"`
	NextPage     int  `json:"next_page,omitempty"`
	PreviousPage int  `json:"previous_page,omitempty"`
}

type service struct {
	client *Client
}

// Response API response.
type Response struct {
	*http.Response
	*Pagination
}

func (r *Response) populatePageValues() error {
	rawPagination := r.Header.Get("Pagination")
	fmt.Println(rawPagination)
	if len(rawPagination) > 0 {
		return json.Unmarshal([]byte(rawPagination), r)
	}
	return nil
}

// ErrorResponse API error response.
type ErrorResponse struct {
	Response *http.Response `json:"-"`     // HTTP response that caused this error
	Message  string         `json:"error"` // error message
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
