// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceZendeskSunshineResourceModel) ToCreateSDKType() *shared.SourceZendeskSunshineCreateRequest {
	var credentials *shared.SourceZendeskSunshineAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceZendeskSunshineOAuth20 *shared.SourceZendeskSunshineOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			sourceZendeskSunshineOAuth20 = &shared.SourceZendeskSunshineOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
		}
		if sourceZendeskSunshineOAuth20 != nil {
			credentials = &shared.SourceZendeskSunshineAuthorizationMethod{
				SourceZendeskSunshineOAuth20: sourceZendeskSunshineOAuth20,
			}
		}
		var sourceZendeskSunshineAPIToken *shared.SourceZendeskSunshineAPIToken
		if r.Configuration.Credentials.APIToken != nil {
			apiToken := r.Configuration.Credentials.APIToken.APIToken.ValueString()
			email := r.Configuration.Credentials.APIToken.Email.ValueString()
			sourceZendeskSunshineAPIToken = &shared.SourceZendeskSunshineAPIToken{
				APIToken: apiToken,
				Email:    email,
			}
		}
		if sourceZendeskSunshineAPIToken != nil {
			credentials = &shared.SourceZendeskSunshineAuthorizationMethod{
				SourceZendeskSunshineAPIToken: sourceZendeskSunshineAPIToken,
			}
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskSunshine{
		Credentials: credentials,
		StartDate:   startDate,
		Subdomain:   subdomain,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskSunshineCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskSunshineResourceModel) ToGetSDKType() *shared.SourceZendeskSunshineCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskSunshineResourceModel) ToUpdateSDKType() *shared.SourceZendeskSunshinePutRequest {
	var credentials *shared.SourceZendeskSunshineUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceZendeskSunshineUpdateOAuth20 *shared.SourceZendeskSunshineUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			sourceZendeskSunshineUpdateOAuth20 = &shared.SourceZendeskSunshineUpdateOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
		}
		if sourceZendeskSunshineUpdateOAuth20 != nil {
			credentials = &shared.SourceZendeskSunshineUpdateAuthorizationMethod{
				SourceZendeskSunshineUpdateOAuth20: sourceZendeskSunshineUpdateOAuth20,
			}
		}
		var sourceZendeskSunshineUpdateAPIToken *shared.SourceZendeskSunshineUpdateAPIToken
		if r.Configuration.Credentials.APIToken != nil {
			apiToken := r.Configuration.Credentials.APIToken.APIToken.ValueString()
			email := r.Configuration.Credentials.APIToken.Email.ValueString()
			sourceZendeskSunshineUpdateAPIToken = &shared.SourceZendeskSunshineUpdateAPIToken{
				APIToken: apiToken,
				Email:    email,
			}
		}
		if sourceZendeskSunshineUpdateAPIToken != nil {
			credentials = &shared.SourceZendeskSunshineUpdateAuthorizationMethod{
				SourceZendeskSunshineUpdateAPIToken: sourceZendeskSunshineUpdateAPIToken,
			}
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskSunshineUpdate{
		Credentials: credentials,
		StartDate:   startDate,
		Subdomain:   subdomain,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskSunshinePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskSunshineResourceModel) ToDeleteSDKType() *shared.SourceZendeskSunshineCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskSunshineResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceZendeskSunshineResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
