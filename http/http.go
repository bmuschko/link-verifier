package http

import (
	"net/http"
	"net/url"
	"time"
)

type HTTP struct {
	client *http.Client
}

// SetTimeout sets timeout for HTTP requests in seconds.
func (h *HTTP) SetTimeout(timeout int) {
	h.client.Timeout = time.Duration(int(time.Second) * timeout)
}

// GetTimeout gets timeout for HTTP requests as duration.
func (h *HTTP) GetTimeout() time.Duration {
	return h.client.Timeout
}

// Get emits a HTTP HEAD request for a given URL. Captures the status code, status and outcome of the call.
// Returns with information about the response.
func (h *HTTP) Head(link string) HttpResponse {
	return sendRequest(link, func(url *url.URL) (resp *http.Response, err error) {
		return h.client.Head(url.String())
	})
}

// Get emits a HTTP GET request for a given URL. Captures the status code, status and outcome of the call.
// Returns with information about the response.
func (h *HTTP) Get(link string) HttpResponse {
	return sendRequest(link, func(url *url.URL) (resp *http.Response, err error) {
		return h.client.Get(url.String())
	})
}

func sendRequest(link string, req func(*url.URL) (resp *http.Response, err error)) HttpResponse {
	result := HttpResponse{Url: link}
	url, err := url.ParseRequestURI(link)
	if err != nil {
		result.Error = err
		return result
	}

	resp, err := req(url)
	if err != nil {
		result.Error = err
		return result
	}

	result.StatusCode = resp.StatusCode
	result.Status = resp.Status

	if resp.StatusCode == 200 {
		result.Success = true
	}

	resp.Body.Close()
	return result
}

// HttpResponse represents HTTP response information.
type HttpResponse struct {
	Url        string
	Success    bool
	StatusCode int
	Status     string
	Error      error
}

func NewHTTP() *HTTP {
	return &HTTP{client: &http.Client{}}
}
