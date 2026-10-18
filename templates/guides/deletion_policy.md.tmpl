---
subcategory: "Guides"
page_title: "Deletion Policy"
description: |-
    Control what happens on the device when a resource is destroyed.
---

# Deletion Policy

Most resources in this provider manage a whole part of the device configuration, for example `nxos_bgp` manages the BGP configuration below `sys/bgp`. Creating and updating a resource only ever changes what is declared in its configuration: configuration that already exists on the device and is not declared is left untouched. Destroying a resource is different, as it has to decide what to do with the objects it managed, which may also contain configuration that is not managed by Terraform.

The deletion policy controls this decision. It applies whenever Terraform destroys a resource: when running `terraform destroy`, when removing the resource from the configuration, and when the resource is replaced because an attribute that requires a replacement changed. The `delete` policy also deletes single child objects on update once none of their attributes are configured anymore, see [Updates](#updates).

| Policy | Effect on destroy | Typical use |
|---|---|---|
| `delete` (default) | Deletes the object managed by the resource, including all configuration below it, whether managed by Terraform or not. | Terraform owns this part of the configuration. |
| `reset` | Deletes the objects that the resource created, and resets the managed attributes of the objects that already existed before. Configuration not managed by Terraform stays in place. | Terraform manages some settings or entries within configuration that also has other owners. |
| `retain` | Leaves the device unchanged and only removes the resource from the Terraform state. | Hand the configuration over to another tool or person, or stop managing it without changing the device. |

## Configuration

The deletion policy of all resources is set in the provider configuration, or with the `NXOS_DELETION_POLICY` environment variable, and defaults to `delete`:

```terraform
provider "nxos" {
  deletion_policy = "reset"
}
```

The `deletion_policy` attribute of a resource overrides the policy of the provider for that resource:

```terraform
resource "nxos_bgp" "example" {
  deletion_policy = "retain"
  # ...
}
```

## Example

Consider a device with an existing BGP configuration, including two peers that are not managed by Terraform:

```
router bgp 65001
  neighbor 10.0.0.1
    remote-as 65002
  neighbor 10.0.0.2
    remote-as 65003
```

Terraform is used to add a third peer. To reach the peer, the configuration has to declare the BGP instance and the default VRF as well:

```terraform
resource "nxos_bgp" "example" {
  asn = "65001"
  vrfs = {
    "default" = {
      peers = {
        "10.0.0.5" = {
          remote_asn = "65004"
        }
      }
    }
  }
}
```

Applying this configuration adds `neighbor 10.0.0.5` and leaves the other two peers as they are. What happens when the resource is destroyed depends on the deletion policy.

### `delete`

The whole BGP configuration is deleted, including `neighbor 10.0.0.1` and `neighbor 10.0.0.2`, because they are part of the object the resource manages. Use this policy when Terraform is the only owner of the configuration the resource manages, for example on a device that is fully managed by Terraform.

### `reset`

```terraform
resource "nxos_bgp" "example" {
  deletion_policy = "reset"
  asn             = "65001"
  vrfs = {
    "default" = {
      peers = {
        "10.0.0.5" = {
          remote_asn = "65004"
        }
      }
    }
  }
}
```

Only `neighbor 10.0.0.5` is deleted. The BGP instance and the default VRF already existed before the resource managed them, so they stay, together with `neighbor 10.0.0.1` and `neighbor 10.0.0.2`.

To be able to tell them apart, the provider records which of the managed objects already exist on the device whenever it creates them. With the `reset` policy, destroying the resource then:

- **Deletes** every managed object that was created by the resource, including any configuration below it.
- **Resets** the managed attributes of every managed object that already existed to their default values, and leaves the object in place.
- **Leaves unchanged** all configuration that is not managed by Terraform, unless it is below a deleted object.

Attributes that identify an object, such as the BGP `asn`, are not reset. Objects that the device does not allow to be deleted (e.g. physical interfaces) are always reset.

### `retain`

```terraform
resource "nxos_bgp" "example" {
  deletion_policy = "retain"
  # ...
}
```

The device is not changed at all: `neighbor 10.0.0.5` stays configured, just like the other two peers. Terraform only forgets about the resource.

## Updates

When some attributes of a single child object are removed from the configuration, for example `graceful_restart_stale_interval` of a VRF in `nxos_bgp`, these attributes are reset to their default values, whatever the policy.

When the last configured attribute or child of a single child object is removed, so that nothing of the object is configured anymore, the policy decides:

- `delete` deletes the object. This also removes objects that the device rejects without the removed attributes, such as a PIM RP without an address. If the device holds list entries below the object that are not managed by the resource, for example IP SLA operations configured elsewhere below the IP SLA sender settings, or entries that NX-OS creates by itself, the removed attributes are reset instead and the object stays.
- `reset` and `retain` reset the removed attributes and leave the object in place.

Objects that the device does not allow to be deleted, and objects that only group other objects without attributes of their own, are always kept. Removing an entry from a list, e.g. a peer from `peers`, deletes that entry on the next apply regardless of the policy, including any configuration below it.

## Things to Know

- **Which policy applies on destroy.** When a resource is destroyed, its `deletion_policy` from the state is used, or the policy of the provider configuration at the time of the destroy if the resource does not set one. When changing the `deletion_policy` of a resource, apply the change first before removing the resource from the configuration, otherwise the previous policy is used.
- **Replacements use the policy too.** When an attribute that requires a replacement changes (e.g. the BGP `asn`), Terraform destroys the resource before it creates it again. With `retain`, the previous configuration stays on the device, and the new object is created on top of it, which the device may reject (e.g. a second BGP instance with a different ASN).
- **Imported configuration counts as existing.** All objects of an imported resource already existed before Terraform managed them, so the `reset` policy only resets them on destroy. Objects added to the configuration after the import are deleted on destroy, unless they already existed on the device when they were added.
- **Resources created by earlier provider versions.** For resources created before the deletion policy was introduced, the provider cannot know which objects already existed before, so the `reset` policy treats all of them as existing and only resets them on destroy.
- **Child objects kept on destroy.** A few child objects that the device removes by itself are never deleted or reset on destroy, whatever the policy. For example, the private VLAN associations and mappings of `nxos_system` are removed by the device once the private VLAN feature is disabled.
