// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceOktaResourceModel) ToSharedSourceOktaCreateRequest() *shared.SourceOktaCreateRequest {
	var credentials *shared.SourceOktaAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceOktaOAuth20 *shared.SourceOktaOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			var clientID string
			clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()

			var clientSecret string
			clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()

			var refreshToken string
			refreshToken = r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()

			sourceOktaOAuth20 = &shared.SourceOktaOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceOktaOAuth20 != nil {
			credentials = &shared.SourceOktaAuthorizationMethod{
				SourceOktaOAuth20: sourceOktaOAuth20,
			}
		}
		var sourceOktaOAuth20WithPrivateKey *shared.SourceOktaOAuth20WithPrivateKey
		if r.Configuration.Credentials.OAuth20WithPrivateKey != nil {
			var clientId1 string
			clientId1 = r.Configuration.Credentials.OAuth20WithPrivateKey.ClientID.ValueString()

			var keyID string
			keyID = r.Configuration.Credentials.OAuth20WithPrivateKey.KeyID.ValueString()

			var privateKey string
			privateKey = r.Configuration.Credentials.OAuth20WithPrivateKey.PrivateKey.ValueString()

			var scope string
			scope = r.Configuration.Credentials.OAuth20WithPrivateKey.Scope.ValueString()

			sourceOktaOAuth20WithPrivateKey = &shared.SourceOktaOAuth20WithPrivateKey{
				ClientID:   clientId1,
				KeyID:      keyID,
				PrivateKey: privateKey,
				Scope:      scope,
			}
		}
		if sourceOktaOAuth20WithPrivateKey != nil {
			credentials = &shared.SourceOktaAuthorizationMethod{
				SourceOktaOAuth20WithPrivateKey: sourceOktaOAuth20WithPrivateKey,
			}
		}
		var sourceOktaAPIToken *shared.SourceOktaAPIToken
		if r.Configuration.Credentials.APIToken != nil {
			var apiToken string
			apiToken = r.Configuration.Credentials.APIToken.APIToken.ValueString()

			sourceOktaAPIToken = &shared.SourceOktaAPIToken{
				APIToken: apiToken,
			}
		}
		if sourceOktaAPIToken != nil {
			credentials = &shared.SourceOktaAuthorizationMethod{
				SourceOktaAPIToken: sourceOktaAPIToken,
			}
		}
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceOkta{
		Credentials: credentials,
		Domain:      domain,
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

	out := shared.SourceOktaCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOktaResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceOktaResourceModel) ToSharedSourceOktaPutRequest() *shared.SourceOktaPutRequest {
	var credentials *shared.SourceOktaUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceOktaUpdateOAuth20 *shared.SourceOktaUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			var clientID string
			clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()

			var clientSecret string
			clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()

			var refreshToken string
			refreshToken = r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()

			sourceOktaUpdateOAuth20 = &shared.SourceOktaUpdateOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceOktaUpdateOAuth20 != nil {
			credentials = &shared.SourceOktaUpdateAuthorizationMethod{
				SourceOktaUpdateOAuth20: sourceOktaUpdateOAuth20,
			}
		}
		var oAuth20WithPrivateKey *shared.OAuth20WithPrivateKey
		if r.Configuration.Credentials.OAuth20WithPrivateKey != nil {
			var clientId1 string
			clientId1 = r.Configuration.Credentials.OAuth20WithPrivateKey.ClientID.ValueString()

			var keyID string
			keyID = r.Configuration.Credentials.OAuth20WithPrivateKey.KeyID.ValueString()

			var privateKey string
			privateKey = r.Configuration.Credentials.OAuth20WithPrivateKey.PrivateKey.ValueString()

			var scope string
			scope = r.Configuration.Credentials.OAuth20WithPrivateKey.Scope.ValueString()

			oAuth20WithPrivateKey = &shared.OAuth20WithPrivateKey{
				ClientID:   clientId1,
				KeyID:      keyID,
				PrivateKey: privateKey,
				Scope:      scope,
			}
		}
		if oAuth20WithPrivateKey != nil {
			credentials = &shared.SourceOktaUpdateAuthorizationMethod{
				OAuth20WithPrivateKey: oAuth20WithPrivateKey,
			}
		}
		var sourceOktaUpdateAPIToken *shared.SourceOktaUpdateAPIToken
		if r.Configuration.Credentials.APIToken != nil {
			var apiToken string
			apiToken = r.Configuration.Credentials.APIToken.APIToken.ValueString()

			sourceOktaUpdateAPIToken = &shared.SourceOktaUpdateAPIToken{
				APIToken: apiToken,
			}
		}
		if sourceOktaUpdateAPIToken != nil {
			credentials = &shared.SourceOktaUpdateAuthorizationMethod{
				SourceOktaUpdateAPIToken: sourceOktaUpdateAPIToken,
			}
		}
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceOktaUpdate{
		Credentials: credentials,
		Domain:      domain,
		StartDate:   startDate,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceOktaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
