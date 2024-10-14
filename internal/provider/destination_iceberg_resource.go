// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_objectplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/airbytehq/terraform-provider-airbyte/internal/provider/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationIcebergResource{}
var _ resource.ResourceWithImportState = &DestinationIcebergResource{}

func NewDestinationIcebergResource() resource.Resource {
	return &DestinationIcebergResource{}
}

// DestinationIcebergResource defines the resource implementation.
type DestinationIcebergResource struct {
	client *sdk.SDK
}

// DestinationIcebergResourceModel describes the resource data model.
type DestinationIcebergResourceModel struct {
	Configuration   tfTypes.DestinationIceberg `tfsdk:"configuration"`
	DefinitionID    types.String               `tfsdk:"definition_id"`
	DestinationID   types.String               `tfsdk:"destination_id"`
	DestinationType types.String               `tfsdk:"destination_type"`
	Name            types.String               `tfsdk:"name"`
	WorkspaceID     types.String               `tfsdk:"workspace_id"`
}

func (r *DestinationIcebergResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_iceberg"
}

func (r *DestinationIcebergResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationIceberg Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"catalog_config": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"glue_catalog": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"catalog_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("Glue"),
										Description: `Default: "Glue"; must be "Glue"`,
										Validators: []validator.String{
											stringvalidator.OneOf("Glue"),
										},
									},
									"database": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("public"),
										Description: `The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public". Default: "public"`,
									},
								},
								Description: `The GlueCatalog connects to a AWS Glue Catalog`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("hadoop_catalog_use_hierarchical_file_systems_as_same_as_storage_config"),
										path.MatchRelative().AtParent().AtName("hive_catalog_use_apache_hive_meta_store"),
										path.MatchRelative().AtParent().AtName("jdbc_catalog_use_relational_database"),
										path.MatchRelative().AtParent().AtName("rest_catalog"),
									}...),
								},
							},
							"hadoop_catalog_use_hierarchical_file_systems_as_same_as_storage_config": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"catalog_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("Hadoop"),
										Description: `Default: "Hadoop"; must be "Hadoop"`,
										Validators: []validator.String{
											stringvalidator.OneOf("Hadoop"),
										},
									},
									"database": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("default"),
										Description: `The default database tables are written to if the source does not specify a namespace. The usual value for this field is "default". Default: "default"`,
									},
								},
								Description: `A Hadoop catalog doesn’t need to connect to a Hive MetaStore, but can only be used with HDFS or similar file systems that support atomic rename.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("glue_catalog"),
										path.MatchRelative().AtParent().AtName("hive_catalog_use_apache_hive_meta_store"),
										path.MatchRelative().AtParent().AtName("jdbc_catalog_use_relational_database"),
										path.MatchRelative().AtParent().AtName("rest_catalog"),
									}...),
								},
							},
							"hive_catalog_use_apache_hive_meta_store": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"catalog_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("Hive"),
										Description: `Default: "Hive"; must be "Hive"`,
										Validators: []validator.String{
											stringvalidator.OneOf("Hive"),
										},
									},
									"database": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("default"),
										Description: `The default database tables are written to if the source does not specify a namespace. The usual value for this field is "default". Default: "default"`,
									},
									"hive_thrift_uri": schema.StringAttribute{
										Required:    true,
										Description: `Hive MetaStore thrift server uri of iceberg catalog.`,
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("glue_catalog"),
										path.MatchRelative().AtParent().AtName("hadoop_catalog_use_hierarchical_file_systems_as_same_as_storage_config"),
										path.MatchRelative().AtParent().AtName("jdbc_catalog_use_relational_database"),
										path.MatchRelative().AtParent().AtName("rest_catalog"),
									}...),
								},
							},
							"jdbc_catalog_use_relational_database": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"catalog_schema": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("public"),
										Description: `Iceberg catalog metadata tables are written to catalog schema. The usual value for this field is "public". Default: "public"`,
									},
									"catalog_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("Jdbc"),
										Description: `Default: "Jdbc"; must be "Jdbc"`,
										Validators: []validator.String{
											stringvalidator.OneOf("Jdbc"),
										},
									},
									"database": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("public"),
										Description: `The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public". Default: "public"`,
									},
									"jdbc_url": schema.StringAttribute{
										Optional: true,
									},
									"password": schema.StringAttribute{
										Optional:    true,
										Sensitive:   true,
										Description: `Password associated with the username.`,
									},
									"ssl": schema.BoolAttribute{
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(false),
										Description: `Encrypt data using SSL. When activating SSL, please select one of the connection modes. Default: false`,
									},
									"username": schema.StringAttribute{
										Optional:    true,
										Description: `Username to use to access the database.`,
									},
								},
								Description: `Using a table in a relational database to manage Iceberg tables through JDBC. Read more <a href="https://iceberg.apache.org/docs/latest/jdbc/">here</a>. Supporting: PostgreSQL`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("glue_catalog"),
										path.MatchRelative().AtParent().AtName("hadoop_catalog_use_hierarchical_file_systems_as_same_as_storage_config"),
										path.MatchRelative().AtParent().AtName("hive_catalog_use_apache_hive_meta_store"),
										path.MatchRelative().AtParent().AtName("rest_catalog"),
									}...),
								},
							},
							"rest_catalog": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"catalog_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("Rest"),
										Description: `Default: "Rest"; must be "Rest"`,
										Validators: []validator.String{
											stringvalidator.OneOf("Rest"),
										},
									},
									"rest_credential": schema.StringAttribute{
										Optional:  true,
										Sensitive: true,
									},
									"rest_token": schema.StringAttribute{
										Optional:  true,
										Sensitive: true,
									},
									"rest_uri": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `The RESTCatalog connects to a REST server at the specified URI`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("glue_catalog"),
										path.MatchRelative().AtParent().AtName("hadoop_catalog_use_hierarchical_file_systems_as_same_as_storage_config"),
										path.MatchRelative().AtParent().AtName("hive_catalog_use_apache_hive_meta_store"),
										path.MatchRelative().AtParent().AtName("jdbc_catalog_use_relational_database"),
									}...),
								},
							},
						},
						Description: `Catalog config of Iceberg.`,
					},
					"format_config": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"auto_compact": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
								Description: `Auto compact data files when stream close. Default: false`,
							},
							"compact_target_file_size_in_mb": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(100),
								Description: `Specify the target size of Iceberg data file when performing a compaction action. Default: 100`,
							},
							"flush_batch_size": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(10000),
								Description: `Iceberg data file flush batch size. Incoming rows write to cache firstly; When cache size reaches this 'batch size', flush into real Iceberg data file. Default: 10000`,
							},
							"format": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString("Parquet"),
								Description: `Default: "Parquet"; must be one of ["Parquet", "Avro"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Parquet",
										"Avro",
									),
								},
							},
						},
						Description: `File format of Iceberg storage.`,
					},
					"storage_config": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"s3": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `The access key ID to access the S3 bucket. Airbyte requires Read and Write permissions to the given bucket. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>.`,
									},
									"s3_bucket_region": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString(""),
										Description: `The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes. Default: ""; must be one of ["", "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-south-2", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-4", "ca-central-1", "ca-west-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-central-2", "eu-north-1", "eu-south-1", "eu-south-2", "eu-west-1", "eu-west-2", "eu-west-3", "il-central-1", "me-central-1", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2"]`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"",
												"af-south-1",
												"ap-east-1",
												"ap-northeast-1",
												"ap-northeast-2",
												"ap-northeast-3",
												"ap-south-1",
												"ap-south-2",
												"ap-southeast-1",
												"ap-southeast-2",
												"ap-southeast-3",
												"ap-southeast-4",
												"ca-central-1",
												"ca-west-1",
												"cn-north-1",
												"cn-northwest-1",
												"eu-central-1",
												"eu-central-2",
												"eu-north-1",
												"eu-south-1",
												"eu-south-2",
												"eu-west-1",
												"eu-west-2",
												"eu-west-3",
												"il-central-1",
												"me-central-1",
												"me-south-1",
												"sa-east-1",
												"us-east-1",
												"us-east-2",
												"us-gov-east-1",
												"us-gov-west-1",
												"us-west-1",
												"us-west-2",
											),
										},
									},
									"s3_endpoint": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString(""),
										Description: `Your S3 endpoint url. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/s3.html#:~:text=Service%20endpoints-,Amazon%20S3%20endpoints,-When%20you%20use">here</a>. Default: ""`,
									},
									"s3_path_style_access": schema.BoolAttribute{
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(true),
										Description: `Use path style access. Default: true`,
									},
									"s3_warehouse_uri": schema.StringAttribute{
										Required:    true,
										Description: `The Warehouse Uri for Iceberg`,
									},
									"secret_access_key": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `The corresponding secret to the access key ID. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>`,
									},
									"storage_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("S3"),
										Description: `Default: "S3"; must be "S3"`,
										Validators: []validator.String{
											stringvalidator.OneOf("S3"),
										},
									},
								},
								Description: `S3 object storage`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("server_managed"),
									}...),
								},
							},
							"server_managed": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"managed_warehouse_name": schema.StringAttribute{
										Required:    true,
										Description: `The name of the managed warehouse`,
									},
									"storage_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("MANAGED"),
										Description: `Default: "MANAGED"; must be "MANAGED"`,
										Validators: []validator.String{
											stringvalidator.OneOf("MANAGED"),
										},
									},
								},
								Description: `Server-managed object storage`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("s3"),
									}...),
								},
							},
						},
						Description: `Storage config of Iceberg.`,
					},
				},
			},
			"definition_id": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.`,
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Name of the destination e.g. dev-mysql-instance.`,
			},
			"workspace_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
		},
	}
}

func (r *DestinationIcebergResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DestinationIcebergResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationIcebergResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToSharedDestinationIcebergCreateRequest()
	res, err := r.client.Destinations.CreateDestinationIceberg(ctx, request)
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
	if !(res.DestinationResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var destinationID string
	destinationID = data.DestinationID.ValueString()

	request1 := operations.GetDestinationIcebergRequest{
		DestinationID: destinationID,
	}
	res1, err := r.client.Destinations.GetDestinationIceberg(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.DestinationResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res1.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationIcebergResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationIcebergResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var destinationID string
	destinationID = data.DestinationID.ValueString()

	request := operations.GetDestinationIcebergRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationIceberg(ctx, request)
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
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.DestinationResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationIcebergResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationIcebergResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationIcebergPutRequest := data.ToSharedDestinationIcebergPutRequest()
	var destinationID string
	destinationID = data.DestinationID.ValueString()

	request := operations.PutDestinationIcebergRequest{
		DestinationIcebergPutRequest: destinationIcebergPutRequest,
		DestinationID:                destinationID,
	}
	res, err := r.client.Destinations.PutDestinationIceberg(ctx, request)
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
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var destinationId1 string
	destinationId1 = data.DestinationID.ValueString()

	request1 := operations.GetDestinationIcebergRequest{
		DestinationID: destinationId1,
	}
	res1, err := r.client.Destinations.GetDestinationIceberg(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.DestinationResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res1.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationIcebergResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationIcebergResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var destinationID string
	destinationID = data.DestinationID.ValueString()

	request := operations.DeleteDestinationIcebergRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationIceberg(ctx, request)
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
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *DestinationIcebergResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("destination_id"), req.ID)...)
}
