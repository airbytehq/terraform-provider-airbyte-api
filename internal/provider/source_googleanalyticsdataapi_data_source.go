// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceGoogleAnalyticsDataAPIDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceGoogleAnalyticsDataAPIDataSource{}

func NewSourceGoogleAnalyticsDataAPIDataSource() datasource.DataSource {
	return &SourceGoogleAnalyticsDataAPIDataSource{}
}

// SourceGoogleAnalyticsDataAPIDataSource is the data source implementation.
type SourceGoogleAnalyticsDataAPIDataSource struct {
	client *sdk.SDK
}

// SourceGoogleAnalyticsDataAPIDataSourceModel describes the data model.
type SourceGoogleAnalyticsDataAPIDataSourceModel struct {
	Configuration SourceGoogleAnalyticsDataAPI `tfsdk:"configuration"`
	Name          types.String                 `tfsdk:"name"`
	SecretID      types.String                 `tfsdk:"secret_id"`
	SourceID      types.String                 `tfsdk:"source_id"`
	WorkspaceID   types.String                 `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceGoogleAnalyticsDataAPIDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_google_analytics_data_api"
}

// Schema defines the schema for the data source.
func (r *SourceGoogleAnalyticsDataAPIDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceGoogleAnalyticsDataAPI DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_google_analytics_data_api_credentials_authenticate_via_google_oauth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Access Token for making authenticated requests.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Client",
											),
										},
										Description: `must be one of ["Client"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of your Google Analytics developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Client Secret of your Google Analytics developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `The token for obtaining a new access token.`,
									},
								},
								Description: `Credentials for the service`,
							},
							"source_google_analytics_data_api_credentials_service_account_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Service",
											),
										},
										Description: `must be one of ["Service"]`,
									},
									"credentials_json": schema.StringAttribute{
										Computed:    true,
										Description: `The JSON key of the service account to use for authorization`,
									},
								},
								Description: `Credentials for the service`,
							},
							"source_google_analytics_data_api_update_credentials_authenticate_via_google_oauth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Access Token for making authenticated requests.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Client",
											),
										},
										Description: `must be one of ["Client"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of your Google Analytics developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Client Secret of your Google Analytics developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `The token for obtaining a new access token.`,
									},
								},
								Description: `Credentials for the service`,
							},
							"source_google_analytics_data_api_update_credentials_service_account_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Service",
											),
										},
										Description: `must be one of ["Service"]`,
									},
									"credentials_json": schema.StringAttribute{
										Computed:    true,
										Description: `The JSON key of the service account to use for authorization`,
									},
								},
								Description: `Credentials for the service`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Credentials for the service`,
					},
					"custom_reports": schema.StringAttribute{
						Computed:    true,
						Description: `A JSON array describing the custom reports you want to sync from Google Analytics. See <a href="https://docs.airbyte.com/integrations/sources/google-analytics-v4/#custom-reports">the docs</a> for more information about the exact format you can use to fill out this field.`,
					},
					"date_ranges_start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `The start date from which to replicate report data in the format YYYY-MM-DD. Data generated before this date will not be included in the report. Not applied to custom Cohort reports.`,
					},
					"property_id": schema.StringAttribute{
						Computed:    true,
						Description: `A Google Analytics GA4 property identifier whose events are tracked. Specified in the URL path and not the body such as "123...". See <a href="https://developers.google.com/analytics/devguides/reporting/data/v1/property-id#what_is_my_property_id">the docs</a> for more details.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"google-analytics-data-api",
							),
						},
						Description: `must be one of ["google-analytics-data-api"]`,
					},
					"window_in_days": schema.Int64Attribute{
						Computed:    true,
						Description: `The time increment used by the connector when requesting data from the Google Analytics API. More information is available in the <a href="https://docs.airbyte.com/integrations/sources/google-analytics-v4/#sampling-in-reports">the docs</a>. The bigger this value is, the faster the sync will be, but the more likely that sampling will be applied to your data, potentially causing inaccuracies in the returned results. We recommend setting this to 1 unless you have a hard requirement to make the sync faster at the expense of accuracy. The minimum allowed value for this field is 1, and the maximum is 364. Not applied to custom Cohort reports.`,
					},
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"secret_id": schema.StringAttribute{
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *SourceGoogleAnalyticsDataAPIDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceGoogleAnalyticsDataAPIDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceGoogleAnalyticsDataAPIDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.GetSourceGoogleAnalyticsDataAPIRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceGoogleAnalyticsDataAPI(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
