---
subcategory: "Guides"
page_title: "Importing Resources"
description: |-
    Importing Resources
---

# Importing Resources

A resource can be imported by using the `terraform import` command or by using an `import` block in the configuration.

Most resources in this provider are singleton resources that manage all objects of a given type on a device (e.g., all VRFs, all loopback interfaces). These resources use an empty import identifier. An optional device name can be appended when managing multiple devices.

## Partial Configuration After Import

Because a singleton resource manages every object of its type on a device, importing one populates state with every existing object — not just the ones you plan to declare in configuration. If your configuration only declares a subset of what was imported (or omits some attributes on an entry you do declare), the first `terraform plan` after import will show those entries and attributes as being removed.

This is expected and safe: Terraform simply stops tracking anything you haven't declared, and nothing is changed on the device for it. Only the objects and attributes you explicitly configure are ever modified. A warning is included in the plan output as a reminder of this, and it will not appear again once you apply.

## String-Based Import

An example for importing a resource using the `terraform import` command with a string-based ID is shown below. Assuming we have the following resource in our configuration:

```terraform
resource "nxos_vrf" "example" {
  vrfs = {
    "VRF1" = {
      description = "My VRF1 Description"
    }
  }
}
```

We could import the resource for the "default" device using the following command:

```
terraform import nxos_vrf.example ""
```

-> **Note:** Terraform's `import` block requires a non-empty `id` value and will error with "Invalid import id argument" if it evaluates to an empty string. Since most resources in this provider use an empty string as their import identifier for the default device, use the `terraform import` command shown above, or identity-based import shown below, instead of an `import` block with `id = ""`.

When managing multiple devices, the device name can be used as the import identifier. Since this identifier isn't empty, it also works with an `import` block:

```terraform
resource "nxos_vrf" "example" {
  device = "LEAF-1"
  vrfs = {
    "VRF1" = {
      description = "My VRF1 Description"
    }
  }
}

import {
  to = nxos_vrf.example
  id = "LEAF-1"
}
```

## Identity-Based Import

Identity-based import requires **Terraform >= 1.12**, the version in which resource identity was introduced. On earlier Terraform versions, fall back to string-based import shown above.

Resources also support identity-based imports using an `identity` block. For singleton resources, the identity block is empty:

```terraform
import {
  to       = nxos_vrf.example
  identity = {}
}
```

When managing multiple devices, the `device` attribute can be included in the identity block:

```terraform
resource "nxos_vrf" "example" {
  device = "LEAF-1"
  vrfs = {
    "VRF1" = {
      description = "My VRF1 Description"
    }
  }
}

import {
  to = nxos_vrf.example
  identity = {
    "device" = "LEAF-1"
  }
}
```

## Importing the DME Resource

The generic `nxos_dme` resource uses a different import format. The import identifier is a comma-separated string of the DN and the NX-API class name:

```terraform
import {
  to = nxos_dme.example
  id = "sys/intf/phys-[eth1/1],l1PhysIf"
}
```

When managing multiple devices, the device name can be appended as a third element:

```terraform
import {
  to = nxos_dme.example
  id = "sys/intf/phys-[eth1/1],l1PhysIf,LEAF-1"
}
```

Or using identity-based import:

```terraform
import {
  to = nxos_dme.example
  identity = {
    "dn"         = "sys/intf/phys-[eth1/1]"
    "class_name" = "l1PhysIf"
    "device"     = "LEAF-1"
  }
}
```
