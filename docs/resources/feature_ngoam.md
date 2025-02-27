---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_feature_ngoam Resource - terraform-provider-nxos"
subcategory: "Feature"
description: |-
  This resource can manage the VXLAN operations, administration, and maintenance feature.
  API Documentation: fmNgoam https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Ngoam/
---

# nxos_feature_ngoam (Resource)

This resource can manage the VXLAN operations, administration, and maintenance feature.

- API Documentation: [fmNgoam](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Ngoam/)

## Example Usage

```terraform
resource "nxos_feature_ngoam" "example" {
  admin_state = "enabled"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `admin_state` (String) Administrative state.
  - Choices: `enabled`, `disabled`

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_feature_ngoam.example "sys/fm/ngoam"
```
