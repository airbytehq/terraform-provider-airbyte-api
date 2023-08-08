---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_salesforce Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSalesforce DataSource
---

# airbyte_source_salesforce (Data Source)

SourceSalesforce DataSource

## Example Usage

```terraform
data "airbyte_source_salesforce" "my_source_salesforce" {
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

- `auth_type` (String) must be one of ["Client"]
- `client_id` (String) Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client ID</a>
- `client_secret` (String) Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client secret</a>
- `is_sandbox` (Boolean) Toggle if you're using a <a href="https://help.salesforce.com/s/articleView?id=sf.deploy_sandboxes_parent.htm&type=5">Salesforce Sandbox</a>
- `refresh_token` (String) Enter your application's <a href="https://developer.salesforce.com/docs/atlas.en-us.mobile_sdk.meta/mobile_sdk/oauth_refresh_token_flow.htm">Salesforce Refresh Token</a> used for Airbyte to access your Salesforce account.
- `source_type` (String) must be one of ["salesforce"]
- `start_date` (String) Enter the date in the YYYY-MM-DD format. Airbyte will replicate the data added on and after this date. If this field is blank, Airbyte will replicate the data for last two years.
- `streams_criteria` (Attributes List) Filter streams relevant to you (see [below for nested schema](#nestedatt--configuration--streams_criteria))

<a id="nestedatt--configuration--streams_criteria"></a>
### Nested Schema for `configuration.streams_criteria`

Read-Only:

- `criteria` (String) must be one of ["starts with", "ends with", "contains", "exacts", "starts not with", "ends not with", "not contains", "not exacts"]
- `value` (String)

