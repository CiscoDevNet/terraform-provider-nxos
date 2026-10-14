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
  tcam_templates = {
    "TPL1" = {
      reference_template = "l3"
      udfs = {
        "racl;UDF1" = {}
      }
    }
  }
  udf_description = "My UDFs"
  udfs = {
    "UDF1" = {
      mpls_stack_1    = true
      mpls_stack_2    = true
      netflow_rtp     = true
      offset_base     = "outer-header-l3"
      offset_length   = 2
      offset_position = 10
    }
  }
}
