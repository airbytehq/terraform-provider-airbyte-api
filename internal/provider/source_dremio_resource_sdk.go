// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceDremioResourceModel) ToSharedSourceDremioCreateRequest() *shared.SourceDremioCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	baseURL := new(string)
	if !r.Configuration.BaseURL.IsUnknown() && !r.Configuration.BaseURL.IsNull() {
		*baseURL = r.Configuration.BaseURL.ValueString()
	} else {
		baseURL = nil
	}
	configuration := shared.SourceDremio{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDremioCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDremioResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceDremioResourceModel) ToSharedSourceDremioPutRequest() *shared.SourceDremioPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	baseURL := new(string)
	if !r.Configuration.BaseURL.IsUnknown() && !r.Configuration.BaseURL.IsNull() {
		*baseURL = r.Configuration.BaseURL.ValueString()
	} else {
		baseURL = nil
	}
	configuration := shared.SourceDremioUpdate{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDremioPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
