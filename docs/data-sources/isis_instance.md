---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_isis_instance Data Source - terraform-provider-nxos"
subcategory: ""
description: |-
  This data source can read the IS-IS instance configuration.
  API Documentation: isisInst https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Inst/
---

# nxos_isis_instance (Data Source)

This data source can read the IS-IS instance configuration.

- API Documentation: [isisInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Inst/)

## Example Usage

```terraform
data "nxos_isis_instance" "example" {
  name = "ISIS1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) IS-IS instance name.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `admin_state` (String) Administrative state.
- `id` (String) The distinguished name of the object.

