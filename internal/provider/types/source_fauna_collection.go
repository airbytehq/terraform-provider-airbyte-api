// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceFaunaCollection struct {
	Deletions SourceFaunaDeletionMode `tfsdk:"deletions"`
	PageSize  types.Int64             `tfsdk:"page_size"`
}
