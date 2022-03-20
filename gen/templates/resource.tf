resource "nxos_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
}
