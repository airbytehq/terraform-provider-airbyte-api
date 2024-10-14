// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationAzureBlobStorage struct {
	AzureBlobStorageAccountKey         types.String                            `tfsdk:"azure_blob_storage_account_key"`
	AzureBlobStorageAccountName        types.String                            `tfsdk:"azure_blob_storage_account_name"`
	AzureBlobStorageContainerName      types.String                            `tfsdk:"azure_blob_storage_container_name"`
	AzureBlobStorageEndpointDomainName types.String                            `tfsdk:"azure_blob_storage_endpoint_domain_name"`
	AzureBlobStorageOutputBufferSize   types.Int64                             `tfsdk:"azure_blob_storage_output_buffer_size"`
	AzureBlobStorageSpillSize          types.Int64                             `tfsdk:"azure_blob_storage_spill_size"`
	Format                             DestinationAzureBlobStorageOutputFormat `tfsdk:"format"`
}
