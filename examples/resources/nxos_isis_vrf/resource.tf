resource "nxos_isis_vrf" "example" {
  instance_name           = "ISIS1"
  name                    = "default"
  admin_state             = "enabled"
  authentication_key_l1   = ""
  authentication_key_l2   = ""
  authentication_type_l1  = "unknown"
  authentication_type_l2  = "unknown"
  bandwidth_reference     = 400000
  banwidth_reference_unit = "mbps"
  is_type                 = "l2"
  metric_type             = "transition"
  mtu                     = 2000
  net                     = "49.0001.0000.0000.3333.00"
  passive_default         = "l12"
}
