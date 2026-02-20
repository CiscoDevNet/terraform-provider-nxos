import {
  to = nxos_{{snakeCase .Name}}.example
  identity = {
    {{- range (importAttributes .)}}
    "{{.TfName}}": "<{{.TfName}}>"
    {{- end}}
  }
}
