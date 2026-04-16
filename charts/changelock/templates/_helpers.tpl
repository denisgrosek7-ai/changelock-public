{{- define "changelock.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "changelock.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name (include "changelock.name" .) | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}

{{- define "changelock.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" -}}
{{- end -}}

{{- define "changelock.labels" -}}
helm.sh/chart: {{ include "changelock.chart" . }}
{{ include "changelock.selectorLabels" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- with .Values.commonLabels }}
{{ toYaml . }}
{{- end }}
{{- end -}}

{{- define "changelock.selectorLabels" -}}
app.kubernetes.io/name: {{ include "changelock.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "changelock.componentFullname" -}}
{{- printf "%s-%s" (include "changelock.fullname" .root) .component | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "changelock.componentLabels" -}}
{{ include "changelock.selectorLabels" .root }}
app.kubernetes.io/component: {{ .component }}
{{- end -}}

{{- define "changelock.serviceAccountName" -}}
{{- if .serviceAccount.create -}}
{{- default (include "changelock.componentFullname" (dict "root" .root "component" .component)) .serviceAccount.name -}}
{{- else -}}
{{- default "default" .serviceAccount.name -}}
{{- end -}}
{{- end -}}

{{- define "changelock.image" -}}
{{- if .image.tag -}}
{{ printf "%s:%s" .image.repository .image.tag }}
{{- else -}}
{{ .image.repository }}
{{- end -}}
{{- end -}}

{{- define "changelock.affinity" -}}
{{- $root := .root -}}
{{- $component := .component -}}
{{- $mode := default "none" .mode -}}
{{- if ne $mode "none" }}
podAntiAffinity:
  {{- if eq $mode "required" }}
  requiredDuringSchedulingIgnoredDuringExecution:
    - labelSelector:
        matchLabels:
          {{- include "changelock.componentLabels" (dict "root" $root "component" $component) | nindent 10 }}
      topologyKey: kubernetes.io/hostname
  {{- else }}
  preferredDuringSchedulingIgnoredDuringExecution:
    - weight: 100
      podAffinityTerm:
        labelSelector:
          matchLabels:
            {{- include "changelock.componentLabels" (dict "root" $root "component" $component) | nindent 12 }}
        topologyKey: kubernetes.io/hostname
  {{- end }}
{{- end }}
{{- end -}}

{{- define "changelock.authSecretName" -}}
{{- if .Values.auth.existingSecret -}}
{{- .Values.auth.existingSecret -}}
{{- else -}}
{{- printf "%s-auth" (include "changelock.fullname" .) -}}
{{- end -}}
{{- end -}}

{{- define "changelock.postgresSecretName" -}}
{{- printf "%s-postgres" (include "changelock.fullname" .) -}}
{{- end -}}

{{- define "changelock.validateValues" -}}
{{- $profile := default "demo" .Values.deploymentProfile -}}
{{- if not (or (eq $profile "demo") (eq $profile "production")) -}}
{{- fail (printf "unsupported deploymentProfile %q (expected demo or production)" $profile) -}}
{{- end -}}
{{- if and .Values.auth.createSecret (ne (trim .Values.auth.existingSecret) "") -}}
{{- fail "set only one of auth.createSecret or auth.existingSecret" -}}
{{- end -}}
{{- if eq $profile "production" -}}
  {{- if eq .Values.auth.mode "disabled" -}}
    {{- fail "production deploymentProfile requires auth.mode to be static-token or oidc-jwt" -}}
  {{- end -}}
  {{- if and .Values.auth.createSecret (eq (trim .Values.auth.internalServiceToken) "service-internal-demo-token") -}}
    {{- fail "production deploymentProfile does not allow the demo internal service token in auth.createSecret" -}}
  {{- end -}}
  {{- if and (eq .Values.auth.mode "static-token") .Values.auth.createSecret (eq (trim .Values.auth.tokensJson) "") -}}
    {{- fail "production static-token mode requires auth.tokensJson when auth.createSecret=true" -}}
  {{- end -}}
  {{- if eq .Values.auth.mode "oidc-jwt" -}}
    {{- if eq (trim .Values.auth.oidc.issuer) "" -}}
      {{- fail "production oidc-jwt mode requires auth.oidc.issuer" -}}
    {{- end -}}
    {{- if eq (trim .Values.auth.oidc.audiences) "" -}}
      {{- fail "production oidc-jwt mode requires auth.oidc.audiences" -}}
    {{- end -}}
    {{- if eq (trim .Values.auth.oidc.jwksUrl) "" -}}
      {{- fail "production oidc-jwt mode requires auth.oidc.jwksUrl" -}}
    {{- end -}}
    {{- if eq (trim .Values.auth.roleBindingsJson) "" -}}
      {{- fail "production oidc-jwt mode requires auth.roleBindingsJson" -}}
    {{- end -}}
  {{- end -}}
  {{- if eq .Values.sync.mode "hub" -}}
    {{- if and .Values.sync.requireClusterId (eq (trim .Values.sync.clusterBindingsExistingSecret) "") (eq (trim .Values.sync.clusterBindingsJson) "") -}}
      {{- fail "production hub mode with sync.requireClusterId=true requires sync.clusterBindingsExistingSecret or sync.clusterBindingsJson" -}}
    {{- end -}}
    {{- if and (eq (trim .Values.auth.existingSecret) "") (not .Values.auth.createSecret) -}}
      {{- fail "production hub mode requires machine auth via auth.existingSecret or auth.createSecret" -}}
    {{- end -}}
  {{- end -}}
  {{- if eq .Values.sync.mode "spoke" -}}
    {{- if eq (trim .Values.sync.clusterId) "" -}}
      {{- fail "production spoke mode requires sync.clusterId" -}}
    {{- end -}}
    {{- if eq (trim .Values.sync.hubUrl) "" -}}
      {{- fail "production spoke mode requires sync.hubUrl" -}}
    {{- end -}}
    {{- if and (eq (trim .Values.sync.tokenExistingSecret) "") (eq (trim .Values.auth.existingSecret) "") (not .Values.auth.createSecret) -}}
      {{- fail "production spoke mode requires sync.tokenExistingSecret or auth secret wiring for machine auth" -}}
    {{- end -}}
  {{- end -}}
  {{- if eq .Values.signer.mode "software" -}}
    {{- if eq (trim .Values.signer.existingSecret) "" -}}
      {{- fail "production software signer mode requires signer.existingSecret" -}}
    {{- end -}}
  {{- end -}}
  {{- if eq .Values.signer.mode "vault-transit" -}}
    {{- if eq (trim .Values.signer.existingSecret) "" -}}
      {{- fail "production vault-transit mode requires signer.existingSecret" -}}
    {{- end -}}
    {{- if eq (trim .Values.signer.vault.addr) "" -}}
      {{- fail "production vault-transit mode requires signer.vault.addr" -}}
    {{- end -}}
    {{- if eq (trim .Values.signer.vault.transitKey) "" -}}
      {{- fail "production vault-transit mode requires signer.vault.transitKey" -}}
    {{- end -}}
  {{- end -}}
{{- end -}}
{{- end -}}
