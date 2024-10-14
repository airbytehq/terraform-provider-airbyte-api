// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationMilvusIndexing struct {
	Auth        DestinationMilvusAuthentication `tfsdk:"auth"`
	Collection  types.String                    `tfsdk:"collection"`
	Db          types.String                    `tfsdk:"db"`
	Host        types.String                    `tfsdk:"host"`
	TextField   types.String                    `tfsdk:"text_field"`
	VectorField types.String                    `tfsdk:"vector_field"`
}
