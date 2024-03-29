---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_feature_hsrp Data Source - terraform-provider-nxos"
subcategory: "Feature"
description: |-
  This data source can read the HSRP feature configuration.
  API Documentation: fmHsrp https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Hsrp/
---

# nxos_feature_hsrp (Data Source)

This data source can read the HSRP feature configuration.

- API Documentation: [fmHsrp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Hsrp/)

## Example Usage

```terraform
data "nxos_feature_hsrp" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `admin_state` (String) Administrative state.
- `id` (String) The distinguished name of the object.
