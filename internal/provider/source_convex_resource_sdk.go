// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceConvexResourceModel) ToSharedSourceConvexCreateRequest() *shared.SourceConvexCreateRequest {
	var accessKey string
	accessKey = r.Configuration.AccessKey.ValueString()

	var deploymentURL string
	deploymentURL = r.Configuration.DeploymentURL.ValueString()

	configuration := shared.SourceConvex{
		AccessKey:     accessKey,
		DeploymentURL: deploymentURL,
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

	out := shared.SourceConvexCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceConvexResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceConvexResourceModel) ToSharedSourceConvexPutRequest() *shared.SourceConvexPutRequest {
	var accessKey string
	accessKey = r.Configuration.AccessKey.ValueString()

	var deploymentURL string
	deploymentURL = r.Configuration.DeploymentURL.ValueString()

	configuration := shared.SourceConvexUpdate{
		AccessKey:     accessKey,
		DeploymentURL: deploymentURL,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceConvexPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
