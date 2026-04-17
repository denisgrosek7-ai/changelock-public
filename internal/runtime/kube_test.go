package runtime

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestKubernetesClientPatchApprovedState(t *testing.T) {
	var gotPath string
	var gotBody map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		if r.Method != http.MethodPatch {
			t.Fatalf("expected PATCH, got %s", r.Method)
		}
		if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
			t.Fatalf("decode patch body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	client := newKubernetesClient(server.URL, "token", server.Client())
	err := client.PatchApprovedState(context.Background(), ApprovedWorkloadState{
		Namespace:          "acme-prod",
		WorkloadKind:       "Deployment",
		Workload:           "booking-api",
		ServiceAccountName: "booking-api",
		ExpectedConfigHash: "cfg-123",
		Containers: []ApprovedContainerState{{
			Name:           "app",
			Image:          "ghcr.io/acme/booking-api@sha256:abc123",
			ApprovedDigest: "sha256:abc123",
			Runtime: SecurityConstraints{
				RunAsNonRoot:             true,
				ReadOnlyRootFilesystem:   true,
				AllowPrivilegeEscalation: false,
				DropAllCapabilities:      true,
				SeccompRuntimeDefault:    true,
				DenyPrivileged:           true,
			},
		}},
	})
	if err != nil {
		t.Fatalf("PatchApprovedState() error = %v", err)
	}
	if gotPath != "/apis/apps/v1/namespaces/acme-prod/deployments/booking-api" {
		t.Fatalf("unexpected patch path %q", gotPath)
	}
	body := gotBody["spec"].(map[string]any)["template"].(map[string]any)["spec"].(map[string]any)
	if body["serviceAccountName"] != "booking-api" {
		t.Fatalf("expected service account patch, got %#v", body)
	}
}

func TestKubernetesClientRestartToApprovedStateDeletesPods(t *testing.T) {
	requests := []string{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.Method+" "+r.URL.Path)
		switch {
		case r.Method == http.MethodGet && strings.Contains(r.URL.Path, "/deployments/booking-api"):
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{
			  "metadata":{"name":"booking-api","namespace":"acme-prod"},
			  "spec":{
			    "selector":{"matchLabels":{"app":"booking-api"}},
			    "template":{"metadata":{"labels":{"app":"booking-api"}},"spec":{"serviceAccountName":"booking-api","containers":[{"name":"app","image":"ghcr.io/acme/booking-api@sha256:abc123"}]}}
			  }
			}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/v1/namespaces/acme-prod/pods":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"items":[{"metadata":{"name":"booking-api-123"}},{"metadata":{"name":"booking-api-456"}}]}`))
		case r.Method == http.MethodDelete:
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := newKubernetesClient(server.URL, "token", server.Client())
	if err := client.RestartToApprovedState(context.Background(), ApprovedWorkloadState{
		Namespace:    "acme-prod",
		WorkloadKind: "Deployment",
		Workload:     "booking-api",
	}); err != nil {
		t.Fatalf("RestartToApprovedState() error = %v", err)
	}
	if len(requests) != 4 {
		t.Fatalf("expected controller read + pod list + pod deletes, got %#v", requests)
	}
}

func TestKubernetesClientApplyQuarantineOverlayCreatesNetworkPolicy(t *testing.T) {
	requests := []string{}
	var created map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.Method+" "+r.URL.Path)
		switch {
		case r.Method == http.MethodGet && strings.Contains(r.URL.Path, "/networkpolicies/changelock-quarantine-booking-api"):
			http.NotFound(w, r)
		case r.Method == http.MethodPost && r.URL.Path == "/apis/networking.k8s.io/v1/namespaces/acme-prod/networkpolicies":
			if err := json.NewDecoder(r.Body).Decode(&created); err != nil {
				t.Fatalf("decode networkpolicy: %v", err)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"status":"created"}`))
		default:
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := newKubernetesClient(server.URL, "token", server.Client())
	err := client.ApplyQuarantineOverlay(context.Background(), ApprovedWorkloadState{
		Namespace:    "acme-prod",
		WorkloadKind: "Deployment",
		Workload:     "booking-api",
		Labels:       map[string]string{"app": "booking-api"},
	}, ObservedWorkloadState{})
	if err != nil {
		t.Fatalf("ApplyQuarantineOverlay() error = %v", err)
	}
	if len(requests) != 2 {
		t.Fatalf("expected lookup + create, got %#v", requests)
	}
	spec := created["spec"].(map[string]any)
	selector := spec["podSelector"].(map[string]any)["matchLabels"].(map[string]any)
	if selector["app"] != "booking-api" {
		t.Fatalf("expected workload label selector, got %#v", selector)
	}
}
