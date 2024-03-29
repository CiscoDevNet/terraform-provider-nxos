---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_bgp_vrf Data Source - terraform-provider-nxos"
subcategory: "BGP"
description: |-
  This data source can read the BGP domain (VRF) configuration.
  API Documentation: bgpDom https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Dom/
---

# nxos_bgp_vrf (Data Source)

This data source can read the BGP domain (VRF) configuration.

- API Documentation: [bgpDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Dom/)

## Example Usage

```terraform
data "nxos_bgp_vrf" "example" {
  asn  = "65001"
  name = "default"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `asn` (String) Autonomous system number.
- `name` (String) VRF name.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.
- `router_id` (String) Router ID.
