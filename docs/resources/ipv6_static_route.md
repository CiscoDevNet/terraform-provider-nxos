---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_ipv6_static_route Resource - terraform-provider-nxos"
subcategory: "IPv6"
description: |-
  This resource can manage an IPv6 static route.
  API Documentation: ipv6Route https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/ipv6:Route/
  Parent resources
  nxos_ipv6_vrf https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/ipv6_vrf
---

# nxos_ipv6_static_route (Resource)

This resource can manage an IPv6 static route.

- API Documentation: [ipv6Route](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/ipv6:Route/)

### Parent resources

- [nxos_ipv6_vrf](https://registry.terraform.io/providers/CiscoDevNet/nxos/latest/docs/resources/ipv6_vrf)

## Example Usage

```terraform
resource "nxos_ipv6_static_route" "example" {
  vrf_name = "default"
  prefix   = "2001:db8:3333:4444:5555:6666:102:304/128"
  next_hops = [{
    interface_id = "unspecified"
    address      = "a:b::c:d/128"
    vrf_name     = "default"
    description  = "My Description"
    object       = 10
    preference   = 123
    tag          = 10
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `next_hops` (Attributes List) List of next hops. (see [below for nested schema](#nestedatt--next_hops))
- `prefix` (String) Prefix.
- `vrf_name` (String) VRF name.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

<a id="nestedatt--next_hops"></a>
### Nested Schema for `next_hops`

Required:

- `address` (String) Nexthop address.
- `interface_id` (String) Must match first field in the output of `show intf brief` or `unspecified`. Example: `eth1/1` or `vlan100`.
- `vrf_name` (String) Nexthop VRF.

Optional:

- `description` (String) Description.
- `object` (Number) Object to be tracked.
  - Range: `0`-`4294967295`
- `preference` (Number) Route preference.
  - Range: `0`-`255`
- `tag` (Number) Tag value.
  - Range: `0`-`4294967295`

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_ipv6_static_route.example "sys/ipv6/inst/dom-[default]/rt-[2001:db8:3333:4444:5555:6666:102:304/128]"
```
