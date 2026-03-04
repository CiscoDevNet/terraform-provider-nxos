import {
  to = nxos_{{snakeCase .Name}}.example
  {{- if importAttributes .}}
  identity = {
    {{- range (importAttributes .)}}
    "{{.TfName}}": "<{{.TfName}}>"
    {{- end}}
  }
  {{- else}}
  identity = {}
  {{- end}}
}
