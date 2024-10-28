---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_ezofficeinventory Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceEzofficeinventory Resource
---

# airbyte_source_ezofficeinventory (Resource)

SourceEzofficeinventory Resource

## Example Usage

```terraform
resource "airbyte_source_ezofficeinventory" "my_source_ezofficeinventory" {
  configuration = {
    api_key    = "...my_api_key..."
    start_date = "2021-07-04T01:47:08.951Z"
    subdomain  = "...my_subdomain..."
  }
  definition_id = "68fdfc06-92b4-4fd6-b3f5-9a8d0acc9948"
  name          = "Thomas Ankunding Jr."
  secret_id     = "...my_secret_id..."
  workspace_id  = "59fac1d6-c9b0-4f0f-b5d9-42704e93ebb3"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the source e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `api_key` (String, Sensitive) Your EZOfficeInventory Access Token. API Access is disabled by default. Enable API Access in Settings > Integrations > API Integration and click on Update to generate a new access token
- `start_date` (String) Earliest date you want to sync historical streams (inventory_histories, asset_histories, asset_stock_histories) from
- `subdomain` (String, Sensitive) The company name used in signup, also visible in the URL when logged in.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_ezofficeinventory.my_airbyte_source_ezofficeinventory ""
```