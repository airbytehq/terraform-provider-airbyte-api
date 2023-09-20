---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_appfollow Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAppfollow DataSource
---

# airbyte_source_appfollow (Data Source)

SourceAppfollow DataSource

## Example Usage

```terraform
data "airbyte_source_appfollow" "my_source_appfollow" {
  secret_id = "...my_secret_id..."
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `api_secret` (String) API Key provided by Appfollow
- `source_type` (String) must be one of ["appfollow"]

