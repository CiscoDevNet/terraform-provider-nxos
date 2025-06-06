---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_isis_overload Data Source - terraform-provider-nxos"
subcategory: "ISIS"
description: |-
  This data source can read the IS-IS overload configuration.
  API Documentation: isisOverload https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Overload/
---

# nxos_isis_overload (Data Source)

This data source can read the IS-IS overload configuration.

- API Documentation: [isisOverload](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Overload/)

## Example Usage

```terraform
data "nxos_isis_overload" "example" {
  instance_name = "ISIS1"
  vrf           = "default"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `instance_name` (String) IS-IS instance name.
- `vrf` (String) VRF name.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.
- `startup_time` (Number) The overload startup time. The overload state begins when the switch boots up and ends at the time specified as the overload startup time.
