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

package helpers

import (
	"fmt"
	"strings"
)

type AttributeDescription struct {
	String string
}

func NewAttributeDescription(s string) *AttributeDescription {
	return &AttributeDescription{s}
}

func (d *AttributeDescription) AddDefaultValueDescription(defaultValue string) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Default value: `%s`", d.String, defaultValue)
	return d
}

func (d *AttributeDescription) AddStringEnumDescription(values ...string) *AttributeDescription {
	v := make([]string, len(values))
	for i, value := range values {
		v[i] = fmt.Sprintf("`%s`", value)
	}
	d.String = fmt.Sprintf("%s\n  - Choices: %s", d.String, strings.Join(v, ", "))
	return d
}

func (d *AttributeDescription) AddIntegerRangeDescription(min, max int64) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Range: `%v`-`%v`", d.String, min, max)
	return d
}

type ResourceDescription struct {
	String string
}

func NewResourceDescription(description string) *ResourceDescription {
	d := description + "\n"
	return &ResourceDescription{d}
}

func (d *ResourceDescription) AddParents(values ...string) *ResourceDescription {
	d.String += "\n### Parent resources\n\n"
	for _, value := range values {
		d.String += fmt.Sprintf("- [nxos_%s](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/%s)\n", value, value)
	}
	return d
}

func (d *ResourceDescription) AddChildren(values ...string) *ResourceDescription {
	d.String += "\n### Child resources\n\n"
	for _, value := range values {
		d.String += fmt.Sprintf("- [nxos_%s](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/%s)\n", value, value)
	}
	return d
}

func (d *ResourceDescription) AddReferences(values ...string) *ResourceDescription {
	d.String += "\n### Referenced resources\n\n"
	for _, value := range values {
		d.String += fmt.Sprintf("- [nxos_%s](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/%s)\n", value, value)
	}
	return d
}

func (d *ResourceDescription) AddApiDocumentation(rootClassName, rootDocPath string, childClassNames, childDocPaths []string) *ResourceDescription {
	var entries []string
	if rootDocPath != "" {
		entries = append(entries, fmt.Sprintf("- [%s](https://pubhub.devnetcloud.com/media/dme-docs-10-5-3/docs/%s)\n", rootClassName, rootDocPath))
	}
	for i, className := range childClassNames {
		if i < len(childDocPaths) && childDocPaths[i] != "" {
			entries = append(entries, fmt.Sprintf("- [%s](https://pubhub.devnetcloud.com/media/dme-docs-10-5-3/docs/%s)\n", className, childDocPaths[i]))
		}
	}
	if len(entries) > 0 {
		d.String += "\n### API Documentation\n\n"
		for _, entry := range entries {
			d.String += entry
		}
	}
	return d
}
