package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Rest struct {
	Id        types.String `tfsdk:"id"`
	Dn        types.String `tfsdk:"dn"`
	ClassName types.String `tfsdk:"class_name"`
	Content   types.Map    `tfsdk:"content"`
}
