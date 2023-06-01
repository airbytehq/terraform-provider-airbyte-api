// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMailgunUpdate struct {
	DomainRegion types.String `tfsdk:"domain_region"`
	PrivateKey   types.String `tfsdk:"private_key"`
	StartDate    types.String `tfsdk:"start_date"`
	SourceType   types.String `tfsdk:"source_type"`
}