---
subcategory: "Guides"
page_title: "Bulk Resources"
description: |-
    Bulk Resources
---

# Bulk Resources

Bulk resources are designed to manage multiple objects of the same type within a single resource block. Instead of creating and managing each object individually, bulk resources allow you to define a set of objects, which are then managed collectively. This approach is particularly useful for scenarios where you need to create, update, or delete many similar objects at once, such as bridge domains, VLANs, etc. Bulk resources significantly reduce the number of resources in Terraform state and require fewer API calls, making them more efficient for large-scale operations.

->Using bulk resources helps reduce the number of resources in Terraform state and API overhead, as all items are managed as a single resource and pushed to the device in a single POST request rather than individual requests per object.

## How Bulk Resources Work

Bulk resources leverage the NX-API REST interface to push all items as children of a parent DME object in a single POST request. For example, when managing bridge domains, all bridge domain objects (`l2BD`) are sent as children of the parent `bdEntity` object. This means creating 100 bridge domains requires a single API call instead of 100 individual calls.

During updates, the provider compares the planned items with the current state to determine which items have been added, modified, or removed. New and modified items are sent in a single POST, while removed items are deleted individually.

## Example Usage

Here is an example of how to use the `nxos_bridge_domains` bulk resource to manage multiple bridge domains:

```hcl
resource "nxos_bridge_domains" "example" {
  items = [
    {
      fabric_encap = "vlan-10"
      name         = "VLAN10"
    },
    {
      fabric_encap = "vlan-20"
      name         = "VLAN20"
    },
    {
      fabric_encap = "vlan-30"
      name         = "VLAN30"
    },
    {
      fabric_encap = "vlan-40"
      name         = "VLAN40"
    }
  ]
}
```

## Changing Objects

The collection of objects (`items`) is managed as a `Set`, a data structure that stores unique elements without ordering. When you modify the `items` collection, you are adding or removing elements from a set rather than changing the order of elements. The provider can efficiently determine which elements have been added, removed, or changed, and only send the necessary updates to the device.

Please note that the use of a `Set` might show unexpected plan outputs, as elements within a set are immutable. Any changes to an element will appear as a new element being added and the old element being removed. For example, changing the name of a bridge domain:

```
Terraform will perform the following actions:

  # nxos_bridge_domains.example will be updated in-place
  ~ resource "nxos_bridge_domains" "example" {
        id    = "sys/bd"
      ~ items = [
          - {
              - fabric_encap = "vlan-30" -> null
              - name         = "VLAN30" -> null
            },
          + {
              + fabric_encap = "vlan-30"
              + name         = "VLAN30-Updated"
            },
            # (3 unchanged elements hidden)
        ]
    }

Plan: 0 to add, 1 to change, 0 to destroy.
```

->The Terraform plan output may look more dramatic than the actual changes performed by the provider. Even if the plan shows an element being "removed" and "added", the provider will only update the existing object in place on the device.

## Importing Bulk Resources

To import all items of a bulk resource:

```hcl
terraform import nxos_bridge_domains.example ""
```

If the bulk resource has parent reference attributes, those values are provided as comma-separated import identifiers. An optional trailing device name can be added for multi-device setups:

```hcl
terraform import nxos_bridge_domains.example "mydevice"
```

The import fetches all children of the parent DN and populates the state with all discovered items.
