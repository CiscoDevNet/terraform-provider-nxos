resource "nxos_user_management" "example" {
  alphabet_sequence             = 3
  description                   = "User management policy"
  keyboard_sequence             = 3
  max_logins                    = 4
  min_unique                    = 3
  password_grace_time           = 5
  password_life_time            = 90
  password_max_length           = 127
  password_min_length           = 10
  password_secure_mode          = "no"
  password_strength_check       = "no"
  password_warning_time         = 10
  service_password_recovery     = "no"
  pre_login_banner_description  = "Pre-login banner"
  pre_login_banner_owner_key    = "owner1"
  pre_login_banner_message      = "Authorized users only."
  pre_login_banner_owner_tag    = "tag1"
  post_login_banner_description = "Post-login banner"
  post_login_banner_owner_key   = "owner1"
  post_login_banner_message     = "Welcome to the system."
  post_login_banner_owner_tag   = "tag1"
  users = {
    "user1" = {
      account_status         = "active"
      allow_expired          = "yes"
      clear_password_history = "no"
      description            = "Test user account"
      email                  = "user1@example.com"
      expiration             = "2030-01-01T00:00:00.000+00:00"
      expires                = "no"
      first_name             = "John"
      force                  = "no"
      last_name              = "Doe"
      phone                  = "1234567890"
      shell_type             = "shellvsh"
      roles = {
        "network-operator" = {
          description    = "Operator role"
          privilege_type = "readPriv"
        }
      }
    }
  }
  tacacs_deadtime         = 5
  tacacs_description      = "TACACS+ settings"
  tacacs_owner_key        = "owner1"
  tacacs_owner_tag        = "tag1"
  tacacs_retries          = 3
  tacacs_source_interface = "unspecified"
  tacacs_timeout          = 10
  tacacs_providers = {
    "10.1.1.1" = {
      authentication_protocol = "chap"
      description             = "TACACS+ provider"
      monitoring_idle_time    = 10
      owner_key               = "owner1"
      owner_tag               = "tag1"
      port                    = 149
      retries                 = 3
      single_connection       = "yes"
      timeout                 = 10
    }
  }
  tacacs_provider_groups = {
    "TACACS_GROUP1" = {
      deadtime         = 5
      description      = "TACACS+ provider group"
      owner_key        = "owner1"
      owner_tag        = "tag1"
      source_interface = "unspecified"
      vrf              = "default"
      servers = {
        "10.1.1.1" = {
          description = "TACACS+ server reference"
          order       = 1
          owner_key   = "owner1"
          owner_tag   = "tag1"
        }
      }
    }
  }
  authentication_realm_default_role_policy     = "no-login"
  authentication_realm_description             = "AAA authentication realm"
  authentication_realm_owner_key               = "owner1"
  authentication_realm_owner_tag               = "tag1"
  authentication_realm_radius_directed_request = "yes"
  authentication_realm_tacacs_directed_request = "yes"
  default_authentication_protocol              = "chap"
  default_authentication_description           = "Default authentication"
  default_authentication_error_enable          = true
  default_authentication_fallback              = "no"
  default_authentication_invalid_user_log      = true
  default_authentication_local                 = "no"
  default_authentication_none                  = "no"
  default_authentication_owner_key             = "owner1"
  default_authentication_owner_tag             = "tag1"
  default_authentication_provider_group        = "TACACS_GROUP1"
  default_authentication_realm                 = "tacacs"
  console_authentication_protocol              = "chap"
  console_authentication_description           = "Console authentication"
  console_authentication_error_enable          = true
  console_authentication_fallback              = "no"
  console_authentication_invalid_user_log      = true
  console_authentication_local                 = "no"
  console_authentication_none                  = "no"
  console_authentication_owner_key             = "owner1"
  console_authentication_owner_tag             = "tag1"
  console_authentication_provider_group        = "TACACS_GROUP1"
  console_authentication_realm                 = "tacacs"
  default_authorizations = {
    "config" = {
      authorization_method_none = false
      description               = "Default authorization"
      local_rbac                = true
      owner_key                 = "owner1"
      owner_tag                 = "tag1"
      provider_group            = "TACACS_GROUP1"
    }
  }
  default_accounting_method_none    = false
  default_accounting_description    = "Default accounting"
  default_accounting_local_rbac     = true
  default_accounting_owner_key      = "owner1"
  default_accounting_owner_tag      = "tag1"
  default_accounting_provider_group = "TACACS_GROUP1"
  default_accounting_realm          = "tacacs"
}
