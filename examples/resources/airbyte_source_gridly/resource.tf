resource "airbyte_source_gridly" "my_source_gridly" {
  configuration = {
    api_key = "...my_api_key..."
    grid_id = "...my_grid_id..."
  }
  definition_id = "4f174cea-63dc-44b8-a077-9af4686f7056"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "c30bfc4e-2cf6-47af-a03b-bb623a3a73f9"
}