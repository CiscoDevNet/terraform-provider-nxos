resource "nxos_platform" "example" {
  nve_interfaces = {
    "1" = {
      ipmc_index_size = 4000
      infra_vlans = {
        "100" = {
          force = "Enable"
        }
      }
    }
  }
  extended_dme_load_interval              = 300
  port_channel_load_balance_concatenation = "enable"
  port_channel_load_balance_fields        = "srcDstIp"
  port_channel_load_balance_rotate        = 3
  port_channel_load_balance_symmetry      = "enable"
}
