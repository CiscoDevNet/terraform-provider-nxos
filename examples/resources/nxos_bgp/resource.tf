resource "nxos_bgp" "example" {
  admin_state             = "enabled"
  instance_admin_state    = "enabled"
  asn                     = "65001"
  enhanced_error_handling = false
  vrfs = [{
    name                               = "default"
    router_id                          = "1.1.1.1"
    route_control_enforce_first_as     = "disabled"
    route_control_fib_accelerate       = "enabled"
    route_control_log_neighbor_changes = "enabled"
    route_control_suppress_routes      = "disabled"
    graceful_restart_interval          = 240
    graceful_restart_stale_interval    = 1800
    address_families = [{
      address_family                         = "ipv4-ucast"
      critical_nexthop_timeout               = "2500"
      non_critical_nexthop_timeout           = "8000"
      advertise_l2vpn_evpn                   = "disabled"
      advertise_physical_ip_for_type5_routes = "disabled"
      max_ecmp_paths                         = 2
      max_external_ecmp_paths                = 1
      max_external_internal_ecmp_paths       = 1
      max_local_ecmp_paths                   = 1
      max_mixed_ecmp_paths                   = 1
      default_information_originate          = "disabled"
      next_hop_route_map_name                = "ROUTEMAP1"
      prefix_priority                        = "none"
      retain_rt_all                          = "disabled"
      advertise_only_active_routes           = "disabled"
      table_map_route_map_name               = "ROUTE_MAP1"
      vni_ethernet_tag                       = "disabled"
      wait_igp_converged                     = "disabled"
      advertised_prefixes = [{
        prefix    = "192.168.1.0/24"
        route_map = "rt-map"
        evpn      = "enabled"
      }]
      redistributions = [{
        protocol          = "ospf"
        protocol_instance = "OSPF1"
        route_map         = "route_map_ospf_1"
        scope             = "inter"
        srv6_prefix_type  = "unspecified"
      }]
    }]
    peer_templates = [{
      name             = "SPINE-PEERS"
      remote_asn       = "65002"
      description      = "My Description"
      peer_type        = "fabric-internal"
      source_interface = "lo0"
      peer_template_address_families = [{
        address_family          = "ipv4-ucast"
        control                 = "nh-self,rr-client"
        send_community_extended = "enabled"
        send_community_standard = "enabled"
        max_prefix_action       = "log"
        max_prefix_number       = 10000
        max_prefix_restart_time = 0
        max_prefix_threshold    = 30
      }]
    }]
    peers = [{
      address               = "192.168.0.1"
      remote_asn            = "65002"
      description           = "My description"
      peer_template         = "SPINE-PEERS"
      peer_type             = "fabric-internal"
      source_interface      = "lo0"
      hold_time             = 45
      keepalive             = 15
      ebgp_multihop_ttl     = 5
      peer_control          = "bfd,dis-conn-check"
      password_type         = "LINE"
      password              = "secret_password"
      local_asn_propagation = "no-prepend"
      local_asn             = "65001"
      peer_address_families = [{
        address_family          = "ipv4-ucast"
        control                 = "nh-self"
        send_community_extended = "enabled"
        send_community_standard = "enabled"
        route_controls = [{
          direction      = "in"
          route_map_name = "ROUTE_MAP1"
        }]
        prefix_list_controls = [{
          direction = "in"
          list      = "PREFIX_LIST1"
        }]
      }]
    }]
  }]
}
