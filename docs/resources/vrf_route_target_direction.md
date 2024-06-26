---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_vrf_route_target_direction Resource - terraform-provider-nxos"
subcategory: "VRF"
description: |-
  This resource can manage a VRF Route Target direction.
  API Documentation: rtctrlRttP https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttP/
  Parent resources
  nxos_vrf_route_target_address_family https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/vrf_route_target_address_family
  Child resources
  nxos_vrf_route_target https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/vrf_route_target
---

# nxos_vrf_route_target_direction (Resource)

This resource can manage a VRF Route Target direction.

- API Documentation: [rtctrlRttP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttP/)

### Parent resources

- [nxos_vrf_route_target_address_family](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/vrf_route_target_address_family)

### Child resources

- [nxos_vrf_route_target](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/vrf_route_target)

## Example Usage

```terraform
resource "nxos_vrf_route_target_direction" "example" {
  vrf                         = "VRF1"
  address_family              = "ipv4-ucast"
  route_target_address_family = "ipv4-ucast"
  direction                   = "import"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address_family` (String) Address family.
  - Choices: `ipv4-ucast`, `ipv6-ucast`
- `direction` (String) Route Target direction.
  - Choices: `import`, `export`
- `route_target_address_family` (String) Route Target Address Family.
  - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
- `vrf` (String) VRF name.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_vrf_route_target_direction.example "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]/ctrl-[ipv4-ucast]/rttp-[import]"
```
