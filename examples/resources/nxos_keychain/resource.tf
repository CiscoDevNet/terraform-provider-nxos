resource "nxos_keychain" "example" {
  admin_state = "enabled"
  keychains = [{
    name = "KEYCHAIN1"
    keys = [{
      key_id                  = 1
      cryptographic_algorithm = "AES"
      encryption_type         = "type7"
      key_string              = "secret_password"
    }]
  }]
}
