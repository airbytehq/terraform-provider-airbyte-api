// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationTeradata struct {
	Host          types.String                 `tfsdk:"host"`
	JdbcURLParams types.String                 `tfsdk:"jdbc_url_params"`
	Password      types.String                 `tfsdk:"password"`
	Schema        types.String                 `tfsdk:"schema"`
	Ssl           types.Bool                   `tfsdk:"ssl"`
	SslMode       *DestinationTeradataSSLModes `tfsdk:"ssl_mode"`
	Username      types.String                 `tfsdk:"username"`
}
