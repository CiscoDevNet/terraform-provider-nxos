resource "nxos_keychain" "example" {
  admin_state = "enabled"
  keychains = [{
    name = "KEYCHAIN1"
    keys = [{
      key_id     = 1
      key_string = "secret_password"
    }]
  }]
}
