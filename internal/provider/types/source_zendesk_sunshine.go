// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceZendeskSunshine struct {
	Credentials *SourceZendeskSunshineAuthorizationMethod `tfsdk:"credentials"`
	StartDate   types.String                              `tfsdk:"start_date"`
	Subdomain   types.String                              `tfsdk:"subdomain"`
}
