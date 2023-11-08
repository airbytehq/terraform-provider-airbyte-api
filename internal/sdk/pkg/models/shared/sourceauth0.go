// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceAuth0SchemasCredentialsAuthenticationMethod string

const (
	SourceAuth0SchemasCredentialsAuthenticationMethodOauth2AccessToken SourceAuth0SchemasCredentialsAuthenticationMethod = "oauth2_access_token"
)

func (e SourceAuth0SchemasCredentialsAuthenticationMethod) ToPointer() *SourceAuth0SchemasCredentialsAuthenticationMethod {
	return &e
}

func (e *SourceAuth0SchemasCredentialsAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2_access_token":
		*e = SourceAuth0SchemasCredentialsAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAuth0SchemasCredentialsAuthenticationMethod: %v", v)
	}
}

type SourceAuth0OAuth2AccessToken struct {
	// Also called <a href="https://auth0.com/docs/secure/tokens/access-tokens/get-management-api-access-tokens-for-testing">API Access Token </a> The access token used to call the Auth0 Management API Token. It's a JWT that contains specific grant permissions knowns as scopes.
	AccessToken string                                            `json:"access_token"`
	authType    SourceAuth0SchemasCredentialsAuthenticationMethod `const:"oauth2_access_token" json:"auth_type"`
}

func (s SourceAuth0OAuth2AccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAuth0OAuth2AccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAuth0OAuth2AccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceAuth0OAuth2AccessToken) GetAuthType() SourceAuth0SchemasCredentialsAuthenticationMethod {
	return SourceAuth0SchemasCredentialsAuthenticationMethodOauth2AccessToken
}

type SourceAuth0SchemasAuthenticationMethod string

const (
	SourceAuth0SchemasAuthenticationMethodOauth2ConfidentialApplication SourceAuth0SchemasAuthenticationMethod = "oauth2_confidential_application"
)

func (e SourceAuth0SchemasAuthenticationMethod) ToPointer() *SourceAuth0SchemasAuthenticationMethod {
	return &e
}

func (e *SourceAuth0SchemasAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2_confidential_application":
		*e = SourceAuth0SchemasAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAuth0SchemasAuthenticationMethod: %v", v)
	}
}

type SourceAuth0OAuth2ConfidentialApplication struct {
	// The audience for the token, which is your API. You can find this in the Identifier field on your  <a href="https://manage.auth0.com/#/apis">API's settings tab</a>
	Audience string                                 `json:"audience"`
	authType SourceAuth0SchemasAuthenticationMethod `const:"oauth2_confidential_application" json:"auth_type"`
	// Your application's Client ID. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.
	ClientID string `json:"client_id"`
	// Your application's Client Secret. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.
	ClientSecret string `json:"client_secret"`
}

func (s SourceAuth0OAuth2ConfidentialApplication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAuth0OAuth2ConfidentialApplication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAuth0OAuth2ConfidentialApplication) GetAudience() string {
	if o == nil {
		return ""
	}
	return o.Audience
}

func (o *SourceAuth0OAuth2ConfidentialApplication) GetAuthType() SourceAuth0SchemasAuthenticationMethod {
	return SourceAuth0SchemasAuthenticationMethodOauth2ConfidentialApplication
}

func (o *SourceAuth0OAuth2ConfidentialApplication) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceAuth0OAuth2ConfidentialApplication) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

type SourceAuth0AuthenticationMethodType string

const (
	SourceAuth0AuthenticationMethodTypeSourceAuth0OAuth2ConfidentialApplication SourceAuth0AuthenticationMethodType = "source-auth0_OAuth2 Confidential Application"
	SourceAuth0AuthenticationMethodTypeSourceAuth0OAuth2AccessToken             SourceAuth0AuthenticationMethodType = "source-auth0_OAuth2 Access Token"
)

type SourceAuth0AuthenticationMethod struct {
	SourceAuth0OAuth2ConfidentialApplication *SourceAuth0OAuth2ConfidentialApplication
	SourceAuth0OAuth2AccessToken             *SourceAuth0OAuth2AccessToken

	Type SourceAuth0AuthenticationMethodType
}

func CreateSourceAuth0AuthenticationMethodSourceAuth0OAuth2ConfidentialApplication(sourceAuth0OAuth2ConfidentialApplication SourceAuth0OAuth2ConfidentialApplication) SourceAuth0AuthenticationMethod {
	typ := SourceAuth0AuthenticationMethodTypeSourceAuth0OAuth2ConfidentialApplication

	return SourceAuth0AuthenticationMethod{
		SourceAuth0OAuth2ConfidentialApplication: &sourceAuth0OAuth2ConfidentialApplication,
		Type:                                     typ,
	}
}

func CreateSourceAuth0AuthenticationMethodSourceAuth0OAuth2AccessToken(sourceAuth0OAuth2AccessToken SourceAuth0OAuth2AccessToken) SourceAuth0AuthenticationMethod {
	typ := SourceAuth0AuthenticationMethodTypeSourceAuth0OAuth2AccessToken

	return SourceAuth0AuthenticationMethod{
		SourceAuth0OAuth2AccessToken: &sourceAuth0OAuth2AccessToken,
		Type:                         typ,
	}
}

func (u *SourceAuth0AuthenticationMethod) UnmarshalJSON(data []byte) error {

	sourceAuth0OAuth2AccessToken := new(SourceAuth0OAuth2AccessToken)
	if err := utils.UnmarshalJSON(data, &sourceAuth0OAuth2AccessToken, "", true, true); err == nil {
		u.SourceAuth0OAuth2AccessToken = sourceAuth0OAuth2AccessToken
		u.Type = SourceAuth0AuthenticationMethodTypeSourceAuth0OAuth2AccessToken
		return nil
	}

	sourceAuth0OAuth2ConfidentialApplication := new(SourceAuth0OAuth2ConfidentialApplication)
	if err := utils.UnmarshalJSON(data, &sourceAuth0OAuth2ConfidentialApplication, "", true, true); err == nil {
		u.SourceAuth0OAuth2ConfidentialApplication = sourceAuth0OAuth2ConfidentialApplication
		u.Type = SourceAuth0AuthenticationMethodTypeSourceAuth0OAuth2ConfidentialApplication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceAuth0AuthenticationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceAuth0OAuth2ConfidentialApplication != nil {
		return utils.MarshalJSON(u.SourceAuth0OAuth2ConfidentialApplication, "", true)
	}

	if u.SourceAuth0OAuth2AccessToken != nil {
		return utils.MarshalJSON(u.SourceAuth0OAuth2AccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Auth0 string

const (
	Auth0Auth0 Auth0 = "auth0"
)

func (e Auth0) ToPointer() *Auth0 {
	return &e
}

func (e *Auth0) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auth0":
		*e = Auth0(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Auth0: %v", v)
	}
}

type SourceAuth0 struct {
	// The Authentication API is served over HTTPS. All URLs referenced in the documentation have the following base `https://YOUR_DOMAIN`
	BaseURL     string                          `json:"base_url"`
	Credentials SourceAuth0AuthenticationMethod `json:"credentials"`
	sourceType  Auth0                           `const:"auth0" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate *string `default:"2023-08-05T00:43:59.244Z" json:"start_date"`
}

func (s SourceAuth0) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAuth0) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAuth0) GetBaseURL() string {
	if o == nil {
		return ""
	}
	return o.BaseURL
}

func (o *SourceAuth0) GetCredentials() SourceAuth0AuthenticationMethod {
	if o == nil {
		return SourceAuth0AuthenticationMethod{}
	}
	return o.Credentials
}

func (o *SourceAuth0) GetSourceType() Auth0 {
	return Auth0Auth0
}

func (o *SourceAuth0) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}
