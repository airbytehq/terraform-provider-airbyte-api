// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type SourceZendeskSunshineAuthorizationMethod struct {
	APIToken *SourceZendeskSunshineAPIToken `tfsdk:"api_token" tfPlanOnly:"true"`
	OAuth20  *SourceNotionOAuth20           `tfsdk:"o_auth20" tfPlanOnly:"true"`
}
