---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_loopback_interface_vrf Data Source - terraform-provider-nxos"
subcategory: "Interface"
description: |-
  This data source can read a loopback interface VRF association.
  API Documentation: nwRtVrfMbr https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/nw:RtVrfMbr/
---

# nxos_loopback_interface_vrf (Data Source)

This data source can read a loopback interface VRF association.

- API Documentation: [nwRtVrfMbr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/nw:RtVrfMbr/)

## Example Usage

```terraform
data "nxos_loopback_interface_vrf" "example" {
  interface_id = "lo123"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_id` (String) Must match first field in the output of `show intf brief`. Example: `lo123`.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.
- `vrf_dn` (String) DN of VRF. For example: `sys/inst-VRF1`.
