---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_smart_license Resource - terraform-provider-fmc"
subcategory: "License"
description: |-
  This resource can manage a Smart License.
---

# fmc_smart_license (Resource)

This resource can manage a Smart License.

## Example Usage

```terraform
resource "fmc_smart_license" "example" {
  registration_type = "REGISTER"
  token             = "X2M3YmJlY..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `registration_type` (String) Action to be executed on the smart license.
  - Choices: `REGISTER`, `DEREGISTER`, `EVALUATION`

### Optional

- `domain` (String) The name of the FMC domain
- `token` (String) Registration token.

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import fmc_smart_license.example "<id>"
```