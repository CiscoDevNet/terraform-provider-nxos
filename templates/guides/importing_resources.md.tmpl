---
subcategory: "Guides"
page_title: "Importing Resources"
description: |-
    Importing Resources
---

# Importing Resources

A resource can be imported by using the `terraform import` command or by using an `import` block in the configuration.

Most resources in this provider are singleton resources that manage all objects of a given type on a device (e.g., all VRFs, all loopback interfaces). These resources use an empty import identifier. An optional device name can be appended when managing multiple devices.

## String-Based Import

An example for importing a resource using the `import` block with a string-based `id` is shown below. Assuming we have the following resource in our configuration:

```terraform
resource "nxos_vrf" "example" {
  vrfs = {
    "VRF1" = {
      description = "My VRF1 Description"
    }
  }
}
```

We could add an import block to import the resource as shown below:

```terraform
import {
  to = nxos_vrf.example
  id = ""
}
```

This will populate the state for the `nxos_vrf` resource for the "default" device. When managing multiple devices, the device name can be used as the import identifier:

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
