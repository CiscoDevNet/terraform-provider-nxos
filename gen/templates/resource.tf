resource "nxos_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- range .ChildClasses}}
{{- if and (not .HideInResource) (eq .Type "single")}}
{{- range .Attributes}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- else if and (not .HideInResource) (eq .Type "list")}}
  {{.TfName}} = [{
  {{- range .Attributes}}
    {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
  {{- end}}
  }]
{{- end}}
{{- end}}
{{- $hasHiddenNestedChildren := false}}
{{- range .ChildClasses}}{{- if and .HideInResource .ChildClasses}}{{$hasHiddenNestedChildren = true}}{{end}}{{end}}
{{- if $hasHiddenNestedChildren}}
{{- range .ChildClasses}}
{{- if .HideInResource}}
{{- range .ChildClasses}}
{{- if eq .Type "list"}}
  {{.TfName}} = [{
  {{- range .Attributes}}
    {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
  {{- end}}
  }]
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
}
