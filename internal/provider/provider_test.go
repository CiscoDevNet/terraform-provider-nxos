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

package provider

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"nxos": providerserver.NewProtocol6WithError(New("test")()),
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

// terraformVersionCapture implements tfversion.TerraformVersionCheck to capture
// the Terraform version detected by the test framework into a local variable.
// This allows SkipFunc on individual test steps to use the version without
// shelling out to the Terraform binary again.
type terraformVersionCapture struct {
	Version **version.Version
}

func (c terraformVersionCapture) CheckTerraformVersion(_ context.Context, req tfversion.CheckTerraformVersionRequest, _ *tfversion.CheckTerraformVersionResponse) {
	*c.Version = req.TerraformVersion
}

// skipBelowTerraformVersion returns a SkipFunc that skips a test step if the
// captured Terraform version is below the given minimum.
func skipBelowTerraformVersion(tfVersion **version.Version, min *version.Version) func() (bool, error) {
	return func() (bool, error) {
		if *tfVersion == nil {
			return true, nil
		}
		return (*tfVersion).LessThan(min), nil
	}
}
