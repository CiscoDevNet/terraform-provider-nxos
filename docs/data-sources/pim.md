---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_pim Data Source - terraform-provider-nxos"
subcategory: ""
description: |-
  This data source can read the global PIM configuration.
  API Documentation: pimEntity https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:Entity/
---

# nxos_pim (Data Source)

This data source can read the global PIM configuration.

- API Documentation: [pimEntity](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:Entity/)

## Example Usage

```terraform
data "nxos_pim" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- **admin_state** (String) Administrative state.
- **id** (String) The distinguished name of the object.

