terraform import nxos_{{snakeCase .BulkName}}.example "{{range $i, $e := (bulkImportAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
