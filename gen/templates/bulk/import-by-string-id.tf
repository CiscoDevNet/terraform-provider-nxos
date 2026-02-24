import {
  to = nxos_{{snakeCase .BulkName}}.example
  id = "{{range $i, $e := (bulkImportAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
}
