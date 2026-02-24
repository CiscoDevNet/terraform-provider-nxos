resource "nxos_{{snakeCase .BulkName}}" "example" {
  items = [{
{{- range .Attributes}}
{{- if not .ReferenceOnly}}
    {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end}}
  }]
}
