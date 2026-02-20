default: test

# Load environment variables from .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Run acceptance tests
# Usage: make test [NAME=TestName] [DEBUG=1]
.PHONY: test
test:
	@echo "========================================="
	@echo "Running acceptance tests..."
	@echo "========================================="
	@if [ -z "$(NXOS_URL)" ]; then \
		echo "SKIPPED: NXOS_URL is not configured"; \
		echo "To enable tests, configure NXOS_URL, NXOS_USERNAME, and NXOS_PASSWORD in your .env file"; \
	else \
		$(if $(DEBUG),echo "Debug mode enabled - logs will be written to test-output.log";) \
		$(if $(NAME),echo "Running tests matching: $(NAME)";) \
		TF_ACC=1 \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $(if $(NAME),-run $(NAME)) $(TESTARGS) -count 1 -timeout 120m ./... $(if $(DEBUG),2>&1 | tee test-output.log); \
	fi

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
