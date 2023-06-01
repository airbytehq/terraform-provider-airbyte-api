// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMyHoursUpdate struct {
	Email         types.String `tfsdk:"email"`
	LogsBatchSize types.Int64  `tfsdk:"logs_batch_size"`
	Password      types.String `tfsdk:"password"`
	StartDate     types.String `tfsdk:"start_date"`
	SourceType    types.String `tfsdk:"source_type"`
}