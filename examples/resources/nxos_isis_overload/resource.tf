resource "nxos_isis_overload" "example" {
  instance_name = "ISIS1"
  vrf           = "default"
  startup_time  = 60
}
