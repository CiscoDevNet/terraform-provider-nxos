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
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
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
	{
		path:   "./gen/templates/import-by-string-id.tf",
		prefix: "./examples/resources/nxos_",
		suffix: "/import-by-string-id.tf",
	},
	{
		path:   "./gen/templates/import-by-identity.tf",
		prefix: "./examples/resources/nxos_",
		suffix: "/import-by-identity.tf",
	},
}

type YamlConfig struct {
	Name              string                 `yaml:"name"`
	ClassName         string                 `yaml:"class_name"`
	Dn                string                 `yaml:"dn"`
	NoDelete          bool                   `yaml:"no_delete"`
	TestTags          []string               `yaml:"test_tags"`
	DsDescription     string                 `yaml:"ds_description"`
	ResDescription    string                 `yaml:"res_description"`
	DocPath           string                 `yaml:"doc_path"`
	Parents           []string               `yaml:"parents"`
	Children          []string               `yaml:"children"`
	References        []string               `yaml:"references"`
	Attributes        []YamlConfigAttribute  `yaml:"attributes"`
	ChildClasses      []YamlConfigChildClass `yaml:"child_classes"`
	TfChildClasses    []YamlConfigChildClass `yaml:"-"`
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
	Sensitive          bool     `yaml:"sensitive"`
	Description        string   `yaml:"description"`
	Example            string   `yaml:"example"`
	EnumValues         []string `yaml:"enum_values"`
	AllowNonEnumValues bool     `yaml:"allow_non_enum_values"`
	AlwaysInclude      bool     `yaml:"always_include"`
	MinInt             int      `yaml:"min_int"`
	MaxInt             int      `yaml:"max_int"`
	DefaultValue       string   `yaml:"default_value"`
	Value              string   `yaml:"value"`
	ExcludeTest        bool     `yaml:"exclude_test"`
	RequiresReplace    bool     `yaml:"requires_replace"`
	TestTags           []string `yaml:"test_tags"`
}

type YamlConfigChildClass struct {
	ClassName      string                 `yaml:"class_name"`
	Rn             string                 `yaml:"rn"`
	Type           string                 `yaml:"type"`
	TfName         string                 `yaml:"tf_name"`
	Description    string                 `yaml:"description"`
	DocPath        string                 `yaml:"doc_path"`
	Mandatory      bool                   `yaml:"mandatory"`
	NoDelete       bool                   `yaml:"no_delete"`
	AlwaysInclude  bool                   `yaml:"always_include"`
	Attributes     []YamlConfigAttribute  `yaml:"attributes"`
	ChildClasses   []YamlConfigChildClass `yaml:"child_classes"`
	TfChildClasses []YamlConfigChildClass `yaml:"-"`
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

func HasWriteOnly(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.WriteOnly && !attr.ExcludeTest {
			return true
		}
	}
	return false
}

func HasSensitiveAttr(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Sensitive {
			return true
		}
	}
	return false
}

func HasSensitiveAttrRecursive(children []YamlConfigChildClass) bool {
	for _, c := range children {
		if HasSensitiveAttr(c.Attributes) {
			return true
		}
		if HasSensitiveAttrRecursive(c.ChildClasses) {
			return true
		}
	}
	return false
}

// Templating helper function to support arithmetic addition
func Add(a, b int) int {
	return a + b
}

// Templating helper function to collect child class names with doc paths (recursive)
func ChildDocClassNames(children []YamlConfigChildClass) []string {
	var names []string
	for _, c := range children {
		if c.DocPath != "" {
			names = append(names, c.ClassName)
		}
		names = append(names, ChildDocClassNames(c.ChildClasses)...)
	}
	return names
}

// Templating helper function to collect child doc paths (recursive)
func ChildDocPaths(children []YamlConfigChildClass) []string {
	var paths []string
	for _, c := range children {
		if c.DocPath != "" {
			paths = append(paths, c.DocPath)
		}
		paths = append(paths, ChildDocPaths(c.ChildClasses)...)
	}
	return paths
}

// Templating helper function to return import attributes (reference_only or id)
func ImportAttributes(config YamlConfig) []YamlConfigAttribute {
	attributes := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if attr.ReferenceOnly || attr.Id {
			attributes = append(attributes, attr)
		}
	}
	return attributes
}

// Templating helper function to check if any attribute is testable (not excluded and not write-only)
func HasTestAttrs(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if !attr.ExcludeTest && !attr.WriteOnly {
			return true
		}
	}
	return false
}

// Templating helper function to check if any child class (recursively) is a list
func HasListChildClasses(children []YamlConfigChildClass) bool {
	for _, c := range children {
		if c.Type == "list" {
			return true
		}
		if HasListChildClasses(c.ChildClasses) {
			return true
		}
	}
	return false
}

// RnHasDynamicSegment returns true if the rn contains a format placeholder ('%').
func RnHasDynamicSegment(rn string) bool {
	return strings.Contains(rn, "%")
}

// RnFormatArgs generates the Go expression list for fmt.Sprintf arguments
// to build the rn from the id attributes with a given variable prefix.
func RnFormatArgs(prefix string, attributes []YamlConfigAttribute) string {
	var args []string
	for _, attr := range attributes {
		if attr.Id {
			switch attr.Type {
			case "Int64":
				args = append(args, prefix+"."+ToGoName(attr.TfName)+".ValueInt64()")
			case "Bool":
				args = append(args, prefix+"."+ToGoName(attr.TfName)+".ValueBool()")
			default:
				args = append(args, prefix+"."+ToGoName(attr.TfName)+".ValueString()")
			}
		}
	}
	return strings.Join(args, ", ")
}

// ChildRn extracts the child RN from a DN pattern.
// e.g., "sys/bd/bd-[%s]" → "bd-[%s]"
func ChildRn(dn string) string {
	idx := strings.LastIndex(dn, "/")
	if idx == -1 {
		return dn
	}
	return dn[idx+1:]
}

// Templating helper function to create a map from key-value pairs for passing multiple values to {{template}}
func MakeMap(pairs ...interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(pairs); i += 2 {
		m[pairs[i].(string)] = pairs[i+1]
	}
	return m
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":                  ToGoName,
	"camelCase":                 CamelCase,
	"snakeCase":                 SnakeCase,
	"hasId":                     HasId,
	"add":                       Add,
	"childDocClassNames":        ChildDocClassNames,
	"childDocPaths":             ChildDocPaths,
	"hasWriteOnly":              HasWriteOnly,
	"hasSensitiveAttr":          HasSensitiveAttr,
	"hasSensitiveAttrRecursive": HasSensitiveAttrRecursive,
	"importAttributes":          ImportAttributes,
	"hasListChildClasses":       HasListChildClasses,
	"hasTestAttrs":              HasTestAttrs,
	"makeMap":                   MakeMap,
	"rnHasDynamicSegment":       RnHasDynamicSegment,
	"rnFormatArgs":              RnFormatArgs,
	"childRn":                   ChildRn,
	"allChildClassNames":        AllChildClassNames,
	"join":                      strings.Join,
	"hasNonIdAttrs":             HasNonIdAttrs,
	"needsPlanItem":             NeedsPlanItem,
	"idCount":                   IdCount,
	"mapKeyExpr":                MapKeyExpr,
	"mapKeyExample":             MapKeyExample,
	"mapKeyParse":               MapKeyParse,
	"mapKeySjsonSetStmts":       MapKeySjsonSetStmts,
	"mapKeyMatchExpr":           MapKeyMatchExpr,
	"mapKeyRnFormatArgs":        MapKeyRnFormatArgs,
	"mapKeyDescription":         MapKeyDescription,
}

// AllChildClassNames recursively collects all child class names.
func AllChildClassNames(children []YamlConfigChildClass) []string {
	var names []string
	for _, c := range children {
		names = append(names, c.ClassName)
		names = append(names, AllChildClassNames(c.ChildClasses)...)
	}
	return names
}

// HasNonIdAttrs returns true if there is at least one non-id attribute (excluding reference_only and value-only).
func HasNonIdAttrs(attrs []YamlConfigAttribute) bool {
	for _, a := range attrs {
		if !a.Id && !a.ReferenceOnly && a.Value == "" {
			return true
		}
	}
	return false
}

// NeedsPlanItem returns true if any child in the list needs a planItem variable
// for delete detection (list children always need it; single children need it only if they have child_classes).
func NeedsPlanItem(children []YamlConfigChildClass) bool {
	for _, c := range children {
		if c.Type == "list" {
			return true
		}
		if c.Type == "single" && len(c.ChildClasses) > 0 {
			return true
		}
	}
	return false
}

// IdCount returns the number of id: true attributes.
func IdCount(attrs []YamlConfigAttribute) int {
	n := 0
	for _, a := range attrs {
		if a.Id {
			n++
		}
	}
	return n
}

// idAttrs returns only the id: true attributes in order.
func idAttrs(attrs []YamlConfigAttribute) []YamlConfigAttribute {
	var result []YamlConfigAttribute
	for _, a := range attrs {
		if a.Id {
			result = append(result, a)
		}
	}
	return result
}

// MapKeyExpr returns a Go expression to compute a map key from ID attributes
// read from a gjson result variable (e.g., "value").
func MapKeyExpr(varName string, attrs []YamlConfigAttribute) string {
	ids := idAttrs(attrs)
	if len(ids) == 0 {
		return `""`
	}
	if len(ids) == 1 {
		return varName + `.Get("attributes.` + ids[0].NxosName + `").String()`
	}
	var parts []string
	for _, a := range ids {
		parts = append(parts, varName+`.Get("attributes.`+a.NxosName+`").String()`)
	}
	return strings.Join(parts, ` + ";" + `)
}

// MapKeyExample returns the map key computed from Example values of ID attributes.
func MapKeyExample(attrs []YamlConfigAttribute) string {
	ids := idAttrs(attrs)
	if len(ids) == 0 {
		return ""
	}
	if len(ids) == 1 {
		return ids[0].Example
	}
	var parts []string
	for _, a := range ids {
		parts = append(parts, a.Example)
	}
	return strings.Join(parts, ";")
}

// MapKeyParse returns Go code to parse a map key back into individual ID values.
// For single ID: no code needed (the key IS the value).
// For composite: declares keyParts via strings.SplitN.
func MapKeyParse(keyVar string, attrs []YamlConfigAttribute) string {
	ids := idAttrs(attrs)
	if len(ids) <= 1 {
		return ""
	}
	return "keyParts := strings.SplitN(" + keyVar + `, ";", ` + fmt.Sprintf("%d", len(ids)) + ")"
}

// MapKeySjsonSetStmts returns Go code with sjson.Set calls to write ID values from the parsed key.
func MapKeySjsonSetStmts(keyVar string, attrs []YamlConfigAttribute) string {
	ids := idAttrs(attrs)
	if len(ids) == 0 {
		return ""
	}
	var lines []string
	if len(ids) == 1 {
		lines = append(lines, `attrs, _ = sjson.Set(attrs, "`+ids[0].NxosName+`", `+keyVar+`)`)
	} else {
		for i, a := range ids {
			lines = append(lines, fmt.Sprintf(`attrs, _ = sjson.Set(attrs, "%s", keyParts[%d])`, a.NxosName, i))
		}
	}
	return strings.Join(lines, "\n\t\t")
}

// MapKeyMatchExpr returns a Go boolean expression to check if a gjson result's
// ID attributes match the map key.
func MapKeyMatchExpr(keyVar string, className string, attrs []YamlConfigAttribute) string {
	ids := idAttrs(attrs)
	if len(ids) == 0 {
		return "true"
	}
	if len(ids) == 1 {
		return `v.Get("` + className + `.attributes.` + ids[0].NxosName + `").String() == ` + keyVar
	}
	var parts []string
	for i, a := range ids {
		parts = append(parts, fmt.Sprintf(`v.Get("%s.attributes.%s").String() == keyParts[%d]`, className, a.NxosName, i))
	}
	return strings.Join(parts, " &&\n\t\t\t\t\t")
}

// MapKeyRnFormatArgs returns Go expression arguments for fmt.Sprintf to build the RN from the parsed key.
func MapKeyRnFormatArgs(keyVar string, attrs []YamlConfigAttribute) string {
	ids := idAttrs(attrs)
	if len(ids) == 0 {
		return ""
	}
	if len(ids) == 1 {
		if ids[0].Type == "Int64" {
			return "helpers.Must(strconv.ParseInt(" + keyVar + ", 10, 64))"
		}
		return keyVar
	}
	var args []string
	for i, a := range ids {
		part := fmt.Sprintf("keyParts[%d]", i)
		if a.Type == "Int64" {
			part = fmt.Sprintf("helpers.Must(strconv.ParseInt(keyParts[%d], 10, 64))", i)
		}
		args = append(args, part)
	}
	return strings.Join(args, ", ")
}

// MapKeyDescription returns a MarkdownDescription suffix documenting the map key.
// The output uses literal backslash-n sequences because it is embedded in a Go string literal.
func MapKeyDescription(attrs []YamlConfigAttribute) string {
	ids := idAttrs(attrs)
	if len(ids) == 0 {
		return ""
	}
	var sb strings.Builder
	if len(ids) == 1 {
		a := ids[0]
		sb.WriteString(fmt.Sprintf("\\n  - Map key: `%s` - %s", a.TfName, a.Description))
		if len(a.EnumValues) > 0 {
			quoted := make([]string, len(a.EnumValues))
			for i, v := range a.EnumValues {
				quoted[i] = fmt.Sprintf("`%s`", v)
			}
			sb.WriteString(fmt.Sprintf("\\n  - Key choices: %s", strings.Join(quoted, ", ")))
		}
		if a.MinInt != 0 || a.MaxInt != 0 {
			sb.WriteString(fmt.Sprintf("\\n  - Key range: `%v`-`%v`", a.MinInt, a.MaxInt))
		}
	} else {
		tfNames := make([]string, len(ids))
		for i, a := range ids {
			tfNames[i] = a.TfName
		}
		sb.WriteString(fmt.Sprintf("\\n  - Map key format: `<%s>`", strings.Join(tfNames, ">;<")))
		for _, a := range ids {
			sb.WriteString(fmt.Sprintf("\\n  - Key component `%s`: %s", a.TfName, a.Description))
			if len(a.EnumValues) > 0 {
				quoted := make([]string, len(a.EnumValues))
				for i, v := range a.EnumValues {
					quoted[i] = fmt.Sprintf("`%s`", v)
				}
				sb.WriteString(fmt.Sprintf(" Choices: %s.", strings.Join(quoted, ", ")))
			}
			if a.MinInt != 0 || a.MaxInt != 0 {
				sb.WriteString(fmt.Sprintf(" Range: `%v`-`%v`.", a.MinInt, a.MaxInt))
			}
		}
	}
	return sb.String()
}

// buildTfChildClasses builds the TfChildClasses list by promoting children
// of auto-hidden classes (single-type classes where all attributes have a value).
func buildTfChildClasses(children []YamlConfigChildClass) []YamlConfigChildClass {
	var result []YamlConfigChildClass
	for _, c := range children {
		if c.Type == "single" && allAttributesHaveValue(c.Attributes) {
			// Promote this class's children up
			promoted := buildTfChildClasses(c.ChildClasses)
			result = append(result, promoted...)
		} else {
			flattened := c
			flattened.TfChildClasses = buildTfChildClasses(c.ChildClasses)
			result = append(result, flattened)
		}
	}
	return result
}

// buildChildTfChildClasses recursively sets TfChildClasses on all ChildClasses at every depth.
func buildChildTfChildClasses(children []YamlConfigChildClass) {
	for i := range children {
		children[i].TfChildClasses = buildTfChildClasses(children[i].ChildClasses)
		buildChildTfChildClasses(children[i].ChildClasses)
	}
}

// allAttributesHaveValue returns true if every attribute has a non-empty Value field.
func allAttributesHaveValue(attrs []YamlConfigAttribute) bool {
	for _, a := range attrs {
		if a.Value == "" {
			return false
		}
	}
	return true
}

func getTemplateSection(content, name string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	result := ""
	foundSection := false
	beginRegex := regexp.MustCompile(`\/\/template:begin\s` + name + `$`)
	endRegex := regexp.MustCompile(`\/\/template:end\s` + name + `$`)
	for scanner.Scan() {
		line := scanner.Text()
		if !foundSection {
			match := beginRegex.MatchString(line)
			if match {
				foundSection = true
				result += line + "\n"
			}
		} else {
			result += line + "\n"
			match := endRegex.MatchString(line)
			if match {
				foundSection = false
			}
		}
	}
	return result
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

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// create output file
	outputFile := filepath.Join(outputPath)
	existingFile, err := os.Open(outputPath)
	if err != nil {
		os.MkdirAll(filepath.Dir(outputFile), 0755)
	} else if strings.HasSuffix(templatePath, ".go") {
		existingScanner := bufio.NewScanner(existingFile)
		var newContent string
		currentSectionName := ""
		foundSections := false
		beginRegex := regexp.MustCompile(`\/\/template:begin\s(.*?)$`)
		endRegex := regexp.MustCompile(`\/\/template:end\s(.*?)$`)
		for existingScanner.Scan() {
			line := existingScanner.Text()
			if currentSectionName == "" {
				matches := beginRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] != "" {
					currentSectionName = matches[1]
					foundSections = true
				} else {
					newContent += line + "\n"
				}
			} else {
				matches := endRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] == currentSectionName {
					currentSectionName = ""
					newSection := getTemplateSection(string(output.Bytes()), matches[1])
					newContent += newSection
				}
			}
		}
		if foundSections && newContent != "" {
			output = bytes.NewBufferString(newContent)
		}
	}
	if existingFile != nil {
		existingFile.Close()
	}

	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	f.Write(output.Bytes())
}

func main() {
	resourceName := ""
	if len(os.Args) == 2 {
		resourceName = os.Args[1]
	}

	items, _ := os.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(items))
	// Iterate over definitions
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
		configs[i].TfChildClasses = buildTfChildClasses(configs[i].ChildClasses)
		buildChildTfChildClasses(configs[i].ChildClasses)
	}

	for _, config := range configs {
		if resourceName != "" && config.Name != resourceName {
			continue
		}
		// Iterate over templates
		for _, t := range templates {
			renderTemplate(t.path, t.prefix+SnakeCase(config.Name)+t.suffix, config)
		}
	}

	renderTemplate(providerTemplate, providerLocation, configs)
	renderTemplate(objectsTemplate, objectsLocation, configs)

	changelog, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading changelog: %v", err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))
}
