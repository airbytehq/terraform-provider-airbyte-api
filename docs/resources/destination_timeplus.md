---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_timeplus Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationTimeplus Resource
---

# airbyte_destination_timeplus (Resource)

DestinationTimeplus Resource

## Example Usage

```terraform
resource "airbyte_destination_timeplus" "my_destination_timeplus" {
  configuration = {
    apikey   = "...my_apikey..."
    endpoint = "https://us-west-2.timeplus.cloud/workspace_id"
  }
  definition_id = "d7fd7136-64c8-4ab0-88c2-48e91396f340"
  name          = "Samantha Gleason"
  workspace_id  = "500686d0-4e60-4803-9bc7-eb0732a47524"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the destination e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.

### Read-Only

- `destination_id` (String)
- `destination_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `apikey` (String, Sensitive) Personal API key

Optional:

- `endpoint` (String) Timeplus workspace endpoint. Default: "https://us-west-2.timeplus.cloud/<workspace_id>"

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_destination_timeplus.my_airbyte_destination_timeplus ""
```