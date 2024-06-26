---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_pim_anycast_rp Resource - terraform-provider-nxos"
subcategory: "PIM"
description: |-
  This resource can manage the PIM Anycast RP configuration.
  API Documentation: pimAcastRPFuncP https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:AcastRPFuncP/
  Parent resources
  nxos_pim_vrf https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/pim_vrf
  Child resources
  nxos_pim_anycast_rp_peer https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/pim_anycast_rp_peer
---

# nxos_pim_anycast_rp (Resource)

This resource can manage the PIM Anycast RP configuration.

- API Documentation: [pimAcastRPFuncP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:AcastRPFuncP/)

### Parent resources

- [nxos_pim_vrf](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/pim_vrf)

### Child resources

- [nxos_pim_anycast_rp_peer](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/pim_anycast_rp_peer)

## Example Usage

```terraform
resource "nxos_pim_anycast_rp" "example" {
  vrf_name         = "default"
  local_interface  = "eth1/10"
  source_interface = "eth1/10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `vrf_name` (String) VRF name.

### Optional

- `device` (String) A device name from the provider configuration.
- `local_interface` (String) Must match first field in the output of `show intf brief`. Example: `eth1/1`.
- `source_interface` (String) Must match first field in the output of `show intf brief`. Example: `eth1/1`.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_pim_anycast_rp.example "sys/pim/inst/dom-[default]/acastrpfunc"
```
