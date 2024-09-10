// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceAirtableResourceModel) ToSharedSourceAirtableCreateRequest() *shared.SourceAirtableCreateRequest {
	var credentials *shared.SourceAirtableAuthentication
	if r.Configuration.Credentials != nil {
		var sourceAirtableOAuth20 *shared.SourceAirtableOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := new(string)
			if !r.Configuration.Credentials.OAuth20.AccessToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.AccessToken.IsNull() {
				*accessToken = r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			} else {
				accessToken = nil
			}
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
			tokenExpiryDate := new(time.Time)
			if !r.Configuration.Credentials.OAuth20.TokenExpiryDate.IsUnknown() && !r.Configuration.Credentials.OAuth20.TokenExpiryDate.IsNull() {
				*tokenExpiryDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
			} else {
				tokenExpiryDate = nil
			}
			sourceAirtableOAuth20 = &shared.SourceAirtableOAuth20{
				AccessToken:     accessToken,
				ClientID:        clientID,
				ClientSecret:    clientSecret,
				RefreshToken:    refreshToken,
				TokenExpiryDate: tokenExpiryDate,
			}
		}
		if sourceAirtableOAuth20 != nil {
			credentials = &shared.SourceAirtableAuthentication{
				SourceAirtableOAuth20: sourceAirtableOAuth20,
			}
		}
		var sourceAirtablePersonalAccessToken *shared.SourceAirtablePersonalAccessToken
		if r.Configuration.Credentials.PersonalAccessToken != nil {
			apiKey := r.Configuration.Credentials.PersonalAccessToken.APIKey.ValueString()
			sourceAirtablePersonalAccessToken = &shared.SourceAirtablePersonalAccessToken{
				APIKey: apiKey,
			}
		}
		if sourceAirtablePersonalAccessToken != nil {
			credentials = &shared.SourceAirtableAuthentication{
				SourceAirtablePersonalAccessToken: sourceAirtablePersonalAccessToken,
			}
		}
	}
	configuration := shared.SourceAirtable{
		Credentials: credentials,
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
	out := shared.SourceAirtableCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAirtableResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceAirtableResourceModel) ToSharedSourceAirtablePutRequest() *shared.SourceAirtablePutRequest {
	var credentials *shared.SourceAirtableUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceAirtableUpdateOAuth20 *shared.SourceAirtableUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := new(string)
			if !r.Configuration.Credentials.OAuth20.AccessToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.AccessToken.IsNull() {
				*accessToken = r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			} else {
				accessToken = nil
			}
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
			tokenExpiryDate := new(time.Time)
			if !r.Configuration.Credentials.OAuth20.TokenExpiryDate.IsUnknown() && !r.Configuration.Credentials.OAuth20.TokenExpiryDate.IsNull() {
				*tokenExpiryDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
			} else {
				tokenExpiryDate = nil
			}
			sourceAirtableUpdateOAuth20 = &shared.SourceAirtableUpdateOAuth20{
				AccessToken:     accessToken,
				ClientID:        clientID,
				ClientSecret:    clientSecret,
				RefreshToken:    refreshToken,
				TokenExpiryDate: tokenExpiryDate,
			}
		}
		if sourceAirtableUpdateOAuth20 != nil {
			credentials = &shared.SourceAirtableUpdateAuthentication{
				SourceAirtableUpdateOAuth20: sourceAirtableUpdateOAuth20,
			}
		}
		var sourceAirtableUpdatePersonalAccessToken *shared.SourceAirtableUpdatePersonalAccessToken
		if r.Configuration.Credentials.PersonalAccessToken != nil {
			apiKey := r.Configuration.Credentials.PersonalAccessToken.APIKey.ValueString()
			sourceAirtableUpdatePersonalAccessToken = &shared.SourceAirtableUpdatePersonalAccessToken{
				APIKey: apiKey,
			}
		}
		if sourceAirtableUpdatePersonalAccessToken != nil {
			credentials = &shared.SourceAirtableUpdateAuthentication{
				SourceAirtableUpdatePersonalAccessToken: sourceAirtableUpdatePersonalAccessToken,
			}
		}
	}
	configuration := shared.SourceAirtableUpdate{
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAirtablePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
