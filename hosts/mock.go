package hosts

import "net/http"

// MockClient implements the Client interface
type MockClient struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

var _ Client = &MockClient{}

// Do executes the DoFunc
func (c *MockClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}
