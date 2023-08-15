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
var _ datasource.DataSource = &SourceGoogleSearchConsoleDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceGoogleSearchConsoleDataSource{}

func NewSourceGoogleSearchConsoleDataSource() datasource.DataSource {
	return &SourceGoogleSearchConsoleDataSource{}
}

// SourceGoogleSearchConsoleDataSource is the data source implementation.
type SourceGoogleSearchConsoleDataSource struct {
	client *sdk.SDK
}

// SourceGoogleSearchConsoleDataSourceModel describes the data model.
type SourceGoogleSearchConsoleDataSourceModel struct {
	Configuration SourceGoogleSearchConsole `tfsdk:"configuration"`
	Name          types.String              `tfsdk:"name"`
	SecretID      types.String              `tfsdk:"secret_id"`
	SourceID      types.String              `tfsdk:"source_id"`
	WorkspaceID   types.String              `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceGoogleSearchConsoleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_google_search_console"
}

// Schema defines the schema for the data source.
func (r *SourceGoogleSearchConsoleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceGoogleSearchConsole DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"authorization": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_google_search_console_authentication_type_o_auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Access token for making authenticated requests. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
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
										Description: `The client ID of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The client secret of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `The token for obtaining a new access token. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
									},
								},
							},
							"source_google_search_console_authentication_type_service_account_key_authentication": schema.SingleNestedAttribute{
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
									"email": schema.StringAttribute{
										Computed:    true,
										Description: `The email of the user which has permissions to access the Google Workspace Admin APIs.`,
									},
									"service_account_info": schema.StringAttribute{
										Computed:    true,
										Description: `The JSON key of the service account to use for authorization. Read more <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys">here</a>.`,
									},
								},
							},
							"source_google_search_console_update_authentication_type_o_auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Access token for making authenticated requests. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
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
										Description: `The client ID of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The client secret of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `The token for obtaining a new access token. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
									},
								},
							},
							"source_google_search_console_update_authentication_type_service_account_key_authentication": schema.SingleNestedAttribute{
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
									"email": schema.StringAttribute{
										Computed:    true,
										Description: `The email of the user which has permissions to access the Google Workspace Admin APIs.`,
									},
									"service_account_info": schema.StringAttribute{
										Computed:    true,
										Description: `The JSON key of the service account to use for authorization. Read more <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys">here</a>.`,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"custom_reports": schema.StringAttribute{
						Computed:    true,
						Description: `A JSON array describing the custom reports you want to sync from Google Search Console. See <a href="https://docs.airbyte.com/integrations/sources/google-search-console#step-2-set-up-the-google-search-console-connector-in-airbyte">the docs</a> for more information about the exact format you can use to fill out this field.`,
					},
					"data_state": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"final",
								"all",
							),
						},
						MarkdownDescription: `must be one of ["final", "all"]` + "\n" +
							`If "final" or if this parameter is omitted, the returned data will include only finalized data. Setting this parameter to "all" should not be used with Incremental Sync mode as it may cause data loss. If "all", data will include fresh data.`,
					},
					"end_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `UTC date in the format 2017-01-25. Any data after this date will not be replicated. Must be greater or equal to the start date field.`,
					},
					"site_urls": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `The URLs of the website property attached to your GSC account. Read more <a href="https://support.google.com/webmasters/answer/34592?hl=en">here</a>.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"google-search-console",
							),
						},
						Description: `must be one of ["google-search-console"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `UTC date in the format 2017-01-25. Any data before this date will not be replicated.`,
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

func (r *SourceGoogleSearchConsoleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceGoogleSearchConsoleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceGoogleSearchConsoleDataSourceModel
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
	request := operations.GetSourceGoogleSearchConsoleRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceGoogleSearchConsole(ctx, request)
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
