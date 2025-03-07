---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_icmpv4 Resource - terraform-provider-nxos"
subcategory: "ICMP"
description: |-
  This resource can manage the global ICMP configuration.
  API Documentation: icmpv4Entity https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/icmpv4:Entity/
  Child resources
  nxos_icmpv4_instance https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/icmpv4_instance
---

# nxos_icmpv4 (Resource)

This resource can manage the global ICMP configuration.

- API Documentation: [icmpv4Entity](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/icmpv4:Entity/)

### Child resources

- [nxos_icmpv4_instance](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/icmpv4_instance)

## Example Usage

```terraform
resource "nxos_icmpv4" "example" {
  admin_state = "enabled"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `admin_state` (String) Administrative state.
  - Choices: `enabled`, `disabled`
  - Default value: `enabled`
- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_icmpv4.example "sys/icmpv4"
```
