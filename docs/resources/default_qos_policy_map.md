---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_default_qos_policy_map Resource - terraform-provider-nxos"
subcategory: ""
description: |-
  This resource can manage the default QoS policy map configuration.
  API Documentation: ipqosPMapInst https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:PMapInst/
  Child resources
  nxosdefaultqospolicymapmatchclass_map https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/default_qos_policy_map_match_class_map
---

# nxos_default_qos_policy_map (Resource)

This resource can manage the default QoS policy map configuration.

- API Documentation: [ipqosPMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:PMapInst/)

### Child resources

- [nxos_default_qos_policy_map_match_class_map](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/default_qos_policy_map_match_class_map)

## Example Usage

```terraform
resource "nxos_default_qos_policy_map" "example" {
  name       = "PM1"
  match_type = "match-any"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) Policy map name.

### Optional

- **match_type** (String) Match type.
  - Default value: `match-all`

### Read-Only

- **id** (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_default_qos_policy_map.example "sys/ipqos/dflt/p/name-[PM1]"
```