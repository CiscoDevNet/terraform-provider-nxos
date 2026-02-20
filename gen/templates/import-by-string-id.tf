import {
  to = nxos_{{snakeCase .Name}}.example
  id = "{{range $i, $e := (importAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
}
