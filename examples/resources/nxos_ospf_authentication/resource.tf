resource "nxos_ospf_authentication" "example" {
  instance_name       = "OSPF1"
  vrf_name            = "VRF1"
  interface_id        = "eth1/10"
  key                 = "0 mykey"
  key_id              = 1
  key_secure_mode     = false
  keychain            = "mykeychain"
  md5_key             = "0 mymd5key"
  md5_key_secure_mode = false
  type                = "none"
}
