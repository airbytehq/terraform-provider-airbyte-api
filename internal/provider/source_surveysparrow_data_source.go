// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceSurveySparrowDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceSurveySparrowDataSource{}

func NewSourceSurveySparrowDataSource() datasource.DataSource {
	return &SourceSurveySparrowDataSource{}
}

// SourceSurveySparrowDataSource is the data source implementation.
type SourceSurveySparrowDataSource struct {
	client *sdk.SDK
}

// SourceSurveySparrowDataSourceModel describes the data model.
type SourceSurveySparrowDataSourceModel struct {
	Configuration SourceSurveySparrow `tfsdk:"configuration"`
	Name          types.String        `tfsdk:"name"`
	SecretID      types.String        `tfsdk:"secret_id"`
	SourceID      types.String        `tfsdk:"source_id"`
	WorkspaceID   types.String        `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceSurveySparrowDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_survey_sparrow"
}

// Schema defines the schema for the data source.
func (r *SourceSurveySparrowDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceSurveySparrow DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"access_token": schema.StringAttribute{
						Computed:    true,
						Description: `Your access token. See <a href="https://developers.surveysparrow.com/rest-apis#authentication">here</a>. The key is case sensitive.`,
					},
					"region": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_survey_sparrow_base_url_eu_based_account": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"url_base": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"https://eu-api.surveysparrow.com/v3",
											),
										},
										Description: `must be one of ["https://eu-api.surveysparrow.com/v3"]`,
									},
								},
								Description: `Is your account location is EU based? If yes, the base url to retrieve data will be different.`,
							},
							"source_survey_sparrow_base_url_global_account": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"url_base": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"https://api.surveysparrow.com/v3",
											),
										},
										Description: `must be one of ["https://api.surveysparrow.com/v3"]`,
									},
								},
								Description: `Is your account location is EU based? If yes, the base url to retrieve data will be different.`,
							},
							"source_survey_sparrow_update_base_url_eu_based_account": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"url_base": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"https://eu-api.surveysparrow.com/v3",
											),
										},
										Description: `must be one of ["https://eu-api.surveysparrow.com/v3"]`,
									},
								},
								Description: `Is your account location is EU based? If yes, the base url to retrieve data will be different.`,
							},
							"source_survey_sparrow_update_base_url_global_account": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"url_base": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"https://api.surveysparrow.com/v3",
											),
										},
										Description: `must be one of ["https://api.surveysparrow.com/v3"]`,
									},
								},
								Description: `Is your account location is EU based? If yes, the base url to retrieve data will be different.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Is your account location is EU based? If yes, the base url to retrieve data will be different.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"survey-sparrow",
							),
						},
						Description: `must be one of ["survey-sparrow"]`,
					},
					"survey_id": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(validators.IsValidJSON()),
						},
						Description: `A List of your survey ids for survey-specific stream`,
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

func (r *SourceSurveySparrowDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceSurveySparrowDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceSurveySparrowDataSourceModel
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
	request := operations.GetSourceSurveySparrowRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceSurveySparrow(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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