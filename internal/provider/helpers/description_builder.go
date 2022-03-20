package helpers

import (
	"fmt"
	"strings"
)

type Description struct {
	String string
}

func NewDescription(s string) *Description {
	return &Description{s}
}

func (d *Description) AddDefaultValueDescription(defaultValue string) *Description {
	d.String = fmt.Sprintf("%s\n  - Default value: `%s`", d.String, defaultValue)
	return d
}

func (d *Description) AddStringEnumDescription(values ...string) *Description {
	v := make([]string, len(values))
	for i, value := range values {
		v[i] = fmt.Sprintf("`%s`", value)
	}
	d.String = fmt.Sprintf("%s\n  - Choices: %s", d.String, strings.Join(v, ", "))
	return d
}

func (d *Description) AddIntegerRangeDescription(min, max int64) *Description {
	d.String = fmt.Sprintf("%s\n  - Range: `%v`-`%v`", d.String, min, max)
	return d
}
