// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type SourceLinkedinAdsAuthentication struct {
	AccessToken *OAuth2AccessToken                                   `tfsdk:"access_token" tfPlanOnly:"true"`
	OAuth20     *DestinationGoogleSheetsAuthenticationViaGoogleOAuth `tfsdk:"o_auth20" tfPlanOnly:"true"`
}
