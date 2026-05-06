resource "nxos_netflow" "example" {
  exporters = {
    "EXPORTER1" = {
      description      = "Netflow exporter"
      destination_ip   = "10.1.1.1"
      destination_port = 9995
      dscp             = 46
      source_interface = "lo0"
      version          = "v9"
      vrf_name         = "default"
    }
  }
  records = {
    "RECORD1" = {
      description        = "Netflow record"
      collect_parameters = "count-bytes"
      match_parameters   = "src-ipv4"
    }
  }
  monitors = {
    "MONITOR1" = {
      description      = "Netflow monitor"
      record_target_dn = "sys/flow/fr-[RECORD1]"
      exporter_buckets = {
        "1" = {
          description         = "Bucket 1"
          hash_high           = 4294967295
          hash_low            = 0
          exporter1_target_dn = "sys/flow/fe-[EXPORTER1]"
          exporter2_target_dn = "sys/flow/fe-[EXPORTER2]"
        }
      }
    }
  }
  hardware_profiles = {
    "HWPROFILE1" = {
      description          = "Hardware profile"
      burst_interval_shift = 10
      export_interval      = 200
      ip_packet_id_shift   = 10
      mtu                  = 1500
      source_port          = 10000
    }
  }
  class_maps = {
    "CMAP1" = {
      match_acls = {
        "ACL1" = {}
      }
    }
  }
}
