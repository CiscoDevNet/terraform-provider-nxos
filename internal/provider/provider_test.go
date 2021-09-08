package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProvider *schema.Provider

func init() {
	testAccProvider = New("dev")()
}

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"nxos": func() (*schema.Provider, error) {
		return testAccProvider, nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
	if v := os.Getenv("NXOS_USERNAME"); v == "" {
		t.Fatal("NXOS_USERNAME env variable must be set for acceptance tests")
	}
	if v := os.Getenv("NXOS_PASSWORD"); v == "" {
		t.Fatal("NXOS_PASSWORD env variable must be set for acceptance tests")
	}
	if v := os.Getenv("NXOS_URL"); v == "" {
		t.Fatal("NXOS_URL env variable must be set for acceptance tests")
	}
}
