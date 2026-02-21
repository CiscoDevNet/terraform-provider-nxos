resource "nxos_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- range .TfChildClasses}}
{{- if eq .Type "single"}}
{{- range .Attributes}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- range .TfChildClasses}}
{{- if eq .Type "single"}}
{{- range .Attributes}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- range .TfChildClasses}}
{{- if eq .Type "list"}}
  {{.TfName}} = [{
  {{- range .Attributes}}
    {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
  {{- end}}
  }]
{{- end}}
{{- end}}
{{- else if eq .Type "list"}}
  {{.TfName}} = [{
  {{- range .Attributes}}
    {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
  {{- end}}
  }]
{{- end}}
{{- end}}
{{- else if eq .Type "list"}}
  {{.TfName}} = [{
  {{- range .Attributes}}
    {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
  {{- end}}
  {{- range .TfChildClasses}}
  {{- if eq .Type "single"}}
  {{- range .Attributes}}
    {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
  {{- end}}
  {{- range .TfChildClasses}}
  {{- if eq .Type "list"}}
    {{.TfName}} = [{
    {{- range .Attributes}}
      {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
    {{- end}}
    }]
  {{- end}}
  {{- end}}
  {{- else if eq .Type "list"}}
    {{.TfName}} = [{
    {{- range .Attributes}}
      {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
    {{- end}}
    }]
  {{- end}}
  {{- end}}
  }]
{{- end}}
{{- end}}
}
