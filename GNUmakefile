default: test

# Load environment variables from .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Run acceptance tests
# Usage: make test [NAME="Definition Name"] [DEBUG=1]
# NAME must be a definition name (e.g., "BGP", "Loopback Interface") matching the name: field in gen/definitions/*.yaml
.PHONY: test
test:
	@echo "========================================="
	@echo "Running acceptance tests..."
	@echo "========================================="
	@if [ -z "$(NXOS_URL)" ]; then \
		echo "SKIPPED: NXOS_URL is not configured"; \
		echo "To enable tests, configure NXOS_URL, NXOS_USERNAME, and NXOS_PASSWORD in your .env file"; \
	else \
		TEST_NAME=""; \
		if [ -n "$(NAME)" ]; then \
			MATCH=$$(grep -rl '^name: $(NAME)$$' gen/definitions/*.yaml | head -1); \
			if [ -n "$${MATCH}" ]; then \
				CAMEL=$$(echo "$(NAME)" | tr -d ' '); \
				TEST_NAME="TestAcc.*Nxos$${CAMEL}"; \
				echo "Resolved definition '$(NAME)' to test pattern: $${TEST_NAME}"; \
			else \
				echo "ERROR: No definition found matching '$(NAME)' in gen/definitions/*.yaml"; \
				exit 1; \
			fi; \
		fi; \
		$(if $(DEBUG),echo "Debug mode enabled - logs will be written to test-output.log";) \
		if [ -n "$${TEST_NAME}" ]; then echo "Running tests matching: $${TEST_NAME}"; fi; \
		TF_ACC=1 \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $${TEST_NAME:+-run "$${TEST_NAME}"} $(TESTARGS) -count 1 -timeout 120m ./internal/provider $(if $(DEBUG),2>&1 | tee test-output.log); \
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

# Preview unreleased changelog entries from .changelog/ fragments
.PHONY: changelog-preview
changelog-preview:
	go run ./gen/changelog.go preview

# Finalize a release: collect changelog fragments into CHANGELOG.md and regenerate
# Usage: make release VERSION=X.Y.Z
.PHONY: release
release:
	@if [ -z "$(VERSION)" ]; then echo "Usage: make release VERSION=X.Y.Z"; exit 1; fi
	go run ./gen/changelog.go release $(VERSION)
	$(MAKE) gen
