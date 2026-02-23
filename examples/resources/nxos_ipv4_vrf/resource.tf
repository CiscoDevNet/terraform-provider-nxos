resource "nxos_ipv4_vrf" "example" {
  name = "VRF1"
  static_routes = [{
    prefix = "1.1.1.0/24"
    next_hops = [{
      interface_id = "unspecified"
      address      = "1.2.3.4"
      vrf_name     = "default"
      description  = "My Description"
      object       = 10
      preference   = 123
      tag          = 10
    }]
  }]
  interfaces = [{
    interface_id = "eth1/10"
    drop_glean   = "disabled"
    forward      = "disabled"
    unnumbered   = "unspecified"
    urpf         = "disabled"
    addresses = [{
      address = "24.63.46.49/30"
      type    = "primary"
      tag     = 1234
    }]
  }]
}
