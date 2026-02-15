---
subcategory: "Guides"
page_title: "Selective Deploy"
description: |-
    Howto selectively deploy to devices.
---

# Selective Deploy

## Overview

Selective deployment allows you to deploy Terraform configurations to a subset of devices from your `devices` list, while keeping the remaining devices in a "frozen" state where no changes are applied. This capability enables staged rollouts, maintenance workflows, and risk mitigation strategies for large-scale network deployments.

### Key Benefits

- **Risk Reduction**: Deploy to a small subset first to validate changes
- **Staged Rollouts**: Progressively deploy across device groups
- **Maintenance Windows**: Exclude devices undergoing maintenance
- **Emergency Procedures**: Quickly isolate or target specific devices

### How It Works

When `selected_devices` is configured, only the specified devices from your `devices` list will be actively managed by Terraform. Non-selected devices maintain their current configuration state and are skipped during plan and apply operations. This setting overrides individual device `managed` attributes.

### Understanding State Behavior
When using selective deployment, Terraform maintains state for ALL devices but only modifies selected devices:

- **Selected Devices**: Configuration changes applied, state updated
- **Non-Selected Devices**: State preserved, no modifications made
- **State File**: Contains all devices regardless of selection

## Configuration Reference

### HCL Configuration

Configure selective deployment using the `selected_devices` attribute in your provider block:

```hcl
provider "nxos" {
  selected_devices = ["LEAF-01", "LEAF-02"]
  devices = [
    { name = "LEAF-01", url = "https://10.1.1.10" },  # Managed
    { name = "LEAF-02", url = "https://10.1.1.20" },  # Managed
    { name = "LEAF-03", url = "https://10.1.1.30" },  # Skipped
    { name = "SPINE-01", url = "https://10.1.1.40" }  # Skipped
  ]
}

```

### Environment Variable Configuration

Alternatively, use the `NXOS_SELECTED_DEVICES` environment variable with comma-separated device names:

```bash
export NXOS_SELECTED_DEVICES="LEAF-01,LEAF-02"
```

```hcl
provider "nxos" {
  devices = [
    { name = "LEAF-01", url = "https://10.1.1.10" },
    { name = "LEAF-02", url = "https://10.1.1.20" },
    { name = "LEAF-03", url = "https://10.1.1.30" },
    { name = "SPINE-01", url = "https://10.1.1.40" }
  ]
}

```

### Default Behavior

When `selected_devices` is not specified, all devices in the `devices` list are managed by default.
