resource "nxos_user_management" "example" {
  alphabet_sequence         = 3
  description               = "User management policy"
  keyboard_sequence         = 3
  max_logins                = 4
  min_unique                = 3
  password_grace_time       = 5
  password_life_time        = 90
  password_max_length       = 127
  password_min_length       = 10
  password_secure_mode      = "no"
  password_strength_check   = "no"
  password_warning_time     = 10
  service_password_recovery = "no"
  users = [{
    name                     = "user1"
    account_status           = "active"
    allow_expired            = "yes"
    clear_password_history   = "no"
    description              = "Test user account"
    email                    = "user1@example.com"
    expiration               = "2030-01-01T00:00:00.000+00:00"
    expires                  = "no"
    first_name               = "John"
    force                    = "no"
    last_name                = "Doe"
    password_hash            = "unspecified"
    phone                    = "1234567890"
    password                 = "password123"
    password_encryption_type = "clear"
    shell_type               = "shellvsh"
    unix_user_id             = 100
    roles = [{
      name           = "network-operator"
      description    = "Operator role"
      privilege_type = "readPriv"
    }]
  }]
}
