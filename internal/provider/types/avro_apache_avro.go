// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AvroApacheAvro struct {
	CompressionCodec DestinationGcsCompressionCodec `tfsdk:"compression_codec"`
	FormatType       types.String                   `tfsdk:"format_type"`
}
