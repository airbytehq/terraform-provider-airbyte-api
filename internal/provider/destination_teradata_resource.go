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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationTeradataResource{}
var _ resource.ResourceWithImportState = &DestinationTeradataResource{}

func NewDestinationTeradataResource() resource.Resource {
	return &DestinationTeradataResource{}
}

// DestinationTeradataResource defines the resource implementation.
type DestinationTeradataResource struct {
	client *sdk.SDK
}

// DestinationTeradataResourceModel describes the resource data model.
type DestinationTeradataResourceModel struct {
	Configuration   tfTypes.DestinationTeradata `tfsdk:"configuration"`
	DefinitionID    types.String                `tfsdk:"definition_id"`
	DestinationID   types.String                `tfsdk:"destination_id"`
	DestinationType types.String                `tfsdk:"destination_type"`
	Name            types.String                `tfsdk:"name"`
	WorkspaceID     types.String                `tfsdk:"workspace_id"`
}

func (r *DestinationTeradataResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_teradata"
}

func (r *DestinationTeradataResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationTeradata Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Required:    true,
						Description: `Hostname of the database.`,
					},
					"jdbc_url_params": schema.StringAttribute{
						Optional:    true,
						Description: `Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).`,
					},
					"password": schema.StringAttribute{
						Optional:    true,
						Sensitive:   true,
						Description: `Password associated with the username.`,
					},
					"schema": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("airbyte_td"),
						Description: `The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public". Default: "airbyte_td"`,
					},
					"ssl": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Encrypt data using SSL. When activating SSL, please select one of the connection modes. Default: false`,
					},
					"ssl_mode": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"allow": schema.SingleNestedAttribute{
								Optional:    true,
								Description: `Allow SSL mode.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("disable"),
										path.MatchRelative().AtParent().AtName("prefer"),
										path.MatchRelative().AtParent().AtName("require"),
										path.MatchRelative().AtParent().AtName("verify_ca"),
										path.MatchRelative().AtParent().AtName("verify_full"),
									}...),
								},
							},
							"disable": schema.SingleNestedAttribute{
								Optional:    true,
								Description: `Disable SSL.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("prefer"),
										path.MatchRelative().AtParent().AtName("require"),
										path.MatchRelative().AtParent().AtName("verify_ca"),
										path.MatchRelative().AtParent().AtName("verify_full"),
									}...),
								},
							},
							"prefer": schema.SingleNestedAttribute{
								Optional:    true,
								Description: `Prefer SSL mode.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("disable"),
										path.MatchRelative().AtParent().AtName("require"),
										path.MatchRelative().AtParent().AtName("verify_ca"),
										path.MatchRelative().AtParent().AtName("verify_full"),
									}...),
								},
							},
							"require": schema.SingleNestedAttribute{
								Optional:    true,
								Description: `Require SSL mode.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("disable"),
										path.MatchRelative().AtParent().AtName("prefer"),
										path.MatchRelative().AtParent().AtName("verify_ca"),
										path.MatchRelative().AtParent().AtName("verify_full"),
									}...),
								},
							},
							"verify_ca": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssl_ca_certificate": schema.StringAttribute{
										Required:  true,
										Sensitive: true,
										MarkdownDescription: `Specifies the file name of a PEM file that contains Certificate Authority (CA) certificates for use with SSLMODE=verify-ca.` + "\n" +
											` See more information - <a href="https://teradata-docs.s3.amazonaws.com/doc/connectivity/jdbc/reference/current/jdbcug_chapter_2.html#URL_SSLCA"> in the docs</a>.`,
									},
								},
								Description: `Verify-ca SSL mode.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("disable"),
										path.MatchRelative().AtParent().AtName("prefer"),
										path.MatchRelative().AtParent().AtName("require"),
										path.MatchRelative().AtParent().AtName("verify_full"),
									}...),
								},
							},
							"verify_full": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssl_ca_certificate": schema.StringAttribute{
										Required:  true,
										Sensitive: true,
										MarkdownDescription: `Specifies the file name of a PEM file that contains Certificate Authority (CA) certificates for use with SSLMODE=verify-full.` + "\n" +
											` See more information - <a href="https://teradata-docs.s3.amazonaws.com/doc/connectivity/jdbc/reference/current/jdbcug_chapter_2.html#URL_SSLCA"> in the docs</a>.`,
									},
								},
								Description: `Verify-full SSL mode.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("disable"),
										path.MatchRelative().AtParent().AtName("prefer"),
										path.MatchRelative().AtParent().AtName("require"),
										path.MatchRelative().AtParent().AtName("verify_ca"),
									}...),
								},
							},
						},
						MarkdownDescription: `SSL connection modes. ` + "\n" +
							` <b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database` + "\n" +
							` <b>allow</b> - Chose this mode to enable encryption only when required by the destination database` + "\n" +
							` <b>prefer</b> - Chose this mode to allow unencrypted connection only if the destination database does not support encryption` + "\n" +
							` <b>require</b> - Chose this mode to always require encryption. If the destination database server does not support encryption, connection will fail` + "\n" +
							`  <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the destination database server has a valid SSL certificate` + "\n" +
							`  <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the destination database server` + "\n" +
							` See more information - <a href="https://teradata-docs.s3.amazonaws.com/doc/connectivity/jdbc/reference/current/jdbcug_chapter_2.html#URL_SSLMODE"> in the docs</a>.`,
					},
					"username": schema.StringAttribute{
						Required:    true,
						Description: `Username to use to access the database.`,
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

func (r *DestinationTeradataResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationTeradataResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationTeradataResourceModel
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

	request := data.ToSharedDestinationTeradataCreateRequest()
	res, err := r.client.Destinations.CreateDestinationTeradata(ctx, request)
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

	request1 := operations.GetDestinationTeradataRequest{
		DestinationID: destinationID,
	}
	res1, err := r.client.Destinations.GetDestinationTeradata(ctx, request1)
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

func (r *DestinationTeradataResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationTeradataResourceModel
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

	request := operations.GetDestinationTeradataRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationTeradata(ctx, request)
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

func (r *DestinationTeradataResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationTeradataResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationTeradataPutRequest := data.ToSharedDestinationTeradataPutRequest()
	var destinationID string
	destinationID = data.DestinationID.ValueString()

	request := operations.PutDestinationTeradataRequest{
		DestinationTeradataPutRequest: destinationTeradataPutRequest,
		DestinationID:                 destinationID,
	}
	res, err := r.client.Destinations.PutDestinationTeradata(ctx, request)
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

	request1 := operations.GetDestinationTeradataRequest{
		DestinationID: destinationId1,
	}
	res1, err := r.client.Destinations.GetDestinationTeradata(ctx, request1)
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

func (r *DestinationTeradataResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationTeradataResourceModel
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

	request := operations.DeleteDestinationTeradataRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationTeradata(ctx, request)
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

func (r *DestinationTeradataResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("destination_id"), req.ID)...)
}
