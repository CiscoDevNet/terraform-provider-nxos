package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &RestResource{}
var _ resource.ResourceWithImportState = &RestResource{}

func NewRestResource() resource.Resource {
	return &RestResource{}
}

type RestResource struct {
	data NxosProviderData
}

func (r *RestResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rest"
}

func (r *RestResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Manages NX-OS DME Objects via REST API calls. This resource can only manage a single API object. It is able to read the state and therefore reconcile configuration drift.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			"dn": {
				MarkdownDescription: "Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].",
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.RequiresReplace(),
				},
			},
			"class_name": {
				MarkdownDescription: "Which class object is being created. (Make sure there is no colon in the classname)",
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.RequiresReplace(),
				},
			},
			"delete": {
				MarkdownDescription: "Delete object during destroy operation. Default value is `true`.",
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.BooleanDefaultModifier(true),
				},
			},
			"content": {
				Type:                types.MapType{ElemType: types.StringType},
				MarkdownDescription: "Map of key-value pairs that need to be passed to the Model object as parameters.",
				Optional:            true,
				Computed:            true,
			},
		},
	}, nil
}

func (r *RestResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RestModel

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.Value))

	body := plan.toBody(ctx)
	_, err := r.data.client.Post(plan.Dn.Value, body.Str, nxos.OverrideUrl(r.data.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	plan.Id = plan.Dn
	plan.Content.Unknown = false

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.Value))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RestResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RestModel

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	res, err := r.data.client.GetDn(state.Dn.Value, nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.data.devices[state.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	// Check if we received an empty response without errors -> object has been deleted
	if !res.Exists() && err == nil {
		state.Id.Value = ""
	} else {
		className := state.ClassName.Value
		state.Id.Value = state.Dn.Value

		stateContent := state.Content
		newContent := make(map[string]attr.Value)
		for attr, value := range res.Get(className + ".attributes").Map() {
			if _, ok := stateContent.Elems[attr]; ok {
				newContent[attr] = types.String{Value: value.Str}
			}
		}

		state.Content.Elems = newContent
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *RestResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RestModel

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.Value))

	body := plan.toBody(ctx)
	_, err := r.data.client.Post(plan.Dn.Value, body.Str, nxos.OverrideUrl(r.data.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.Value))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RestResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RestModel

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.Value))

	if state.Delete.Value {
		res, err := r.data.client.DeleteDn(state.Dn.Value, nxos.OverrideUrl(r.data.devices[state.Device.Value]))
		if err != nil {
			errCode := res.Get("imdata.0.error.attributes.code").Str
			// Ignore errors of type "Cannot delete object"
			if errCode != "1" && errCode != "107" {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.Value))

	resp.State.RemoveResource(ctx)
}

func (r *RestResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ":")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Unexpected format of ID (%s), expected class_name:dn", req.ID),
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Import", idParts[1]))

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("dn"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("class_name"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Import finished successfully", idParts[1]))
}
