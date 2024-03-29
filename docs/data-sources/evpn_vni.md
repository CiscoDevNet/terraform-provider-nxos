---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_evpn_vni Data Source - terraform-provider-nxos"
subcategory: "EVPN"
description: |-
  This data source can read a EVPN VNI Route Distinguisher.
  API Documentation: rtctrlBDEvi https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:BDEvi/
---

# nxos_evpn_vni (Data Source)

This data source can read a EVPN VNI Route Distinguisher.

- API Documentation: [rtctrlBDEvi](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:BDEvi/)

## Example Usage

```terraform
data "nxos_evpn_vni" "example" {
  encap = "vxlan-123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `encap` (String) Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.
- `route_distinguisher` (String) Route Distinguisher value in NX-OS DME format.
