// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationGcsJSONLinesNewlineDelimitedJSON struct {
	Compression *DestinationGcsCompression `tfsdk:"compression"`
	FormatType  types.String               `tfsdk:"format_type"`
}
