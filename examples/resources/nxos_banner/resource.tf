resource "nxos_banner" "example" {
  motd_banner = "My MOTD Banner"
  exec_banner = "My Exec Banner"
}
