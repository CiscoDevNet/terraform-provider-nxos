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
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const definitionsPath = "./gen/definitions/"

type YamlConfigAttribute struct {
	NxosName      string  `yaml:"nxos_name"`
	TfName        string  `yaml:"tf_name"`
	Id            bool    `yaml:"id"`
	ReferenceOnly bool    `yaml:"reference_only"`
	DeleteValue   *string `yaml:"delete_value"`
}

type YamlConfigChildClass struct {
	ClassName    string                 `yaml:"class_name"`
	NoDelete     bool                   `yaml:"no_delete"`
	Attributes   []YamlConfigAttribute  `yaml:"attributes"`
	ChildClasses []YamlConfigChildClass `yaml:"child_classes"`
}

type YamlConfig struct {
	Name         string                 `yaml:"name"`
	ClassName    string                 `yaml:"class_name"`
	NoDelete     bool                   `yaml:"no_delete"`
	Attributes   []YamlConfigAttribute  `yaml:"attributes"`
	ChildClasses []YamlConfigChildClass `yaml:"child_classes"`
}

type issue struct {
	file      string
	classPath string
	attribute string
	nxosName  string
}

func checkAttributes(attrs []YamlConfigAttribute, fileName, classPath string) []issue {
	var issues []issue
	for _, attr := range attrs {
		if attr.Id || attr.ReferenceOnly {
			continue
		}
		if attr.DeleteValue == nil {
			name := attr.TfName
			if name == "" {
				name = attr.NxosName
			}
			issues = append(issues, issue{
				file:      fileName,
				classPath: classPath,
				attribute: name,
				nxosName:  attr.NxosName,
			})
		}
	}
	return issues
}

func checkClass(className string, noDelete bool, attrs []YamlConfigAttribute, children []YamlConfigChildClass, fileName, parentPath string) []issue {
	var issues []issue
	classPath := className
	if parentPath != "" {
		classPath = parentPath + " > " + className
	}

	if noDelete {
		issues = append(issues, checkAttributes(attrs, fileName, classPath)...)
	}

	for _, child := range children {
		issues = append(issues, checkClass(child.ClassName, child.NoDelete, child.Attributes, child.ChildClasses, fileName, classPath)...)
	}
	return issues
}

func main() {
	items, err := os.ReadDir(definitionsPath)
	if err != nil {
		log.Fatalf("Error reading definitions directory: %v", err)
	}

	var allIssues []issue

	for _, item := range items {
		if item.IsDir() {
			continue
		}
		yamlFile, err := os.ReadFile(filepath.Join(definitionsPath, item.Name()))
		if err != nil {
			log.Fatalf("Error reading file %s: %v", item.Name(), err)
		}

		var config YamlConfig
		if err := yaml.Unmarshal(yamlFile, &config); err != nil {
			log.Fatalf("Error parsing %s: %v", item.Name(), err)
		}

		issues := checkClass(config.ClassName, config.NoDelete, config.Attributes, config.ChildClasses, item.Name(), "")
		allIssues = append(allIssues, issues...)
	}

	if len(allIssues) > 0 {
		fmt.Printf("Found %d attribute(s) missing 'delete_value' in no_delete classes:\n\n", len(allIssues))
		for _, i := range allIssues {
			fmt.Printf("  File:      %s\n", i.file)
			fmt.Printf("  Class:     %s\n", i.classPath)
			fmt.Printf("  Attribute: %s (nxos_name: %s)\n\n", i.attribute, i.nxosName)
		}
		os.Exit(1)
	}

	fmt.Println("OK: All attributes in no_delete classes have 'delete_value' defined.")
}
