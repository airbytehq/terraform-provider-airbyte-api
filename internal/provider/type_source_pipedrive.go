// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourcePipedrive struct {
	Authorization        *SourcePipedriveAPIKeyAuthentication `tfsdk:"authorization"`
	ReplicationStartDate types.String                         `tfsdk:"replication_start_date"`
	SourceType           types.String                         `tfsdk:"source_type"`
}