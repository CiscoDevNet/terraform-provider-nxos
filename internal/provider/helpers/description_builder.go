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

func NewResourceDescription(description, className, docPath string) *ResourceDescription {
	d := fmt.Sprintf("%s\n\n- API Documentation: [%s](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/%s)\n", description, className, docPath)
	return &ResourceDescription{d}
}

func (d *ResourceDescription) AddParents(values ...string) *ResourceDescription {
	d.String += "\n### Parent resources\n\n"
	for _, value := range values {
		d.String += fmt.Sprintf("- [nxos_%s](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/%s)\n", value, value)
	}
	return d
}

func (d *ResourceDescription) AddChildren(values ...string) *ResourceDescription {
	d.String += "\n### Child resources\n\n"
	for _, value := range values {
		d.String += fmt.Sprintf("- [nxos_%s](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/%s)\n", value, value)
	}
	return d
}

func (d *ResourceDescription) AddReferences(values ...string) *ResourceDescription {
	d.String += "\n### Referenced resources\n\n"
	for _, value := range values {
		d.String += fmt.Sprintf("- [nxos_%s](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/%s)\n", value, value)
	}
	return d
}
