---
subcategory: "Guides"
page_title: "Getting Started"
description: |-
    Getting Started
---

# Getting Started

This example demonstrates how the provider can be used to configure the system name, enable a feature, create a loopback interface, and add a VRF. The full example can be found here: [https://github.com/CiscoDevNet/terraform-provider-nxos/tree/main/examples/basic/getting_started](https://github.com/CiscoDevNet/terraform-provider-nxos/tree/main/examples/basic/getting_started)

First of all we need to add the necessary provider configuration to the Terraform configuration file:

```hcl
terraform {
  required_providers {
    nxos = {
      source = "CiscoDevNet/nxos"
    }
  }
}

provider "nxos" {
  username = "admin"
  password = "password"
  url      = "https://10.1.1.1"
}
```

Next we configure the system hostname.

```hcl
resource "nxos_system" "example" {
  name = "LEAF1"
}
```

We can then enable the BGP feature, which is a prerequisite for any BGP configuration.

```hcl
resource "nxos_feature_bgp" "example" {
  admin_state = "enabled"
}
```

Now we add a loopback interface.

```hcl
resource "nxos_loopback_interface" "example" {
  interface_id = "lo0"
  description  = "Loopback0"
}
```

And finally we create a VRF.

```hcl
resource "nxos_vrf" "example" {
  name        = "VRF1"
  description = "My VRF1 Description"
}
```
