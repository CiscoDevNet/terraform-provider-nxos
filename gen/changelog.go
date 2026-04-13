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
	"sort"
	"strings"
)

const (
	fragmentDir   = ".changelog"
	changelogFile = "CHANGELOG.md"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run gen/changelog.go <preview|release> [version]")
	}

	switch os.Args[1] {
	case "preview":
		entries := readFragments()
		if len(entries) == 0 {
			fmt.Println("No changelog fragments found.")
			return
		}
		fmt.Println("## Unreleased")
		fmt.Println()
		for _, entry := range entries {
			fmt.Println(entry)
		}
	case "release":
		if len(os.Args) < 3 {
			log.Fatal("Usage: go run gen/changelog.go release <version>")
		}
		version := os.Args[2]
		release(version)
	default:
		log.Fatalf("Unknown command: %s (expected 'preview' or 'release')", os.Args[1])
	}
}

// readFragments reads all .md files from the fragment directory (excluding README.md),
// collects lines starting with "- ", and returns them sorted by source filename.
func readFragments() []string {
	entries, err := os.ReadDir(fragmentDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		log.Fatalf("Error reading fragment directory: %v", err)
	}

	// Collect and sort fragment filenames for deterministic output
	var filenames []string
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".md") || strings.EqualFold(name, "README.md") {
			continue
		}
		filenames = append(filenames, name)
	}
	sort.Strings(filenames)

	var lines []string
	for _, name := range filenames {
		content, err := os.ReadFile(filepath.Join(fragmentDir, name))
		if err != nil {
			log.Fatalf("Error reading fragment %s: %v", name, err)
		}
		for _, line := range strings.Split(strings.TrimSpace(string(content)), "\n") {
			line = strings.TrimRight(line, " \t\r")
			if strings.HasPrefix(line, "- ") {
				lines = append(lines, line)
			}
		}
	}
	return lines
}

// release collects fragments, prepends a new version section to CHANGELOG.md,
// and deletes the fragment files.
func release(version string) {
	entries := readFragments()
	if len(entries) == 0 {
		log.Fatal("No changelog fragments found in .changelog/")
	}

	existing, err := os.ReadFile(changelogFile)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error reading %s: %v", changelogFile, err)
	}

	var sb strings.Builder
	sb.WriteString("## ")
	sb.WriteString(version)
	sb.WriteString("\n\n")
	for _, entry := range entries {
		sb.WriteString(entry)
		sb.WriteString("\n")
	}
	sb.WriteString("\n")
	sb.WriteString(strings.TrimLeft(string(existing), "\n"))

	if err := os.WriteFile(changelogFile, []byte(sb.String()), 0644); err != nil {
		log.Fatalf("Error writing %s: %v", changelogFile, err)
	}

	// Delete fragment files
	dirEntries, _ := os.ReadDir(fragmentDir)
	for _, entry := range dirEntries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".md") || strings.EqualFold(name, "README.md") {
			continue
		}
		if err := os.Remove(filepath.Join(fragmentDir, name)); err != nil {
			log.Printf("Warning: could not delete fragment %s: %v", name, err)
		}
	}

	fmt.Printf("Released %s: %d entries added to %s, fragments deleted.\n", version, len(entries), changelogFile)
}
