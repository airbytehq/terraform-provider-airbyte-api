// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceBambooHrResourceModel) ToSharedSourceBambooHrCreateRequest() *shared.SourceBambooHrCreateRequest {
	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	customReportsFields := new(string)
	if !r.Configuration.CustomReportsFields.IsUnknown() && !r.Configuration.CustomReportsFields.IsNull() {
		*customReportsFields = r.Configuration.CustomReportsFields.ValueString()
	} else {
		customReportsFields = nil
	}
	customReportsIncludeDefaultFields := new(bool)
	if !r.Configuration.CustomReportsIncludeDefaultFields.IsUnknown() && !r.Configuration.CustomReportsIncludeDefaultFields.IsNull() {
		*customReportsIncludeDefaultFields = r.Configuration.CustomReportsIncludeDefaultFields.ValueBool()
	} else {
		customReportsIncludeDefaultFields = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var subdomain string
	subdomain = r.Configuration.Subdomain.ValueString()

	configuration := shared.SourceBambooHr{
		APIKey:                            apiKey,
		CustomReportsFields:               customReportsFields,
		CustomReportsIncludeDefaultFields: customReportsIncludeDefaultFields,
		StartDate:                         startDate,
		Subdomain:                         subdomain,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	var name string
	name = r.Name.ValueString()

	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceBambooHrCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBambooHrResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceBambooHrResourceModel) ToSharedSourceBambooHrPutRequest() *shared.SourceBambooHrPutRequest {
	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	customReportsFields := new(string)
	if !r.Configuration.CustomReportsFields.IsUnknown() && !r.Configuration.CustomReportsFields.IsNull() {
		*customReportsFields = r.Configuration.CustomReportsFields.ValueString()
	} else {
		customReportsFields = nil
	}
	customReportsIncludeDefaultFields := new(bool)
	if !r.Configuration.CustomReportsIncludeDefaultFields.IsUnknown() && !r.Configuration.CustomReportsIncludeDefaultFields.IsNull() {
		*customReportsIncludeDefaultFields = r.Configuration.CustomReportsIncludeDefaultFields.ValueBool()
	} else {
		customReportsIncludeDefaultFields = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var subdomain string
	subdomain = r.Configuration.Subdomain.ValueString()

	configuration := shared.SourceBambooHrUpdate{
		APIKey:                            apiKey,
		CustomReportsFields:               customReportsFields,
		CustomReportsIncludeDefaultFields: customReportsIncludeDefaultFields,
		StartDate:                         startDate,
		Subdomain:                         subdomain,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceBambooHrPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
