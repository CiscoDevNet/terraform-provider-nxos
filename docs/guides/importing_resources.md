---
subcategory: "Guides"
page_title: "Importing Resources"
description: |-
    Importing Resources
---

# Importing Resources

A resource can be imported by using the `terraform import` command or by using an `import` block in the configuration. The resource documentation has an example for the `terraform import` command. The import identifier is a comma-separated list of the resource's identifying attributes (`reference_only` and `id` attributes from the resource documentation).

## String-Based Import

An example for importing a resource using the `import` block with a string-based `id` is shown below. Assuming we have the following resource in our configuration:

```terraform
resource "nxos_vrf" "vrf1" {
  name        = "VRF1"
  description = "My VRF1 Description"
}
```

We could add an import block to import the resource as shown below:

```terraform
import {
  to = nxos_vrf.vrf1
  id = "VRF1"
}
```

This will populate the state for the `nxos_vrf` resource for the "default" device. When managing multiple devices, the device name can be appended to the import identifier using a comma as a separator:

```terraform
resource "nxos_vrf" "vrf1" {
  device      = "LEAF-1"
  name        = "VRF1"
  description = "My VRF1 Description"
}

import {
  to = nxos_vrf.vrf1
  id = "VRF1,LEAF-1"
}
```

## Identity-Based Import

Resources also support identity-based imports using an `identity` block. This provides a structured alternative to string-based imports where each attribute is specified by name:

```terraform
import {
  to = nxos_vrf.vrf1
  identity = {
    "name" = "VRF1"
  }
}
```

When managing multiple devices, the `device` attribute can be included in the identity block:

```terraform
resource "nxos_vrf" "vrf1" {
  device      = "LEAF-1"
  name        = "VRF1"
  description = "My VRF1 Description"
}

import {
  to = nxos_vrf.vrf1
  identity = {
    "name"   = "VRF1"
    "device" = "LEAF-1"
  }
}
```
