// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceZendeskChatUpdateSchemasCredentials string

const (
	SourceZendeskChatUpdateSchemasCredentialsAccessToken SourceZendeskChatUpdateSchemasCredentials = "access_token"
)

func (e SourceZendeskChatUpdateSchemasCredentials) ToPointer() *SourceZendeskChatUpdateSchemasCredentials {
	return &e
}
func (e *SourceZendeskChatUpdateSchemasCredentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceZendeskChatUpdateSchemasCredentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskChatUpdateSchemasCredentials: %v", v)
	}
}

type SourceZendeskChatUpdateAccessToken struct {
	// The Access Token to make authenticated requests.
	AccessToken string                                    `json:"access_token"`
	credentials SourceZendeskChatUpdateSchemasCredentials `const:"access_token" json:"credentials"`
}

func (s SourceZendeskChatUpdateAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskChatUpdateAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskChatUpdateAccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceZendeskChatUpdateAccessToken) GetCredentials() SourceZendeskChatUpdateSchemasCredentials {
	return SourceZendeskChatUpdateSchemasCredentialsAccessToken
}

type SourceZendeskChatUpdateCredentials string

const (
	SourceZendeskChatUpdateCredentialsOauth20 SourceZendeskChatUpdateCredentials = "oauth2.0"
)

func (e SourceZendeskChatUpdateCredentials) ToPointer() *SourceZendeskChatUpdateCredentials {
	return &e
}
func (e *SourceZendeskChatUpdateCredentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceZendeskChatUpdateCredentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskChatUpdateCredentials: %v", v)
	}
}

type SourceZendeskChatUpdateOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken *string `json:"access_token,omitempty"`
	// The Client ID of your OAuth application
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your OAuth application.
	ClientSecret *string                            `json:"client_secret,omitempty"`
	credentials  SourceZendeskChatUpdateCredentials `const:"oauth2.0" json:"credentials"`
	// Refresh Token to obtain new Access Token, when it's expired.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

func (s SourceZendeskChatUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskChatUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskChatUpdateOAuth20) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *SourceZendeskChatUpdateOAuth20) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *SourceZendeskChatUpdateOAuth20) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *SourceZendeskChatUpdateOAuth20) GetCredentials() SourceZendeskChatUpdateCredentials {
	return SourceZendeskChatUpdateCredentialsOauth20
}

func (o *SourceZendeskChatUpdateOAuth20) GetRefreshToken() *string {
	if o == nil {
		return nil
	}
	return o.RefreshToken
}

type SourceZendeskChatUpdateAuthorizationMethodType string

const (
	SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateOAuth20     SourceZendeskChatUpdateAuthorizationMethodType = "source-zendesk-chat-update_OAuth2.0"
	SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateAccessToken SourceZendeskChatUpdateAuthorizationMethodType = "source-zendesk-chat-update_Access Token"
)

type SourceZendeskChatUpdateAuthorizationMethod struct {
	SourceZendeskChatUpdateOAuth20     *SourceZendeskChatUpdateOAuth20
	SourceZendeskChatUpdateAccessToken *SourceZendeskChatUpdateAccessToken

	Type SourceZendeskChatUpdateAuthorizationMethodType
}

func CreateSourceZendeskChatUpdateAuthorizationMethodSourceZendeskChatUpdateOAuth20(sourceZendeskChatUpdateOAuth20 SourceZendeskChatUpdateOAuth20) SourceZendeskChatUpdateAuthorizationMethod {
	typ := SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateOAuth20

	return SourceZendeskChatUpdateAuthorizationMethod{
		SourceZendeskChatUpdateOAuth20: &sourceZendeskChatUpdateOAuth20,
		Type:                           typ,
	}
}

func CreateSourceZendeskChatUpdateAuthorizationMethodSourceZendeskChatUpdateAccessToken(sourceZendeskChatUpdateAccessToken SourceZendeskChatUpdateAccessToken) SourceZendeskChatUpdateAuthorizationMethod {
	typ := SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateAccessToken

	return SourceZendeskChatUpdateAuthorizationMethod{
		SourceZendeskChatUpdateAccessToken: &sourceZendeskChatUpdateAccessToken,
		Type:                               typ,
	}
}

func (u *SourceZendeskChatUpdateAuthorizationMethod) UnmarshalJSON(data []byte) error {

	var sourceZendeskChatUpdateAccessToken SourceZendeskChatUpdateAccessToken = SourceZendeskChatUpdateAccessToken{}
	if err := utils.UnmarshalJSON(data, &sourceZendeskChatUpdateAccessToken, "", true, true); err == nil {
		u.SourceZendeskChatUpdateAccessToken = &sourceZendeskChatUpdateAccessToken
		u.Type = SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateAccessToken
		return nil
	}

	var sourceZendeskChatUpdateOAuth20 SourceZendeskChatUpdateOAuth20 = SourceZendeskChatUpdateOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceZendeskChatUpdateOAuth20, "", true, true); err == nil {
		u.SourceZendeskChatUpdateOAuth20 = &sourceZendeskChatUpdateOAuth20
		u.Type = SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateOAuth20
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceZendeskChatUpdateAuthorizationMethod", string(data))
}

func (u SourceZendeskChatUpdateAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceZendeskChatUpdateOAuth20 != nil {
		return utils.MarshalJSON(u.SourceZendeskChatUpdateOAuth20, "", true)
	}

	if u.SourceZendeskChatUpdateAccessToken != nil {
		return utils.MarshalJSON(u.SourceZendeskChatUpdateAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type SourceZendeskChatUpdateAuthorizationMethod: all fields are null")
}

type SourceZendeskChatUpdate struct {
	Credentials *SourceZendeskChatUpdateAuthorizationMethod `json:"credentials,omitempty"`
	// The date from which you'd like to replicate data for Zendesk Chat API, in the format YYYY-MM-DDT00:00:00Z.
	StartDate time.Time `json:"start_date"`
	// Required if you access Zendesk Chat from a Zendesk Support subdomain.
	Subdomain *string `default:"" json:"subdomain"`
}

func (s SourceZendeskChatUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskChatUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskChatUpdate) GetCredentials() *SourceZendeskChatUpdateAuthorizationMethod {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceZendeskChatUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceZendeskChatUpdate) GetSubdomain() *string {
	if o == nil {
		return nil
	}
	return o.Subdomain
}
