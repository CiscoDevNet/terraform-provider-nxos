---
subcategory: "guides"
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
