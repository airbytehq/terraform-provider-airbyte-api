// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationSnowflakeCortex struct {
	Embedding   DestinationAstraEmbedding                     `tfsdk:"embedding"`
	Indexing    DestinationSnowflakeCortexSnowflakeConnection `tfsdk:"indexing"`
	OmitRawText types.Bool                                    `tfsdk:"omit_raw_text"`
	Processing  DestinationAstraProcessingConfigModel         `tfsdk:"processing"`
}
