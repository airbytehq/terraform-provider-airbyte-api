resource "airbyte_destination_astra" "my_destination_astra" {
  configuration = {
    embedding = {
      azure_open_ai = {
        api_base   = "https://your-resource-name.openai.azure.com"
        deployment = "your-resource-name"
        openai_key = "...my_openai_key..."
      }
      cohere = {
        cohere_key = "...my_cohere_key..."
      }
      fake = {
        # ...
      }
      open_ai = {
        openai_key = "...my_openai_key..."
      }
      open_ai_compatible = {
        api_key    = "...my_api_key..."
        base_url   = "https://your-service-name.com"
        dimensions = 1536
        model_name = "text-embedding-ada-002"
      }
    }
    indexing = {
      astra_db_app_token = "...my_astra_db_app_token..."
      astra_db_endpoint  = "https://8292d414-dd1b-4c33-8431-e838bedc04f7-us-east1.apps.astra.datastax.com"
      astra_db_keyspace  = "...my_astra_db_keyspace..."
      collection         = "...my_collection..."
    }
    omit_raw_text = false
    processing = {
      chunk_overlap = 6
      chunk_size    = 2127
      field_name_mappings = [
        {
          from_field = "...my_from_field..."
          to_field   = "...my_to_field..."
        }
      ]
      metadata_fields = [
        "..."
      ]
      text_fields = [
        "..."
      ]
      text_splitter = {
        by_markdown_header = {
          split_level = 1
        }
        by_programming_language = {
          language = "js"
        }
        by_separator = {
          keep_separator = false
          separators = [
            "..."
          ]
        }
      }
    }
  }
  definition_id = "79152260-aed1-4b65-bb98-3dd0b8ec05bd"
  name          = "...my_name..."
  workspace_id  = "543d6f27-bf11-4034-a571-6e04a190e68b"
}