// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceZendeskSunshineUpdate struct {
	Credentials *SourceZendeskSunshineUpdateAuthorizationMethod `tfsdk:"credentials"`
	StartDate   types.String                                    `tfsdk:"start_date"`
	Subdomain   types.String                                    `tfsdk:"subdomain"`
	SourceType  types.String                                    `tfsdk:"source_type"`
}