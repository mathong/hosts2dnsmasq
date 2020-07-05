package hosts

import "net/http"

// Client is just here to enable mocking HTTP calls
type Client interface {
	Do(*http.Request) (*http.Response, error)
}
