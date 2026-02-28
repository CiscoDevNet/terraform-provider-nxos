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

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &DmeResource{}
var _ resource.ResourceWithIdentity = &DmeResource{}

func NewDmeResource() resource.Resource {
	return &DmeResource{}
}

type DmeResource struct {
	data *NxosProviderData
}

func (r *DmeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dme"
}

func (r *DmeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Manages NX-OS DME Objects via REST API calls. This resource can manage a single API object and its children. It is able to read the state and therefore reconcile configuration drift.",

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
			"dn": schema.StringAttribute{
				MarkdownDescription: "Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"class_name": schema.StringAttribute{
				MarkdownDescription: "Which class object is being created. (Make sure there is no colon in the classname)",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"delete": schema.BoolAttribute{
				MarkdownDescription: "Delete object during destroy operation. Default value is `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"content": schema.MapAttribute{
				MarkdownDescription: "Map of key-value pairs that need to be passed to the Model object as parameters.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"children": schema.ListNestedAttribute{
				MarkdownDescription: "List of children.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rn": schema.StringAttribute{
							MarkdownDescription: "The relative name of the child object.",
							Required:            true,
						},
						"class_name": schema.StringAttribute{
							MarkdownDescription: "Class name of the child object.",
							Required:            true,
						},
						"content": schema.MapAttribute{
							MarkdownDescription: "Map of key-value pairs which represents the attributes of the child object.",
							Optional:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (r *DmeResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"dn": identityschema.StringAttribute{
				Description:       "Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].",
				RequiredForImport: true,
			},
			"class_name": identityschema.StringAttribute{
				Description:       "Which class object is being created.",
				RequiredForImport: true,
			},
			"device": identityschema.StringAttribute{
				Description:       "A device name from the provider configuration.",
				OptionalForImport: true,
			},
		},
	}
}

func (r *DmeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*NxosProviderData)
}

func (r *DmeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DmeModel

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Dn.ValueString()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", plan.Device.ValueString()))
		return
	}

	if device.Managed {
		body := plan.toBody(ctx)
		_, err := device.Client.Post(plan.Dn.ValueString(), body.Str)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
			return
		}
	}

	plan.Id = plan.Dn
	var identity DmeIdentity
	identity.toIdentity(ctx, &plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *DmeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DmeModel

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read identity if available (requires Terraform >= 1.12.0)
	if req.Identity != nil && !req.Identity.Raw.IsNull() {
		var identity DmeIdentity
		diags = req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		state.fromIdentity(ctx, &identity)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", state.Device.ValueString()))
		return
	}

	if device.Managed {
		queries := []func(*nxos.Req){nxos.Query("rsp-prop-include", "config-only")}
		if len(state.Children) > 0 {
			queries = append(queries, nxos.Query("rsp-subtree", "children"))
		}
		res, err := device.Client.GetDn(state.Dn.ValueString(), queries...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		imp, diags := helpers.IsFlagImporting(ctx, req)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		state.fromBody(ctx, res, imp)
	}

	var identity DmeIdentity
	identity.toIdentity(ctx, &state)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *DmeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DmeModel

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", plan.Device.ValueString()))
		return
	}

	if device.Managed {
		body := plan.toBody(ctx)
		_, err := device.Client.Post(plan.Dn.ValueString(), body.Str)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	var identity DmeIdentity
	identity.toIdentity(ctx, &plan)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)
}

func (r *DmeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DmeModel

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", state.Device.ValueString()))
		return
	}

	if device.Managed {
		if state.Delete.ValueBool() {
			body := helpers.DeleteBody(state.ClassName.ValueString())
			res, err := device.Client.Post(state.Dn.ValueString(), body)
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *DmeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != "" || req.Identity == nil || req.Identity.Raw.IsNull() {
		idParts := strings.Split(req.ID, ",")
		idParts = helpers.RemoveEmptyStrings(idParts)

		if len(idParts) != 2 && len(idParts) != 3 {
			expectedIdentifier := "Expected import identifier with format: '<dn>,<class_name>'"
			expectedIdentifier += " or '<dn>,<class_name>,<device>'"
			resp.Diagnostics.AddError(
				"Unexpected Import Identifier",
				fmt.Sprintf("%s. Got: %q", expectedIdentifier, req.ID),
			)
			return
		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Import", idParts[0]))

		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("dn"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("class_name"), idParts[1])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("delete"), true)...)
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("dn"), idParts[0])...)
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("class_name"), idParts[1])...)
		if len(idParts) == 3 {
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device"), idParts[2])...)
			resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("device"), idParts[2])...)
		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Import finished successfully", idParts[0]))
	} else {
		var identity DmeIdentity
		diags := req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Import", identity.Dn.ValueString()))

		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("dn"), identity.Dn.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("class_name"), identity.ClassName.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), identity.Dn.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("delete"), true)...)
		if !identity.Device.IsNull() && !identity.Device.IsUnknown() && identity.Device.ValueString() != "" {
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device"), identity.Device.ValueString())...)
		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Import finished successfully", identity.Dn.ValueString()))
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
