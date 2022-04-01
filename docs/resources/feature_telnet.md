---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_feature_telnet Resource - terraform-provider-nxos"
subcategory: ""
description: |-
  This resource can manage the Telnet feature configuration.
  API Documentation: fmTelnet https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Telnet/
---

# nxos_feature_telnet (Resource)

This resource can manage the Telnet feature configuration.

- API Documentation: [fmTelnet](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Telnet/)

## Example Usage

```terraform
resource "nxos_feature_telnet" "example" {
  admin_state = "enabled"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `admin_state` (String) Administrative state.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_feature_telnet.example "sys/fm/telnet"
```