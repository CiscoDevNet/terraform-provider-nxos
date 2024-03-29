---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_pim_ssm_range Data Source - terraform-provider-nxos"
subcategory: "PIM"
description: |-
  This data source can read the PIM SSM range configuration.
  API Documentation: pimSSMRangeP https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:SSMRangeP/
---

# nxos_pim_ssm_range (Data Source)

This data source can read the PIM SSM range configuration.

- API Documentation: [pimSSMRangeP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:SSMRangeP/)

## Example Usage

```terraform
data "nxos_pim_ssm_range" "example" {
  vrf_name = "default"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `vrf_name` (String) VRF name.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `group_list_1` (String) Group list 1.
- `group_list_2` (String) Group list 2.
- `group_list_3` (String) Group list 3.
- `group_list_4` (String) Group list 4.
- `id` (String) The distinguished name of the object.
- `prefix_list` (String) Prefix list name.
- `route_map` (String) Route map name.
- `ssm_none` (Boolean) Exclude standard SSM range (232.0.0.0/8).
