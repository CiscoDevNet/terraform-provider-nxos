---
subcategory: "Guides"
page_title: "Manage Multiple Devices"
description: |-
    Howto manage multiple devices.
---

# Manage Multiple Devices

When it comes to managing multiple NX-OS devices, one can create multiple provider configurations and distinguish them by `alias` ([documentation](https://www.terraform.io/language/providers/configuration#alias-multiple-provider-configurations)).

```terraform
provider "nxos" {
  alias    = "LEAF-1"
  username = "admin"
  password = "Cisco123"
  url      = "https://10.1.1.1"
}

provider "nxos" {
  alias    = "LEAF-2"
  username = "admin"
  password = "Cisco123"
  url      = "https://10.1.1.2"
}

resource "nxos_bridge_domain" "LEAF-1-VLAN10" {
  provider     = nxos.LEAF-1
  fabric_encap = "vlan-10"
  name         = "VLAN10"
}

resource "nxos_bridge_domain" "LEAF-2-VLAN10" {
  provider     = nxos.LEAF-2
  fabric_encap = "vlan-10"
  name         = "VLAN10"
}
```

The disadvantages here is that the `provider` attribute of resources cannot be dynamic and therefore cannot be used in combination with `for_each` as an example. The issue is being tracked [here](https://github.com/hashicorp/terraform/issues/24476).

This provider offers an alternative approach where mutliple devices can be managed by a single provider configuration and the optional `device` attribute, which is available in every resource and data source, can then be used to select the respective device. This assumes that every device uses the same credentials.

```terraform
locals {
  leafs = [
    {
      name = "LEAF-1"
      url  = "https://10.1.1.1"
    },
    {
      name = "LEAF-2"
      url  = "https://10.1.1.2"
    },
  ]
  spines = [
    {
      name = "SPINE-1"
      url  = "https://10.1.1.3"
    },
  ]
}

provider "nxos" {
  username = "admin"
  password = "Cisco123"
  devices  = concat(local.leafs, local.spines)
}

resource "nxos_bridge_domain" "VLAN10" {
  for_each     = toset([for leaf in local.leafs : leaf.name])
  device       = each.key
  fabric_encap = "vlan-10"
  name         = "VLAN010"
}
```

## Device-Level Management Control

### The "managed" Flag

Each device in the `devices` list supports an optional `managed` attribute that controls whether the device is actively managed by Terraform:

- **`managed = true`** (default): Device is actively managed - configurations are applied
- **`managed = false`**: Device is "frozen" - state preserved but no changes applied
- **Not specified**: Defaults to `true`

### Basic Managed Flag Usage

```hcl
locals {
  devices = [
    {
      name    = "LEAF-01"
      url     = "https://10.1.1.10"
      managed = true   # Actively managed
    },
    {
      name    = "LEAF-02"
      url     = "https://10.1.1.20"
      managed = false  # Temporarily frozen for maintenance
    },
    {
      name    = "SPINE-01"
      url     = "https://10.1.1.30"
      # managed defaults to true when not specified
    }
  ]
}

provider "nxos" {
  username = "admin"
  password = "Cisco123"
  devices  = local.devices
}

resource "nxos_bridge_domain" "VLAN10" {
  for_each     = toset([for device in local.devices : device.name])
  device       = each.key
  fabric_encap = "vlan-10"
  name         = "VLAN010"
}
```

**Result**:
- `LEAF-01` and `SPINE-01`: Configuration applied
- `LEAF-02`: Configuration skipped, existing state preserved

### Relationship with `selected_devices`

**Important**: The [`selected_devices` provider attribute](selective_deploy.md) takes precedence over individual `managed` flags:

- **When `selected_devices` is specified**: Individual `managed` flags are ignored
- **When `selected_devices` is not specified**: Individual `managed` flags are respected

#### Example: selected_devices Override
```hcl
provider "nxos" {
  username         = "admin"
  password         = "Cisco123"
  selected_devices = ["LEAF-01", "SPINE-01"]  # Only these devices managed
  devices = [
    { name = "LEAF-01", url = "https://10.1.1.10", managed = false },  # Overridden to managed=true
    { name = "LEAF-02", url = "https://10.1.1.20", managed = true },   # Overridden to managed=false
    { name = "SPINE-01", url = "https://10.1.1.30", managed = true }   # Remains managed=true
  ]
}
```

**Result**: Only `LEAF-01` and `SPINE-01` are managed, regardless of their individual `managed` flags.
