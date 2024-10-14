// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type SourceOktaAuthorizationMethod struct {
	APIToken              *SourceK6Cloud                                       `tfsdk:"api_token" tfPlanOnly:"true"`
	OAuth20WithPrivateKey *OAuth20WithPrivateKey                               `tfsdk:"o_auth20_with_private_key" tfPlanOnly:"true"`
	OAuth20               *DestinationGoogleSheetsAuthenticationViaGoogleOAuth `tfsdk:"o_auth20" tfPlanOnly:"true"`
}
