resource "nxos_ospfv3_interface_authentication" "example" {
  interface_id                       = "eth1/4"
  authentication_key                 = "10 MY_SECRET_KEY"
  authentication_key_encryption_type = "cisco-type-7"
  security_parameter_index           = 444
  authentication_type                = "md5"
}
