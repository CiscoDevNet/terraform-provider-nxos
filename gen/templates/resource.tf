{{- /* renderChildren recursively renders TfChildClasses.
       Context: map with "Children" ([]TfChildClass) and "Indent" (string). */ -}}
{{- define "renderChildren" -}}
{{- $indent := .Indent -}}
{{- range .Children -}}
{{- if eq .Type "single" -}}
{{- range .Attributes}}
{{$indent}}{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end -}}
{{- if .TfChildClasses -}}
{{- template "renderChildren" (makeMap "Children" .TfChildClasses "Indent" $indent) -}}
{{- end -}}
{{- else if eq .Type "list"}}
{{$indent}}{{.TfName}} = [{
{{- range .Attributes}}
{{$indent}}  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end -}}
{{- if .TfChildClasses -}}
{{- template "renderChildren" (makeMap "Children" .TfChildClasses "Indent" (printf "%s  " $indent)) -}}
{{- end}}
{{$indent}}}]
{{- end -}}
{{- end -}}
{{- end -}}
{{- /* ==================== end renderChildren ==================== */ -}}
resource "nxos_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end -}}
{{- template "renderChildren" (makeMap "Children" .TfChildClasses "Indent" "  ") }}
}
