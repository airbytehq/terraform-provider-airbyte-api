// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceBuzzsprout struct {
	APIKey    types.String `tfsdk:"api_key"`
	PodcastID types.String `tfsdk:"podcast_id"`
	StartDate types.String `tfsdk:"start_date"`
}
