// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleDrive struct {
	Credentials SourceGoogleDriveAuthentication `tfsdk:"credentials"`
	FolderURL   types.String                    `tfsdk:"folder_url"`
	StartDate   types.String                    `tfsdk:"start_date"`
	Streams     []FileBasedStreamConfig         `tfsdk:"streams"`
}
