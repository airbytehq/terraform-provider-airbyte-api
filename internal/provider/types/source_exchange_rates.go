// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceExchangeRates struct {
	AccessKey      types.String `tfsdk:"access_key"`
	Base           types.String `tfsdk:"base"`
	IgnoreWeekends types.Bool   `tfsdk:"ignore_weekends"`
	StartDate      types.String `tfsdk:"start_date"`
}
