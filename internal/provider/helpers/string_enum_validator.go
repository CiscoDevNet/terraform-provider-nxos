package helpers

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	stringEnumValidatorErr = "Value must be one of [%s]"
)

type stringEnumValidator struct {
	values []string
}

func (v stringEnumValidator) Description(context.Context) string {
	return fmt.Sprintf(stringEnumValidatorErr, strings.Join(v.values, ", "))
}

func (v stringEnumValidator) MarkdownDescription(context.Context) string {
	return fmt.Sprintf(stringEnumValidatorErr, strings.Join(v.values, ", "))
}

func (v stringEnumValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var str types.String
	{
		diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &str)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() {
			return
		}
	}

	if str.Unknown || str.Null {
		return
	}

	for _, val := range v.values {
		if val == str.Value {
			return
		}
	}

	resp.Diagnostics.AddAttributeError(
		req.AttributePath,
		"Invalid String",
		fmt.Sprintf(stringEnumValidatorErr, strings.Join(v.values, ", ")),
	)
}

func StringEnumValidator(values ...string) tfsdk.AttributeValidator {
	return stringEnumValidator{
		values: values,
	}
}
