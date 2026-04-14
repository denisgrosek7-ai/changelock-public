package audit

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHTTPExceptionClientSendsBearerToken(t *testing.T) {
	var gotAuthorization string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuthorization = r.Header.Get("Authorization")
		if r.URL.Path != "/v1/exceptions/validate" {
			t.Fatalf("unexpected path %q", r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(ExceptionValidationResult{Valid: true})
	}))
	defer server.Close()

	client := NewHTTPExceptionClient(server.URL, time.Second, "service-internal-demo-token")
	if client == nil {
		t.Fatal("expected HTTP exception client")
	}

	if _, err := client.Validate(context.Background(), ExceptionValidationRequest{ExceptionID: "EX-1"}); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if gotAuthorization != "Bearer service-internal-demo-token" {
		t.Fatalf("expected bearer token header, got %q", gotAuthorization)
	}
}
