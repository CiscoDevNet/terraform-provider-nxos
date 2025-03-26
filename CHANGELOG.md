## 0.5.9

- Add `managed` flag to provider device configuration to allow temporarily skipping a device due to maintenance
- Add `evpn` attribute to `nxos_bgp_advertised_prefix` resource and data source

## 0.5.8

- Check if a device referenced in a resource or data source exists in the provider configuration
- When importing a resource, also populate attributes referencing parent resources
- Add `nxos_isis_address_family` resource and data source
- Add `nxos_isis_overload` resource and data source
- Use JSON-RPC API to save config, [link](https://github.com/CiscoDevNet/terraform-provider-nxos/issues/298)

## 0.5.7

- Allow empty router ID in `nxos_gbp_vrf` resource, [link](https://github.com/CiscoDevNet/terraform-provider-nxos/issues/290)
- Add more attributes to `nxos_bgp_address_family` resource and data source
- Add `nxos_route_map_rule_entry_match_tag` resource and data source
- Add `nxos_feature_ngoam` resource and data source
- Add `tag` attribute to `nxos_ipv4_interface_address` resource and data source
- Add `nxos_icmpv4`, `nxos_icmpv4_instance`, `nxos_icmpv4_vrf` and `nxos_icmpv4_interface` resources and data sources

## 0.5.6

- Fix value range of `ebgp_multihop_ttl` attribute of `nxos_bgp_peer` resource
- Add `nxos_feature_ngmvpn` resource and data source
- Fix issue when updating `nxos_ipv4_access_list_entry` resource, [link](https://github.com/CiscoDevNet/terraform-provider-nxos/issues/204)

## 0.5.5

- Fix importing of resources

## 0.5.4

- Remove `control` default value of `nxos_ospf_vrf` resource
- Fix `peer_control` attribute of `nxos_bgp_peer` resource
- Fix `control` attribute of `nxos_ospf_vrf` resource

## 0.5.3

- Add `enable_ipv4` and `instance_name` attributes to `nxos_isis_interface` resource and data source
- Add `nxos_bgp_route_redistribution` resource and data source
- Add capability to reset object to default values when destroying resource, e.g. `nxos_physical_interface`
- Add `ebgp_multihop_ttl`, `peer_control`, `password_type` and `password` attributes to `nxos_bgp_peer` resource and data source
- Add `nxos_bgp_peer_local_asn` resource and data source
- Add `default_admin_status` attribute to `nxos_ethernet` resource and data source
- Add `control` attribute to `nxos_ospf_vrf` resource and data source

## 0.5.2

- Add `nxos_vpc_keepalive` resource and data source
- Add `nxos_vpc_peerlink` resource and data source
- Fix `bandwidth_reference_unit` of `nxos_ospf_vrf` resource and data source
- Add `force` attribute to `nxos_port_channel_interface_member` resource and data source
- Add `nxos_save_config` resource

## 0.5.1

- Add `nxos_bgp_advertised_prefix` resource and data source

## 0.5.0

- Migrate to `CiscoDevNet` registry namespace

## 0.4.2

- Fix issue with device authentication

## 0.4.1

- Fix issue with idempotency when multiple devices are used

## 0.4.0

- Add option to manage child objects with `nxos_rest` resource
- Add `access_list_name` attribute to `nxos_ipv4_access_list_policy_ingress_interface` resource and data source
- BREAKING CHANGE: Remove `nxos_ipv4_access_list_policy_ingress_interface_instance` resource and data source
- Add `access_list_name` attribute to `nxos_ipv4_access_list_policy_egress_interface` resource and data source
- BREAKING CHANGE: Remove  `nxos_ipv4_access_list_policy_egress_interface_instance` resource and data source
- Add `nxos_vpc_instance` resource and data source
- Add `nxos_vpc_domain` resource and data source
- Add `nxos_vpc_interface` resource and data source
- Add `nxos_port_channel_interface` resource and data source
- Add `nxos_port_channel_interface_vrf` resource and data source
- BREAKING CHANGE: Remove `ctrl` attribute of `nxos_spanning_tree_interface` resource and data source
- Add `nxos_port_channel_interface_member` resource and data source
- Add `nxos_ipv4_static_route` resource and data source

## 0.3.23

- Add `nxos_ntp_server` resource and data source
- Add `nxos_route_map_rule` resource and data source
- Add `nxos_route_map_rule_entry` resource and data source
- Add `nxos_ipv4_prefix_list_rule` resource and data source
- Add `nxos_ipv4_prefix_list_rule_entry` resource and data source
- Add `hold_time` and `keepalive` attributes to `nxos_bgp_peer` resource
- Add `nxos_bgp_peer_address_family_prefix_list_control` resource and data source
- Add `nxos_bgp_peer_address_family_route_control` resource and data source
- Add `nxos_route_map_rule_entry_set_regular_community` resource and data source
- Add `nxos_route_map_rule_entry_set_regular_community_item` resource and data source
- Add `nxos_route_map_rule_entry_match_route` resource and data source
- Add `nxos_route_map_rule_entry_match_route_prefix_list` resource and data source
- BREAKING CHANGE: Rename `nxos_ipv4_access_list_policy_egress_interface_instace` resource and data source to `nxos_ipv4_access_list_policy_egress_interface_instance`
- BREAKING CHANGE: Rename `nxos_ipv4_access_list_policy_ingress_interface_instace` resource and data source to `nxos_ipv4_access_list_policy_ingress_interface_instance`

## 0.3.22

- Add `nxos_spanning_tree` resource and data source

## 0.3.21

- Fix handling named ports in `nxos_ipv4_access_list_entry` resource

## 0.3.20

- Add `nxos_ipv4_access_list` resource and data source
- Add `nxos_ipv4_access_list_entry` resource and data source
- Add `nxos_ipv4_access_list_policy_ingress_interface` resource and data source
- Add `nxos_ipv4_access_list_policy_ingress_interface_instance` resource and data source
- Add `nxos_ipv4_access_list_policy_egress_interface` resource and data source
- Add `nxos_ipv4_access_list_policy_egress_interface_instance` resource and data source

## 0.3.19

- Change option (`unspecified` -> `none`) of `bfd` attribute of `pim_interface` resource

## 0.3.18

- Add `unspecified` option to `bfd` attribute of `pim_interface` resource

## 0.3.17

- Fix provider config url attribute

## 0.3.16

- Allow provider config without url attribute in case devices attribute is being used

## 0.3.15

- Fix `asn` attribute of BGP data sources

## 0.3.14

- Allow multiple values for `control` attribute of `nxos_bgp_peer_address_family` and `nxos_bgp_peer_template_address_family` resources
- Handle deletion of non-existent objects gracefully
- BREAKING CHANGE: Add `asn` attribute to `nxos_bgp_address_family`, `nxos_bgp_graceful_restart`, `nxos_bgp_peer_address_family`, `nxos_bgp_peer_template_address_family`, `nxos_bgp_peer_template_max_prefix`, `nxos_bgp_peer_template`, `nxos_bgp_peer`, `nxos_route_control` and `nxos_bgp_vrf` resources.
- BREAKING CHANGE: Rename `asn` attribute of `nxos_bgp_peer` resource to `remote_asn`
- BREAKING CHANGE: Rename `asn` attribute of `nxos_bgp_peer_template` resource to `remote_asn`

## 0.3.13

- Add `nxos_pim_anycast_rp` resource and data source
- Add `nxos_pim_anycast_rp_peer` resource and data source

## 0.3.12

- Fix `type` attribute default value of `nxos_ipv4_interface_address` resource

## 0.3.11

- Add `type` attribute to `nxos_ipv4_interface_address` resource
- Add `nxos_evpn` resource and data source
- Add `nxos_evpn_vni` resource and data source
- Add `nxos_evpn_vni_route_target` resource and data source
- Add `nxos_evpn_vni_route_target_direction` resource and data source

## 0.3.10

- Ignore error codes 1 and 107 ("Cannot delete object...") when deleting objects

## 0.3.9

- Update documentation

## 0.3.8

- BREAKING CHANGE: Fix typo in `nxos_nve_vni` resource (`multisite_ingrress_replication` -> `multisite_ingress_replication`)

## 0.3.7

- Fix `nxos_rest` resource object deletion

## 0.3.6

- Add delete attribute to `nxos_rest` resource
- Add attributes to `nxos_physical_interface` resource
- Add attributes to `nxos_ipv4_interface` resource
- Add `nxos_hmm_interface` resource and data source

## 0.3.5

- URL provider configuration attribute is now optional
- Add a list of supported DME objects and its corresponding resources and data sources to documentation

## 0.3.4

- Add `nxos_bgp_peer_address_family` resource and data source
- Add `nxos_bgp_route_control` resource and data source
- Add `nxos_nve_interface` resource and data source
- Add `nxos_nve_vni_container` resource and data source
- Add `nxos_nve_vni resource` and data source
- Add `nxos_nve_vni_ingress_replication` resource and data source
- Add `device` attribute to `nxos_rest` resource and data source

## 0.3.3

- Add `nxos_feature_hmm` resource and data source
- Add `nxos_bgp` resource and data source
- Add `nxos_bgp_instance` resource and data source
- Add `nxos_hmm` resource and data source
- Add `nxos_hmm_instance` resource and data source
- Rename `nxos_vrf_container` resource and data source to `nxos_ipv4_vrf`
- Add `nxos_vrf_routing` resource and data source
- Add `nxos_vrf_address_family` resource and data source
- Add `nxos_vrf_route_target_address_family` resource and data source
- Add `nxos_vrf_route_target_direction` resource and data source
- Add `nxos_vrf_route_target` resource and data source
- Add `nxos_bgp_vrf` resource and data source
- Add `nxos_bgp_address_family` resource and data source
- Add `nxos_bgp_graceful_restart` resource and data source
- Add `nxos_bgp_peer` resource and data source
- Add `nxos_bgp_peer_template` resource and data source
- Add `nxos_bgp_peer_template_address_family` resource and data source
- Add `nxos_bgp_peer_template_max_prefix` resource and data source

## 0.3.2

- Fix documentation

## 0.3.1

- Add `nxos_feature_bgp` resource and data source
- Add `nxos_feature_evpn` resource and data source
- Add `nxos_feature_hsrp` resource and data source
- Add `nxos_feature_isis` resource and data source
- Add `nxos_feature_lacp` resource and data source
- Add `nxos_feature_macsec` resource and data source
- Add `nxos_feature_netflow` resource and data source
- Add `nxos_feature_ospfv3` resource and data source
- Add `nxos_feature_ptp` resource and data source
- Add `nxos_feature_pvlan` resource and data source
- Add `nxos_feature_ssh` resource and data source
- Add `nxos_feature_tacacs` resource and data source
- Add `nxos_feature_telnet` resource and data source
- Add `nxos_feature_udld` resource and data source
- Add `nxos_feature_vpc` resource and data source
- Add `nxos_feature_vn_segment` resource and data source
- Add `nxos_feature_nv_overlay` resource and data source
- Add `nxos_system` resource and data source
- Add `nxos_isis` resource and data source
- Add `nxos_isis_instance` resource and data source
- Add `nxos_isis_vrf` resource and data source
- Add `nxos_isis_interface` resource and data source

## 0.3.0

- Add option to manage multiple devices
- Add `nxos_feature_bfd` resource and data source

## 0.2.5

- Fix redundant authentications
- Fix default retries value (3)
- Make `admin_state` attribute of `nxos_feature_dhcp` resource mandatory
- Make `admin_state` attribute of `nxos_feature_interface_vlan` resource mandatory
- Make `admin_state` attribute of `nxos_feature_lldp` resource mandatory
- Make `admin_state` attribute of `nxos_feature_ospf` resource mandatory
- Make `admin_state` attribute of `nxos_feature_pim` resource mandatory

## 0.2.4

- Add `nxos_svi_interface` resource and data source
- Add `nxos_svi_interface_vrf` resource and data source
- Add `nxos_feature_ospf` resource and data source
- Add `nxos_feature_interface_vlan` resource and data source
- Add `nxos_feature_pim` resource and data source
- Add `nxos_feature_dhcp` resource and data source
- Add `nxos_feature_lldp` resource and data source
- Add `nxos_loopback_interface` resource and data source
- Add `nxos_loopback_interface_vrf` resource and data source
- Add `nxos_physical_interface_vrf` resource and data source
- Add `layer` attribute to `nxos_physical_interface` resource and data source
- Add `nxos_subinterface` resource and data source
- Add `nxos_subinterface_vrf` resource and data source
- Add `nxos_dhcp_relay_interface` resource and data source
- Add `nxos_dhcp_relay_address` resource and data source
- Add `nxos_ethernet` resource and data source
- Add `nxos_ospf` resource and data source
- Add `nxos_ospf_instance` resource and data source
- Add `nxos_ospf_vrf` resource and data source
- Add `nxos_ospf_area` resource and data source
- Add `nxos_ospf_interface` resource and data source
- Add `nxos_ospf_authentication` resource and data source
- Add `nxos_pim resource` and data source
- Add `nxos_pim_instance` resource and data source
- Add `nxos_pim_vrf` resource and data source
- Add `nxos_pim_interface` resource and data source
- Add `nxos_pim_static_rp_policy` resource and data source
- Add `nxos_pim_static_rp` resource and data source
- Add `nxos_pim_static_rp_group_list` resource and data source
- Add `nxos_pim_ssm_policy` resource and data source
- Add `nxos_pim_ssm_range` resource and data source
- Add `nxos_default_qos_class_map` resource and data source
- Add `nxos_default_qos_class_map_dscp` resource and data source
- Add `nxos_default_qos_policy_map` resource and data source
- Add `nxos_default_qos_policy_map_match_class_map` resource and data source
- Add `nxos_default_qos_policy_map_match_class_map_set_qos_group` resource and data source
- Add `nxos_default_qos_policy_map_match_class_map_police` resource and data source
- Add `nxos_default_qos_policy_interface_in` resource and data source
- Add `nxos_default_qos_policy_interface_in_policy_map` resource and data source
- Add `nxos_queuing_qos_policy_map` resource and data source
- Add `nxos_queuing_qos_policy_map_match_class_map` resource and data source
- Add `nxos_queuing_qos_policy_map_match_class_map_priority` resource and data source
- Add `nxos_queuing_qos_policy_map_match_class_map_remaining_bandwidth` resource and data source
- Add `nxos_queuing_qos_policy_system_out` resource and data source
- Add `nxos_queuing_qos_policy_system_out_policy_map` resource and data source

## 0.2.3

- Fix links in resource documentation

## 0.2.2

- Add `nxos_ipv4_interface` resource and data source
- Add `nxos_ipv4_interface_address` resource and data source
- Add `nxos_vrf` resource and data source
- Add `nxos_vrf_container` resource and data source
- Add `nxos_bridge_domain` resource and data source

## 0.2.1

- Add `nxos_pyhsical_interface` resource and data source

## 0.2.0

- Transition to Terraform Plugin Framework

## 0.1.1

- FIX: Render valid request body for resource without content

## 0.1.0

- Initial release
