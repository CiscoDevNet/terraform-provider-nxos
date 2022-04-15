resource "nxos_bgp_instance" "example" {
  admin_state             = "enabled"
  asn                     = "65001"
  enhanced_error_handling = "no"
}
