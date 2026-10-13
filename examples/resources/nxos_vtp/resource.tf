resource "nxos_vtp" "example" {
  admin_state          = "enabled"
  instance_admin_state = "enabled"
  domain               = "DOMAIN1"
  file_path            = "bootflash:/vtp.dat"
  mode                 = "server"
  password             = "secret_password"
  version              = 2
  interfaces = {
    "eth1/9" = {
      admin_state = "disabled"
    }
  }
}
