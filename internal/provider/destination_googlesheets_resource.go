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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationGoogleSheetsResource{}
var _ resource.ResourceWithImportState = &DestinationGoogleSheetsResource{}

func NewDestinationGoogleSheetsResource() resource.Resource {
	return &DestinationGoogleSheetsResource{}
}

// DestinationGoogleSheetsResource defines the resource implementation.
type DestinationGoogleSheetsResource struct {
	client *sdk.SDK
}

// DestinationGoogleSheetsResourceModel describes the resource data model.
type DestinationGoogleSheetsResourceModel struct {
	Configuration   tfTypes.DestinationGoogleSheets `tfsdk:"configuration"`
	DefinitionID    types.String                    `tfsdk:"definition_id"`
	DestinationID   types.String                    `tfsdk:"destination_id"`
	DestinationType types.String                    `tfsdk:"destination_type"`
	Name            types.String                    `tfsdk:"name"`
	WorkspaceID     types.String                    `tfsdk:"workspace_id"`
}

func (r *DestinationGoogleSheetsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_google_sheets"
}

func (r *DestinationGoogleSheetsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationGoogleSheets Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"client_id": schema.StringAttribute{
								Required:    true,
								Sensitive:   true,
								Description: `The Client ID of your Google Sheets developer application.`,
							},
							"client_secret": schema.StringAttribute{
								Required:    true,
								Sensitive:   true,
								Description: `The Client Secret of your Google Sheets developer application.`,
							},
							"refresh_token": schema.StringAttribute{
								Required:    true,
								Sensitive:   true,
								Description: `The token for obtaining new access token.`,
							},
						},
						Description: `Google API Credentials for connecting to Google Sheets and Google Drive APIs`,
					},
					"spreadsheet_id": schema.StringAttribute{
						Required:    true,
						Description: `The link to your spreadsheet. See <a href='https://docs.airbyte.com/integrations/destinations/google-sheets#sheetlink'>this guide</a> for more details.`,
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

func (r *DestinationGoogleSheetsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationGoogleSheetsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationGoogleSheetsResourceModel
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

	request := data.ToSharedDestinationGoogleSheetsCreateRequest()
	res, err := r.client.Destinations.CreateDestinationGoogleSheets(ctx, request)
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

	request1 := operations.GetDestinationGoogleSheetsRequest{
		DestinationID: destinationID,
	}
	res1, err := r.client.Destinations.GetDestinationGoogleSheets(ctx, request1)
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

func (r *DestinationGoogleSheetsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationGoogleSheetsResourceModel
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

	request := operations.GetDestinationGoogleSheetsRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationGoogleSheets(ctx, request)
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

func (r *DestinationGoogleSheetsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationGoogleSheetsResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationGoogleSheetsPutRequest := data.ToSharedDestinationGoogleSheetsPutRequest()
	var destinationID string
	destinationID = data.DestinationID.ValueString()

	request := operations.PutDestinationGoogleSheetsRequest{
		DestinationGoogleSheetsPutRequest: destinationGoogleSheetsPutRequest,
		DestinationID:                     destinationID,
	}
	res, err := r.client.Destinations.PutDestinationGoogleSheets(ctx, request)
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

	request1 := operations.GetDestinationGoogleSheetsRequest{
		DestinationID: destinationId1,
	}
	res1, err := r.client.Destinations.GetDestinationGoogleSheets(ctx, request1)
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

func (r *DestinationGoogleSheetsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationGoogleSheetsResourceModel
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

	request := operations.DeleteDestinationGoogleSheetsRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationGoogleSheets(ctx, request)
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

func (r *DestinationGoogleSheetsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("destination_id"), req.ID)...)
}
