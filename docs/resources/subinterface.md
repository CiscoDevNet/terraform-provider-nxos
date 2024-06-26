---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_subinterface Resource - terraform-provider-nxos"
subcategory: "Interface"
description: |-
  This resource can manage a subinterface.
  API Documentation: l3EncRtdIf https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/l3:EncRtdIf/
  Child resources
  nxos_subinterface_vrf https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/subinterface_vrf
---

# nxos_subinterface (Resource)

This resource can manage a subinterface.

- API Documentation: [l3EncRtdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/l3:EncRtdIf/)

### Child resources

- [nxos_subinterface_vrf](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/subinterface_vrf)

## Example Usage

```terraform
resource "nxos_subinterface" "example" {
  interface_id = "eth1/10.124"
  admin_state  = "down"
  bandwidth    = 1000
  delay        = 10
  description  = "My Description"
  encap        = "vlan-124"
  link_logging = "enable"
  medium       = "broadcast"
  mtu          = 1500
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_id` (String) Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.

### Optional

- `admin_state` (String) Administrative state.
  - Choices: `up`, `down`
  - Default value: `up`
- `bandwidth` (Number) Specifies the administrative port bandwidth.
  - Range: `0`-`3200000000`
  - Default value: `0`
- `delay` (Number) Specifies the administrative port delay.
  - Range: `1`-`16777215`
  - Default value: `1`
- `description` (String) Interface description.
- `device` (String) A device name from the provider configuration.
- `encap` (String) Subinterface encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
  - Default value: `unknown`
- `link_logging` (String) Administrative link logging.
  - Choices: `default`, `enable`, `disable`
  - Default value: `default`
- `medium` (String) The administrative port medium type.
  - Choices: `broadcast`, `p2p`
  - Default value: `broadcast`
- `mtu` (Number) Administrative port MTU.
  - Range: `576`-`9216`
  - Default value: `1500`

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_subinterface.example "sys/intf/encrtd-[eth1/10.124]"
```
