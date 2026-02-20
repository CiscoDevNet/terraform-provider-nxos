terraform {
  required_providers {
    nxos = {
      source = "CiscoDevNet/nxos"
    }
  }
}

provider "nxos" {
  // By default uses env $NXOS_USERNAME $NXOS_PASSWORD $NXOS_URL
}
