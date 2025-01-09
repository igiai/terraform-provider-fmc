resource "fmc_device_ha_pair" "example" {
  name                             = "FTD_HA"
  primary_device_id                = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  secondary_device_id              = "96d24097-41c4-4332-a4d0-a8c07ac08482"
  ha_link_interface_id             = "96d24097-41c4-4332-a4d0-a8c07ac08482"
  ha_link_interface_name           = "GigabitEthernet0/0"
  ha_link_interface_type           = ""
  ha_link_logical_name             = "LAN-INTERFACE"
  ha_link_use_ipv6                 = false
  ha_link_primary_ip               = "1.1.1.1"
  ha_link_secondary_ip             = "1.1.1.2"
  ha_link_netmask                  = "255.255.255.0"
  state_link_use_same_as_ha        = false
  state_link_interface_id          = "76d24097-hj7r-7786-a4d0-a8c07ac08470"
  state_link_interface_name        = "GigabitEthernet0/0"
  state_link_interface_type        = "PhysicalInterface"
  state_link_logical_name          = "Stateful-INTERFACE"
  state_link_use_ipv6              = false
  state_link_primary_ip            = "10.10.10.1"
  state_link_secondary_ip          = "10.10.10.2"
  state_link_netmask               = "255.255.255.0"
  encryption_enabled               = true
  encryption_key_generation_scheme = "AUTO"
  failed_interfaces_limit          = 1
  peer_poll_time                   = 1
  peer_poll_time_unit              = "SEC"
  peer_hold_time                   = 15
  peer_hold_time_unit              = "SEC"
  interface_poll_time              = 5
  interface_poll_time_unit         = "SEC"
  interface_hold_time              = 25
}
