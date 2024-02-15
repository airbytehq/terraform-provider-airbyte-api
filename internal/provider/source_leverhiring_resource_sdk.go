// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceLeverHiringResourceModel) ToSharedSourceLeverHiringCreateRequest() *shared.SourceLeverHiringCreateRequest {
	var credentials *shared.SourceLeverHiringAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceLeverHiringAuthenticateViaLeverOAuth *shared.SourceLeverHiringAuthenticateViaLeverOAuth
		if r.Configuration.Credentials.AuthenticateViaLeverOAuth != nil {
			clientID := new(string)
			if !r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientID.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientSecret.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			refreshToken := r.Configuration.Credentials.AuthenticateViaLeverOAuth.RefreshToken.ValueString()
			sourceLeverHiringAuthenticateViaLeverOAuth = &shared.SourceLeverHiringAuthenticateViaLeverOAuth{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLeverHiringAuthenticateViaLeverOAuth != nil {
			credentials = &shared.SourceLeverHiringAuthenticationMechanism{
				SourceLeverHiringAuthenticateViaLeverOAuth: sourceLeverHiringAuthenticateViaLeverOAuth,
			}
		}
		var sourceLeverHiringAuthenticateViaLeverAPIKey *shared.SourceLeverHiringAuthenticateViaLeverAPIKey
		if r.Configuration.Credentials.AuthenticateViaLeverAPIKey != nil {
			apiKey := r.Configuration.Credentials.AuthenticateViaLeverAPIKey.APIKey.ValueString()
			sourceLeverHiringAuthenticateViaLeverAPIKey = &shared.SourceLeverHiringAuthenticateViaLeverAPIKey{
				APIKey: apiKey,
			}
		}
		if sourceLeverHiringAuthenticateViaLeverAPIKey != nil {
			credentials = &shared.SourceLeverHiringAuthenticationMechanism{
				SourceLeverHiringAuthenticateViaLeverAPIKey: sourceLeverHiringAuthenticateViaLeverAPIKey,
			}
		}
	}
	environment := new(shared.SourceLeverHiringEnvironment)
	if !r.Configuration.Environment.IsUnknown() && !r.Configuration.Environment.IsNull() {
		*environment = shared.SourceLeverHiringEnvironment(r.Configuration.Environment.ValueString())
	} else {
		environment = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceLeverHiring{
		Credentials: credentials,
		Environment: environment,
		StartDate:   startDate,
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
	out := shared.SourceLeverHiringCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLeverHiringResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceLeverHiringResourceModel) ToSharedSourceLeverHiringPutRequest() *shared.SourceLeverHiringPutRequest {
	var credentials *shared.SourceLeverHiringUpdateAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var authenticateViaLeverOAuth *shared.AuthenticateViaLeverOAuth
		if r.Configuration.Credentials.AuthenticateViaLeverOAuth != nil {
			clientID := new(string)
			if !r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientID.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientSecret.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.AuthenticateViaLeverOAuth.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			refreshToken := r.Configuration.Credentials.AuthenticateViaLeverOAuth.RefreshToken.ValueString()
			authenticateViaLeverOAuth = &shared.AuthenticateViaLeverOAuth{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if authenticateViaLeverOAuth != nil {
			credentials = &shared.SourceLeverHiringUpdateAuthenticationMechanism{
				AuthenticateViaLeverOAuth: authenticateViaLeverOAuth,
			}
		}
		var authenticateViaLeverAPIKey *shared.AuthenticateViaLeverAPIKey
		if r.Configuration.Credentials.AuthenticateViaLeverAPIKey != nil {
			apiKey := r.Configuration.Credentials.AuthenticateViaLeverAPIKey.APIKey.ValueString()
			authenticateViaLeverAPIKey = &shared.AuthenticateViaLeverAPIKey{
				APIKey: apiKey,
			}
		}
		if authenticateViaLeverAPIKey != nil {
			credentials = &shared.SourceLeverHiringUpdateAuthenticationMechanism{
				AuthenticateViaLeverAPIKey: authenticateViaLeverAPIKey,
			}
		}
	}
	environment := new(shared.SourceLeverHiringUpdateEnvironment)
	if !r.Configuration.Environment.IsUnknown() && !r.Configuration.Environment.IsNull() {
		*environment = shared.SourceLeverHiringUpdateEnvironment(r.Configuration.Environment.ValueString())
	} else {
		environment = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceLeverHiringUpdate{
		Credentials: credentials,
		Environment: environment,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLeverHiringPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
