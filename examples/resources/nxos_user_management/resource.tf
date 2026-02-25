resource "nxos_user_management" "example" {
  users = [{
    name                     = "user1"
    allow_expired            = "yes"
    password                 = "password123"
    password_encryption_type = "clear"
    roles = [{
      name = "network-operator"
    }]
  }]
}
