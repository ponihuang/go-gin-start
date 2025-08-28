package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHelloWorld(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	if got := w.Body.String(); got != "Hello, World!" {
		t.Fatalf("expected body %q, got %q", "Hello, World!", got)
	}
}
