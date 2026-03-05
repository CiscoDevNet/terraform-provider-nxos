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
{{- $mapKey := mapKeyExample .Attributes}}
{{- if and (not (hasNonIdAttrs .Attributes)) (not .TfChildClasses)}}
{{$indent}}{{.TfName}} = {
{{$indent}}  "{{$mapKey}}" = {}
{{$indent}}}
{{- else}}
{{$indent}}{{.TfName}} = {
{{$indent}}  "{{$mapKey}}" = {
{{- range .Attributes}}
{{- if not .Id}}
{{$indent}}    {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end -}}
{{- if .TfChildClasses -}}
{{- template "renderChildren" (makeMap "Children" .TfChildClasses "Indent" (printf "%s    " $indent)) -}}
{{- end}}
{{$indent}}  }
{{$indent}}}
{{- end}}
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
