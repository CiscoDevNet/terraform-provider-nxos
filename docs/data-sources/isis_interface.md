---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_isis_interface Data Source - terraform-provider-nxos"
subcategory: "ISIS"
description: |-
  This data source can read the IS-IS interface configuration.
  API Documentation: isisInternalIf https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:InternalIf/
---

# nxos_isis_interface (Data Source)

This data source can read the IS-IS interface configuration.

- API Documentation: [isisInternalIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:InternalIf/)

## Example Usage

```terraform
data "nxos_isis_interface" "example" {
  interface_id = "eth1/10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_id` (String) Must match first field in the output of `show intf brief`. Example: `eth1/1`.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `authentication_check` (Boolean) Authentication Check for ISIS without specific level.
- `authentication_check_l1` (Boolean) Authentication Check for ISIS on Level-1.
- `authentication_check_l2` (Boolean) Authentication Check for ISIS on Level-2.
- `authentication_key` (String) Authentication Key for IS-IS without specific level.
- `authentication_key_l1` (String) Authentication Key for IS-IS on Level-1.
- `authentication_key_l2` (String) Authentication Key for IS-IS on Level-2.
- `authentication_type` (String) IS-IS Authentication-Type without specific level.
- `authentication_type_l1` (String) IS-IS Authentication-Type for Level-1.
- `authentication_type_l2` (String) IS-IS Authentication-Type for Level-2.
- `circuit_type` (String) Circuit type.
- `enable_ipv4` (Boolean) Enabling ISIS router tag on Interface's IPV4 family.
- `hello_interval` (Number) Hello interval.
- `hello_interval_l1` (Number) Hello interval Level-1.
- `hello_interval_l2` (Number) Hello interval Level-2.
- `hello_multiplier` (Number) Hello multiplier.
- `hello_multiplier_l1` (Number) Hello multiplier Level-1.
- `hello_multiplier_l2` (Number) Hello multiplier Level-2.
- `hello_padding` (String) Hello padding.
- `id` (String) The distinguished name of the object.
- `instance_name` (String) Instance to which the interface belongs to.
- `metric_l1` (Number) Interface metric Level-1.
- `metric_l2` (Number) Interface metric Level-2.
- `mtu_check` (Boolean) MTU Check for IS-IS without specific level.
- `mtu_check_l1` (Boolean) MTU Check for IS-IS on Level-1.
- `mtu_check_l2` (Boolean) MTU Check for IS-IS on Level-2.
- `network_type_p2p` (String) Enabling Point-to-Point Network Type on IS-IS Interface.
- `passive` (String) IS-IS Passive Interface Info.
- `priority_l1` (Number) Circuit priority.
- `priority_l2` (Number) Circuit priority.
- `vrf` (String) VRF.
