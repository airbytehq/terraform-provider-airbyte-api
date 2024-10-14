// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceKyveResourceModel) ToSharedSourceKyveCreateRequest() *shared.SourceKyveCreateRequest {
	var poolIds string
	poolIds = r.Configuration.PoolIds.ValueString()

	var startIds string
	startIds = r.Configuration.StartIds.ValueString()

	urlBase := new(string)
	if !r.Configuration.URLBase.IsUnknown() && !r.Configuration.URLBase.IsNull() {
		*urlBase = r.Configuration.URLBase.ValueString()
	} else {
		urlBase = nil
	}
	configuration := shared.SourceKyve{
		PoolIds:  poolIds,
		StartIds: startIds,
		URLBase:  urlBase,
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

	out := shared.SourceKyveCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceKyveResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceKyveResourceModel) ToSharedSourceKyvePutRequest() *shared.SourceKyvePutRequest {
	var poolIds string
	poolIds = r.Configuration.PoolIds.ValueString()

	var startIds string
	startIds = r.Configuration.StartIds.ValueString()

	urlBase := new(string)
	if !r.Configuration.URLBase.IsUnknown() && !r.Configuration.URLBase.IsNull() {
		*urlBase = r.Configuration.URLBase.ValueString()
	} else {
		urlBase = nil
	}
	configuration := shared.SourceKyveUpdate{
		PoolIds:  poolIds,
		StartIds: startIds,
		URLBase:  urlBase,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceKyvePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
