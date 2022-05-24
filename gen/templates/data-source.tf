data "nxos_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if or (eq .Id true) (eq .ReferenceOnly true)}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end}}
}
