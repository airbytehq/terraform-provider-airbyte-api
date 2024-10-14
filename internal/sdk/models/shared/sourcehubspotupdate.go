// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

// SourceHubspotUpdateSchemasAuthType - Name of the credentials set
type SourceHubspotUpdateSchemasAuthType string

const (
	SourceHubspotUpdateSchemasAuthTypePrivateAppCredentials SourceHubspotUpdateSchemasAuthType = "Private App Credentials"
)

func (e SourceHubspotUpdateSchemasAuthType) ToPointer() *SourceHubspotUpdateSchemasAuthType {
	return &e
}
func (e *SourceHubspotUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Private App Credentials":
		*e = SourceHubspotUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHubspotUpdateSchemasAuthType: %v", v)
	}
}

type PrivateApp struct {
	// HubSpot Access token. See the <a href="https://developers.hubspot.com/docs/api/private-apps">Hubspot docs</a> if you need help finding this token.
	AccessToken string `json:"access_token"`
	// Name of the credentials set
	credentialsTitle SourceHubspotUpdateSchemasAuthType `const:"Private App Credentials" json:"credentials_title"`
}

func (p PrivateApp) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PrivateApp) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *PrivateApp) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *PrivateApp) GetCredentialsTitle() SourceHubspotUpdateSchemasAuthType {
	return SourceHubspotUpdateSchemasAuthTypePrivateAppCredentials
}

// SourceHubspotUpdateAuthType - Name of the credentials
type SourceHubspotUpdateAuthType string

const (
	SourceHubspotUpdateAuthTypeOAuthCredentials SourceHubspotUpdateAuthType = "OAuth Credentials"
)

func (e SourceHubspotUpdateAuthType) ToPointer() *SourceHubspotUpdateAuthType {
	return &e
}
func (e *SourceHubspotUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAuth Credentials":
		*e = SourceHubspotUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHubspotUpdateAuthType: %v", v)
	}
}

type SourceHubspotUpdateOAuth struct {
	// The Client ID of your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this ID.
	ClientID string `json:"client_id"`
	// The client secret for your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this secret.
	ClientSecret string `json:"client_secret"`
	// Name of the credentials
	credentialsTitle SourceHubspotUpdateAuthType `const:"OAuth Credentials" json:"credentials_title"`
	// Refresh token to renew an expired access token. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this token.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceHubspotUpdateOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceHubspotUpdateOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceHubspotUpdateOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceHubspotUpdateOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceHubspotUpdateOAuth) GetCredentialsTitle() SourceHubspotUpdateAuthType {
	return SourceHubspotUpdateAuthTypeOAuthCredentials
}

func (o *SourceHubspotUpdateOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceHubspotUpdateAuthenticationType string

const (
	SourceHubspotUpdateAuthenticationTypeSourceHubspotUpdateOAuth SourceHubspotUpdateAuthenticationType = "source-hubspot-update_OAuth"
	SourceHubspotUpdateAuthenticationTypePrivateApp               SourceHubspotUpdateAuthenticationType = "Private App"
)

// SourceHubspotUpdateAuthentication - Choose how to authenticate to HubSpot.
type SourceHubspotUpdateAuthentication struct {
	SourceHubspotUpdateOAuth *SourceHubspotUpdateOAuth
	PrivateApp               *PrivateApp

	Type SourceHubspotUpdateAuthenticationType
}

func CreateSourceHubspotUpdateAuthenticationSourceHubspotUpdateOAuth(sourceHubspotUpdateOAuth SourceHubspotUpdateOAuth) SourceHubspotUpdateAuthentication {
	typ := SourceHubspotUpdateAuthenticationTypeSourceHubspotUpdateOAuth

	return SourceHubspotUpdateAuthentication{
		SourceHubspotUpdateOAuth: &sourceHubspotUpdateOAuth,
		Type:                     typ,
	}
}

func CreateSourceHubspotUpdateAuthenticationPrivateApp(privateApp PrivateApp) SourceHubspotUpdateAuthentication {
	typ := SourceHubspotUpdateAuthenticationTypePrivateApp

	return SourceHubspotUpdateAuthentication{
		PrivateApp: &privateApp,
		Type:       typ,
	}
}

func (u *SourceHubspotUpdateAuthentication) UnmarshalJSON(data []byte) error {

	var privateApp PrivateApp = PrivateApp{}
	if err := utils.UnmarshalJSON(data, &privateApp, "", true, true); err == nil {
		u.PrivateApp = &privateApp
		u.Type = SourceHubspotUpdateAuthenticationTypePrivateApp
		return nil
	}

	var sourceHubspotUpdateOAuth SourceHubspotUpdateOAuth = SourceHubspotUpdateOAuth{}
	if err := utils.UnmarshalJSON(data, &sourceHubspotUpdateOAuth, "", true, true); err == nil {
		u.SourceHubspotUpdateOAuth = &sourceHubspotUpdateOAuth
		u.Type = SourceHubspotUpdateAuthenticationTypeSourceHubspotUpdateOAuth
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceHubspotUpdateAuthentication", string(data))
}

func (u SourceHubspotUpdateAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceHubspotUpdateOAuth != nil {
		return utils.MarshalJSON(u.SourceHubspotUpdateOAuth, "", true)
	}

	if u.PrivateApp != nil {
		return utils.MarshalJSON(u.PrivateApp, "", true)
	}

	return nil, errors.New("could not marshal union type SourceHubspotUpdateAuthentication: all fields are null")
}

type SourceHubspotUpdate struct {
	// Choose how to authenticate to HubSpot.
	Credentials SourceHubspotUpdateAuthentication `json:"credentials"`
	// If enabled then experimental streams become available for sync.
	EnableExperimentalStreams *bool `default:"false" json:"enable_experimental_streams"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. If not set, "2006-06-01T00:00:00Z" (Hubspot creation date) will be used as start date. It's recommended to provide relevant to your data start date value to optimize synchronization.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceHubspotUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceHubspotUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceHubspotUpdate) GetCredentials() SourceHubspotUpdateAuthentication {
	if o == nil {
		return SourceHubspotUpdateAuthentication{}
	}
	return o.Credentials
}

func (o *SourceHubspotUpdate) GetEnableExperimentalStreams() *bool {
	if o == nil {
		return nil
	}
	return o.EnableExperimentalStreams
}

func (o *SourceHubspotUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
