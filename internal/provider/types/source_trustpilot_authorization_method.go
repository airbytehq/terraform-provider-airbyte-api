// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type SourceTrustpilotAuthorizationMethod struct {
	APIKey  *SourceTrustpilotAPIKey `tfsdk:"api_key" tfPlanOnly:"true"`
	OAuth20 *SourceGitlabOAuth20    `tfsdk:"o_auth20" tfPlanOnly:"true"`
}
