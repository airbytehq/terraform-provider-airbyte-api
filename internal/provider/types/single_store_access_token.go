// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SingleStoreAccessToken struct {
	AccessToken types.String `tfsdk:"access_token"`
	StoreName   types.String `tfsdk:"store_name"`
}
