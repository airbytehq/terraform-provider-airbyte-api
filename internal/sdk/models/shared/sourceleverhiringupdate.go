// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceLeverHiringUpdateSchemasAuthType string

const (
	SourceLeverHiringUpdateSchemasAuthTypeAPIKey SourceLeverHiringUpdateSchemasAuthType = "Api Key"
)

func (e SourceLeverHiringUpdateSchemasAuthType) ToPointer() *SourceLeverHiringUpdateSchemasAuthType {
	return &e
}
func (e *SourceLeverHiringUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Api Key":
		*e = SourceLeverHiringUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLeverHiringUpdateSchemasAuthType: %v", v)
	}
}

type AuthenticateViaLeverAPIKey struct {
	// The Api Key of your Lever Hiring account.
	APIKey   string                                  `json:"api_key"`
	authType *SourceLeverHiringUpdateSchemasAuthType `const:"Api Key" json:"auth_type,omitempty"`
}

func (a AuthenticateViaLeverAPIKey) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticateViaLeverAPIKey) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticateViaLeverAPIKey) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *AuthenticateViaLeverAPIKey) GetAuthType() *SourceLeverHiringUpdateSchemasAuthType {
	return SourceLeverHiringUpdateSchemasAuthTypeAPIKey.ToPointer()
}

type SourceLeverHiringUpdateAuthType string

const (
	SourceLeverHiringUpdateAuthTypeClient SourceLeverHiringUpdateAuthType = "Client"
)

func (e SourceLeverHiringUpdateAuthType) ToPointer() *SourceLeverHiringUpdateAuthType {
	return &e
}
func (e *SourceLeverHiringUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceLeverHiringUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLeverHiringUpdateAuthType: %v", v)
	}
}

type AuthenticateViaLeverOAuth struct {
	authType *SourceLeverHiringUpdateAuthType `const:"Client" json:"auth_type,omitempty"`
	// The Client ID of your Lever Hiring developer application.
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your Lever Hiring developer application.
	ClientSecret *string `json:"client_secret,omitempty"`
	// The token for obtaining new access token.
	RefreshToken string `json:"refresh_token"`
}

func (a AuthenticateViaLeverOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticateViaLeverOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticateViaLeverOAuth) GetAuthType() *SourceLeverHiringUpdateAuthType {
	return SourceLeverHiringUpdateAuthTypeClient.ToPointer()
}

func (o *AuthenticateViaLeverOAuth) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *AuthenticateViaLeverOAuth) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *AuthenticateViaLeverOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceLeverHiringUpdateAuthenticationMechanismType string

const (
	SourceLeverHiringUpdateAuthenticationMechanismTypeAuthenticateViaLeverOAuth  SourceLeverHiringUpdateAuthenticationMechanismType = "Authenticate via Lever (OAuth)"
	SourceLeverHiringUpdateAuthenticationMechanismTypeAuthenticateViaLeverAPIKey SourceLeverHiringUpdateAuthenticationMechanismType = "Authenticate via Lever (Api Key)"
)

// SourceLeverHiringUpdateAuthenticationMechanism - Choose how to authenticate to Lever Hiring.
type SourceLeverHiringUpdateAuthenticationMechanism struct {
	AuthenticateViaLeverOAuth  *AuthenticateViaLeverOAuth
	AuthenticateViaLeverAPIKey *AuthenticateViaLeverAPIKey

	Type SourceLeverHiringUpdateAuthenticationMechanismType
}

func CreateSourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth(authenticateViaLeverOAuth AuthenticateViaLeverOAuth) SourceLeverHiringUpdateAuthenticationMechanism {
	typ := SourceLeverHiringUpdateAuthenticationMechanismTypeAuthenticateViaLeverOAuth

	return SourceLeverHiringUpdateAuthenticationMechanism{
		AuthenticateViaLeverOAuth: &authenticateViaLeverOAuth,
		Type:                      typ,
	}
}

func CreateSourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey(authenticateViaLeverAPIKey AuthenticateViaLeverAPIKey) SourceLeverHiringUpdateAuthenticationMechanism {
	typ := SourceLeverHiringUpdateAuthenticationMechanismTypeAuthenticateViaLeverAPIKey

	return SourceLeverHiringUpdateAuthenticationMechanism{
		AuthenticateViaLeverAPIKey: &authenticateViaLeverAPIKey,
		Type:                       typ,
	}
}

func (u *SourceLeverHiringUpdateAuthenticationMechanism) UnmarshalJSON(data []byte) error {

	var authenticateViaLeverAPIKey AuthenticateViaLeverAPIKey = AuthenticateViaLeverAPIKey{}
	if err := utils.UnmarshalJSON(data, &authenticateViaLeverAPIKey, "", true, true); err == nil {
		u.AuthenticateViaLeverAPIKey = &authenticateViaLeverAPIKey
		u.Type = SourceLeverHiringUpdateAuthenticationMechanismTypeAuthenticateViaLeverAPIKey
		return nil
	}

	var authenticateViaLeverOAuth AuthenticateViaLeverOAuth = AuthenticateViaLeverOAuth{}
	if err := utils.UnmarshalJSON(data, &authenticateViaLeverOAuth, "", true, true); err == nil {
		u.AuthenticateViaLeverOAuth = &authenticateViaLeverOAuth
		u.Type = SourceLeverHiringUpdateAuthenticationMechanismTypeAuthenticateViaLeverOAuth
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceLeverHiringUpdateAuthenticationMechanism", string(data))
}

func (u SourceLeverHiringUpdateAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.AuthenticateViaLeverOAuth != nil {
		return utils.MarshalJSON(u.AuthenticateViaLeverOAuth, "", true)
	}

	if u.AuthenticateViaLeverAPIKey != nil {
		return utils.MarshalJSON(u.AuthenticateViaLeverAPIKey, "", true)
	}

	return nil, errors.New("could not marshal union type SourceLeverHiringUpdateAuthenticationMechanism: all fields are null")
}

// SourceLeverHiringUpdateEnvironment - The environment in which you'd like to replicate data for Lever. This is used to determine which Lever API endpoint to use.
type SourceLeverHiringUpdateEnvironment string

const (
	SourceLeverHiringUpdateEnvironmentProduction SourceLeverHiringUpdateEnvironment = "Production"
	SourceLeverHiringUpdateEnvironmentSandbox    SourceLeverHiringUpdateEnvironment = "Sandbox"
)

func (e SourceLeverHiringUpdateEnvironment) ToPointer() *SourceLeverHiringUpdateEnvironment {
	return &e
}
func (e *SourceLeverHiringUpdateEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Production":
		fallthrough
	case "Sandbox":
		*e = SourceLeverHiringUpdateEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLeverHiringUpdateEnvironment: %v", v)
	}
}

type SourceLeverHiringUpdate struct {
	// Choose how to authenticate to Lever Hiring.
	Credentials *SourceLeverHiringUpdateAuthenticationMechanism `json:"credentials,omitempty"`
	// The environment in which you'd like to replicate data for Lever. This is used to determine which Lever API endpoint to use.
	Environment *SourceLeverHiringUpdateEnvironment `default:"Sandbox" json:"environment"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. Note that it will be used only in the following incremental streams: comments, commits, and issues.
	StartDate string `json:"start_date"`
}

func (s SourceLeverHiringUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLeverHiringUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLeverHiringUpdate) GetCredentials() *SourceLeverHiringUpdateAuthenticationMechanism {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceLeverHiringUpdate) GetEnvironment() *SourceLeverHiringUpdateEnvironment {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *SourceLeverHiringUpdate) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
