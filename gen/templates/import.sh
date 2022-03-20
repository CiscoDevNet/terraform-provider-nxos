{{- $name := .Name }}
{{- $dn := .Dn }}
{{- range  .Attributes}}
{{- if eq .Id true -}}
terraform import nxos_{{snakeCase $name}}.example "{{sprintf $dn .Example}}"
{{- end}}
{{- end}}