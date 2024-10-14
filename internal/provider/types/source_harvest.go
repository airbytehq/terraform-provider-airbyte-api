// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceHarvest struct {
	AccountID            types.String                          `tfsdk:"account_id"`
	Credentials          *SourceHarvestAuthenticationMechanism `tfsdk:"credentials"`
	ReplicationStartDate types.String                          `tfsdk:"replication_start_date"`
}
