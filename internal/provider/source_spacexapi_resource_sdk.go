// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSpacexAPIResourceModel) ToSharedSourceSpacexAPICreateRequest() *shared.SourceSpacexAPICreateRequest {
	id := new(string)
	if !r.Configuration.ID.IsUnknown() && !r.Configuration.ID.IsNull() {
		*id = r.Configuration.ID.ValueString()
	} else {
		id = nil
	}
	optionsVar := new(string)
	if !r.Configuration.Options.IsUnknown() && !r.Configuration.Options.IsNull() {
		*optionsVar = r.Configuration.Options.ValueString()
	} else {
		optionsVar = nil
	}
	configuration := shared.SourceSpacexAPI{
		ID:      id,
		Options: optionsVar,
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

	out := shared.SourceSpacexAPICreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSpacexAPIResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceSpacexAPIResourceModel) ToSharedSourceSpacexAPIPutRequest() *shared.SourceSpacexAPIPutRequest {
	id := new(string)
	if !r.Configuration.ID.IsUnknown() && !r.Configuration.ID.IsNull() {
		*id = r.Configuration.ID.ValueString()
	} else {
		id = nil
	}
	optionsVar := new(string)
	if !r.Configuration.Options.IsUnknown() && !r.Configuration.Options.IsNull() {
		*optionsVar = r.Configuration.Options.ValueString()
	} else {
		optionsVar = nil
	}
	configuration := shared.SourceSpacexAPIUpdate{
		ID:      id,
		Options: optionsVar,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceSpacexAPIPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
