// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Parquet struct {
	BatchSize  types.Int64    `tfsdk:"batch_size"`
	BufferSize types.Int64    `tfsdk:"buffer_size"`
	Columns    []types.String `tfsdk:"columns"`
}