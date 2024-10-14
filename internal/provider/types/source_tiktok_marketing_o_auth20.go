// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceTiktokMarketingOAuth20 struct {
	AccessToken  types.String `tfsdk:"access_token"`
	AdvertiserID types.String `tfsdk:"advertiser_id"`
	AppID        types.String `tfsdk:"app_id"`
	Secret       types.String `tfsdk:"secret"`
}
