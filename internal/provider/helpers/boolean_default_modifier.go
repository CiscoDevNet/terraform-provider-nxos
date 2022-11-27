package helpers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// booleanDefaultModifier is a plan modifier that sets a default value for a
// types.BoolType attribute when it is not configured. The attribute must be
// marked as Optional and Computed. When setting the state during the resource
// Create, Read, or Update methods, this default value must also be included or
// the Terraform CLI will generate an error.
type booleanDefaultModifier struct {
	Default bool
}

// Description returns a plain text description of the validator's behavior, suitable for a practitioner to understand its impact.
func (m booleanDefaultModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to %v", m.Default)
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior, suitable for a practitioner to understand its impact.
func (m booleanDefaultModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%v`", m.Default)
}

// Modify runs the logic of the plan modifier.
func (m booleanDefaultModifier) Modify(ctx context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	var bool types.Bool
	diags := tfsdk.ValueAs(ctx, req.AttributePlan, &bool)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	if !bool.IsUnknown() {
		return
	}

	resp.AttributePlan = types.BoolValue(m.Default)
}

func BooleanDefaultModifier(defaultValue bool) booleanDefaultModifier {
	return booleanDefaultModifier{
		Default: defaultValue,
	}
}
