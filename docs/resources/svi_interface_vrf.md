---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_svi_interface_vrf Resource - terraform-provider-nxos"
subcategory: "Interface"
description: |-
  This resource can manage an SVI interface VRF association.
  API Documentation: nwRtVrfMbr https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/nw:RtVrfMbr/
  Parent resources
  nxos_svi_interface https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/svi_interface
  Referenced resources
  nxos_vrf https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/vrf
---

# nxos_svi_interface_vrf (Resource)

This resource can manage an SVI interface VRF association.

- API Documentation: [nwRtVrfMbr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/nw:RtVrfMbr/)

### Parent resources

- [nxos_svi_interface](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/svi_interface)

### Referenced resources

- [nxos_vrf](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/vrf)

## Example Usage

```terraform
resource "nxos_svi_interface_vrf" "example" {
  interface_id = "vlan293"
  vrf_dn       = "sys/inst-VRF123"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_id` (String) Must match first field in the output of `show intf brief`. Example: `vlan100`.
- `vrf_dn` (String) DN of VRF. For example: `sys/inst-VRF1`.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_svi_interface_vrf.example "sys/intf/svi-[vlan293]/rtvrfMbr"
```
