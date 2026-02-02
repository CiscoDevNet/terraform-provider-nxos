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
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

const (
	definitionsPath   = "./gen/definitions/"
	providerTemplate  = "./gen/templates/provider.go"
	providerLocation  = "./internal/provider/provider.go"
	objectsTemplate   = "./gen/templates/supported_objects.md.tmpl"
	objectsLocation   = "./templates/guides/supported_objects.md.tmpl"
	changelogTemplate = "./gen/templates/changelog.md.tmpl"
	changelogLocation = "./templates/guides/changelog.md.tmpl"
	changelogOriginal = "./CHANGELOG.md"
)

type t struct {
	path   string
	prefix string
	suffix string
}

var templates = []t{
	{
		path:   "./gen/templates/model.go",
		prefix: "./internal/provider/model_nxos_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source.go",
		prefix: "./internal/provider/data_source_nxos_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source_test.go",
		prefix: "./internal/provider/data_source_nxos_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/resource.go",
		prefix: "./internal/provider/resource_nxos_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/resource_test.go",
		prefix: "./internal/provider/resource_nxos_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/data-source.tf",
		prefix: "./examples/data-sources/nxos_",
		suffix: "/data-source.tf",
	},
	{
		path:   "./gen/templates/resource.tf",
		prefix: "./examples/resources/nxos_",
		suffix: "/resource.tf",
	},
	{
		path:   "./gen/templates/import.sh",
		prefix: "./examples/resources/nxos_",
		suffix: "/import.sh",
	},
}

type YamlConfig struct {
	Name              string                 `yaml:"name"`
	ClassName         string                 `yaml:"class_name"`
	Dn                string                 `yaml:"dn"`
	NoDelete          bool                   `yaml:"no_delete"`
	StatusReplace     bool                   `yaml:"status_replace"`
	TestTags          []string               `yaml:"test_tags"`
	DsDescription     string                 `yaml:"ds_description"`
	ResDescription    string                 `yaml:"res_description"`
	DocPath           string                 `yaml:"doc_path"`
	Parents           []string               `yaml:"parents"`
	Children          []string               `yaml:"children"`
	References        []string               `yaml:"references"`
	Attributes        []YamlConfigAttribute  `yaml:"attributes"`
	ChildClasses      []YamlConfigChildClass `yaml:"child_classes"`
	TestPrerequisites []YamlTest             `yaml:"test_prerequisites"`
}

type YamlConfigAttribute struct {
	NxosName           string   `yaml:"nxos_name"`
	TfName             string   `yaml:"tf_name"`
	Type               string   `yaml:"type"`
	Id                 bool     `yaml:"id"`
	ReferenceOnly      bool     `yaml:"reference_only"`
	Mandatory          bool     `yaml:"mandatory"`
	WriteOnly          bool     `yaml:"write_only"`
	Description        string   `yaml:"description"`
	Example            string   `yaml:"example"`
	EnumValues         []string `yaml:"enum_values"`
	AllowNonEnumValues bool     `yaml:"allow_non_enum_values"`
	OmitEmptyValue     bool     `yaml:"omit_empty_value"`
	MinInt             int      `yaml:"min_int"`
	MaxInt             int      `yaml:"max_int"`
	DefaultValue       string   `yaml:"default_value"`
	DeleteValue        string   `yaml:"delete_value"`
	ExcludeTest        bool     `yaml:"exclude_test"`
	RequiresReplace    bool     `yaml:"requires_replace"`
}

type YamlConfigChildClass struct {
	ClassName        string                     `yaml:"class_name"`
	Rn               string                     `yaml:"rn"`
	Type             string                     `yaml:"type"`
	TfName           string                     `yaml:"tf_name"`
	Description      string                     `yaml:"description"`
	Mandatory        bool                       `yaml:"mandatory"`
	HideInResource   bool                       `yaml:"hide_in_resource"`
	Attributes       []YamlConfigAttribute      `yaml:"attributes"`
	ChildClasses     []YamlConfigChildChildClass `yaml:"child_classes"`
}

type YamlConfigChildChildClass struct {
	ClassName        string                `yaml:"class_name"`
	Rn               string                `yaml:"rn"`
	Type             string                `yaml:"type"`
	TfName           string                `yaml:"tf_name"`
	Description      string                `yaml:"description"`
	Mandatory        bool                  `yaml:"mandatory"`
	HideInResource   bool                  `yaml:"hide_in_resource"`
	Attributes       []YamlConfigAttribute `yaml:"attributes"`
}

type YamlTest struct {
	Dn           string              `yaml:"dn"`
	ClassName    string              `yaml:"class_name"`
	NoDelete     bool                `yaml:"no_delete"`
	Attributes   []YamlTestAttribute `yaml:"attributes"`
	Dependencies []string            `yaml:"dependencies"`
}

type YamlTestAttribute struct {
	Name      string `yaml:"name"`
	Value     string `yaml:"value"`
	Reference string `yaml:"reference"`
}

// Templating helper function to convert TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// Templating helper function to convert string to camel case
func CamelCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// Templating helper function to convert string to snake case
func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// Templating helper function to return true if id included in attributes
func HasId(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Id {
			return true
		}
	}
	return false
}

// Templating helper function to get example dn
func GetExampleDn(dn string, attributes []YamlConfigAttribute) string {
	a := make([]interface{}, 0, len(attributes))
	for _, attr := range attributes {
		if attr.Id {
			a = append(a, attr.Example)
		}
	}
	return fmt.Sprintf(dn, a...)
}

// Templating helper function to identify last element of list
func IsLast(index int, len int) bool {
	return index+1 == len
}

// Templating helper function to count non-reference attributes
func LenNoRef(attributes []YamlConfigAttribute) int {
	count := 0
	for _, attr := range attributes {
		if !attr.ReferenceOnly {
			count += 1
		}
	}
	return count
}

// Templating helper function to support arithmetic addition
func Add(a, b int) int {
	return a + b
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":     ToGoName,
	"camelCase":    CamelCase,
	"snakeCase":    SnakeCase,
	"hasId":        HasId,
	"getExampleDn": GetExampleDn,
	"isLast":       IsLast,
	"sprintf":      fmt.Sprintf,
	"lenNoRef":     LenNoRef,
	"add":          Add,
}

func renderTemplate(templatePath, outputPath string, config interface{}) {
	file, err := os.Open(templatePath)
	if err != nil {
		log.Fatalf("Error opening template: %v", err)
	}
	defer file.Close()

	// skip first line with 'build-ignore' directive for go files
	scanner := bufio.NewScanner(file)
	if strings.HasSuffix(templatePath, ".go") {
		scanner.Scan()
	}
	var temp string
	for scanner.Scan() {
		temp = temp + scanner.Text() + "\n"
	}

	template, err := template.New(path.Base(templatePath)).Funcs(functions).Parse(temp)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// create output file
	outputFile := filepath.Join(outputPath)
	os.MkdirAll(filepath.Dir(outputFile), 0755)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
	f.Write(output.Bytes())
}

func main() {
	items, _ := ioutil.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(items))
	// Iterate over definitions
	for i, filename := range items {
		yamlFile, err := ioutil.ReadFile(filepath.Join(definitionsPath, filename.Name()))
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

	for _, config := range configs {
		// Iterate over templates
		for _, t := range templates {
			renderTemplate(t.path, t.prefix+SnakeCase(config.Name)+t.suffix, config)
		}
	}

	renderTemplate(providerTemplate, providerLocation, configs)
	renderTemplate(objectsTemplate, objectsLocation, configs)

	changelog, err := ioutil.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading changelog: %v", err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))
}
