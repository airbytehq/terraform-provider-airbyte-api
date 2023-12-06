resource "airbyte_destination_s3" "my_destination_s3" {
  configuration = {
    access_key_id     = "A012345678910EXAMPLE"
    file_name_pattern = "{date}"
    format = {
      destination_s3_avro_apache_avro = {
        compression_codec = {
          destination_s3_bzip2 = {
            codec = "bzip2"
          }
        }
        format_type = "Avro"
      }
    }
    s3_bucket_name    = "airbyte_sync"
    s3_bucket_path    = "data_sync/test"
    s3_bucket_region  = "ap-southeast-1"
    s3_endpoint       = "http://localhost:9000"
    s3_path_format    = "$${NAMESPACE}/$${STREAM_NAME}/$${YEAR}_$${MONTH}_$${DAY}_$${EPOCH}_"
    secret_access_key = "a012345678910ABCDEFGH/AbCdEfGhEXAMPLEKEY"
  }
  definition_id = "b1d5b002-89a0-4dc0-a329-a5cae9f38884"
  name          = "Lloyd Watsica"
  workspace_id  = "20ebb305-f362-44c4-b900-725fa3e33722"
}
