//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

const definitionsPath = "./gen/definitions/"

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
	Name           string                `yaml:"name"`
	ClassName      string                `yaml:"class_name"`
	Dn             string                `yaml:"dn"`
	DsDescription  string                `yaml:"ds_description"`
	ResDescription string                `yaml:"res_description"`
	Parents        []string              `yaml:"parents"`
	Attributes     []YamlConfigAttribute `yaml:"attributes"`
}

type YamlConfigAttribute struct {
	NxosName      string   `yaml:"nxos_name"`
	TfName        string   `yaml:"tf_name"`
	Type          string   `yaml:"type"`
	Id            bool     `yaml:"id"`
	ReferenceOnly bool     `yaml:"reference_only"`
	Mandatory     bool     `yaml:"mandatory"`
	ReadOnly      bool     `yaml:"read_only"`
	Description   string   `yaml:"description"`
	Example       string   `yaml:"example"`
	EnumValues    []string `yaml:"enum_values"`
	MinInt        int      `yaml:"min_int"`
	MaxInt        int      `yaml:"max_int"`
	DefaultValue  string   `yaml:"default_value"`
}

// Templating helper function to convert first character in string to uppercase
func ToTitle(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + string(s[1:])
	}
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

// Map of templating functions
var functions = template.FuncMap{
	"toTitle":      ToTitle,
	"camelCase":    CamelCase,
	"snakeCase":    SnakeCase,
	"hasId":        HasId,
	"getExampleDn": GetExampleDn,
	"isLast":       IsLast,
	"sprintf":      fmt.Sprintf,
}

func main() {
	items, _ := ioutil.ReadDir(definitionsPath)
	// Iterate over definitions
	for _, filename := range items {
		yamlFile, err := ioutil.ReadFile(filepath.Join(definitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}

		// Iterate over templates
		for _, t := range templates {
			file, err := os.Open(t.path)
			if err != nil {
				log.Fatalf("Error opening template: %v", err)
			}
			defer file.Close()

			// skip first line with 'build-ignore' directive for go files
			scanner := bufio.NewScanner(file)
			if strings.HasSuffix(t.path, ".go") {
				scanner.Scan()
			}
			var temp string
			for scanner.Scan() {
				temp = temp + scanner.Text() + "\n"
			}

			template, err := template.New(path.Base(t.path)).Funcs(functions).Parse(temp)
			if err != nil {
				log.Fatalf("Error parsing template: %v", err)
			}

			// create output file
			outputFile := filepath.Join(t.prefix + SnakeCase(config.Name) + t.suffix)
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

			if strings.HasSuffix(t.path, ".g") {
				// format go code
				fOutput, err := format.Source(output.Bytes())
				if err != nil {
					log.Fatalf("Error formatting go in %s: %v", t.path, err)
				}
				f.Write(fOutput)
			} else {
				f.Write(output.Bytes())
			}
		}
	}
}
