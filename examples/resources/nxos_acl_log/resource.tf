resource "nxos_acl_log" "example" {
  detailed        = true
  entries         = 8000
  include_mac     = true
  include_sgt     = true
  interval        = 300
  match_log_level = 3
  threshold       = 100
}
