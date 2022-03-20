package helpers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// integerDefaultModifier is a plan modifier that sets a default value for a
// types.Int64Type attribute when it is not configured. The attribute must be
// marked as Optional and Computed. When setting the state during the resource
// Create, Read, or Update methods, this default value must also be included or
// the Terraform CLI will generate an error.
type integerDefaultModifier struct {
	Default int64
}

// Description returns a plain text description of the validator's behavior, suitable for a practitioner to understand its impact.
func (m integerDefaultModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to %v", m.Default)
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior, suitable for a practitioner to understand its impact.
func (m integerDefaultModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to %v", m.Default)
}

// Modify runs the logic of the plan modifier.
func (m integerDefaultModifier) Modify(ctx context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	var i types.Int64
	diags := tfsdk.ValueAs(ctx, req.AttributePlan, &i)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	if !i.Unknown {
		return
	}

	resp.AttributePlan = types.Int64{Value: m.Default}
}

func IntegerDefaultModifier(defaultValue int64) integerDefaultModifier {
	return integerDefaultModifier{
		Default: defaultValue,
	}
}
