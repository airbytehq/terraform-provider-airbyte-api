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
var _ datasource.DataSource = &SourceYandexMetricaDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceYandexMetricaDataSource{}

func NewSourceYandexMetricaDataSource() datasource.DataSource {
	return &SourceYandexMetricaDataSource{}
}

// SourceYandexMetricaDataSource is the data source implementation.
type SourceYandexMetricaDataSource struct {
	client *sdk.SDK
}

// SourceYandexMetricaDataSourceModel describes the data model.
type SourceYandexMetricaDataSourceModel struct {
	Configuration SourceYandexMetrica `tfsdk:"configuration"`
	Name          types.String        `tfsdk:"name"`
	SecretID      types.String        `tfsdk:"secret_id"`
	SourceID      types.String        `tfsdk:"source_id"`
	WorkspaceID   types.String        `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceYandexMetricaDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_yandex_metrica"
}

// Schema defines the schema for the data source.
func (r *SourceYandexMetricaDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceYandexMetrica DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"auth_token": schema.StringAttribute{
						Computed:    true,
						Description: `Your Yandex Metrica API access token`,
					},
					"counter_id": schema.StringAttribute{
						Computed:    true,
						Description: `Counter ID`,
					},
					"end_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `Starting point for your data replication, in format of "YYYY-MM-DD". If not provided will sync till most recent date.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"yandex-metrica",
							),
						},
						Description: `must be one of ["yandex-metrica"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `Starting point for your data replication, in format of "YYYY-MM-DD".`,
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

func (r *SourceYandexMetricaDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceYandexMetricaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceYandexMetricaDataSourceModel
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
	request := operations.GetSourceYandexMetricaRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceYandexMetrica(ctx, request)
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