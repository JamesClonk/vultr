package lib

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestServer(code int, body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, body)
	}))
}

func getTestClient(t *testing.T, endpoint string) *Client {
	options := Options{
		Endpoint: endpoint,
	}
	client, err := NewClient("test-key", &options)
	if err != nil {
		t.Fatal(err)
	}
	return client
}
