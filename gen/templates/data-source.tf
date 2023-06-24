data "nxos_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if or .Id .ReferenceOnly}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end}}
}
