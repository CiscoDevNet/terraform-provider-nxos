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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RestModel struct {
	Device    types.String     `tfsdk:"device"`
	Id        types.String     `tfsdk:"id"`
	Dn        types.String     `tfsdk:"dn"`
	ClassName types.String     `tfsdk:"class_name"`
	Delete    types.Bool       `tfsdk:"delete"`
	Content   types.Map        `tfsdk:"content"`
	Children  []RestModelChild `tfsdk:"children"`
}

type RestModelChild struct {
	Rn        types.String `tfsdk:"rn"`
	ClassName types.String `tfsdk:"class_name"`
	Content   types.Map    `tfsdk:"content"`
}

type RestDataSourceModel struct {
	Device    types.String `tfsdk:"device"`
	Id        types.String `tfsdk:"id"`
	Dn        types.String `tfsdk:"dn"`
	ClassName types.String `tfsdk:"class_name"`
	Content   types.Map    `tfsdk:"content"`
}

func (data RestModel) toBody(ctx context.Context) nxos.Body {
	body := nxos.Body{}

	className := data.ClassName.ValueString()

	var content map[string]string
	data.Content.ElementsAs(ctx, &content, false)

	body.Str = fmt.Sprintf("{\"%s\":{\"attributes\":{}}}", className)
	for attr, value := range content {
		body = body.Set(className+".attributes."+attr, value)
	}

	for _, child := range data.Children {
		var childContent map[string]string
		child.Content.ElementsAs(ctx, &childContent, false)
		attrs := ""
		for attr, value := range childContent {
			attrs, _ = sjson.Set(attrs, attr, value)
		}
		body = body.SetRaw(className+".children.-1."+child.ClassName.ValueString()+".attributes", attrs)
	}

	return body
}

func (data *RestModel) fromBody(ctx context.Context, res gjson.Result) {
	if !res.Exists() {
		data.Id = types.StringNull()
		return
	}

	content := data.Content.Elements()
	for attr := range content {
		value := res.Get(data.ClassName.ValueString() + ".attributes." + attr)
		content[attr] = types.StringValue(value.String())
	}

	if len(content) > 0 {
		data.Content = types.MapValueMust(types.StringType, content)
	} else {
		data.Content = types.MapNull(types.StringType)
	}

	for c := range data.Children {
		var r gjson.Result
		res.Get(data.ClassName.ValueString() + ".children").ForEach(
			func(_, v gjson.Result) bool {
				key := v.Get(data.Children[c].ClassName.ValueString() + ".attributes.rn").String()
				if key == data.Children[c].Rn.ValueString() {
					r = v
					return false
				}
				return true
			},
		)

		childContent := data.Children[c].Content.Elements()
		for attr := range childContent {
			value := r.Get(data.Children[c].ClassName.ValueString() + ".attributes." + attr)
			childContent[attr] = types.StringValue(value.String())
		}

		if len(content) > 0 {
			data.Children[c].Content = types.MapValueMust(types.StringType, childContent)
		} else {
			data.Children[c].Content = types.MapNull(types.StringType)
		}
	}
}
