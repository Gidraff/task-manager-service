{{/*
Database URL template
*/}}
{{- define "chart.database.url" -}}
{{- $fakeRootContext := $ -}}
{{- $databaseSettings := mergeOverwrite (dict) (mergeOverwrite (dict) .Values.database) -}}
{{- $_ := set $fakeRootContext "database" $databaseSettings -}}
{{- $databaseUrlTemplate := "" -}}
{{- if $databaseSettings.urlTemplate -}}
{{- $databaseUrlTemplate = $databaseSettings.urlTemplate }}
{{- else }}
{{- $databaseEnvVarsMap := mergeOverwrite (dict) (mergeOverwrite (dict) .Values.databaseEnvVarsMap) -}}
{{- $databaseUrlTemplate = "host={{.database.host}} user={{.database.username}} password={{.database.password}} dbname={{.database.dbname}} port={{.database.port}} sslmode={{.database.sslmode}} TimeZone={{.database.timeZone}}" -}}
{{- end -}}
{{- printf "%s" (tpl $databaseUrlTemplate $fakeRootContext) -}}
{{- end -}}

{{- define "chart.server.envVars" -}}
{{- $serverSettings := mergeOverwrite (dict) (mergeOverwrite (dict) .Values.server) -}}
{{- $serverEnvVarsMap := mergeOverwrite (dict) (mergeOverwrite (dict) .Values.serverEnvVarsMap) -}}
- name: {{ $serverEnvVarsMap.PORT }}
  value: {{ $serverSettings.port | quote }}
{{- end -}}

{{/*
Postgre Env variables
*/}}
{{- define "chart.database.envVars" -}}
{{- $databaseSettings := mergeOverwrite (dict) (mergeOverwrite (dict) .Values.database) -}}
{{- $databaseEnvVarsMap := mergeOverwrite (dict) (mergeOverwrite (dict) .Values.databaseEnvVarsMap) -}}
{{- if $databaseSettings.enabled -}}
- name: {{ $databaseEnvVarsMap.USERNAME }}
  value: {{ $databaseSettings.username | quote }}
- name: {{ $databaseEnvVarsMap.PASSWORD }}
{{- if $databaseSettings.existingSecret }}
  valueFrom:
    secretKeyRef:
      name: {{ $databaseSettings.existingSecret }}
      key: {{ $databaseSettings.existingSecretKey }}
{{- else }}
  value: {{ required "A valid database.password entry required!" $databaseSettings.password | quote }}
{{- end }}
- name: {{ $databaseEnvVarsMap.URL }}
  value: {{ include "chart.database.url" . | quote }}
{{- end -}}
{{- end -}}