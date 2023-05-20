resource "nxos_ipv4_static_route" "example" {
  vrf_name = "default"
  prefix   = "1.1.1.0/24"
  next_hops = [{
    interface_id = "unspecified"
    address      = "1.2.3.4"
    vrf_name     = "default"
    description  = "My Description"
    object       = 10
    preference   = 123
    tag          = 10
  }]
}
