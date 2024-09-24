// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceHardcodedRecordsResourceModel) ToSharedSourceHardcodedRecordsCreateRequest() *shared.SourceHardcodedRecordsCreateRequest {
	count := new(int64)
	if !r.Configuration.Count.IsUnknown() && !r.Configuration.Count.IsNull() {
		*count = r.Configuration.Count.ValueInt64()
	} else {
		count = nil
	}
	configuration := shared.SourceHardcodedRecords{
		Count: count,
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
	out := shared.SourceHardcodedRecordsCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceHardcodedRecordsResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceHardcodedRecordsResourceModel) ToSharedSourceHardcodedRecordsPutRequest() *shared.SourceHardcodedRecordsPutRequest {
	count := new(int64)
	if !r.Configuration.Count.IsUnknown() && !r.Configuration.Count.IsNull() {
		*count = r.Configuration.Count.ValueInt64()
	} else {
		count = nil
	}
	configuration := shared.SourceHardcodedRecordsUpdate{
		Count: count,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceHardcodedRecordsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}