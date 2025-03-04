//go:build ignore
// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &{{camelCase .Name}}Resource{}
var _ resource.ResourceWithImportState = &{{camelCase .Name}}Resource{}

func New{{camelCase .Name}}Resource() resource.Resource {
	return &{{camelCase .Name}}Resource{}
}

type {{camelCase .Name}}Resource struct {
	data *NxosProviderData
}

func (r *{{camelCase .Name}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .Name}}"
}

func (r *{{camelCase .Name}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("{{.ResDescription}}", "{{.ClassName}}", "{{.DocPath}}")
			{{- if len .Parents -}}
			.AddParents({{range .Parents}}"{{snakeCase .}}", {{end}})
			{{- end -}}
			{{- if len .Children -}}
			.AddChildren({{range .Children}}"{{snakeCase .}}", {{end}})
			{{- end -}}
			{{- if len .References -}}
			.AddReferences({{range .References}}"{{snakeCase .}}", {{end}})
			{{- end -}}
			.String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			{{- range  .Attributes}}
			"{{.TfName}}": schema.{{.Type}}Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
					{{- if len .EnumValues -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
					.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
					{{- end -}}
					{{- if len .DefaultValue -}}
					.AddDefaultValueDescription("{{.DefaultValue}}")
					{{- end -}}
					.String,
				{{- if or .Id .ReferenceOnly .Mandatory}}
				Required:            true,
				{{- else}}
				Optional:            true,
				{{- if len .DefaultValue}}
				Computed:            true,
				{{- end}}
				{{- end}}
				{{- if and (len .DefaultValue) (not .Id) (not .ReferenceOnly) (not .Mandatory)}}
				{{- if eq .Type "Int64"}}
				Default: int64default.StaticInt64({{.DefaultValue}}),
				{{- else if eq .Type "Bool"}}
				Default: booldefault.StaticBool({{.DefaultValue}}),
				{{- else if eq .Type "String"}}
				Default: stringdefault.StaticString("{{.DefaultValue}}"),
				{{- end}}
				{{- end}}
				{{- if and (len .EnumValues) (not .AllowNonEnumValues) }}
				Validators: []validator.String{
					stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
				Validators: []validator.Int64{
					int64validator.Between({{.MinInt}}, {{.MaxInt}}),
				},
				{{- end}}
				{{- if or .Id .ReferenceOnly .RequiresReplace}}
				PlanModifiers: []planmodifier.{{.Type}}{
					{{snakeCase .Type}}planmodifier.RequiresReplace(),
				},
				{{- end}}
			},
			{{- end}}
			{{- range .ChildClasses}}
			{{- if eq .Type "single"}}
			{{- range  .Attributes}}
			"{{.TfName}}": schema.{{.Type}}Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
					{{- if len .EnumValues -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
					.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
					{{- end -}}
					{{- if len .DefaultValue -}}
					.AddDefaultValueDescription("{{.DefaultValue}}")
					{{- end -}}
					.String,
				{{- if or .Id .ReferenceOnly .Mandatory}}
				Required:            true,
				{{- else}}
				Optional:            true,
				{{- if len .DefaultValue}}
				Computed:            true,
				{{- end}}
				{{- end}}
				{{- if and (len .DefaultValue) (not .Id) (not .ReferenceOnly) (not .Mandatory)}}
				{{- if eq .Type "Int64"}}
				Default: int64default.StaticInt64({{.DefaultValue}}),
				{{- else if eq .Type "Bool"}}
				Default: booldefault.StaticBool({{.DefaultValue}}),
				{{- else if eq .Type "String"}}
				Default: stringdefault.StaticString("{{.DefaultValue}}"),
				{{- end}}
				{{- end}}
				{{- if and (len .EnumValues) (not .AllowNonEnumValues) }}
				Validators: []validator.String{
					stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
				Validators: []validator.Int64{
					int64validator.Between({{.MinInt}}, {{.MaxInt}}),
				},
				{{- end}}
				{{- if or .Id .ReferenceOnly .RequiresReplace}}
				PlanModifiers: []planmodifier.{{.Type}}{
					{{snakeCase .Type}}planmodifier.RequiresReplace(),
				},
				{{- end}}
			},
			{{- end}}
			{{- else if eq .Type "list"}}
			"{{.TfName}}": schema.ListNestedAttribute{
				MarkdownDescription: "{{.Description}}",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range  .Attributes}}
						"{{.TfName}}": schema.{{.Type}}Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
								{{- if len .EnumValues -}}
								.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
								{{- end -}}
								{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
								.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
								{{- end -}}
								{{- if len .DefaultValue -}}
								.AddDefaultValueDescription("{{.DefaultValue}}")
								{{- end -}}
								.String,
							{{- if or .Id .ReferenceOnly .Mandatory}}
							Required:            true,
							{{- else}}
							Optional:            true,
							{{- if len .DefaultValue}}
							Computed:            true,
							{{- end}}
							{{- end}}
							{{- if and (len .DefaultValue) (not .Id) (not .ReferenceOnly) (not .Mandatory)}}
							{{- if eq .Type "Int64"}}
							Default: int64default.StaticInt64({{.DefaultValue}}),
							{{- else if eq .Type "Bool"}}
							Default: booldefault.StaticBool({{.DefaultValue}}),
							{{- else if eq .Type "String"}}
							Default: stringdefault.StaticString("{{.DefaultValue}}"),
							{{- end}}
							{{- end}}
							{{- if and (len .EnumValues) (not .AllowNonEnumValues) }}
							Validators: []validator.String{
								stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
							},
							{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
							Validators: []validator.Int64{
								int64validator.Between({{.MinInt}}, {{.MaxInt}}),
							},
							{{- end}}
							{{- if or .Id .ReferenceOnly .RequiresReplace}}
							PlanModifiers: []planmodifier.{{.Type}}{
								{{snakeCase .Type}}planmodifier.RequiresReplace(),
							},
							{{- end}}
						},
						{{- end}}
					},
				},
			},
			{{- end}}
			{{- end}}
		},
	}
}

func (r *{{camelCase .Name}}Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*NxosProviderData)
}

func (r *{{camelCase .Name}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", plan.Device.ValueString()))
		return
	}

	// Post object
	if device.Managed {
		body := plan.toBody(false)
		_, err := device.Client.Post(plan.getDn(), body.Str)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
			return
		}
	}

	plan.Dn = types.StringValue(plan.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *{{camelCase .Name}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", state.Device.ValueString()))
		return
	}

	if device.Managed {
		queries := []func(*nxos.Req){nxos.Query("rsp-prop-include", "config-only")}
		{{- if .ChildClasses}}
		queries = append(queries, nxos.Query("rsp-subtree", "children"))
		{{- end}}
		res, err := device.Client.GetDn(state.Dn.ValueString(), queries...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		imp, diags := helpers.IsFlagImporting(ctx, req)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		state.fromBody(res, imp)
		{{- if hasId .Attributes}}
		if imp {
			state.getIdsFromDn()
		}
		{{- end}}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *{{camelCase .Name}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", plan.Device.ValueString()))
		return
	}

	if device.Managed {
	{{ if .StatusReplace}}
		body := plan.toBody(true)
	{{ else}}
		body := plan.toBody(false)
	{{ end}}
		_, err := device.Client.Post(plan.getDn(), body.Str)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *{{camelCase .Name}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.ValueString()))
{{ if not .NoDelete}}
	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", state.Device.ValueString()))
		return
	}

	if device.Managed {
		body := state.toDeleteBody()

		if (len(body.Str) > 0) {
			_, err := device.Client.Post(state.getDn(), body.Str)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
				return
			}
		} else {
			res, err := device.Client.DeleteDn(state.Dn.ValueString())
			if err != nil {
				errCode := res.Get("imdata.0.error.attributes.code").Str
				// Ignore errors of type "Cannot delete object"
				if errCode != "1" && errCode != "107" {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
		}
	}
{{ end}}
	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *{{camelCase .Name}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
