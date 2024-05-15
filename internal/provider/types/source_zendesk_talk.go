// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceZendeskTalk struct {
	Credentials *SourceZendeskSupportAuthentication `tfsdk:"credentials"`
	StartDate   types.String                        `tfsdk:"start_date"`
	Subdomain   types.String                        `tfsdk:"subdomain"`
}