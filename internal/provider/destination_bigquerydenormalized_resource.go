// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	speakeasy_int64planmodifier "airbyte/internal/planmodifiers/int64planmodifier"
	speakeasy_objectplanmodifier "airbyte/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "airbyte/internal/planmodifiers/stringplanmodifier"
	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationBigqueryDenormalizedResource{}
var _ resource.ResourceWithImportState = &DestinationBigqueryDenormalizedResource{}

func NewDestinationBigqueryDenormalizedResource() resource.Resource {
	return &DestinationBigqueryDenormalizedResource{}
}

// DestinationBigqueryDenormalizedResource defines the resource implementation.
type DestinationBigqueryDenormalizedResource struct {
	client *sdk.SDK
}

// DestinationBigqueryDenormalizedResourceModel describes the resource data model.
type DestinationBigqueryDenormalizedResourceModel struct {
	Configuration   DestinationBigqueryDenormalized `tfsdk:"configuration"`
	DestinationID   types.String                    `tfsdk:"destination_id"`
	DestinationType types.String                    `tfsdk:"destination_type"`
	Name            types.String                    `tfsdk:"name"`
	WorkspaceID     types.String                    `tfsdk:"workspace_id"`
}

func (r *DestinationBigqueryDenormalizedResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_bigquery_denormalized"
}

func (r *DestinationBigqueryDenormalizedResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationBigqueryDenormalized Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"big_query_client_buffer_size_mb": schema.Int64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Int64{
							speakeasy_int64planmodifier.TrustRead(),
						},
						Optional: true,
					},
					"credentials_json": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.TrustRead(),
						},
						Optional: true,
					},
					"dataset_id": schema.StringAttribute{
						Required: true,
					},
					"dataset_location": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.TrustRead(),
						},
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"US",
								"EU",
								"asia-east1",
								"asia-east2",
								"asia-northeast1",
								"asia-northeast2",
								"asia-northeast3",
								"asia-south1",
								"asia-south2",
								"asia-southeast1",
								"asia-southeast2",
								"australia-southeast1",
								"australia-southeast2",
								"europe-central1",
								"europe-central2",
								"europe-north1",
								"europe-southwest1",
								"europe-west1",
								"europe-west2",
								"europe-west3",
								"europe-west4",
								"europe-west6",
								"europe-west7",
								"europe-west8",
								"europe-west9",
								"me-west1",
								"northamerica-northeast1",
								"northamerica-northeast2",
								"southamerica-east1",
								"southamerica-west1",
								"us-central1",
								"us-east1",
								"us-east2",
								"us-east3",
								"us-east4",
								"us-east5",
								"us-west1",
								"us-west2",
								"us-west3",
								"us-west4",
							),
						},
						Description: `The location of the dataset. Warning: Changes made after creation will not be applied. The default "US" value is used if not set explicitly. Read more <a href="https://cloud.google.com/bigquery/docs/locations">here</a>.`,
					},
					"destination_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"bigquery-denormalized",
							),
						},
					},
					"loading_method": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							speakeasy_objectplanmodifier.TrustRead(),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"destination_bigquery_denormalized_loading_method_gcs_staging": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									speakeasy_objectplanmodifier.TrustRead(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"credential": schema.SingleNestedAttribute{
										Required: true,
										Attributes: map[string]schema.Attribute{
											"destination_bigquery_denormalized_loading_method_gcs_staging_credential_hmac_key": schema.SingleNestedAttribute{
												Computed: true,
												PlanModifiers: []planmodifier.Object{
													speakeasy_objectplanmodifier.TrustRead(),
												},
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"credential_type": schema.StringAttribute{
														Required: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"HMAC_KEY",
															),
														},
													},
													"hmac_key_access_id": schema.StringAttribute{
														Required: true,
													},
													"hmac_key_secret": schema.StringAttribute{
														Required: true,
													},
												},
												Description: `An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"file_buffer_count": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											speakeasy_int64planmodifier.TrustRead(),
										},
										Optional: true,
									},
									"gcs_bucket_name": schema.StringAttribute{
										Required: true,
									},
									"gcs_bucket_path": schema.StringAttribute{
										Required: true,
									},
									"keep_files_in_gcs_bucket": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.TrustRead(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Delete all tmp files from GCS",
												"Keep all tmp files in GCS",
											),
										},
										Description: `This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.`,
									},
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"GCS Staging",
											),
										},
									},
								},
								Description: `Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.`,
							},
							"destination_bigquery_denormalized_loading_method_standard_inserts": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									speakeasy_objectplanmodifier.TrustRead(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
									},
								},
								Description: `Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.`,
							},
							"destination_bigquery_denormalized_update_loading_method_gcs_staging": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									speakeasy_objectplanmodifier.TrustRead(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"credential": schema.SingleNestedAttribute{
										Required: true,
										Attributes: map[string]schema.Attribute{
											"destination_bigquery_denormalized_update_loading_method_gcs_staging_credential_hmac_key": schema.SingleNestedAttribute{
												Computed: true,
												PlanModifiers: []planmodifier.Object{
													speakeasy_objectplanmodifier.TrustRead(),
												},
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"credential_type": schema.StringAttribute{
														Required: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"HMAC_KEY",
															),
														},
													},
													"hmac_key_access_id": schema.StringAttribute{
														Required: true,
													},
													"hmac_key_secret": schema.StringAttribute{
														Required: true,
													},
												},
												Description: `An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"file_buffer_count": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											speakeasy_int64planmodifier.TrustRead(),
										},
										Optional: true,
									},
									"gcs_bucket_name": schema.StringAttribute{
										Required: true,
									},
									"gcs_bucket_path": schema.StringAttribute{
										Required: true,
									},
									"keep_files_in_gcs_bucket": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.TrustRead(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Delete all tmp files from GCS",
												"Keep all tmp files in GCS",
											),
										},
										Description: `This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.`,
									},
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"GCS Staging",
											),
										},
									},
								},
								Description: `Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.`,
							},
							"destination_bigquery_denormalized_update_loading_method_standard_inserts": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									speakeasy_objectplanmodifier.TrustRead(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
									},
								},
								Description: `Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"project_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.TrustRead(),
				},
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.TrustRead(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *DestinationBigqueryDenormalizedResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationBigqueryDenormalizedResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationBigqueryDenormalizedResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	request := *data.ToCreateSDKType()
	res, err := r.client.Destinations.CreateDestinationBigqueryDenormalized(ctx, request)
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
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationBigqueryDenormalizedResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationBigqueryDenormalizedResourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationBigqueryDenormalizedRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationBigqueryDenormalized(ctx, request)
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
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationBigqueryDenormalizedResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationBigqueryDenormalizedResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationBigqueryDenormalizedPutRequest := data.ToUpdateSDKType()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationBigqueryDenormalizedRequest{
		DestinationBigqueryDenormalizedPutRequest: destinationBigqueryDenormalizedPutRequest,
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.PutDestinationBigqueryDenormalized(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	destinationId1 := data.DestinationID.ValueString()
	getRequest := operations.GetDestinationBigqueryDenormalizedRequest{
		DestinationID: destinationId1,
	}
	getResponse, err := r.client.Destinations.GetDestinationBigqueryDenormalized(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if getResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", getResponse))
		return
	}
	if getResponse.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", getResponse.StatusCode), debugResponse(getResponse.RawResponse))
		return
	}
	if getResponse.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationBigqueryDenormalizedResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationBigqueryDenormalizedResourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.DeleteDestinationBigqueryDenormalizedRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationBigqueryDenormalized(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *DestinationBigqueryDenormalizedResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
