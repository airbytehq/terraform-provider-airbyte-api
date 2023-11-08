// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceNotionResourceModel) ToCreateSDKType() *shared.SourceNotionCreateRequest {
	var credentials *shared.SourceNotionAuthenticateUsing
	if r.Configuration.Credentials != nil {
		var sourceNotionOAuth20 *shared.SourceNotionOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			sourceNotionOAuth20 = &shared.SourceNotionOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
		}
		if sourceNotionOAuth20 != nil {
			credentials = &shared.SourceNotionAuthenticateUsing{
				SourceNotionOAuth20: sourceNotionOAuth20,
			}
		}
		var sourceNotionAccessToken *shared.SourceNotionAccessToken
		if r.Configuration.Credentials.AccessToken != nil {
			token := r.Configuration.Credentials.AccessToken.Token.ValueString()
			sourceNotionAccessToken = &shared.SourceNotionAccessToken{
				Token: token,
			}
		}
		if sourceNotionAccessToken != nil {
			credentials = &shared.SourceNotionAuthenticateUsing{
				SourceNotionAccessToken: sourceNotionAccessToken,
			}
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceNotion{
		Credentials: credentials,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceNotionCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceNotionResourceModel) ToGetSDKType() *shared.SourceNotionCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceNotionResourceModel) ToUpdateSDKType() *shared.SourceNotionPutRequest {
	var credentials *shared.AuthenticateUsing
	if r.Configuration.Credentials != nil {
		var sourceNotionUpdateOAuth20 *shared.SourceNotionUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			sourceNotionUpdateOAuth20 = &shared.SourceNotionUpdateOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
		}
		if sourceNotionUpdateOAuth20 != nil {
			credentials = &shared.AuthenticateUsing{
				SourceNotionUpdateOAuth20: sourceNotionUpdateOAuth20,
			}
		}
		var sourceNotionUpdateAccessToken *shared.SourceNotionUpdateAccessToken
		if r.Configuration.Credentials.AccessToken != nil {
			token := r.Configuration.Credentials.AccessToken.Token.ValueString()
			sourceNotionUpdateAccessToken = &shared.SourceNotionUpdateAccessToken{
				Token: token,
			}
		}
		if sourceNotionUpdateAccessToken != nil {
			credentials = &shared.AuthenticateUsing{
				SourceNotionUpdateAccessToken: sourceNotionUpdateAccessToken,
			}
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceNotionUpdate{
		Credentials: credentials,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceNotionPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceNotionResourceModel) ToDeleteSDKType() *shared.SourceNotionCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceNotionResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceNotionResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
