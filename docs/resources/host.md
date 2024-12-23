---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_host Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource can manage a Host.
---

# fmc_host (Resource)

This resource can manage a Host.

## Example Usage

```terraform
resource "fmc_host" "example" {
  name        = "HOST1"
  description = "My host object"
  ip          = "10.1.1.1"
  overridable = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ip` (String) IP of the host.
- `name` (String) The name of the host object.

### Optional

- `description` (String) Description
- `domain` (String) The name of the FMC domain
- `overridable` (Boolean) Whether the object values can be overridden.

### Read-Only

- `id` (String) The id of the object
- `type` (String) Type of the object; this value is always 'Host'.

## Import

Import is supported using the following syntax:

```shell
terraform import fmc_host.example "<id>"
```
