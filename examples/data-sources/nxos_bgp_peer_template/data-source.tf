data "nxos_bgp_peer_template" "example" {
  vrf           = "default"
  template_name = "SPINE-PEERS"
}
