// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// SourceAsanaSchemasCredentialsTitle - PAT Credentials
type SourceAsanaSchemasCredentialsTitle string

const (
	SourceAsanaSchemasCredentialsTitlePatCredentials SourceAsanaSchemasCredentialsTitle = "PAT Credentials"
)

func (e SourceAsanaSchemasCredentialsTitle) ToPointer() *SourceAsanaSchemasCredentialsTitle {
	return &e
}

func (e *SourceAsanaSchemasCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PAT Credentials":
		*e = SourceAsanaSchemasCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAsanaSchemasCredentialsTitle: %v", v)
	}
}

// SourceAsanaAuthenticateWithPersonalAccessToken - Choose how to authenticate to Github
type SourceAsanaAuthenticateWithPersonalAccessToken struct {
	// PAT Credentials
	optionTitle *SourceAsanaSchemasCredentialsTitle `const:"PAT Credentials" json:"option_title,omitempty"`
	// Asana Personal Access Token (generate yours <a href="https://app.asana.com/0/developer-console">here</a>).
	PersonalAccessToken string `json:"personal_access_token"`
}

func (s SourceAsanaAuthenticateWithPersonalAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAsanaAuthenticateWithPersonalAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAsanaAuthenticateWithPersonalAccessToken) GetOptionTitle() *SourceAsanaSchemasCredentialsTitle {
	return SourceAsanaSchemasCredentialsTitlePatCredentials.ToPointer()
}

func (o *SourceAsanaAuthenticateWithPersonalAccessToken) GetPersonalAccessToken() string {
	if o == nil {
		return ""
	}
	return o.PersonalAccessToken
}

// SourceAsanaCredentialsTitle - OAuth Credentials
type SourceAsanaCredentialsTitle string

const (
	SourceAsanaCredentialsTitleOAuthCredentials SourceAsanaCredentialsTitle = "OAuth Credentials"
)

func (e SourceAsanaCredentialsTitle) ToPointer() *SourceAsanaCredentialsTitle {
	return &e
}

func (e *SourceAsanaCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAuth Credentials":
		*e = SourceAsanaCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAsanaCredentialsTitle: %v", v)
	}
}

// SourceAsanaAuthenticateViaAsanaOauth - Choose how to authenticate to Github
type SourceAsanaAuthenticateViaAsanaOauth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	// OAuth Credentials
	optionTitle  *SourceAsanaCredentialsTitle `const:"OAuth Credentials" json:"option_title,omitempty"`
	RefreshToken string                       `json:"refresh_token"`
}

func (s SourceAsanaAuthenticateViaAsanaOauth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAsanaAuthenticateViaAsanaOauth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAsanaAuthenticateViaAsanaOauth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceAsanaAuthenticateViaAsanaOauth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceAsanaAuthenticateViaAsanaOauth) GetOptionTitle() *SourceAsanaCredentialsTitle {
	return SourceAsanaCredentialsTitleOAuthCredentials.ToPointer()
}

func (o *SourceAsanaAuthenticateViaAsanaOauth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceAsanaAuthenticationMechanismType string

const (
	SourceAsanaAuthenticationMechanismTypeSourceAsanaAuthenticateViaAsanaOauth           SourceAsanaAuthenticationMechanismType = "source-asana_Authenticate via Asana (Oauth)"
	SourceAsanaAuthenticationMechanismTypeSourceAsanaAuthenticateWithPersonalAccessToken SourceAsanaAuthenticationMechanismType = "source-asana_Authenticate with Personal Access Token"
)

type SourceAsanaAuthenticationMechanism struct {
	SourceAsanaAuthenticateViaAsanaOauth           *SourceAsanaAuthenticateViaAsanaOauth
	SourceAsanaAuthenticateWithPersonalAccessToken *SourceAsanaAuthenticateWithPersonalAccessToken

	Type SourceAsanaAuthenticationMechanismType
}

func CreateSourceAsanaAuthenticationMechanismSourceAsanaAuthenticateViaAsanaOauth(sourceAsanaAuthenticateViaAsanaOauth SourceAsanaAuthenticateViaAsanaOauth) SourceAsanaAuthenticationMechanism {
	typ := SourceAsanaAuthenticationMechanismTypeSourceAsanaAuthenticateViaAsanaOauth

	return SourceAsanaAuthenticationMechanism{
		SourceAsanaAuthenticateViaAsanaOauth: &sourceAsanaAuthenticateViaAsanaOauth,
		Type:                                 typ,
	}
}

func CreateSourceAsanaAuthenticationMechanismSourceAsanaAuthenticateWithPersonalAccessToken(sourceAsanaAuthenticateWithPersonalAccessToken SourceAsanaAuthenticateWithPersonalAccessToken) SourceAsanaAuthenticationMechanism {
	typ := SourceAsanaAuthenticationMechanismTypeSourceAsanaAuthenticateWithPersonalAccessToken

	return SourceAsanaAuthenticationMechanism{
		SourceAsanaAuthenticateWithPersonalAccessToken: &sourceAsanaAuthenticateWithPersonalAccessToken,
		Type: typ,
	}
}

func (u *SourceAsanaAuthenticationMechanism) UnmarshalJSON(data []byte) error {

	sourceAsanaAuthenticateWithPersonalAccessToken := new(SourceAsanaAuthenticateWithPersonalAccessToken)
	if err := utils.UnmarshalJSON(data, &sourceAsanaAuthenticateWithPersonalAccessToken, "", true, true); err == nil {
		u.SourceAsanaAuthenticateWithPersonalAccessToken = sourceAsanaAuthenticateWithPersonalAccessToken
		u.Type = SourceAsanaAuthenticationMechanismTypeSourceAsanaAuthenticateWithPersonalAccessToken
		return nil
	}

	sourceAsanaAuthenticateViaAsanaOauth := new(SourceAsanaAuthenticateViaAsanaOauth)
	if err := utils.UnmarshalJSON(data, &sourceAsanaAuthenticateViaAsanaOauth, "", true, true); err == nil {
		u.SourceAsanaAuthenticateViaAsanaOauth = sourceAsanaAuthenticateViaAsanaOauth
		u.Type = SourceAsanaAuthenticationMechanismTypeSourceAsanaAuthenticateViaAsanaOauth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceAsanaAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.SourceAsanaAuthenticateViaAsanaOauth != nil {
		return utils.MarshalJSON(u.SourceAsanaAuthenticateViaAsanaOauth, "", true)
	}

	if u.SourceAsanaAuthenticateWithPersonalAccessToken != nil {
		return utils.MarshalJSON(u.SourceAsanaAuthenticateWithPersonalAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Asana string

const (
	AsanaAsana Asana = "asana"
)

func (e Asana) ToPointer() *Asana {
	return &e
}

func (e *Asana) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asana":
		*e = Asana(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Asana: %v", v)
	}
}

type SourceAsana struct {
	// Choose how to authenticate to Github
	Credentials *SourceAsanaAuthenticationMechanism `json:"credentials,omitempty"`
	sourceType  *Asana                              `const:"asana" json:"sourceType,omitempty"`
}

func (s SourceAsana) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAsana) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAsana) GetCredentials() *SourceAsanaAuthenticationMechanism {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceAsana) GetSourceType() *Asana {
	return AsanaAsana.ToPointer()
}
