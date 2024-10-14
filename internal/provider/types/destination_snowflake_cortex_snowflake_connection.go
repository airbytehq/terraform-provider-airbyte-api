// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationSnowflakeCortexSnowflakeConnection struct {
	Credentials   DestinationPgvectorCredentials `tfsdk:"credentials"`
	Database      types.String                   `tfsdk:"database"`
	DefaultSchema types.String                   `tfsdk:"default_schema"`
	Host          types.String                   `tfsdk:"host"`
	Role          types.String                   `tfsdk:"role"`
	Username      types.String                   `tfsdk:"username"`
	Warehouse     types.String                   `tfsdk:"warehouse"`
}
