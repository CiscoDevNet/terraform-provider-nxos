resource "nxos_telemetry" "example" {
  admin_state                       = "enabled"
  batch_dme_events                  = false
  merge_subscriptions               = true
  destination_profile_admin_state   = "enabled"
  destination_profile_vrf           = "management"
  destination_profile_compression   = "gzip"
  destination_profile_grpc_deadline = 60
  destination_profile_node_id       = "leaf1"
  destination_profile_retry_size    = 100
  destination_groups = {
    "1" = {
      chunking_size    = 128
      compression      = "gzip"
      max_message_size = 16384
      source_interface = "lo0"
      vrf              = "default"
      destinations = {
        "192.168.1.1;50001" = {
          encoding = "GPB"
          node_id  = "host1"
          protocol = "gRPC"
        }
      }
      hosts = {
        "192.168.1.2;8088" = {
          dial_in_stream                 = false
          encoding                       = "JSON"
          endpoint                       = "services/collector/raw"
          grpc_async                     = false
          index                          = "network"
          node_id                        = "host1"
          protocol                       = "HEC"
          proxy_password                 = "secret_password"
          proxy_port                     = 3128
          proxy_password_encryption_type = "0"
          proxy_server                   = "proxy.example.com"
          proxy_username                 = "proxyuser"
          token                          = "secret_token"
          token_encryption_type          = "0"
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
