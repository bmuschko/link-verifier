package http

import (
	"net/http"
	"net/url"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

// Get emits a HTTP GET request for a given URL. Captures the status code, status and outcome of the call.
// Returns with information about the response.
func Get(link string) HttpResponse {
	result := HttpResponse{Url: link, Success: true}
	url, err := url.ParseRequestURI(link)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get(url.String())

	if err != nil {
		panic(err)
	}

	result.StatusCode = resp.StatusCode
	result.Status = resp.Status

	if resp.StatusCode != 200 {
		result.Success = false
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
}
