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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceMssqlResource{}
var _ resource.ResourceWithImportState = &SourceMssqlResource{}

func NewSourceMssqlResource() resource.Resource {
	return &SourceMssqlResource{}
}

// SourceMssqlResource defines the resource implementation.
type SourceMssqlResource struct {
	client *sdk.SDK
}

// SourceMssqlResourceModel describes the resource data model.
type SourceMssqlResourceModel struct {
	Configuration tfTypes.SourceMssql `tfsdk:"configuration"`
	DefinitionID  types.String        `tfsdk:"definition_id"`
	Name          types.String        `tfsdk:"name"`
	SecretID      types.String        `tfsdk:"secret_id"`
	SourceID      types.String        `tfsdk:"source_id"`
	SourceType    types.String        `tfsdk:"source_type"`
	WorkspaceID   types.String        `tfsdk:"workspace_id"`
}

func (r *SourceMssqlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_mssql"
}

func (r *SourceMssqlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceMssql Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Required:    true,
						Description: `The name of the database.`,
					},
					"host": schema.StringAttribute{
						Required:    true,
						Description: `The hostname of the database.`,
					},
					"jdbc_url_params": schema.StringAttribute{
						Optional:    true,
						Description: `Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).`,
					},
					"password": schema.StringAttribute{
						Required:    true,
						Sensitive:   true,
						Description: `The password associated with the username.`,
					},
					"port": schema.Int64Attribute{
						Required:    true,
						Description: `The port of the database.`,
						Validators: []validator.Int64{
							int64validator.AtMost(65536),
						},
					},
					"replication_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"read_changes_using_change_data_capture_cdc": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"initial_load_timeout_hours": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(8),
										Description: `The amount of time an initial load is allowed to continue for before catching up on CDC logs. Default: 8`,
									},
									"initial_waiting_seconds": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(300),
										Description: `The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 3600 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">initial waiting time</a>. Default: 300`,
									},
									"invalid_cdc_cursor_position_behavior": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("Fail sync"),
										Description: `Determines whether Airbyte should fail or re-sync data in case of an stale/invalid cursor value into the WAL. If 'Fail sync' is chosen, a user will have to manually reset the connection before being able to continue syncing data. If 'Re-sync data' is chosen, Airbyte will automatically trigger a refresh but could lead to higher cloud costs and data loss. Default: "Fail sync"; must be one of ["Fail sync", "Re-sync data"]`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Fail sync",
												"Re-sync data",
											),
										},
									},
									"queue_size": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(10000),
										Description: `The size of the internal queue. This may interfere with memory consumption and efficiency of the connector, please be careful. Default: 10000`,
									},
								},
								Description: `<i>Recommended</i> - Incrementally reads new inserts, updates, and deletes using the SQL Server's <a href="https://docs.airbyte.com/integrations/sources/mssql/#change-data-capture-cdc">change data capture feature</a>. This must be enabled on your database.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("scan_changes_with_user_defined_cursor"),
									}...),
								},
							},
							"scan_changes_with_user_defined_cursor": schema.SingleNestedAttribute{
								Optional:    true,
								Description: `Incrementally detects new inserts and updates using the <a href="https://docs.airbyte.com/understanding-airbyte/connections/incremental-append/#user-defined-cursor">cursor column</a> chosen when configuring a connection (e.g. created_at, updated_at).`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("read_changes_using_change_data_capture_cdc"),
									}...),
								},
							},
						},
						Description: `Configures how data is extracted from the database.`,
					},
					"schemas": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
						Description: `The list of schemas to sync from. Defaults to user. Case sensitive.`,
						Validators: []validator.List{
							listvalidator.UniqueValues(),
						},
					},
					"ssl_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"encrypted_trust_server_certificate": schema.SingleNestedAttribute{
								Optional:    true,
								Description: `Use the certificate provided by the server without verification. (For testing purposes only!)`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("encrypted_verify_certificate"),
										path.MatchRelative().AtParent().AtName("unencrypted"),
									}...),
								},
							},
							"encrypted_verify_certificate": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"certificate": schema.StringAttribute{
										Optional:    true,
										Sensitive:   true,
										Description: `certificate of the server, or of the CA that signed the server certificate`,
									},
									"host_name_in_certificate": schema.StringAttribute{
										Optional:    true,
										Description: `Specifies the host name of the server. The value of this property must match the subject property of the certificate.`,
									},
								},
								Description: `Verify and use the certificate provided by the server.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("encrypted_trust_server_certificate"),
										path.MatchRelative().AtParent().AtName("unencrypted"),
									}...),
								},
							},
							"unencrypted": schema.SingleNestedAttribute{
								Optional:    true,
								Description: `Data transfer will not be encrypted.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("encrypted_trust_server_certificate"),
										path.MatchRelative().AtParent().AtName("encrypted_verify_certificate"),
									}...),
								},
							},
						},
						Description: `The encryption method which is used when communicating with the database.`,
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"no_tunnel": schema.SingleNestedAttribute{
								Optional: true,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("password_authentication"),
										path.MatchRelative().AtParent().AtName("ssh_key_authentication"),
									}...),
								},
							},
							"password_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Required:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(22),
										Description: `Port on the proxy/jump server that accepts inbound ssh connections. Default: 22`,
										Validators: []validator.Int64{
											int64validator.AtMost(65536),
										},
									},
									"tunnel_user": schema.StringAttribute{
										Required:    true,
										Description: `OS-level username for logging into the jump server host`,
									},
									"tunnel_user_password": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("no_tunnel"),
										path.MatchRelative().AtParent().AtName("ssh_key_authentication"),
									}...),
								},
							},
							"ssh_key_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )`,
									},
									"tunnel_host": schema.StringAttribute{
										Required:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(22),
										Description: `Port on the proxy/jump server that accepts inbound ssh connections. Default: 22`,
										Validators: []validator.Int64{
											int64validator.AtMost(65536),
										},
									},
									"tunnel_user": schema.StringAttribute{
										Required:    true,
										Description: `OS-level username for logging into the jump server host.`,
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("no_tunnel"),
										path.MatchRelative().AtParent().AtName("password_authentication"),
									}...),
								},
							},
						},
						Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
					},
					"username": schema.StringAttribute{
						Required:    true,
						Description: `The username which is used to access the database.`,
					},
				},
			},
			"definition_id": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Name of the source e.g. dev-mysql-instance.`,
			},
			"secret_id": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.`,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"source_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
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

func (r *SourceMssqlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceMssqlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceMssqlResourceModel
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

	request := data.ToSharedSourceMssqlCreateRequest()
	res, err := r.client.Sources.CreateSourceMssql(ctx, request)
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
	if !(res.SourceResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var sourceID string
	sourceID = data.SourceID.ValueString()

	request1 := operations.GetSourceMssqlRequest{
		SourceID: sourceID,
	}
	res1, err := r.client.Sources.GetSourceMssql(ctx, request1)
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
	if !(res1.SourceResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res1.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceMssqlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceMssqlResourceModel
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

	var sourceID string
	sourceID = data.SourceID.ValueString()

	request := operations.GetSourceMssqlRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceMssql(ctx, request)
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
	if !(res.SourceResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceMssqlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceMssqlResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceMssqlPutRequest := data.ToSharedSourceMssqlPutRequest()
	var sourceID string
	sourceID = data.SourceID.ValueString()

	request := operations.PutSourceMssqlRequest{
		SourceMssqlPutRequest: sourceMssqlPutRequest,
		SourceID:              sourceID,
	}
	res, err := r.client.Sources.PutSourceMssql(ctx, request)
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
	var sourceId1 string
	sourceId1 = data.SourceID.ValueString()

	request1 := operations.GetSourceMssqlRequest{
		SourceID: sourceId1,
	}
	res1, err := r.client.Sources.GetSourceMssql(ctx, request1)
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
	if !(res1.SourceResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res1.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceMssqlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceMssqlResourceModel
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

	var sourceID string
	sourceID = data.SourceID.ValueString()

	request := operations.DeleteSourceMssqlRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceMssql(ctx, request)
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

func (r *SourceMssqlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("source_id"), req.ID)...)
}
