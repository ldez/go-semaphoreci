package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// hook type
const (
	PostBuild  = "post_build"
	PostDeploy = "post_deploy"
	All        = "all"
)

// WebHooksService https://semaphoreci.com/docs/webhooks-api.html
type WebHooksService service

// GetByProject https://semaphoreci.com/docs/webhooks-api.html#list_hooks
func (c *WebHooksService) GetByProject(projectHashID string) ([]WebHook, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/hooks", projectHashID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]WebHook)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// Create https://semaphoreci.com/docs/webhooks-api.html#create_hook
func (c *WebHooksService) Create(projectHashID string, hook WebHook) (*WebHook, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/hooks", projectHashID)

	body, err := json.Marshal(hook)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest(http.MethodPost, urlStr, bytes.NewReader(body))
	if err != nil {
		return nil, nil, err
	}

	v := new(WebHook)

	resp, err := c.client.Do(req, v)

	return v, resp, err
}

// Update https://semaphoreci.com/docs/webhooks-api.html#update_hook
func (c *WebHooksService) Update(projectHashID string, webHookID int, hook WebHook) (*WebHook, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/hooks/%v", projectHashID, webHookID)

	body, err := json.Marshal(hook)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest(http.MethodPut, urlStr, bytes.NewReader(body))
	if err != nil {
		return nil, nil, err
	}

	v := new(WebHook)

	resp, err := c.client.Do(req, v)

	return v, resp, err
}

// Delete https://semaphoreci.com/docs/webhooks-api.html#delete_hook
func (c *WebHooksService) Delete(projectHashID string, webHookID int) (*Response, error) {
	urlStr := fmt.Sprintf("projects/%s/hooks/%v", projectHashID, webHookID)

	req, err := c.client.NewRequest(http.MethodDelete, urlStr, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req, nil)
}
