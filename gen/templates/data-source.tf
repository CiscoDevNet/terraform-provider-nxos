data "nxos_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if eq .Id true}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end}}
}
