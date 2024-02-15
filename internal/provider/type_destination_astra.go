// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationAstra struct {
	Embedding   DestinationAstraEmbedding             `tfsdk:"embedding"`
	Indexing    DestinationAstraIndexing              `tfsdk:"indexing"`
	OmitRawText types.Bool                            `tfsdk:"omit_raw_text"`
	Processing  DestinationAstraProcessingConfigModel `tfsdk:"processing"`
}
