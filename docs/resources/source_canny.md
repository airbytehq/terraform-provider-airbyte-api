---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_canny Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceCanny Resource
---

# airbyte_source_canny (Resource)

SourceCanny Resource

## Example Usage

```terraform
resource "airbyte_source_canny" "my_source_canny" {
  configuration = {
    api_key = "...my_api_key..."
  }
  definition_id = "16442d85-f5b6-4382-a70e-18a8172f9322"
  name          = "Rebecca Satterfield"
  secret_id     = "...my_secret_id..."
  workspace_id  = "9cbaa542-e6e0-4809-a1d8-4c3fbc24f860"
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

- `api_key` (String, Sensitive) You can find your secret API key in Your Canny Subdomain > Settings > API

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_canny.my_airbyte_source_canny ""
```