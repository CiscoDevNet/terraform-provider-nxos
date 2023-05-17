---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_route_map_rule_entry_set_regular_community_item Data Source - terraform-provider-nxos"
subcategory: "Routing"
description: |-
  This data source can read a Set Community item in a Route-Map Rule Entry.
  API Documentation: rtregcomItem https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtregcom:Item/
---

# nxos_route_map_rule_entry_set_regular_community_item (Data Source)

This data source can read a Set Community item in a Route-Map Rule Entry.

- API Documentation: [rtregcomItem](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtregcom:Item/)

## Example Usage

```terraform
data "nxos_route_map_rule_entry_set_regular_community_item" "example" {
  rule_name = "RULE1"
  order     = 10
  community = "regular:as2-nn2:65001:123"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `community` (String) Set Community.
- `order` (Number) Route-Map Rule Entry order.
- `rule_name` (String) Route Map rule name.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

