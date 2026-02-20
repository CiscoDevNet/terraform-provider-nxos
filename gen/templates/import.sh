terraform import nxos_{{snakeCase .Name}}.example "{{range $i, $e := (importAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
