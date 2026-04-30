resource "nxos_telemetry" "example" {
  admin_state         = "enabled"
  batch_dme_events    = false
  merge_subscriptions = true
  destination_groups = {
    "1" = {
      destinations = {
        "192.168.1.1;50001" = {
          encoding = "GPB"
          node_id  = "host1"
          protocol = "gRPC"
        }
      }
    }
  }
  sensor_groups = {
    "1" = {
      data_source = "DME"
      sensor_paths = {
        "sys/intf" = {
          alias            = "mypath"
          depth            = 1
          filter_condition = "gt(operBitrate,0)"
          query_condition  = "query-condition-example"
        }
      }
    }
  }
  subscriptions = {
    "1" = {
      sensor_group_relationships = {
        "sys/tm/sensor-1" = {
          sample_interval = 5000
        }
      }
      destination_group_relationships = {
        "sys/tm/dest-1" = {}
      }
    }
  }
}
