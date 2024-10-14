// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceRetentlyUpdateSchemasAuthType string

const (
	SourceRetentlyUpdateSchemasAuthTypeToken SourceRetentlyUpdateSchemasAuthType = "Token"
)

func (e SourceRetentlyUpdateSchemasAuthType) ToPointer() *SourceRetentlyUpdateSchemasAuthType {
	return &e
}
func (e *SourceRetentlyUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Token":
		*e = SourceRetentlyUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRetentlyUpdateSchemasAuthType: %v", v)
	}
}

type AuthenticateWithAPIToken struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// Retently API Token. See the <a href="https://app.retently.com/settings/api/tokens">docs</a> for more information on how to obtain this key.
	APIKey   string                               `json:"api_key"`
	authType *SourceRetentlyUpdateSchemasAuthType `const:"Token" json:"auth_type,omitempty"`
}

func (a AuthenticateWithAPIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticateWithAPIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticateWithAPIToken) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AuthenticateWithAPIToken) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *AuthenticateWithAPIToken) GetAuthType() *SourceRetentlyUpdateSchemasAuthType {
	return SourceRetentlyUpdateSchemasAuthTypeToken.ToPointer()
}

type SourceRetentlyUpdateAuthType string

const (
	SourceRetentlyUpdateAuthTypeClient SourceRetentlyUpdateAuthType = "Client"
)

func (e SourceRetentlyUpdateAuthType) ToPointer() *SourceRetentlyUpdateAuthType {
	return &e
}
func (e *SourceRetentlyUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceRetentlyUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRetentlyUpdateAuthType: %v", v)
	}
}

type AuthenticateViaRetentlyOAuth struct {
	AdditionalProperties any                           `additionalProperties:"true" json:"-"`
	authType             *SourceRetentlyUpdateAuthType `const:"Client" json:"auth_type,omitempty"`
	// The Client ID of your Retently developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Retently developer application.
	ClientSecret string `json:"client_secret"`
	// Retently Refresh Token which can be used to fetch new Bearer Tokens when the current one expires.
	RefreshToken string `json:"refresh_token"`
}

func (a AuthenticateViaRetentlyOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticateViaRetentlyOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticateViaRetentlyOAuth) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AuthenticateViaRetentlyOAuth) GetAuthType() *SourceRetentlyUpdateAuthType {
	return SourceRetentlyUpdateAuthTypeClient.ToPointer()
}

func (o *AuthenticateViaRetentlyOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *AuthenticateViaRetentlyOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *AuthenticateViaRetentlyOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceRetentlyUpdateAuthenticationMechanismType string

const (
	SourceRetentlyUpdateAuthenticationMechanismTypeAuthenticateViaRetentlyOAuth SourceRetentlyUpdateAuthenticationMechanismType = "Authenticate via Retently (OAuth)"
	SourceRetentlyUpdateAuthenticationMechanismTypeAuthenticateWithAPIToken     SourceRetentlyUpdateAuthenticationMechanismType = "Authenticate with API Token"
)

// SourceRetentlyUpdateAuthenticationMechanism - Choose how to authenticate to Retently
type SourceRetentlyUpdateAuthenticationMechanism struct {
	AuthenticateViaRetentlyOAuth *AuthenticateViaRetentlyOAuth
	AuthenticateWithAPIToken     *AuthenticateWithAPIToken

	Type SourceRetentlyUpdateAuthenticationMechanismType
}

func CreateSourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth(authenticateViaRetentlyOAuth AuthenticateViaRetentlyOAuth) SourceRetentlyUpdateAuthenticationMechanism {
	typ := SourceRetentlyUpdateAuthenticationMechanismTypeAuthenticateViaRetentlyOAuth

	return SourceRetentlyUpdateAuthenticationMechanism{
		AuthenticateViaRetentlyOAuth: &authenticateViaRetentlyOAuth,
		Type:                         typ,
	}
}

func CreateSourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken(authenticateWithAPIToken AuthenticateWithAPIToken) SourceRetentlyUpdateAuthenticationMechanism {
	typ := SourceRetentlyUpdateAuthenticationMechanismTypeAuthenticateWithAPIToken

	return SourceRetentlyUpdateAuthenticationMechanism{
		AuthenticateWithAPIToken: &authenticateWithAPIToken,
		Type:                     typ,
	}
}

func (u *SourceRetentlyUpdateAuthenticationMechanism) UnmarshalJSON(data []byte) error {

	var authenticateWithAPIToken AuthenticateWithAPIToken = AuthenticateWithAPIToken{}
	if err := utils.UnmarshalJSON(data, &authenticateWithAPIToken, "", true, true); err == nil {
		u.AuthenticateWithAPIToken = &authenticateWithAPIToken
		u.Type = SourceRetentlyUpdateAuthenticationMechanismTypeAuthenticateWithAPIToken
		return nil
	}

	var authenticateViaRetentlyOAuth AuthenticateViaRetentlyOAuth = AuthenticateViaRetentlyOAuth{}
	if err := utils.UnmarshalJSON(data, &authenticateViaRetentlyOAuth, "", true, true); err == nil {
		u.AuthenticateViaRetentlyOAuth = &authenticateViaRetentlyOAuth
		u.Type = SourceRetentlyUpdateAuthenticationMechanismTypeAuthenticateViaRetentlyOAuth
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceRetentlyUpdateAuthenticationMechanism", string(data))
}

func (u SourceRetentlyUpdateAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.AuthenticateViaRetentlyOAuth != nil {
		return utils.MarshalJSON(u.AuthenticateViaRetentlyOAuth, "", true)
	}

	if u.AuthenticateWithAPIToken != nil {
		return utils.MarshalJSON(u.AuthenticateWithAPIToken, "", true)
	}

	return nil, errors.New("could not marshal union type SourceRetentlyUpdateAuthenticationMechanism: all fields are null")
}

type SourceRetentlyUpdate struct {
	// Choose how to authenticate to Retently
	Credentials *SourceRetentlyUpdateAuthenticationMechanism `json:"credentials,omitempty"`
}

func (o *SourceRetentlyUpdate) GetCredentials() *SourceRetentlyUpdateAuthenticationMechanism {
	if o == nil {
		return nil
	}
	return o.Credentials
}
