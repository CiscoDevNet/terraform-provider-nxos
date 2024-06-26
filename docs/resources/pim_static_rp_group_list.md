---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_pim_static_rp_group_list Resource - terraform-provider-nxos"
subcategory: "PIM"
description: |-
  This resource can manage the PIM Static RP group list configuration.
  API Documentation: pimRPGrpList https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:RPGrpList/
  Parent resources
  nxos_pim_static_rp https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/pim_static_rp
---

# nxos_pim_static_rp_group_list (Resource)

This resource can manage the PIM Static RP group list configuration.

- API Documentation: [pimRPGrpList](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:RPGrpList/)

### Parent resources

- [nxos_pim_static_rp](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/pim_static_rp)

## Example Usage

```terraform
resource "nxos_pim_static_rp_group_list" "example" {
  vrf_name   = "default"
  rp_address = "1.2.3.4"
  address    = "224.0.0.0/4"
  bidir      = true
  override   = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String) Group list address information.
- `rp_address` (String) RP address.
- `vrf_name` (String) VRF name.

### Optional

- `bidir` (Boolean) Flag to treat Group Ranges as BiDir.
  - Default value: `false`
- `device` (String) A device name from the provider configuration.
- `override` (Boolean) Flag to override RP preference to use Static over Dynamic RP.
  - Default value: `false`

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_pim_static_rp_group_list.example "sys/pim/inst/dom-[default]/staticrp/rp-[1.2.3.4]/rpgrplist-[224.0.0.0/4]"
```
