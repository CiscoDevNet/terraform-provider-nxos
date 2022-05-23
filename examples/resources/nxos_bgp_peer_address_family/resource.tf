resource "nxos_bgp_peer_address_family" "example" {
  vrf                     = "default"
  address                 = "192.168.0.1"
  address_family          = "ipv4-ucast"
  control                 = "nh-self,rr-client"
  send_community_extended = "enabled"
  send_community_standard = "enabled"
}
