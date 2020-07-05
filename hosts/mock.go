package hosts

import "net/http"

type MockClient struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

var _ Client = &MockClient{}

func (c *MockClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}
