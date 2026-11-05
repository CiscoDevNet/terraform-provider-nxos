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

//go:build ignore

package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	definitionsPath = "./gen/definitions/"
)

type YamlConfig struct {
	Name        string `yaml:"name"`
	DocCategory string `yaml:"doc_category"`
}

var docPaths = []string{"./docs/data-sources/", "./docs/resources/", "./docs/actions/"}

var extraDocs = map[string]string{
	"dme":         "General",
	"save_config": "General",
	"cli":         "General",
}

func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// pageTitleRegex matches the default tfplugindocs page title, e.g. "nxos_bgp Resource - terraform-provider-nxos"
var pageTitleRegex = regexp.MustCompile(`(?m)^page_title: "(\S+) (Resource|Data Source|Action) - terraform-provider-nxos"$`)

// cleanFrontmatter shortens the page title to e.g. "nxos_bgp (Resource)" and reduces the
// description to its first line, which drops the list of API documentation links. Both are
// only used for page titles and meta descriptions, e.g. by the OpenTofu registry.
func cleanFrontmatter(s string) string {
	s = pageTitleRegex.ReplaceAllString(s, `page_title: "$1 ($2)"`)

	lines := strings.Split(s, "\n")
	out := make([]string, 0, len(lines))
	inDescription, keptLine := false, false
	for i, line := range lines {
		if i > 0 && line == "---" {
			// end of frontmatter, keep the rest unchanged
			out = append(out, lines[i:]...)
			break
		}
		if line == "description: |-" {
			inDescription, keptLine = true, false
			out = append(out, line)
			continue
		}
		if inDescription && strings.HasPrefix(line, "  ") {
			if !keptLine {
				out = append(out, line)
				keptLine = true
			}
			continue
		}
		inDescription = false
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}

func main() {
	items, _ := os.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(items))

	// Load configs
	for i, filename := range items {
		yamlFile, err := os.ReadFile(filepath.Join(definitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		configs[i] = config
	}

	for i := range configs {
		name := SnakeCase(configs[i].Name)
		for _, path := range docPaths {
			filename := path + name + ".md"
			content, err := os.ReadFile(filename)
			if err != nil {
				continue
			}

			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+configs[i].DocCategory+`"`)

			os.WriteFile(filename, []byte(s), 0644)
		}
	}

	// Update extra doc categories
	for doc, cat := range extraDocs {
		for _, path := range docPaths {
			filename := path + doc + ".md"
			content, err := os.ReadFile(filename)
			if err == nil {
				s := string(content)
				s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+cat+`"`)

				os.WriteFile(filename, []byte(s), 0644)
			}
		}
	}

	// Clean up frontmatter of all docs
	for _, path := range docPaths {
		files, _ := filepath.Glob(filepath.Join(path, "*.md"))
		for _, filename := range files {
			content, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error reading file: %v", err)
			}
			os.WriteFile(filename, []byte(cleanFrontmatter(string(content))), 0644)
		}
	}
}
