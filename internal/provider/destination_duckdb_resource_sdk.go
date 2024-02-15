// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationDuckdbResourceModel) ToSharedDestinationDuckdbCreateRequest() *shared.DestinationDuckdbCreateRequest {
	destinationPath := r.Configuration.DestinationPath.ValueString()
	motherduckAPIKey := new(string)
	if !r.Configuration.MotherduckAPIKey.IsUnknown() && !r.Configuration.MotherduckAPIKey.IsNull() {
		*motherduckAPIKey = r.Configuration.MotherduckAPIKey.ValueString()
	} else {
		motherduckAPIKey = nil
	}
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	configuration := shared.DestinationDuckdb{
		DestinationPath:  destinationPath,
		MotherduckAPIKey: motherduckAPIKey,
		Schema:           schema,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDuckdbCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDuckdbResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationDuckdbResourceModel) ToSharedDestinationDuckdbPutRequest() *shared.DestinationDuckdbPutRequest {
	destinationPath := r.Configuration.DestinationPath.ValueString()
	motherduckAPIKey := new(string)
	if !r.Configuration.MotherduckAPIKey.IsUnknown() && !r.Configuration.MotherduckAPIKey.IsNull() {
		*motherduckAPIKey = r.Configuration.MotherduckAPIKey.ValueString()
	} else {
		motherduckAPIKey = nil
	}
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	configuration := shared.DestinationDuckdbUpdate{
		DestinationPath:  destinationPath,
		MotherduckAPIKey: motherduckAPIKey,
		Schema:           schema,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDuckdbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
