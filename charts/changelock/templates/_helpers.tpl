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

