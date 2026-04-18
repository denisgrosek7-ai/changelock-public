package runtime

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	inClusterTokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	inClusterCAPath    = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
)

var (
	ErrKubernetesUnavailable   = errors.New("kubernetes runtime client unavailable")
	ErrUnsupportedWorkloadKind = errors.New("unsupported workload kind")
)

type RemediationClient interface {
	PatchApprovedState(ctx context.Context, desired ApprovedWorkloadState) error
	RestartToApprovedState(ctx context.Context, desired ApprovedWorkloadState) error
	ApplyQuarantineOverlay(ctx context.Context, desired ApprovedWorkloadState, observed ObservedWorkloadState) error
}

type NoopRemediationClient struct{}

func (NoopRemediationClient) PatchApprovedState(context.Context, ApprovedWorkloadState) error {
	return ErrKubernetesUnavailable
}

func (NoopRemediationClient) RestartToApprovedState(context.Context, ApprovedWorkloadState) error {
	return ErrKubernetesUnavailable
}

func (NoopRemediationClient) ApplyQuarantineOverlay(context.Context, ApprovedWorkloadState, ObservedWorkloadState) error {
	return ErrKubernetesUnavailable
}

type KubernetesClient struct {
	baseURL string
	client  *http.Client
	token   string
}

func newKubernetesClient(baseURL, token string, client *http.Client) *KubernetesClient {
	if client == nil {
		client = &http.Client{Timeout: 5 * time.Second}
	}
	return &KubernetesClient{
		baseURL: strings.TrimRight(strings.TrimSpace(baseURL), "/"),
		client:  client,
		token:   strings.TrimSpace(token),
	}
}

func NewKubernetesClientFromInCluster() (*KubernetesClient, error) {
	host := strings.TrimSpace(os.Getenv("KUBERNETES_SERVICE_HOST"))
	port := strings.TrimSpace(os.Getenv("KUBERNETES_SERVICE_PORT"))
	if host == "" || port == "" {
		return nil, ErrKubernetesUnavailable
	}

	tokenBytes, err := os.ReadFile(inClusterTokenPath)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrKubernetesUnavailable, err)
	}
	token := strings.TrimSpace(string(tokenBytes))
	if token == "" {
		return nil, fmt.Errorf("%w: empty service account token", ErrKubernetesUnavailable)
	}

	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}
	if caBytes, err := os.ReadFile(inClusterCAPath); err == nil {
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(caBytes)
		tlsConfig.RootCAs = pool
	}

	return newKubernetesClient(
		"https://"+host+":"+port,
		token,
		&http.Client{
			Timeout: 5 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
	), nil
}

func (c *KubernetesClient) ReadObservedWorkload(ctx context.Context, target WorkloadTarget) (ObservedWorkloadState, error) {
	resource, err := c.getController(ctx, target)
	if err != nil {
		return ObservedWorkloadState{}, err
	}

	observed := ObservedWorkloadState{
		ClusterID:          strings.TrimSpace(target.ClusterID),
		Namespace:          resource.Metadata.Namespace,
		WorkloadKind:       normalizeWorkloadKind(target.Kind),
		Workload:           resource.Metadata.Name,
		ServiceAccountName: resource.Spec.Template.Spec.ServiceAccountName,
		ActualConfigHash:   firstNonEmpty(resource.Spec.Template.Metadata.Labels["changelock/config-hash"], resource.Spec.Template.Metadata.Labels["changelock.dev/config-hash"]),
		PodLabels:          cloneStringMap(resource.Spec.Template.Metadata.Labels),
	}

	digests, err := c.readRunningDigests(ctx, target, resource.Spec.Selector.MatchLabels)
	if err != nil && !errors.Is(err, ErrKubernetesUnavailable) {
		return ObservedWorkloadState{}, err
	}

	for _, container := range resource.Spec.Template.Spec.Containers {
		observed.Containers = append(observed.Containers, ObservedContainerState{
			Name:          container.Name,
			Image:         container.Image,
			RunningDigest: firstNonEmpty(digests[container.Name], digestFromImageReference(container.Image)),
			Runtime:       convertObservedSecurity(container.SecurityContext),
		})
	}

	return observed, nil
}

func (c *KubernetesClient) PatchApprovedState(ctx context.Context, desired ApprovedWorkloadState) error {
	resourcePath, err := workloadResourcePath(desired.Namespace, desired.WorkloadKind, desired.Workload)
	if err != nil {
		return err
	}

	containers := make([]map[string]any, 0, len(desired.Containers))
	for _, container := range desired.Containers {
		containers = append(containers, map[string]any{
			"name":            container.Name,
			"image":           container.Image,
			"securityContext": buildSecurityContextPatch(container.Runtime),
		})
	}

	patch := map[string]any{
		"spec": map[string]any{
			"template": map[string]any{
				"metadata": map[string]any{
					"labels": cloneStringMap(desired.Labels),
				},
				"spec": map[string]any{
					"serviceAccountName": desired.ServiceAccountName,
					"containers":         containers,
				},
			},
		},
	}

	if desired.ExpectedConfigHash != "" {
		template := patch["spec"].(map[string]any)["template"].(map[string]any)
		metadata := template["metadata"].(map[string]any)
		labels, _ := metadata["labels"].(map[string]string)
		if labels == nil {
			labels = map[string]string{}
		}
		labels["changelock/config-hash"] = desired.ExpectedConfigHash
		metadata["labels"] = labels
	}

	body, err := json.Marshal(patch)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.baseURL+resourcePath, bytes.NewReader(body))
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", "Bearer "+c.token)
	request.Header.Set("Content-Type", "application/merge-patch+json")
	request.Header.Set("Accept", "application/json")
	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		payload, _ := io.ReadAll(io.LimitReader(response.Body, 1024))
		return fmt.Errorf("kubernetes patch returned %d: %s", response.StatusCode, strings.TrimSpace(string(payload)))
	}
	return nil
}

func (c *KubernetesClient) RestartToApprovedState(ctx context.Context, desired ApprovedWorkloadState) error {
	resource, err := c.getController(ctx, WorkloadTarget{
		ClusterID: desired.ClusterID,
		Namespace: desired.Namespace,
		Kind:      desired.WorkloadKind,
		Workload:  desired.Workload,
	})
	if err != nil {
		return err
	}
	pods, err := c.listPods(ctx, desired.Namespace, resource.Spec.Selector.MatchLabels)
	if err != nil {
		return err
	}
	for _, pod := range pods.Items {
		path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s", url.PathEscape(desired.Namespace), url.PathEscape(pod.Metadata.Name))
		request, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseURL+path, nil)
		if err != nil {
			return err
		}
		request.Header.Set("Authorization", "Bearer "+c.token)
		response, err := c.client.Do(request)
		if err != nil {
			return err
		}
		response.Body.Close()
		if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
			return fmt.Errorf("kubernetes pod delete returned %d", response.StatusCode)
		}
	}
	return nil
}

func (c *KubernetesClient) ApplyQuarantineOverlay(ctx context.Context, desired ApprovedWorkloadState, observed ObservedWorkloadState) error {
	labels := cloneStringMap(desired.Labels)
	if len(labels) == 0 {
		labels = cloneStringMap(observed.PodLabels)
	}
	if len(labels) == 0 {
		return fmt.Errorf("%w: quarantine overlay requires workload labels", ErrKubernetesUnavailable)
	}

	name := quarantinePolicyName(desired.Workload)
	resourcePath := fmt.Sprintf("/apis/networking.k8s.io/v1/namespaces/%s/networkpolicies/%s", url.PathEscape(desired.Namespace), url.PathEscape(name))
	collectionPath := fmt.Sprintf("/apis/networking.k8s.io/v1/namespaces/%s/networkpolicies", url.PathEscape(desired.Namespace))
	body, err := json.Marshal(map[string]any{
		"apiVersion": "networking.k8s.io/v1",
		"kind":       "NetworkPolicy",
		"metadata": map[string]any{
			"name": name,
			"labels": map[string]string{
				"app.kubernetes.io/managed-by": "changelock",
				"changelock.io/quarantine":     "true",
			},
		},
		"spec": map[string]any{
			"podSelector": map[string]any{"matchLabels": labels},
			"policyTypes": []string{"Ingress", "Egress"},
			"ingress":     []any{},
			"egress":      []any{},
		},
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+resourcePath, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		return c.doJSON(ctx, http.MethodPut, resourcePath, "application/json", body, nil)
	case http.StatusNotFound:
		return c.doJSON(ctx, http.MethodPost, collectionPath, "application/json", body, nil)
	default:
		return fmt.Errorf("kubernetes networkpolicy lookup returned %d", resp.StatusCode)
	}
}

func (c *KubernetesClient) getController(ctx context.Context, target WorkloadTarget) (controllerResource, error) {
	path, err := workloadResourcePath(target.Namespace, target.Kind, target.Workload)
	if err != nil {
		return controllerResource{}, err
	}
	var resource controllerResource
	if err := c.doJSON(ctx, http.MethodGet, path, "", nil, &resource); err != nil {
		return controllerResource{}, err
	}
	return resource, nil
}

func (c *KubernetesClient) readRunningDigests(ctx context.Context, target WorkloadTarget, selector map[string]string) (map[string]string, error) {
	pods, err := c.listPods(ctx, target.Namespace, selector)
	if err != nil {
		return nil, err
	}

	digests := map[string]string{}
	for _, pod := range pods.Items {
		for _, status := range pod.Status.ContainerStatuses {
			digest := parseDigestFromImageID(status.ImageID)
			if digest == "" {
				digest = digestFromImageReference(status.Image)
			}
			if digest != "" {
				digests[status.Name] = digest
			}
		}
	}
	return digests, nil
}

func (c *KubernetesClient) listPods(ctx context.Context, namespace string, selector map[string]string) (podList, error) {
	path := fmt.Sprintf("/api/v1/namespaces/%s/pods", url.PathEscape(namespace))
	if len(selector) > 0 {
		values := url.Values{}
		values.Set("labelSelector", buildLabelSelector(selector))
		path += "?" + values.Encode()
	}
	var pods podList
	if err := c.doJSON(ctx, http.MethodGet, path, "", nil, &pods); err != nil {
		return podList{}, err
	}
	return pods, nil
}

func (c *KubernetesClient) doJSON(ctx context.Context, method, path, contentType string, body []byte, out any) error {
	request, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bytes.NewReader(body))
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", "Bearer "+c.token)
	request.Header.Set("Accept", "application/json")
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		payload, _ := io.ReadAll(io.LimitReader(response.Body, 1024))
		return fmt.Errorf("kubernetes request returned %d: %s", response.StatusCode, strings.TrimSpace(string(payload)))
	}
	if out == nil {
		return nil
	}
	return json.NewDecoder(response.Body).Decode(out)
}

func workloadResourcePath(namespace, kind, name string) (string, error) {
	switch normalizeWorkloadKind(kind) {
	case "Deployment":
		return fmt.Sprintf("/apis/apps/v1/namespaces/%s/deployments/%s", url.PathEscape(namespace), url.PathEscape(name)), nil
	case "DaemonSet":
		return fmt.Sprintf("/apis/apps/v1/namespaces/%s/daemonsets/%s", url.PathEscape(namespace), url.PathEscape(name)), nil
	case "StatefulSet":
		return fmt.Sprintf("/apis/apps/v1/namespaces/%s/statefulsets/%s", url.PathEscape(namespace), url.PathEscape(name)), nil
	default:
		return "", fmt.Errorf("%w: %s", ErrUnsupportedWorkloadKind, kind)
	}
}

func buildLabelSelector(values map[string]string) string {
	if len(values) == 0 {
		return ""
	}
	parts := make([]string, 0, len(values))
	for key, value := range values {
		parts = append(parts, key+"="+value)
	}
	sortStrings(parts)
	return strings.Join(parts, ",")
}

func parseDigestFromImageID(imageID string) string {
	trimmed := strings.TrimSpace(imageID)
	if trimmed == "" {
		return ""
	}
	if idx := strings.Index(trimmed, "@"); idx >= 0 {
		return trimmed[idx+1:]
	}
	return ""
}

func buildSecurityContextPatch(expected SecurityConstraints) map[string]any {
	return map[string]any{
		"runAsNonRoot":             expected.RunAsNonRoot,
		"readOnlyRootFilesystem":   expected.ReadOnlyRootFilesystem,
		"allowPrivilegeEscalation": expected.AllowPrivilegeEscalation,
		"privileged":               !expected.DenyPrivileged,
		"capabilities": map[string]any{
			"drop": []string{"ALL"},
		},
		"seccompProfile": map[string]any{
			"type": "RuntimeDefault",
		},
	}
}

func convertObservedSecurity(value *containerSecurityContext) SecurityPosture {
	if value == nil {
		return SecurityPosture{}
	}
	posture := SecurityPosture{
		RunAsNonRoot:             value.RunAsNonRoot,
		ReadOnlyRootFilesystem:   value.ReadOnlyRootFilesystem,
		AllowPrivilegeEscalation: value.AllowPrivilegeEscalation,
		SeccompRuntimeDefault:    strings.EqualFold(value.SeccompProfile.Type, "RuntimeDefault"),
		Privileged:               value.Privileged,
	}
	for _, dropped := range value.Capabilities.Drop {
		if dropped == "ALL" {
			posture.DropAllCapabilities = true
			break
		}
	}
	return posture
}

func cloneStringMap(source map[string]string) map[string]string {
	if len(source) == 0 {
		return nil
	}
	cloned := make(map[string]string, len(source))
	for key, value := range source {
		cloned[key] = value
	}
	return cloned
}

func sortStrings(values []string) {
	if len(values) < 2 {
		return
	}
	for idx := 0; idx < len(values)-1; idx++ {
		for jdx := idx + 1; jdx < len(values); jdx++ {
			if values[jdx] < values[idx] {
				values[idx], values[jdx] = values[jdx], values[idx]
			}
		}
	}
}

type controllerResource struct {
	Metadata struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		Selector struct {
			MatchLabels map[string]string `json:"matchLabels"`
		} `json:"selector"`
		Template struct {
			Metadata struct {
				Labels map[string]string `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				ServiceAccountName string                    `json:"serviceAccountName"`
				Containers         []controllerContainerSpec `json:"containers"`
			} `json:"spec"`
		} `json:"template"`
	} `json:"spec"`
}

type controllerContainerSpec struct {
	Name            string                    `json:"name"`
	Image           string                    `json:"image"`
	SecurityContext *containerSecurityContext `json:"securityContext,omitempty"`
}

type containerSecurityContext struct {
	RunAsNonRoot             bool `json:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem   bool `json:"readOnlyRootFilesystem,omitempty"`
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation,omitempty"`
	Privileged               bool `json:"privileged,omitempty"`
	Capabilities             struct {
		Drop []string `json:"drop,omitempty"`
	} `json:"capabilities,omitempty"`
	SeccompProfile struct {
		Type string `json:"type,omitempty"`
	} `json:"seccompProfile,omitempty"`
}

type podList struct {
	Items []podResource `json:"items"`
}

type podResource struct {
	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Status struct {
		ContainerStatuses []struct {
			Name    string `json:"name"`
			Image   string `json:"image"`
			ImageID string `json:"imageID"`
		} `json:"containerStatuses"`
	} `json:"status"`
}

func digestFromImageReference(image string) string {
	parts := strings.SplitN(strings.TrimSpace(image), "@", 2)
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func quarantinePolicyName(workload string) string {
	trimmed := strings.ToLower(strings.TrimSpace(workload))
	if trimmed == "" {
		trimmed = "workload"
	}
	trimmed = strings.ReplaceAll(trimmed, "_", "-")
	return "changelock-quarantine-" + trimmed
}
