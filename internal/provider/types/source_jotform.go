// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceJotform struct {
	APIEndpoint SourceJotformAPIEndpoint `tfsdk:"api_endpoint"`
	APIKey      types.String             `tfsdk:"api_key"`
	EndDate     types.String             `tfsdk:"end_date"`
	StartDate   types.String             `tfsdk:"start_date"`
}
