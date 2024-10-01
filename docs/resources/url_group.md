---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_url_group Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource can manage an URL Group.
---

# fmc_url_group (Resource)

This resource can manage an URL Group.

## Example Usage

```terraform
resource "fmc_url_group" "example" {
  name        = "url_group_1"
  description = "My URL group"
  urls = [
    {
      id = "0050568A-FAC7-0ed3-0000-004294987896"
    }
  ]
  literals = [
    {
      url = "https://www.example.com/app"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) User-created name of the object.
- `urls` (Attributes Set) (see [below for nested schema](#nestedatt--urls))

### Optional

- `description` (String) Optional user-created description.
- `domain` (String) The name of the FMC domain
- `literals` (Attributes Set) (see [below for nested schema](#nestedatt--literals))
- `overridable` (Boolean) Indicates whether object values can be overridden.

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--urls"></a>
### Nested Schema for `urls`

Optional:

- `id` (String) UUID of the object (such as fmc_url.example.id, etc.).


<a id="nestedatt--literals"></a>
### Nested Schema for `literals`

Optional:

- `url` (String) URL literal value.

## Import

Import is supported using the following syntax:

```shell
terraform import fmc_url_group.example "<id>"
```