// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceYotpoResourceModel) ToSharedSourceYotpoCreateRequest() *shared.SourceYotpoCreateRequest {
	var accessToken string
	accessToken = r.Configuration.AccessToken.ValueString()

	var appKey string
	appKey = r.Configuration.AppKey.ValueString()

	email := new(string)
	if !r.Configuration.Email.IsUnknown() && !r.Configuration.Email.IsNull() {
		*email = r.Configuration.Email.ValueString()
	} else {
		email = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceYotpo{
		AccessToken: accessToken,
		AppKey:      appKey,
		Email:       email,
		StartDate:   startDate,
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

	out := shared.SourceYotpoCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceYotpoResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceYotpoResourceModel) ToSharedSourceYotpoPutRequest() *shared.SourceYotpoPutRequest {
	var accessToken string
	accessToken = r.Configuration.AccessToken.ValueString()

	var appKey string
	appKey = r.Configuration.AppKey.ValueString()

	email := new(string)
	if !r.Configuration.Email.IsUnknown() && !r.Configuration.Email.IsNull() {
		*email = r.Configuration.Email.ValueString()
	} else {
		email = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceYotpoUpdate{
		AccessToken: accessToken,
		AppKey:      appKey,
		Email:       email,
		StartDate:   startDate,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceYotpoPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
