resource "nxos_bgp_peer_template_address_family" "example" {
  template_name           = "SPINE-PEERS"
  address_family          = "ipv4-ucast"
  control                 = "rr-client"
  send_community_extended = "enabled"
  send_community_standard = "enabled"
}
