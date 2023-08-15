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
var _ datasource.DataSource = &DestinationKinesisDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationKinesisDataSource{}

func NewDestinationKinesisDataSource() datasource.DataSource {
	return &DestinationKinesisDataSource{}
}

// DestinationKinesisDataSource is the data source implementation.
type DestinationKinesisDataSource struct {
	client *sdk.SDK
}

// DestinationKinesisDataSourceModel describes the data model.
type DestinationKinesisDataSourceModel struct {
	Configuration DestinationKinesis `tfsdk:"configuration"`
	DestinationID types.String       `tfsdk:"destination_id"`
	Name          types.String       `tfsdk:"name"`
	WorkspaceID   types.String       `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationKinesisDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_kinesis"
}

// Schema defines the schema for the data source.
func (r *DestinationKinesisDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationKinesis DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"access_key": schema.StringAttribute{
						Computed:    true,
						Description: `Generate the AWS Access Key for current user.`,
					},
					"buffer_size": schema.Int64Attribute{
						Computed:    true,
						Description: `Buffer size for storing kinesis records before being batch streamed.`,
					},
					"destination_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"kinesis",
							),
						},
						Description: `must be one of ["kinesis"]`,
					},
					"endpoint": schema.StringAttribute{
						Computed:    true,
						Description: `AWS Kinesis endpoint.`,
					},
					"private_key": schema.StringAttribute{
						Computed:    true,
						Description: `The AWS Private Key - a string of numbers and letters that are unique for each account, also known as a "recovery phrase".`,
					},
					"region": schema.StringAttribute{
						Computed:    true,
						Description: `AWS region. Your account determines the Regions that are available to you.`,
					},
					"shard_count": schema.Int64Attribute{
						Computed:    true,
						Description: `Number of shards to which the data should be streamed.`,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *DestinationKinesisDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DestinationKinesisDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationKinesisDataSourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationKinesisRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationKinesis(ctx, request)
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
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
