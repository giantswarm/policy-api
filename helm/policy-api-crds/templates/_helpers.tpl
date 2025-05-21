{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "labels.selector" -}}
app.kubernetes.io/name: "policy-api-crds"
app.kubernetes.io/instance: {{ .Release.Name | quote }}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "labels.common" -}}
{{ include "labels.selector" . }}
application.giantswarm.io/team: "shield"
helm.sh/chart: {{ include "chart" . | quote }}
{{- end -}}

{{- define "crds.labels" -}}
{{ include "labels.common" . }}
{{- end -}}

{{- define "labels.dummy" -}}
application.giantswarm.io/team: {{ index .Chart.Annotations "application.giantswarm.io/team" | quote }}
{{- end -}}

{{/*
Common annotations
*/}}
{{- define "annotations.helm" -}}
{{- if .Values.chartNameOverride }}
meta.helm.sh/release-name: {{ .Values.chartNameOverride }}
{{- end }}
{{- if .Values.chartNamespaceOverride }}
meta.helm.sh/release-namespace: {{ .Values.chartNamespaceOverride }}
{{- end }}
helm.sh/resource-policy: keep
{{- end -}}

{{- define "crds.annotations" -}}
{{ include "annotations.helm" . }}
{{- end -}}
