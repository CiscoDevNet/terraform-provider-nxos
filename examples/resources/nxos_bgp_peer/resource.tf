resource "nxos_bgp_peer" "example" {
  asn               = "65001"
  vrf               = "default"
  address           = "192.168.0.1"
  remote_asn        = "65002"
  description       = "My description"
  peer_template     = "SPINE-PEERS"
  peer_type         = "fabric-internal"
  source_interface  = "lo0"
  hold_time         = 45
  keepalive         = 15
  ebgp_multihop_ttl = 5
  peer_control      = "bfd,dis-conn-check"
  password_type     = "LINE"
  password          = "secret_password"
}
