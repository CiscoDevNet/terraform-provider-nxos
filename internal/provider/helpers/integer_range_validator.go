package helpers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	integerRangeValidatorErr = "Value must be between [%v] and [%v]"
)

type integerRangeValidator struct {
	min int64
	max int64
}

func (v integerRangeValidator) Description(context.Context) string {
	return fmt.Sprintf(integerRangeValidatorErr, v.min, v.max)
}

func (v integerRangeValidator) MarkdownDescription(context.Context) string {
	return fmt.Sprintf(integerRangeValidatorErr, v.min, v.max)
}

func (v integerRangeValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var int types.Int64
	{
		diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &int)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() {
			return
		}
	}

	if int.Unknown || int.Null {
		return
	}

	if int.Value >= v.min && int.Value <= v.max {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.AttributePath,
		"Invalid value",
		fmt.Sprintf(integerRangeValidatorErr, v.min, v.max),
	)
}

func IntegerRangeValidator(min, max int64) tfsdk.AttributeValidator {
	return integerRangeValidator{
		min: min,
		max: max,
	}
}
