// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type SourceNotionAuthenticationMethod struct {
	AccessToken *DestinationMilvusAPIToken `tfsdk:"access_token" tfPlanOnly:"true"`
	OAuth20     *SourceNotionOAuth20       `tfsdk:"o_auth20" tfPlanOnly:"true"`
}
