resource "nxos_isis_interface" "example" {
  interface_id            = "eth1/10"
  authentication_check    = false
  authentication_check_l1 = false
  authentication_check_l2 = false
  authentication_key      = ""
  authentication_key_l1   = ""
  authentication_key_l2   = ""
  authentication_type     = "unknown"
  authentication_type_l1  = "unknown"
  authentication_type_l2  = "unknown"
  circuit_type            = "l2"
  vrf                     = "default"
  hello_interval          = 20
  hello_interval_l1       = 20
  hello_interval_l2       = 20
  hello_multiplier        = 4
  hello_multiplier_l1     = 4
  hello_multiplier_l2     = 4
  hello_padding           = "never"
  instance_name           = "ISIS1"
  metric_l1               = 1000
  metric_l2               = 1000
  mtu_check               = true
  mtu_check_l1            = true
  mtu_check_l2            = true
  network_type_p2p        = "on"
  passive                 = "l1"
  priority_l1             = 80
  priority_l2             = 80
  enable_ipv4             = true
}
