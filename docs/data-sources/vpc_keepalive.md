---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_vpc_keepalive Data Source - terraform-provider-nxos"
subcategory: "vPC"
description: |-
  This data source can read the vPC keepalive configuration.
  API Documentation: vpcKeepalive https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/vpc:Keepalive/
---

# nxos_vpc_keepalive (Data Source)

This data source can read the vPC keepalive configuration.

- API Documentation: [vpcKeepalive](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/vpc:Keepalive/)

## Example Usage

```terraform
data "nxos_vpc_keepalive" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `destination_ip` (String) vPC Keepalive destination address.
- `flush_timeout` (Number) vPC Keepalive flush timeout.
- `id` (String) The distinguished name of the object.
- `interval` (Number) vPC Keepalive interval.
- `precedence_type` (Number) vPC Keepalive precedence type. `0` - network, `1` - internet, `2` - critical, `3` flash-override, `4` - flash, `5` - immediate, `6` - prioriy, `7` - routine.
- `precedence_value` (Number) vPC Keepalive precedence value.
- `source_ip` (String) vPC Keepalive source address.
- `timeout` (Number) vPC Keepalive timeout.
- `type_of_service_byte` (Number) vPC Keepalive type of service (ToS) byte.
- `type_of_service_configuration_type` (Number) vPC Keepalive type of service (ToS) configuration type. `0` - noCfg, `1` - tos-byte, `2` - tos-value, `3` - tos-type, `4` -  precedence-type, `5` - precedence-value.
- `type_of_service_type` (Number) vPC Keepalive type of service (ToS) type. `0` - min-delay, `1` - max-throughput, `2` - max-reliability, `3` - min-monetary-cost, `4` -  normal.
- `type_of_service_value` (Number) vPC Keepalive type of service (ToS) value.
- `udp_port` (Number) vPC Keepalive UDP port.
- `vrf` (String) vPC Keepalive VRF.