---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_ipv4_interface_address Resource - terraform-provider-nxos"
subcategory: "IPv4"
description: |-
  This resource can manage an IPv4 interface address.
  API Documentation: ipv4Addr https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/ipv4:Addr/
  Parent resources
  nxos_ipv4_interface https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/ipv4_interface
---

# nxos_ipv4_interface_address (Resource)

This resource can manage an IPv4 interface address.

- API Documentation: [ipv4Addr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/ipv4:Addr/)

### Parent resources

- [nxos_ipv4_interface](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/ipv4_interface)

## Example Usage

```terraform
resource "nxos_ipv4_interface_address" "example" {
  vrf          = "default"
  interface_id = "eth1/10"
  address      = "24.63.46.49/30"
  type         = "primary"
  tag          = 1234
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String) IPv4 address.
- `interface_id` (String) Must match first field in the output of `show intf brief`. Example: `eth1/1`.
- `vrf` (String) VRF name.

### Optional

- `device` (String) A device name from the provider configuration.
- `tag` (Number) Route Tag
  - Default value: `0`
- `type` (String) Address type.
  - Choices: `primary`, `secondary`
  - Default value: `primary`

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_ipv4_interface_address.example "sys/ipv4/inst/dom-[default]/if-[eth1/10]/addr-[24.63.46.49/30]"
```
