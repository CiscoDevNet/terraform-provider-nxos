## 0.19.0

- Add CoS, precedence, QoS group, ACL and outer (tunnel) DSCP match criteria, as well as negated DSCP, CoS and precedence match criteria, to `class_maps` of `nxos_default_qos` resource and data source
- Add DSCP, outer (tunnel) DSCP, CoS and precedence marking (`set_dscp_*`, `set_cos_value`, `set_precedence_*` attributes) to `match_class_maps` of `nxos_default_qos` resource and data source
- Add dynamic load balancing (DLB) configuration (`dlb`), including static pinning (`dlb_static_pinnings`), to `nxos_platform` resource and data source
- Add `set_dlb_mode` attribute to `match_class_maps` of `nxos_default_qos` resource and data source
- Add global hardware QoS configuration (`qos_*` attributes), including lossless buffer tuning (PFC static, no-drop PG thresholds, pool sizes), AFD profile and ETRAP settings, to `nxos_platform` resource and data source
- Add global priority flow control (PFC) watchdog configuration (`pfc_watchdog_*` attributes) to `nxos_platform` resource and data source
- Add per-interface egress queuing policy map assignments (`policy_interface_out`) to `nxos_queuing_qos` resource and data source
- Add NVE interface ingress QoS policy map assignments (`policy_nve_interface_in`) to `nxos_default_qos` resource and data source
- Fix destroy of `nxos_physical_interface`, `nxos_platform` and `nxos_system` (and removal of list items) not removing nested configuration of non-deletable objects, e.g. storm control, PFC and load interval settings of physical interfaces or NVE infra VLANs are now removed or reset to their defaults
- Add bandwidth, shaping, queue limit and micro burst detection configuration to `match_class_maps` of `nxos_queuing_qos` resource and data source
- Add WRED/ECN random detect and Approximate Fair Discard (AFD) configuration to `match_class_maps` of `nxos_queuing_qos` resource and data source
- Add global QoS configuration (`qos_dynamic_buffer_sharing`, `qos_statistics`, `qos_policies_statistics`, `qos_reset_dei` and `qos_mode` attributes) to `nxos_system` resource and data source
- Fix destroy of `nxos_system` hanging the NX-OS private VLAN process when `pvlans`, `pvlan_interfaces` or `pvlan_svis` are configured: they are no longer deleted on destroy (disabling the `private-vlan` feature removes them), as NX-OS processes their removal asynchronously and the following deletion of the referenced interfaces and VLANs could block it

## 0.18.4

- Retry requests rejected by NX-OS due to frequent invocation of a command (e.g. changing the `shutdown` attribute of `nxos_macsec` or destroying it within 30 seconds of a previous MACsec status change)

## 0.18.3

- Add IPv6 DHCP relay addresses (`ipv6_addresses`) to `relay_interfaces` of `nxos_dhcp` resource and data source

## 0.18.2

- Fix creating and deleting UDFs with TCAM region qualifiers (`qualify_*` attributes) in `nxos_platform` resource

## 0.18.1

- Fix destroy of `nxos_feature`, `nxos_platform`, `nxos_snmp`, `nxos_system` and `nxos_user_management` not resetting attributes of nested non-deletable objects to their defaults, e.g. features enabled by `nxos_feature` (including `vtp`) are now disabled on destroy
- Keep NX-API enabled when the `nxapi` attribute of `nxos_feature` is removed or the resource is destroyed, to avoid losing access to the device

## 0.18.0

- Add `nxos_epbr` resource and data source
- Add ePBR feature configuration to `nxos_feature` resource and data source
- Add policy-based routing (PBR) configuration to `nxos_route_policy` resource and data source, including set next-hop addresses (with track and VRF), set interfaces, IPv6 access list matches, and IPv4/IPv6 policy route-map interface assignments
- Add PBR feature configuration to `nxos_feature` resource and data source
- Add TCAM templates (including UDF qualifiers) and user-defined fields (UDFs) configuration to `nxos_platform` resource and data source
- Add `dot1q_tunnel_transit_range` attribute to `nxos_platform` resource and data source
- Add `nxos_poe` resource and data source
- Add PoE feature configuration to `nxos_feature` resource and data source
- Add `nxos_sla` resource and data source
- Add SLA sender, SLA responder and SLA TWAMP server feature configuration to `nxos_feature` resource and data source
- Add `nxos_vtp` resource and data source
- Add VTP feature configuration to `nxos_feature` resource and data source

## 0.17.2

- Fix `nxos_save_config` action panic (nil pointer dereference) by setting `ActionData` in the provider `Configure` method
- Fix VLAN mapping configuration of `nxos_physical_interface` and `nxos_port_channel_interface` not being applied by the device when configured together with the switchport mode, VLAN mapping objects are now sent in a separate request after the rest of the interface configuration
- BREAKING CHANGE: Replace `selective_qinq_vlans` attribute of `nxos_physical_interface` and `nxos_port_channel_interface` with `selective_qinq_mappings` (keyed by outer VLAN ID, with `vlans`), as the previous attribute was accepted but not applied by the device

## 0.17.1

- Fix permanent diff and update error of `nxos_ntp` authentication keys when the device encrypts a key, `encryption_type` and `encrypted` are now write-only attributes

## 0.17.0

- Add BGP bestpath options (AS-path multipath relax, always compare MED, compare router ID, compare neighbor ID, MED missing as worst, MED confederation, MED non-deterministic, ignore AS-path, ignore cost community, ignore IGP metric, ignore AIGP) configuration to `nxos_bgp` VRFs
- Add BGP peer-session templates (including local AS) and peer-policy templates (including route-map, prefix-list, filter-list, maximum-prefix and peer-policy inheritance) configuration to `nxos_bgp` VRFs
- Add peer-policy inheritance configuration to `nxos_bgp` peer address families and peer template address families
- Add route-map and prefix-list in/out configuration to `nxos_bgp` peer template address families
- Add DHCP and DHCPv6 relay source interface configuration, both global and per interface, to `nxos_dhcp` resource and data source
- Add DNS name servers (including source interface), domain lists, static IPv4 and IPv6 hosts, and per-VRF DNS (domain name, domain list and name servers) configuration to `nxos_system` resource and data source
- Add errdisable recovery interval and per-cause errdisable detect and recovery configuration to `nxos_system` resource and data source
- Add security contracts (source and destination security group with policy map) configuration to `nxos_esg` domains
- Add protocol, source/destination port range, and TCP session rules configuration to `nxos_esg` class map filter entries
- BREAKING CHANGE: Change `match_dscp` attribute of `nxos_esg` class map filter entries from a number to a string, using the DSCP names returned by the device (e.g., `AF11`, `EF`, `CS1`) and plain numbers for values without a name (e.g., `'0'`, `'56'`), to fix permanent diffs when a named DSCP value is configured
- Add intra-group isolation configuration to `nxos_esg` security groups (requires NX-OS 10.6(3) or later)
- Add layer 2 isolation configuration to `nxos_bridge_domain` (requires NX-OS 10.6(3) or later)
- Add anycast gateway proxy forwarding mode configuration to `nxos_hmm` interfaces (requires NX-OS 10.6(3) or later)
- Add VLAN mapping (enable, and mappings with optional inner VLAN and receive-only mode) configuration to `nxos_physical_interface` and `nxos_port_channel_interface`, receive-only mapping requires NX-OS 10.6(3) or later
- Add selective QinQ (VLAN range or all VLANs to dot1q tunnel outer VLAN) configuration to `nxos_physical_interface` and `nxos_port_channel_interface`
- Add external subnet selectors (VRF, IPv4/IPv6 prefix and route inject) configuration to `nxos_esg` security groups
- Add external subnet with next-hop selectors (VRF, IPv4/IPv6 prefix, next-hop address and MPLS VPN encapsulation type) configuration to `nxos_esg` security groups
- Add interface selectors configuration to `nxos_esg` security groups
- Add VLAN and MAC selectors configuration to `nxos_esg` security groups
- Add VLAN and interface selectors configuration to `nxos_esg` security groups (requires NX-OS 10.6(1) or later)
- Add layer 4-7 type configuration to `nxos_esg` security groups
- Add default-information originate (including always and route-map), SPF throttle timers, LSA timers (throttle, arrival and group pacing), graceful restart (including grace period, helper disable and planned only) and summary address configuration to `nxos_ospf` VRFs
- Add area range (including not-advertise and cost) configuration to `nxos_ospf` areas
- Add LSA timers (throttle, arrival and group pacing) and graceful restart (including grace period, helper disable and planned only) configuration to `nxos_ospfv3` VRFs
- Add redistribution (protocol with route-map), default-information originate (including always and route-map), SPF throttle timers and summary prefix configuration to `nxos_ospfv3` address families
- Add area address family default cost and area range (including not-advertise and cost) configuration to `nxos_ospfv3` areas
- Add redistribution, per-level summary address and inter-level route distribution (route-map or all) configuration to `nxos_isis` address families
- Add per-level SPF interval, LSP generation interval and level controls, and graceful restart (including T3 manual) configuration to `nxos_isis` VRFs
- Add ECMP load-sharing (source/destination address and port hashing, concatenation, rotate, universal ID, GRE outer hash, UDF offset and length, and hardware load-balancing key shift) configuration to `nxos_ipv4` resource and data source
- Add LACP configuration (system priority, system MAC and role, front-panel port numbering, rapid transition, and per-interface port priority and rate) to `nxos_system` resource and data source
- Add load interval counters configuration to `nxos_physical_interface` and `nxos_port_channel_interface` resources and data sources
- Add sFlow data source interfaces configuration to `nxos_hardware_telemetry` resource and data source
- Add port-channel load-balance hash fields, rotate, concatenation, symmetric, encapsulation, FEX and MPLS label options configuration to `nxos_platform` resource and data source
- Add MAC address table configuration (aging time, MAC move notification, MAC limits and thresholds, static MAC clean-up, and static entries with VLAN and interface or drop) to `nxos_system` resource and data source
- Add `nxos_macsec` resource and data source (MACsec instance shutdown, MACsec policies, and per-interface MACsec keychain, fallback keychain, and policy settings)
- Add MACsec keychains, including keys, key octet strings, cryptographic algorithms, and send lifetimes, and MACsec PSK no-show configuration to `nxos_keychain` resource and data source
- Add NTP authentication keys and trusted keys configuration to `nxos_ntp` resource and data source
- Add RADIUS global settings (key, timeout, retransmit, deadtime and source interface), server monitoring (test username, password and idle time), servers (key, authentication and accounting ports, and authentication/accounting usage) and server groups (including servers, VRF, source interface and deadtime) configuration to `nxos_user_management` resource and data source
- Add set AS-path prepend (AS numbers and last AS), set tag, set weight, set origin, set extended community route target, and set community list delete configuration to `nxos_route_policy` route map entries
- Add match interface and match route type (including EVPN route types) configuration to `nxos_route_policy` route map entries
- Add SNMP community (group or access level, and IPv4/IPv6 ACLs) configuration to `nxos_snmp` resource and data source
- Add SNMP host source interface and filter VRF configuration to `nxos_snmp` resource and data source
- Add SNMP informs source interface configuration to `nxos_snmp` resource and data source
- Add SNMP context (instance, VRF and topology) configuration to `nxos_snmp` resource and data source
- Add vPC orphan port suspend configuration to `nxos_vpc` resource and data source
- Fix resources and data sources failing to read with `Unknown class` on NX-OS releases that do not support some of their objects (e.g. releases older than the tested ones); objects not supported by the device are now ignored when reading, configuring them still fails on apply

## 0.16.0

- Add `nxos_eigrp` resource and data source (EIGRP instances, VRFs, address families, authentication, default metric, graceful restart, metric weights, default-information originate, redistribution, redistribute maximum-prefix, table-map, stub, and per-interface address family settings including timers, metrics, passive, split-horizon, next-hop-self, authentication, distribute-lists, offset-lists, and summary addresses)
- Add EIGRP feature (`feature eigrp`) configuration to `nxos_feature` resource and data source
- Add `nxos_track` resource and data source
- Add tracked objects configuration to `nxos_hsrp` groups

## 0.15.0

- Add global BFD timers (keepalive policy: detect multiplier, echo receive interval, minimum receive interval, minimum transmit interval) configuration to `nxos_bfd` resource and data source
- Add AS-Path access list (and their entries) configuration to `nxos_route_policy` resource and data source
- Add match AS-Path access list and match AS number AS-Path access list configuration to `nxos_route_policy` route map entries
- Fix `terraform destroy` failing for objects belonging to a feature that is not enabled on the device, even though they were never configured
- Fix `nxos_access_list` update failing with HTTP 400 or `child (Rn) of class aclPolicy is already attached` when removing interface bindings (`ingress_interfaces`, `egress_interfaces`, and IPv6 variants)
- Fix removed attributes of nested child objects (for example `ingress_vty_access_list_name` in `nxos_access_list`) not being unset on the device
- Fix `nxos_pim` update failing with `child (Rn) cannot be added to deleted object` when all group lists of a static RP are replaced at once

## 0.14.0

- Add IPv4 address, IPv6 address, and port object group (and their members) configuration to `nxos_access_list` resource and data source
- Add `nxos_igmp_snooping` resource and data source (global IGMP snooping configuration including VXLAN snooping, NVE static router port, querier, opt-flood, routing, and stateful-HA controls)
- Add SSH source interface configuration (per VRF) to `nxos_system` resource and data source
- Add FTP source interface configuration (per VRF) to `nxos_system` resource and data source
- Add AES password encryption (`feature password encryption aes`) configuration to `nxos_system` resource and data source
- Add ACL logging configuration (match log level, cache entries, interval, threshold, detailed, include-mac, include-sgt) to `nxos_system` resource and data source
- Add Private VLAN (PVLAN) configuration (VLAN type/association, per-interface promiscuous/host/trunk mappings, and SVI secondary VLAN propagation) to `nxos_system` resource and data source
- Add `vlan configuration` (per-VLAN MAC learning) support to `nxos_bridge_domain` resource and data source
- Fix issue where importing a resource that manages multiple objects (e.g. `nxos_physical_interface`, `nxos_port_channel_interface`, `nxos_system`) could cause the next apply to reset objects, attributes, or entries that were never declared in configuration, which could fail outright for objects NX-OS doesn't allow configuring directly, such as physical interfaces that are port-channel members
- Add a warning to the first plan/apply after importing a resource, noting that entries or attributes shown as removed are expected and will not change anything on the device
- Fix the "Importing Resources" guide and every affected resource's generated string-based import example/docs, which showed an empty `id` value that Terraform's `import` block does not accept
- Mark `password_type` attribute on `nxos_bgp` peers and interface_peers as write-only to prevent perpetual drift when the device round-trips the value differently than supplied
- Mark `key_encryption` and `monitoring_password_type` attributes on `nxos_user_management` tacacs_providers as write-only for the same reason

## 0.13.1

- Fix issue where removing an optional attribute from the configuration would not reset the value on the device during updates

## 0.13.0

- Fix nested child list entries not being deleted from the device when a parent list item key changes during updates
- BREAKING CHANGE: Extract platform configuration from `nxos_system` into a new `nxos_platform` resource and data source. All `platform_*`, `platform_extended_*`, `platform_nve_*`, and `tcam_region_*` attributes have been removed from `nxos_system`. Users must create a new `nxos_platform` resource and import existing state with `terraform import nxos_platform.example ""`. Attribute names on the new resource drop the `platform_` prefix (e.g., `platform_access_list_match_inner_header` becomes `access_list_match_inner_header`, `platform_extended_dme_load_interval` becomes `extended_dme_load_interval`). The `platform_nve_interfaces` attribute is renamed to `nve_interfaces`.

## 0.12.1

- Fix `nxos_hardware_telemetry`: destroying sFlow receiver configuration could fail due to malformed DN

## 0.12.0

- BREAKING CHANGE: Replace `sflow_receiver_address`, `sflow_receiver_source_address`, and `sflow_receiver_vrf_name` attributes on `nxos_hardware_telemetry` resource and data source with a `receivers` list supporting multiple sFlow collectors and full read support
- Remove eager device login during provider configuration so that `terraform plan` no longer fails when a device is unreachable

## 0.11.2

- Add flow monitor, record, collector, and collector bucket configuration to `nxos_analytics` resource and data source
- Add forward instance target configuration to `nxos_analytics` resource and data source
- Add support for analytics match ACL configuration to `nxos_analytics` resource and data source
- Add monitor record reference, exporter bucket exporter references to `nxos_netflow` resource and data source
- Add extended platform entity configuration to `nxos_system` resource and data source
- Add set policy tag configuration to `nxos_route_policy` resource and data source
- Add destination profile and use-vrf configuration to `nxos_telemetry` resource and data source
- Add route map policy configuration to `nxos_vrf` resource and data source

## 0.11.1

- Add Priority Flow Control (PFC) configuration to `nxos_physical_interface` resource and data source
- Add Priority Flow Control Watchdog configuration to `nxos_physical_interface` resource and data source
- Add Priority Flow Control (PFC) configuration to `nxos_port_channel_interface` resource and data source
- Add Priority Flow Control Watchdog configuration to `nxos_port_channel_interface` resource and data source
- Add new `nxos_network_qos` resource and data source

## 0.11.0

- Add `nxos_analytics` resource and data source
- Add additional paths configuration to `nxos_bgp` resource and data source
- Add aggregate address configuration to `nxos_bgp` resource and data source
- Add interface peer configuration to `nxos_bgp` resource and data source
- Add `default_information_originate_route_distinguisher`, `default_information_originate_route_target`, `load_balance_egress_filter_policy_route_map`, and `load_balance_egress_multipath_auto_policy_route_map` attributes to `nxos_bgp` resource and data source
- Add `nxos_esg` resource and data source
- Add SCP Server, SFTP Server, gRPC, Analytics, Security Group, and Telemetry feature configuration to `nxos_feature` resource and data source
- Add IPv6 prefix list configuration to `nxos_route_policy` resource and data source
- Add `nxos_netflow` resource and data source
- Add NVE platform and InfraVlan configuration to `nxos_system` resource and data source
- Add EVPN Multisite Border Gateway configuration to `nxos_nvo` resource and data source
- Add TCAM region and extended TCAM region configuration to `nxos_system` resource and data source
- Add `nxos_ptp` resource and data source
- Add set local preference, set path selection, set evpn, match next hop, match regular community, and match route access list configuration to `nxos_route_policy` resource and data source
- Add `nxos_span` resource and data source
- Add `erspan_origin_ip` attributes to `nxos_system` resource and data source
- Add SSH server configuration to `nxos_system` resource and data source
- Add SSH server key configuration to `nxos_system` resource and data source
- Add `nxos_telemetry` resource and data source
- Add PTP ttag configuration to `nxos_system` resource and data source
- Add multisite interface tracking configuration to `nxos_physical_interface`, `nxos_port_channel_interface`, and `nxos_svi_interface` resources and data sources

## 0.10.2

- Fix `nxos_port_channel_interface`: extended attributes (storm control, switchport block, MAC learn, etc.) were not read back from the device due to incorrect child class RN (`physExtd` → `aggrExtd`)

## 0.10.1

- Add community list configuration to `nxos_route_policy` resource and data source
- Add `nxos_hardware_telemetry` resource and data source

## 0.10.0

- Add feature-sets (`fex`, `mpls`, `virtualization`) configuration to `nxos_feature` resource and data source
- BREAKING CHANGE: The `nxos_feature` resource is now rooted at `sys` instead of `sys/fm`. Existing state must be removed and re-imported (`terraform state rm nxos_feature.<name>` followed by `terraform import nxos_feature.<name> sys`). The Terraform attribute surface (e.g. `bgp`, `ospf`, `vpc`) is unchanged.
- Add extended configuration to `nxos_port_channel_interface` resource and data source
- Add set next hop peer address and unchanged configuration to `nxos_route_policy` resource and data source
- Fix resource updates sending attributes not present in the Terraform configuration, causing failures on certain NX-OS versions (#460)

## 0.9.6

- Fix `nxos_access_list`: updating an existing access list entry no longer fails with `Duplicate sequence number`/duplicate object errors (regression of #204, re-reported in #453)

## 0.9.5

- Add interface breakout configuration to `nxos_system` resource and data source
- Add Hypershield (Service Acceleration) support to `nxos_system` resource and data source
- Add Service Acceleration feature to `nxos_feature` resource and data source

## 0.9.4

- Add NXAPI feature configuration to `nxos_feature` resource and data source

## 0.9.3

- Add `nxos_management_interface` resource and data source
- Add NX-API configuration to `nxos_system` resource and data source
- Add set metric type configuration to `nxos_route_policy` resource and data source
- Add SNMP host VRF configuration to `nxos_snmp` resource and data source
- Add SNMP AAA sync disable and global enforce privacy attributes to `nxos_snmp` resource and data source
- Add SNMP trap enable configuration to `nxos_snmp` resource and data source

## 0.9.2

- Add `file_severity` attribute to `nxos_logging` resource and data source

## 0.9.1

- Add `l1PhysIfExtended` child class to `nxos_physical_interface` resource and data source (storm control, switchport, flow control, chassis, and extended interface attributes)

## 0.9.0

- Add CDP configuration to `nxos_system` resource and data source
- Add `nxos_bfd` resource and data source
- Add CFS distribution configuration to `nxos_system` resource and data source
- Add per-VLAN STP configuration to `nxos_spanning_tree` resource and data source
- Add route redistribution configuration to `nxos_ospf` resource and data source
- Add `nxos_icmpv6` resource and data source
- Add IPv6 ND VPC domain synchronization configuration to `nxos_system` resource and data source
- Add set metric configuration to `nxos_route_policy` resource and data source
- Add ICAM scale monitoring configuration to `nxos_system` resource and data source
- Add terminal console exec timeout configuration to `nxos_system` resource and data source
- Add terminal VTY exec timeout configuration to `nxos_system` resource and data source
- Add terminal VTY session limit configuration to `nxos_system` resource and data source
- Add VTY ingress and egress access list policy to `nxos_access_list` resource and data source
- Add CoPP configuration to `nxos_system` resource and data source
- Add `extchp-edge`, `extchp-bpdu-filter`, and `extchp-bpdu-guard` enum values to spanning tree instance `control` attribute in `nxos_spanning_tree` resource and data source
- Add TACACS+ provider group server members to `nxos_user_management` resource and data source
- Add RMON event configuration to `nxos_snmp` resource and data source
- Add management interface configuration to `nxos_system` resource and data source
- Add LLDP configuration to `nxos_system` resource and data source
- Add platform entity configuration to `nxos_system` resource and data source
- Add `nxos_snmp` resource and data source
- Add UDLD configuration to `nxos_system` resource and data source
- Add NTP source interface configuration to `nxos_ntp` resource and data source
- Add NTP access group configuration to `nxos_ntp` resource and data source
- Add CLI alias configuration to `nxos_system` resource and data source
- Add storm control policy attributes to `nxos_physical_interface` resource and data source
- Add storm control policy attributes to `nxos_port_channel_interface` resource and data source
- Add boot configuration and boot image attributes to `nxos_system` resource and data source
- Add Smart Licensing transport mode and CSLU URL attributes to `nxos_system` resource and data source
- Add AAA authentication realm configuration to `nxos_user_management` resource and data source
- Add default authentication configuration to `nxos_user_management` resource and data source
- Add console authentication configuration to `nxos_user_management` resource and data source
- Add default authorization configuration to `nxos_user_management` resource and data source
- Add default accounting configuration to `nxos_user_management` resource and data source
- Add VDC and resource limit attributes to `nxos_system` resource and data source
- Add DNS configuration support to `nxos_system` resource
- Add optional per-device `username` and `password` attributes to override provider-level credentials
- Add clock timezone and summer time attributes to `nxos_system` resource and data source
- Add `otm` to logging facilities in `nxos_logging` resource and data source
- Add syslog file configuration to `nxos_logging` resource and data source
- Add syslog remote destinations to `nxos_logging` resource and data source
- Add syslog source interface to `nxos_logging` resource and data source
- Add syslog timestamp format to `nxos_logging` resource and data source
- Add syslog terminal monitor to `nxos_logging` resource and data source
- Add syslog console to `nxos_logging` resource and data source
- Add syslog origin ID to `nxos_logging` resource and data source

## 0.8.1

- Fix `nxos_dme` resource not detecting out-of-band deletion of managed objects

## 0.8.0

- BREAKING CHANGE: Complete provider rewrite — all resources and data sources have been redesigned with breaking changes and no backwards compatibility with 0.7.x . Related NX-API objects that were previously managed as separate resources are now consolidated into single, unified resources (e.g., `nxos_bgp` replaces 15 individual BGP resources, `nxos_features` replaces 27 individual feature resources)
- BREAKING CHANGE: Rename `nxos_rest` resource and data source to `nxos_dme`
- BREAKING CHANGE: Empty or null attribute values are now omitted from API payloads by default

## 0.7.0

- BREAKING CHANGE: Resource imports now use comma-separated attribute values instead of full DN strings (e.g., `terraform import nxos_bgp_peer.example "<asn>,<vrf>,<address>"` instead of `terraform import nxos_bgp_peer.example "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]"`)
- BREAKING CHANGE: `nxos_rest` import ID format changed from `<class_name>:<dn>` to `<dn>,<class_name>` (with optional `,<device>` suffix)
- Add identity-based import support for all resources
- Add `nxos_save_config` action
- Add `nxos_cli` resource

## 0.6.0

- BREAKING CHANGE: Remove `nxos_ipv4_access_list_entry` resource and data source, entries are now managed as child classes of `nxos_ipv4_access_list`
- Add `nxos_keychain` resource and data source
- Add `nxos_keychain_key` resource and data source
- Add `nxos_keychain_manager` resource and data source
- Add `nxos_ospf_max_metric` resource and data source
- Add `log_adjacency_changes` attribute to `nxos_ospf_vrf` resource and data source
- Add `nxos_feature_bash_shell` resource and data source
- Add `nxos_feature_sflow` resource and data source
- Add `nxos_logging` resource and data source
- Add `nxos_user` resource and data source
- Fix `nxos_icmpv4_interface` controls
- Fix reading nested child classes during import
- Add `selected_devices` provider configuration attribute
- Add support for device-specific imports using comma separator (e.g., `sys/bgp,LEAF-1`)
- Update `go-nxos` client to improve HTTP connection reuse under Terraform parallelism

## 0.5.10

- Add `nxos_ipv6_interface` resource and data source
- Add `nxos_ipv6_interface_address` resource and data source
- Add `nxos_ipv6_static_route` resource and data source
- Add `nxos_ipv6_vrf` resource and data source
- Add `nxos_ospfv3` resource and data source
- Add `nxos_ospfv3_area` resource and data source
- Add `nxos_ospfv3_instance` resource and data source
- Add `nxos_ospfv3_interface` resource and data source
- Add `nxos_ospfv3_vrf` resource and data source
- Add `nxos_ospfv3_vrf_address_family` resource and data source
- Honor proxy settings (`HTTP_PROXY`, `HTTPS_PROXY`, `NO_PROXY` environment variables)

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
