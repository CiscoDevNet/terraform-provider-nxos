//go:build tools

package tools

import (
	// Documentation generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	// Code generation
	_ "golang.org/x/tools/cmd/goimports"
	_ "gopkg.in/yaml.v3"
)
