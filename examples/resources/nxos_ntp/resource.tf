resource "nxos_ntp" "example" {
  servers = [{
    name     = "1.2.3.4"
    vrf      = "management"
    type     = "server"
    key_id   = 10
    min_poll = 4
    max_poll = 6
  }]
}
