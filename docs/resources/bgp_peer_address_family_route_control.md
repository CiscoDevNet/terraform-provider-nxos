---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_bgp_peer_address_family_route_control Resource - terraform-provider-nxos"
subcategory: "BGP"
description: |-
  This resource can manage the BGP peer address family route control configuration.
  API Documentation: bgpRtCtrlP https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:RtCtrlP/
  Parent resources
  nxos_bgp_peer_address_family https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/bgp_peer_address_family
---

# nxos_bgp_peer_address_family_route_control (Resource)

This resource can manage the BGP peer address family route control configuration.

- API Documentation: [bgpRtCtrlP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:RtCtrlP/)

### Parent resources

- [nxos_bgp_peer_address_family](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/bgp_peer_address_family)

## Example Usage

```terraform
resource "nxos_bgp_peer_address_family_route_control" "example" {
  asn            = "65001"
  vrf            = "default"
  address        = "192.168.0.1"
  address_family = "ipv4-ucast"
  direction      = "in"
  route_map_name = "ROUTE_MAP1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String) Peer address.
- `address_family` (String) Address Family.
  - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate`
  - Default value: `ipv4-ucast`
- `asn` (String) Autonomous system number.
- `direction` (String) Route Control direction.
  - Choices: `in`, `out`
  - Default value: `in`
- `vrf` (String) VRF name.

### Optional

- `device` (String) A device name from the provider configuration.
- `route_map_name` (String) Route Control Route-Map name.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_bgp_peer_address_family_route_control.example "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]/af-[ipv4-ucast]/rtctrl-[in]"
```
