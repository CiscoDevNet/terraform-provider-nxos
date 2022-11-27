//go:build ignore
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	{{- if hasId .Attributes }}
	"fmt"
	{{- end}}
	{{ $strconv := false }}{{ range .Attributes}}{{ if or (and (eq .Type "Int64") (ne .ReferenceOnly true)) (eq .Type "Bool")}}{{ $strconv = true }}{{ end}}{{- end}}
	{{- if $strconv }}
	"strconv"
	{{- end}}

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	{{- $helpers := false }}{{ range .Attributes}}{{ if eq .Type "Bool"}}{{ $helpers = true }}{{ end}}{{- end}}
	{{- if $helpers }}
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
	{{- end}}
)

type {{camelCase .Name}} struct {
	Device types.String `tfsdk:"device"`
	Dn types.String `tfsdk:"id"`
{{- range .Attributes}}
	{{toTitle .NxosName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
}

func (data {{camelCase .Name}}) getDn() string {
{{- if hasId .Attributes}}
	return fmt.Sprintf("{{.Dn}}"{{range .Attributes}}{{if eq .Id true}}, data.{{toTitle .NxosName}}.Value{{.Type}}(){{end}}{{end}})
{{- else}}
	return "{{.Dn}}"
{{- end}}
}

func (data {{camelCase .Name}}) getClassName() string {
	return "{{.ClassName}}"
}

func (data {{camelCase .Name}}) toBody() nxos.Body {
	{{- $lenAttr := len .Attributes}}
	{{- if ge $lenAttr 1 }}
	attrs := nxos.Body{}.
	{{- end}}
	{{- $lenAttr := len .Attributes}}
	{{- range $index, $item := .Attributes}}
	{{- if ne .ReferenceOnly true}}
		{{- if eq .Type "Int64"}}
		Set("{{.NxosName}}", strconv.FormatInt(data.{{toTitle .NxosName}}.ValueInt64(), 10))
		{{- else if eq .Type "Bool"}}
		Set("{{.NxosName}}", strconv.FormatBool(data.{{toTitle .NxosName}}.ValueBool()))
		{{- else if eq .Type "String"}}
		Set("{{.NxosName}}", data.{{toTitle .NxosName}}.ValueString())
		{{- end}}
		{{- if not (isLast $index $lenAttr)}}.{{- end}}
	{{- end}}
	{{- end}}
	{{- range .Attributes}}
	{{- if eq .OmitEmptyValue true}}
	if data.{{toTitle .NxosName}}.IsUnknown() || data.{{toTitle .NxosName}}.IsNull() {
		attrs = attrs.Delete("{{.NxosName}}")
	}
	{{- end}}
	{{- end}}
	{{- if ge $lenAttr 1 }}
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
	{{- else}}
	return nxos.Body{Str: `{"` + data.getClassName() + `":{"attributes":{}}}`}
	{{- end}}
}

func (data *{{camelCase .Name}}) fromBody(res gjson.Result) {
	{{- range .Attributes}}
	{{- if and (ne .ReferenceOnly true) (ne .WriteOnly true)}}
	{{- if eq .Type "Int64"}}
	data.{{toTitle .NxosName}} = types.Int64Value(res.Get("*.attributes.{{.NxosName}}").Int())
	{{- else if eq .Type "Bool"}}
	data.{{toTitle .NxosName}} = types.BoolValue(helpers.ParseNxosBoolean(res.Get("*.attributes.{{.NxosName}}").String()))
	{{- else if eq .Type "String"}}
	data.{{toTitle .NxosName}} = types.StringValue(res.Get("*.attributes.{{.NxosName}}").String())
	{{- end}}
	{{- end}}
	{{- end}}
}

func (data *{{camelCase .Name}}) fromPlan(plan {{camelCase .Name}}) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	{{- range .Attributes}}
	{{- if or (eq .ReferenceOnly true) (eq .WriteOnly true)}}
	data.{{toTitle .NxosName}} = plan.{{toTitle .NxosName}}
	{{- end}}
	{{- end}}
}