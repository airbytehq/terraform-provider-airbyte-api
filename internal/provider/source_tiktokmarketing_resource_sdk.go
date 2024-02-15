// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	customTypes "github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceTiktokMarketingResourceModel) ToSharedSourceTiktokMarketingCreateRequest() *shared.SourceTiktokMarketingCreateRequest {
	attributionWindow := new(int64)
	if !r.Configuration.AttributionWindow.IsUnknown() && !r.Configuration.AttributionWindow.IsNull() {
		*attributionWindow = r.Configuration.AttributionWindow.ValueInt64()
	} else {
		attributionWindow = nil
	}
	var credentials *shared.SourceTiktokMarketingAuthenticationMethod
	if r.Configuration.Credentials != nil {
		var sourceTiktokMarketingOAuth20 *shared.SourceTiktokMarketingOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			advertiserID := new(string)
			if !r.Configuration.Credentials.OAuth20.AdvertiserID.IsUnknown() && !r.Configuration.Credentials.OAuth20.AdvertiserID.IsNull() {
				*advertiserID = r.Configuration.Credentials.OAuth20.AdvertiserID.ValueString()
			} else {
				advertiserID = nil
			}
			appID := r.Configuration.Credentials.OAuth20.AppID.ValueString()
			secret := r.Configuration.Credentials.OAuth20.Secret.ValueString()
			sourceTiktokMarketingOAuth20 = &shared.SourceTiktokMarketingOAuth20{
				AccessToken:  accessToken,
				AdvertiserID: advertiserID,
				AppID:        appID,
				Secret:       secret,
			}
		}
		if sourceTiktokMarketingOAuth20 != nil {
			credentials = &shared.SourceTiktokMarketingAuthenticationMethod{
				SourceTiktokMarketingOAuth20: sourceTiktokMarketingOAuth20,
			}
		}
		var sourceTiktokMarketingSandboxAccessToken *shared.SourceTiktokMarketingSandboxAccessToken
		if r.Configuration.Credentials.SandboxAccessToken != nil {
			accessToken1 := r.Configuration.Credentials.SandboxAccessToken.AccessToken.ValueString()
			advertiserId1 := r.Configuration.Credentials.SandboxAccessToken.AdvertiserID.ValueString()
			sourceTiktokMarketingSandboxAccessToken = &shared.SourceTiktokMarketingSandboxAccessToken{
				AccessToken:  accessToken1,
				AdvertiserID: advertiserId1,
			}
		}
		if sourceTiktokMarketingSandboxAccessToken != nil {
			credentials = &shared.SourceTiktokMarketingAuthenticationMethod{
				SourceTiktokMarketingSandboxAccessToken: sourceTiktokMarketingSandboxAccessToken,
			}
		}
	}
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	includeDeleted := new(bool)
	if !r.Configuration.IncludeDeleted.IsUnknown() && !r.Configuration.IncludeDeleted.IsNull() {
		*includeDeleted = r.Configuration.IncludeDeleted.ValueBool()
	} else {
		includeDeleted = nil
	}
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceTiktokMarketing{
		AttributionWindow: attributionWindow,
		Credentials:       credentials,
		EndDate:           endDate,
		IncludeDeleted:    includeDeleted,
		StartDate:         startDate,
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
	out := shared.SourceTiktokMarketingCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTiktokMarketingResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceTiktokMarketingResourceModel) ToSharedSourceTiktokMarketingPutRequest() *shared.SourceTiktokMarketingPutRequest {
	attributionWindow := new(int64)
	if !r.Configuration.AttributionWindow.IsUnknown() && !r.Configuration.AttributionWindow.IsNull() {
		*attributionWindow = r.Configuration.AttributionWindow.ValueInt64()
	} else {
		attributionWindow = nil
	}
	var credentials *shared.SourceTiktokMarketingUpdateAuthenticationMethod
	if r.Configuration.Credentials != nil {
		var sourceTiktokMarketingUpdateOAuth20 *shared.SourceTiktokMarketingUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			advertiserID := new(string)
			if !r.Configuration.Credentials.OAuth20.AdvertiserID.IsUnknown() && !r.Configuration.Credentials.OAuth20.AdvertiserID.IsNull() {
				*advertiserID = r.Configuration.Credentials.OAuth20.AdvertiserID.ValueString()
			} else {
				advertiserID = nil
			}
			appID := r.Configuration.Credentials.OAuth20.AppID.ValueString()
			secret := r.Configuration.Credentials.OAuth20.Secret.ValueString()
			sourceTiktokMarketingUpdateOAuth20 = &shared.SourceTiktokMarketingUpdateOAuth20{
				AccessToken:  accessToken,
				AdvertiserID: advertiserID,
				AppID:        appID,
				Secret:       secret,
			}
		}
		if sourceTiktokMarketingUpdateOAuth20 != nil {
			credentials = &shared.SourceTiktokMarketingUpdateAuthenticationMethod{
				SourceTiktokMarketingUpdateOAuth20: sourceTiktokMarketingUpdateOAuth20,
			}
		}
		var sandboxAccessToken *shared.SandboxAccessToken
		if r.Configuration.Credentials.SandboxAccessToken != nil {
			accessToken1 := r.Configuration.Credentials.SandboxAccessToken.AccessToken.ValueString()
			advertiserId1 := r.Configuration.Credentials.SandboxAccessToken.AdvertiserID.ValueString()
			sandboxAccessToken = &shared.SandboxAccessToken{
				AccessToken:  accessToken1,
				AdvertiserID: advertiserId1,
			}
		}
		if sandboxAccessToken != nil {
			credentials = &shared.SourceTiktokMarketingUpdateAuthenticationMethod{
				SandboxAccessToken: sandboxAccessToken,
			}
		}
	}
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	includeDeleted := new(bool)
	if !r.Configuration.IncludeDeleted.IsUnknown() && !r.Configuration.IncludeDeleted.IsNull() {
		*includeDeleted = r.Configuration.IncludeDeleted.ValueBool()
	} else {
		includeDeleted = nil
	}
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceTiktokMarketingUpdate{
		AttributionWindow: attributionWindow,
		Credentials:       credentials,
		EndDate:           endDate,
		IncludeDeleted:    includeDeleted,
		StartDate:         startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTiktokMarketingPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
