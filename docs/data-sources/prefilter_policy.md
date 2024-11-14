---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_prefilter_policy Data Source - terraform-provider-fmc"
subcategory: "Policy"
description: |-
  This data source can read the Prefilter Policy.
---

# fmc_prefilter_policy (Data Source)

This data source can read the Prefilter Policy.

## Example Usage

```terraform
data "fmc_prefilter_policy" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) The name of the FMC domain
- `id` (String) The id of the object
- `name` (String) The name of the prefilter policy.

### Read-Only

- `default_action` (String) Specifies the default action to take when none of the rules meet the conditions.
- `default_action_log_begin` (Boolean) Indicating whether the device will log events at the beginning of the connection.
- `default_action_log_end` (Boolean) Indicating whether the device will log events at the end of the connection.
- `default_action_send_events_to_fmc` (Boolean) Indicating whether the device will send events to the Firepower Management Center event viewer.
- `default_action_snmp_config_id` (String) UUID of the SNMP alert. Can be set only when either default_action_log_begin or default_action_log_end is true.
- `default_action_syslog_config_id` (String) UUID of the syslog config. Can be set only when either default_action_log_begin or default_action_log_end is true.
- `description` (String) Description