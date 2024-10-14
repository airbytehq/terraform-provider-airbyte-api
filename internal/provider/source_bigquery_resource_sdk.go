// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceBigqueryResourceModel) ToSharedSourceBigqueryCreateRequest() *shared.SourceBigqueryCreateRequest {
	var credentialsJSON string
	credentialsJSON = r.Configuration.CredentialsJSON.ValueString()

	datasetID := new(string)
	if !r.Configuration.DatasetID.IsUnknown() && !r.Configuration.DatasetID.IsNull() {
		*datasetID = r.Configuration.DatasetID.ValueString()
	} else {
		datasetID = nil
	}
	var projectID string
	projectID = r.Configuration.ProjectID.ValueString()

	configuration := shared.SourceBigquery{
		CredentialsJSON: credentialsJSON,
		DatasetID:       datasetID,
		ProjectID:       projectID,
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

	out := shared.SourceBigqueryCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBigqueryResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceBigqueryResourceModel) ToSharedSourceBigqueryPutRequest() *shared.SourceBigqueryPutRequest {
	var credentialsJSON string
	credentialsJSON = r.Configuration.CredentialsJSON.ValueString()

	datasetID := new(string)
	if !r.Configuration.DatasetID.IsUnknown() && !r.Configuration.DatasetID.IsNull() {
		*datasetID = r.Configuration.DatasetID.ValueString()
	} else {
		datasetID = nil
	}
	var projectID string
	projectID = r.Configuration.ProjectID.ValueString()

	configuration := shared.SourceBigqueryUpdate{
		CredentialsJSON: credentialsJSON,
		DatasetID:       datasetID,
		ProjectID:       projectID,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceBigqueryPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
