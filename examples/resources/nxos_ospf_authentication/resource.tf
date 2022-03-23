resource "nxos_ospf_authentication" "example" {
  instance_name       = "OSPF1"
  vrf_name            = "default"
  interface_id        = "eth1/10"
  key                 = "mykey"
  key_id              = 1
  key_secure_mode     = false
  keychain            = "mykeychain"
  md5_key             = "mymdkey"
  md5_key_secure_mode = false
  type                = "none"
}
