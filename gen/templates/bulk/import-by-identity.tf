import {
  to = nxos_{{snakeCase .BulkName}}.example
  identity = {
    {{- range (bulkImportAttributes .)}}
    "{{.TfName}}": "<{{.TfName}}>"
    {{- end}}
  }
}
