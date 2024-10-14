// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type ReportConfig struct {
	AttributionTypes     []types.String `tfsdk:"attribution_types"`
	ClickWindowDays      types.Int64    `tfsdk:"click_window_days"`
	Columns              []types.String `tfsdk:"columns"`
	ConversionReportTime types.String   `tfsdk:"conversion_report_time"`
	EngagementWindowDays types.Int64    `tfsdk:"engagement_window_days"`
	Granularity          types.String   `tfsdk:"granularity"`
	Level                types.String   `tfsdk:"level"`
	Name                 types.String   `tfsdk:"name"`
	StartDate            types.String   `tfsdk:"start_date"`
	ViewWindowDays       types.Int64    `tfsdk:"view_window_days"`
}
