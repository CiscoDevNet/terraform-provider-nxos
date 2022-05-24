resource "nxos_bgp_peer_template_address_family" "example" {
  asn                     = "65001"
  template_name           = "SPINE-PEERS"
  address_family          = "ipv4-ucast"
  control                 = "nh-self,rr-client"
  send_community_extended = "enabled"
  send_community_standard = "enabled"
}
