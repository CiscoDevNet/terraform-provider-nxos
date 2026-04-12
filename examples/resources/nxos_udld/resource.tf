resource "nxos_udld" "example" {
  admin_state      = "enabled"
  aggressive       = "enabled"
  message_interval = 20
  interfaces = {
    "eth1/9" = {
      admin_state             = "port-enabled"
      aggressive              = "enabled"
      bidirectional_detection = "port-enabled"
    }
  }
}
