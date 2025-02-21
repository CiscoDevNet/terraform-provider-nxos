resource "nxos_isis_vrf_overload" "example" {
  instance_name = "ISIS1"
  vrf           = "default"
  startup_time  = 60
}
