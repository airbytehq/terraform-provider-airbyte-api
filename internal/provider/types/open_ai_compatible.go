// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OpenAICompatible struct {
	APIKey     types.String `tfsdk:"api_key"`
	BaseURL    types.String `tfsdk:"base_url"`
	Dimensions types.Int64  `tfsdk:"dimensions"`
	ModelName  types.String `tfsdk:"model_name"`
}
