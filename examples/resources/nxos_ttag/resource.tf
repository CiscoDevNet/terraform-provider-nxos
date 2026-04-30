resource "nxos_ttag" "example" {
  ttag_marker_interval = 120
  ttag_interfaces = {
    "eth1/10" = {
      ttag        = true
      ttag_inner  = false
      ttag_marker = false
      ttag_strip  = false
    }
  }
}
