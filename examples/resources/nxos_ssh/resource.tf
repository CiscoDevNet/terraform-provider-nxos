resource "nxos_ssh" "example" {
  ssh_admin_state                  = "enabled"
  ssh_ciphers                      = "no"
  ssh_description                  = "SSH access"
  ssh_enable_weak_ciphers          = "no"
  ssh_key_exchange_algorithms      = "no"
  ssh_key_types                    = "no"
  ssh_login_attempts               = 4
  ssh_login_grace_time             = 60
  ssh_message_authentication_codes = "no"
  ssh_port                         = 22
  keys = {
    "rsa" = {
      key_length = 2048
    }
  }
}
