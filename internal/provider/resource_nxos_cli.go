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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &CliResource{}

func NewCliResource() resource.Resource {
	return &CliResource{}
}

type CliResource struct {
	data *NxosProviderData
}

func (r *CliResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cli"
}

func (r *CliResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource is used to send arbitrary CLI commands to NX-OS devices via NX-API JSON-RPC. This should be considered a last resort in case DME objects are not available, as it cannot read the state and therefore cannot reconcile changes.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"cli": schema.StringAttribute{
				MarkdownDescription: "The CLI commands to send to the device. Multi-line strings are supported, with each line being sent as a separate command.",
				Required:            true,
			},
		},
	}
}

func (r *CliResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*NxosProviderData)
}

func (r *CliResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var device types.String
	var cli types.String

	diags := req.Plan.GetAttribute(ctx, path.Root("device"), &device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.GetAttribute(ctx, path.Root("cli"), &cli)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning to send CLI commands")

	d, ok := r.data.Devices[device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", device.ValueString()))
		return
	}

	if d.Managed {
		commands := parseCliCommands(cli.ValueString())
		_, err := d.Client.JsonRpc(commands)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to send CLI commands, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, "Send CLI commands finished successfully")

	diags = resp.State.SetAttribute(ctx, path.Root("device"), device)
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(ctx, path.Root("cli"), cli)
	resp.Diagnostics.Append(diags...)
}

func (r *CliResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *CliResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var device types.String
	var cli types.String

	diags := req.Plan.GetAttribute(ctx, path.Root("device"), &device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.GetAttribute(ctx, path.Root("cli"), &cli)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning to send CLI commands")

	d, ok := r.data.Devices[device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", device.ValueString()))
		return
	}

	if d.Managed {
		commands := parseCliCommands(cli.ValueString())
		_, err := d.Client.JsonRpc(commands)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to send CLI commands, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, "Send CLI commands finished successfully")

	diags = resp.State.SetAttribute(ctx, path.Root("device"), device)
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(ctx, path.Root("cli"), cli)
	resp.Diagnostics.Append(diags...)
}

func (r *CliResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

// parseCliCommands splits a multi-line CLI string into individual commands,
// trimming whitespace and skipping empty lines.
func parseCliCommands(cli string) []string {
	var commands []string
	for _, line := range strings.Split(cli, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			commands = append(commands, line)
		}
	}
	return commands
}
