// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type SourceSalesloftCredentials struct {
	AuthenticateViaAPIKey *APIKeyAuth          `tfsdk:"authenticate_via_api_key" tfPlanOnly:"true"`
	AuthenticateViaOAuth  *SourceGitlabOAuth20 `tfsdk:"authenticate_via_o_auth" tfPlanOnly:"true"`
}
