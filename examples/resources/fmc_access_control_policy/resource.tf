resource "fmc_access_control_policy" "example" {
  name                              = "POLICY1"
  description                       = "My access control policy"
  default_action                    = "BLOCK"
  default_action_log_begin          = true
  default_action_log_end            = false
  default_action_send_events_to_fmc = true
  default_action_send_syslog        = true
  default_action_syslog_config_id   = "35e197ca-33a8-11ef-b2d1-d98ae17766e7"
  default_action_syslog_severity    = "DEBUG"
  categories = [
    {
      name = "cat1"
    }
  ]
  rules = [
    {
      action = "ALLOW"
      name   = "rule1"
      source_network_literals = [
        {
          value = "10.1.1.0/24"
        }
      ]
      destination_network_literals = [
        {
          value = "10.2.2.0/24"
        }
      ]
      source_network_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "Network"
        }
      ]
      destination_network_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "Network"
        }
      ]
      source_dynamic_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      destination_dynamic_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      source_port_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      destination_port_objects = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      source_security_group_tag_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "SecurityGroupTag"
        }
      ]
      destination_security_group_tag_objects = [
        {
          id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
          type = "SecurityGroupTag"
        }
      ]
      source_zones = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      destination_zones = [
        {
          id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
        }
      ]
      log_begin           = true
      log_end             = true
      send_events_to_fmc  = true
      send_syslog         = true
      syslog_config_id    = "35e197ca-33a8-11ef-b2d1-d98ae17766e7"
      syslog_severity     = "DEBUG"
      snmp_config_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      file_policy_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      intrusion_policy_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
}
