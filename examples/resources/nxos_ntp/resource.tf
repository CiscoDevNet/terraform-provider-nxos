resource "nxos_ntp" "example" {
  admin_state          = "enabled"
  allow_control        = "enabled"
  allow_private        = "enabled"
  authentication_state = "enabled"
  logging              = "enabled"
  logging_level        = "error"
  master               = "enabled"
  master_stratum       = 4
  passive              = "enabled"
  rate_limit           = 5
  servers = [{
    name      = "1.2.3.4"
    vrf       = "management"
    type      = "server"
    key_id    = 10
    min_poll  = 4
    max_poll  = 6
    preferred = true
  }]
}
