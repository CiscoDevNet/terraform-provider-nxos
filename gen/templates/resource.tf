{{- /* renderChildren recursively renders TfChildClasses.
       Context: map with "Children" ([]TfChildClass) and "Indent" (string). */ -}}
{{- define "renderChildren" -}}
{{- $indent := .Indent -}}
{{- range (testChildClasses .Children) -}}
{{- if eq .Type "single" -}}
{{- range .Attributes}}
{{- if not .ExcludeTest}}
{{$indent}}{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end -}}
{{- if .TfChildClasses -}}
{{- template "renderChildren" (makeMap "Children" .TfChildClasses "Indent" $indent) -}}
{{- end -}}
{{- else if eq .Type "list"}}
{{- $mapKey := mapKeyExample .Attributes}}
{{- if and (not (hasNonIdAttrs .Attributes)) (not .TfChildClasses)}}
{{$indent}}{{.TfName}} = {
{{$indent}}  "{{$mapKey}}" = {}
{{- range exampleEntries .}}
{{$indent}}  "{{.Key}}" = {
{{- range .Attributes}}
{{$indent}}    {{.TfName}} = {{.Hcl}}
{{- end}}
{{$indent}}  }
{{- end}}
{{$indent}}}
{{- else}}
{{$indent}}{{.TfName}} = {
{{$indent}}  "{{$mapKey}}" = {
{{- range .Attributes}}
{{- if and (not .Id) (not .ExcludeTest)}}
{{$indent}}    {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end -}}
{{- if .TfChildClasses -}}
{{- template "renderChildren" (makeMap "Children" .TfChildClasses "Indent" (printf "%s    " $indent)) -}}
{{- end}}
{{$indent}}  }
{{- range exampleEntries .}}
{{$indent}}  "{{.Key}}" = {
{{- range .Attributes}}
{{$indent}}    {{.TfName}} = {{.Hcl}}
{{- end}}
{{$indent}}  }
{{- end}}
{{$indent}}}
{{- end}}
{{- end -}}
{{- end -}}
{{- end -}}
{{- /* ==================== end renderChildren ==================== */ -}}
resource "nxos_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if not .ExcludeTest}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end -}}
{{- template "renderChildren" (makeMap "Children" .TfChildClasses "Indent" "  ") }}
}
