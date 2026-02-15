default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

# Update all files from a single definition
# Usage: make gen NAME="Logging"
# NAME: The name of the definition, e.g. "Logging"
.PHONY: gen
gen:
	go run ./gen/generator.go "$(NAME)"
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go run gen/doc_category.go
