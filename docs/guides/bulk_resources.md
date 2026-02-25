---
subcategory: "Guides"
page_title: "Bulk Resources"
description: |-
    Bulk Resources
---

# Bulk Resources

Bulk resources are designed to manage multiple objects of the same type within a single resource block. Instead of creating and managing each object individually, bulk resources allow you to define a list of objects, which are then managed collectively. This approach is particularly useful for scenarios where you need to create, update, or delete many similar objects at once, such as VRFs, Interfaces, etc. Bulk resources significantly reduce the number of resources in Terraform state and require fewer API calls, making them more efficient for large-scale operations.

->Using bulk resources helps reduce the number of resources in Terraform state and API overhead, as all items are managed as a single resource and pushed to the device in a single POST request rather than individual requests per object.

## How Bulk Resources Work

Bulk resources leverage the NX-API REST interface to push all items as children of a parent DME object in a single POST request. For example, when managing VRFs, all VRF objects (`l3Inst`) are sent as children of the parent `topSystem` object. This means creating 100 VRFs requires a single API call instead of 100 individual calls.

During updates, the provider compares the planned items with the current state to determine which items have been added, modified, or removed. New and modified items are sent in a single POST, while removed items are deleted individually.

## Example Usage

Here is an example of how to use the `nxos_vrfs` bulk resource to manage multiple VRFs:

```hcl
resource "nxos_vrfs" "example" {
  items = [
    {
      name        = "VRF1"
      description = "My VRF1 Description"
      encap       = "vxlan-103901"
    },
    {
      name        = "VRF2"
      description = "My VRF2 Description"
      encap       = "vxlan-103902"
    },
    {
      name        = "VRF3"
      description = "My VRF3 Description"
      encap       = "vxlan-103903"
    },
    {
      name        = "VRF4"
      description = "My VRF4 Description"
      encap       = "vxlan-103904"
    }
  ]
}
```

## Changing Objects

The collection of objects (`items`) is managed as a `List`, an ordered data structure. When you modify the `items` collection, the provider compares the planned items with the current state using the identity attributes (e.g., `name` for VRFs) to efficiently determine which elements have been added, removed, or changed, and only sends the necessary updates to the device.

For example, changing the description of a VRF:

```
Terraform will perform the following actions:

  # nxos_vrfs.example will be updated in-place
  ~ resource "nxos_vrfs" "example" {
        id    = "sys"
      ~ items = [
          ~ {
              ~ description = "My VRF3 Description" -> "My VRF3 Updated Description"
                name        = "VRF3"
                # (1 unchanged attribute hidden)
            },
            # (3 unchanged elements hidden)
        ]
    }

Plan: 0 to add, 1 to change, 0 to destroy.
```

## Importing Bulk Resources

To import all items of a bulk resource:

```hcl
terraform import nxos_vrfs.example ""
```

If the bulk resource has parent reference attributes, those values are provided as comma-separated import identifiers. An optional trailing device name can be added for multi-device setups:

```hcl
terraform import nxos_vrfs.example "mydevice"
```

The import fetches all children of the parent DN and populates the state with all discovered items.
