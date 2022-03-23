---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_subinterface Data Source - terraform-provider-nxos"
subcategory: ""
description: |-
  This data source can read a subinterface.
---

# nxos_subinterface (Data Source)

This data source can read a subinterface.

## Example Usage

```terraform
data "nxos_subinterface" "example" {
  interface_id = "eth1/10.124"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **interface_id** (String) Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.

### Read-Only

- **admin_state** (String) Administrative state.
- **bandwidth** (Number) Specifies the administrative port bandwidth.
- **delay** (Number) Specifies the administrative port delay.
- **description** (String) Interface description.
- **encap** (String) Subinterface encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
- **id** (String) The distinguished name of the object.
- **link_logging** (String) Administrative link logging.
- **medium** (String) The administrative port medium type.
- **mtu** (Number) Administrative port MTU.

