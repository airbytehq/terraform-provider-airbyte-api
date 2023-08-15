// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceZuoraDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceZuoraDataSource{}

func NewSourceZuoraDataSource() datasource.DataSource {
	return &SourceZuoraDataSource{}
}

// SourceZuoraDataSource is the data source implementation.
type SourceZuoraDataSource struct {
	client *sdk.SDK
}

// SourceZuoraDataSourceModel describes the data model.
type SourceZuoraDataSourceModel struct {
	Configuration SourceZuora  `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceZuoraDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_zuora"
}

// Schema defines the schema for the data source.
func (r *SourceZuoraDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceZuora DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Computed:    true,
						Description: `Your OAuth user Client ID`,
					},
					"client_secret": schema.StringAttribute{
						Computed:    true,
						Description: `Your OAuth user Client Secret`,
					},
					"data_query": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Live",
								"Unlimited",
							),
						},
						MarkdownDescription: `must be one of ["Live", "Unlimited"]` + "\n" +
							`Choose between ` + "`" + `Live` + "`" + `, or ` + "`" + `Unlimited` + "`" + ` - the optimized, replicated database at 12 hours freshness for high volume extraction <a href="https://knowledgecenter.zuora.com/Central_Platform/Query/Data_Query/A_Overview_of_Data_Query#Query_Processing_Limitations">Link</a>`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"zuora",
							),
						},
						Description: `must be one of ["zuora"]`,
					},
					"start_date": schema.StringAttribute{
						Computed:    true,
						Description: `Start Date in format: YYYY-MM-DD`,
					},
					"tenant_endpoint": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"US Production",
								"US Cloud Production",
								"US API Sandbox",
								"US Cloud API Sandbox",
								"US Central Sandbox",
								"US Performance Test",
								"EU Production",
								"EU API Sandbox",
								"EU Central Sandbox",
							),
						},
						MarkdownDescription: `must be one of ["US Production", "US Cloud Production", "US API Sandbox", "US Cloud API Sandbox", "US Central Sandbox", "US Performance Test", "EU Production", "EU API Sandbox", "EU Central Sandbox"]` + "\n" +
							`Please choose the right endpoint where your Tenant is located. More info by this <a href="https://www.zuora.com/developer/api-reference/#section/Introduction/Access-to-the-API">Link</a>`,
					},
					"window_in_days": schema.StringAttribute{
						Computed:    true,
						Description: `The amount of days for each data-chunk begining from start_date. Bigger the value - faster the fetch. (0.1 - as for couple of hours, 1 - as for a Day; 364 - as for a Year).`,
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

func (r *SourceZuoraDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceZuoraDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceZuoraDataSourceModel
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
	request := operations.GetSourceZuoraRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceZuora(ctx, request)
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
