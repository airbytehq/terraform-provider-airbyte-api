// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type BySeparator struct {
	KeepSeparator types.Bool     `tfsdk:"keep_separator"`
	Separators    []types.String `tfsdk:"separators"`
}
