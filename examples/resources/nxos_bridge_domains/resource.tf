resource "nxos_bridge_domains" "example" {
  svi_autostate = "disable"
  bridge_domains = [{
    fabric_encap        = "vlan-10"
    access_encap        = "unknown"
    name                = "VLAN10"
    bridge_domain_state = "suspend"
    admin_state         = "active"
    bridge_mode         = "mac"
    control             = "untagged"
    forwarding_control  = "mdst-flood"
    forwarding_mode     = "bridge"
    long_name           = false
    mac_packet_classify = "enable"
    mode                = "CE"
    vrf_name            = "default"
    cross_connect       = "disable"
  }]
}
