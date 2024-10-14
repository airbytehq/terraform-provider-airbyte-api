// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationGoogleSheets struct {
	Credentials   DestinationGoogleSheetsAuthenticationViaGoogleOAuth `tfsdk:"credentials"`
	SpreadsheetID types.String                                        `tfsdk:"spreadsheet_id"`
}
