resource "nxos_bgp_instance_domain_peercontainer_af" "example" {
  vrf                     = "default"
  template_name           = "SPINE-PEERS"
  af                      = "ipv4-ucast"
  control                 = "rr-client"
  send_community_extended = "enabled"
  send_community_standard = "enabled"
}
