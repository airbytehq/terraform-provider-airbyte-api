// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceNorthpassLmsResourceModel) ToSharedSourceNorthpassLmsCreateRequest() *shared.SourceNorthpassLmsCreateRequest {
	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	configuration := shared.SourceNorthpassLms{
		APIKey: apiKey,
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

	out := shared.SourceNorthpassLmsCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceNorthpassLmsResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceNorthpassLmsResourceModel) ToSharedSourceNorthpassLmsPutRequest() *shared.SourceNorthpassLmsPutRequest {
	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	configuration := shared.SourceNorthpassLmsUpdate{
		APIKey: apiKey,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceNorthpassLmsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
