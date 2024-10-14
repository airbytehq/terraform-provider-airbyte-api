resource "airbyte_source_s3" "my_source_s3" {
  configuration = {
    aws_access_key_id     = "...my_aws_access_key_id..."
    aws_secret_access_key = "...my_aws_secret_access_key..."
    bucket                = "...my_bucket..."
    endpoint              = "my-s3-endpoint.com"
    region_name           = "...my_region_name..."
    role_arn              = "...my_role_arn..."
    start_date            = "2021-01-01T00:00:00.000000Z"
    streams = [
      {
        days_to_sync_if_history_is_full = 5
        format = {
          avro_format = {
            double_as_string = true
          }
          csv_format = {
            delimiter    = "...my_delimiter..."
            double_quote = false
            encoding     = "...my_encoding..."
            escape_char  = "...my_escape_char..."
            false_values = [
              "..."
            ]
            header_definition = {
              autogenerated = {
                # ...
              }
              from_csv = {
                # ...
              }
              user_provided = {
                column_names = [
                  "..."
                ]
              }
            }
            ignore_errors_on_fields_mismatch = true
            null_values = [
              "..."
            ]
            quote_char              = "...my_quote_char..."
            skip_rows_after_header  = 4
            skip_rows_before_header = 7
            strings_can_be_null     = true
            true_values = [
              "..."
            ]
          }
          jsonl_format = {
            # ...
          }
          parquet_format = {
            decimal_as_float = false
          }
          unstructured_document_format = {
            processing = {
              local = {
                # ...
              }
            }
            skip_unprocessable_files = true
            strategy                 = "auto"
          }
        }
        globs = [
          "..."
        ]
        input_schema                                = "...my_input_schema..."
        name                                        = "...my_name..."
        recent_n_files_to_read_for_schema_discovery = 10
        schemaless                                  = true
        validation_policy                           = "Wait for Discover"
      }
    ]
  }
  definition_id = "07ef8ae4-b6a4-4fd9-99ea-a368c6fc144c"
  name          = "...my_name..."
  secret_id     = "...my_secret_id..."
  workspace_id  = "bba7dce0-5020-4916-bbd7-be8f298d5f78"
}