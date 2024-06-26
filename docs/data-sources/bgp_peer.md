---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_bgp_peer Data Source - terraform-provider-nxos"
subcategory: "BGP"
description: |-
  This data source can read the BGP peer configuration.
  API Documentation: bgpPeer https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Peer/
---

# nxos_bgp_peer (Data Source)

This data source can read the BGP peer configuration.

- API Documentation: [bgpPeer](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Peer/)

## Example Usage

```terraform
data "nxos_bgp_peer" "example" {
  asn     = "65001"
  vrf     = "default"
  address = "192.168.0.1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String) Peer address.
- `asn` (String) Autonomous system number.
- `vrf` (String) VRF name.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `description` (String) Peer description.
- `ebgp_multihop_ttl` (Number) eBGP Multihop TTL
- `hold_time` (Number) BGP Hold Timer in seconds. The value must be greater than the keepalive timer
- `id` (String) The distinguished name of the object.
- `keepalive` (Number) BGP Keepalive Timer in seconds
- `password` (String) Password.
- `password_type` (String) Password Encryption Type.
- `peer_control` (String) Peer Controls. Choices: `bfd`, `dis-conn-check`, `cap-neg-off`, `no-dyn-cap`. Can be an empty string. Allowed formats:
  - Single value. Example: `bfd`
  - Multiple values (comma-separated). Example: `bfd,dis-conn-check`. In this case values must be in alphabetical order.
- `peer_template` (String) Peer template name.
- `peer_type` (String) Neighbor Fabric Type.
- `remote_asn` (String) Peer autonomous system number.
- `source_interface` (String) Source Interface. Must match first field in the output of `show intf brief`.
