{{- if importAttributes .}}
import {
  to = nxos_{{snakeCase .Name}}.example
  id = "{{range $i, $e := (importAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
}
{{- else}}
# Terraform's `import` block requires a non-empty `id` value, so this format cannot be used to
# import nxos_{{snakeCase .Name}} for the default device. Use the `terraform import` command
# below, or identity-based import shown above, instead.
{{- end}}
