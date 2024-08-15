package server

import (
	"net/http"
	"net/url"
	"testing"
)

func TestNew(t *testing.T) {
	server := New()

	if server.Addr != ":8080" {
		t.Errorf("Expected server address to be :8080, got %s", server.Addr)
	}

	if server.Handler == nil {
		t.Error("Expected server handler to be set, but it's nil")
	}

	mux, ok := server.Handler.(*http.ServeMux)
	if !ok {
		t.Error("Expected server handler to be of type *http.ServeMux")
	}

	routes := []string{"/health", "/"}
	for _, route := range routes {
		if _, pattern := mux.Handler(&http.Request{URL: &url.URL{Path: route}}); pattern != route {
			t.Errorf("Expected route %s to be registered, but it's not", route)
		}
	}
}
