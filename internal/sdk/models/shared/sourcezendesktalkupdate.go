// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceZendeskTalkUpdateSchemasAuthType string

const (
	SourceZendeskTalkUpdateSchemasAuthTypeAPIToken SourceZendeskTalkUpdateSchemasAuthType = "api_token"
)

func (e SourceZendeskTalkUpdateSchemasAuthType) ToPointer() *SourceZendeskTalkUpdateSchemasAuthType {
	return &e
}
func (e *SourceZendeskTalkUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceZendeskTalkUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskTalkUpdateSchemasAuthType: %v", v)
	}
}

type SourceZendeskTalkUpdateAPIToken struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// The value of the API token generated. See the <a href="https://docs.airbyte.com/integrations/sources/zendesk-talk">docs</a> for more information.
	APIToken string                                  `json:"api_token"`
	authType *SourceZendeskTalkUpdateSchemasAuthType `const:"api_token" json:"auth_type,omitempty"`
	// The user email for your Zendesk account.
	Email string `json:"email"`
}

func (s SourceZendeskTalkUpdateAPIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskTalkUpdateAPIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskTalkUpdateAPIToken) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceZendeskTalkUpdateAPIToken) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceZendeskTalkUpdateAPIToken) GetAuthType() *SourceZendeskTalkUpdateSchemasAuthType {
	return SourceZendeskTalkUpdateSchemasAuthTypeAPIToken.ToPointer()
}

func (o *SourceZendeskTalkUpdateAPIToken) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

type SourceZendeskTalkUpdateAuthType string

const (
	SourceZendeskTalkUpdateAuthTypeOauth20 SourceZendeskTalkUpdateAuthType = "oauth2.0"
)

func (e SourceZendeskTalkUpdateAuthType) ToPointer() *SourceZendeskTalkUpdateAuthType {
	return &e
}
func (e *SourceZendeskTalkUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceZendeskTalkUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskTalkUpdateAuthType: %v", v)
	}
}

type SourceZendeskTalkUpdateOAuth20 struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// The value of the API token generated. See the <a href="https://docs.airbyte.com/integrations/sources/zendesk-talk">docs</a> for more information.
	AccessToken string                           `json:"access_token"`
	authType    *SourceZendeskTalkUpdateAuthType `const:"oauth2.0" json:"auth_type,omitempty"`
	// Client ID
	ClientID *string `json:"client_id,omitempty"`
	// Client Secret
	ClientSecret *string `json:"client_secret,omitempty"`
}

func (s SourceZendeskTalkUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskTalkUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskTalkUpdateOAuth20) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceZendeskTalkUpdateOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceZendeskTalkUpdateOAuth20) GetAuthType() *SourceZendeskTalkUpdateAuthType {
	return SourceZendeskTalkUpdateAuthTypeOauth20.ToPointer()
}

func (o *SourceZendeskTalkUpdateOAuth20) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *SourceZendeskTalkUpdateOAuth20) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

type SourceZendeskTalkUpdateAuthenticationType string

const (
	SourceZendeskTalkUpdateAuthenticationTypeSourceZendeskTalkUpdateOAuth20  SourceZendeskTalkUpdateAuthenticationType = "source-zendesk-talk-update_OAuth2.0"
	SourceZendeskTalkUpdateAuthenticationTypeSourceZendeskTalkUpdateAPIToken SourceZendeskTalkUpdateAuthenticationType = "source-zendesk-talk-update_API Token"
)

// SourceZendeskTalkUpdateAuthentication - Zendesk service provides two authentication methods. Choose between: `OAuth2.0` or `API token`.
type SourceZendeskTalkUpdateAuthentication struct {
	SourceZendeskTalkUpdateOAuth20  *SourceZendeskTalkUpdateOAuth20
	SourceZendeskTalkUpdateAPIToken *SourceZendeskTalkUpdateAPIToken

	Type SourceZendeskTalkUpdateAuthenticationType
}

func CreateSourceZendeskTalkUpdateAuthenticationSourceZendeskTalkUpdateOAuth20(sourceZendeskTalkUpdateOAuth20 SourceZendeskTalkUpdateOAuth20) SourceZendeskTalkUpdateAuthentication {
	typ := SourceZendeskTalkUpdateAuthenticationTypeSourceZendeskTalkUpdateOAuth20

	return SourceZendeskTalkUpdateAuthentication{
		SourceZendeskTalkUpdateOAuth20: &sourceZendeskTalkUpdateOAuth20,
		Type:                           typ,
	}
}

func CreateSourceZendeskTalkUpdateAuthenticationSourceZendeskTalkUpdateAPIToken(sourceZendeskTalkUpdateAPIToken SourceZendeskTalkUpdateAPIToken) SourceZendeskTalkUpdateAuthentication {
	typ := SourceZendeskTalkUpdateAuthenticationTypeSourceZendeskTalkUpdateAPIToken

	return SourceZendeskTalkUpdateAuthentication{
		SourceZendeskTalkUpdateAPIToken: &sourceZendeskTalkUpdateAPIToken,
		Type:                            typ,
	}
}

func (u *SourceZendeskTalkUpdateAuthentication) UnmarshalJSON(data []byte) error {

	var sourceZendeskTalkUpdateAPIToken SourceZendeskTalkUpdateAPIToken = SourceZendeskTalkUpdateAPIToken{}
	if err := utils.UnmarshalJSON(data, &sourceZendeskTalkUpdateAPIToken, "", true, true); err == nil {
		u.SourceZendeskTalkUpdateAPIToken = &sourceZendeskTalkUpdateAPIToken
		u.Type = SourceZendeskTalkUpdateAuthenticationTypeSourceZendeskTalkUpdateAPIToken
		return nil
	}

	var sourceZendeskTalkUpdateOAuth20 SourceZendeskTalkUpdateOAuth20 = SourceZendeskTalkUpdateOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceZendeskTalkUpdateOAuth20, "", true, true); err == nil {
		u.SourceZendeskTalkUpdateOAuth20 = &sourceZendeskTalkUpdateOAuth20
		u.Type = SourceZendeskTalkUpdateAuthenticationTypeSourceZendeskTalkUpdateOAuth20
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceZendeskTalkUpdateAuthentication", string(data))
}

func (u SourceZendeskTalkUpdateAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceZendeskTalkUpdateOAuth20 != nil {
		return utils.MarshalJSON(u.SourceZendeskTalkUpdateOAuth20, "", true)
	}

	if u.SourceZendeskTalkUpdateAPIToken != nil {
		return utils.MarshalJSON(u.SourceZendeskTalkUpdateAPIToken, "", true)
	}

	return nil, errors.New("could not marshal union type SourceZendeskTalkUpdateAuthentication: all fields are null")
}

type SourceZendeskTalkUpdate struct {
	// Zendesk service provides two authentication methods. Choose between: `OAuth2.0` or `API token`.
	Credentials *SourceZendeskTalkUpdateAuthentication `json:"credentials,omitempty"`
	// The date from which you'd like to replicate data for Zendesk Talk API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
	// This is your Zendesk subdomain that can be found in your account URL. For example, in https://{MY_SUBDOMAIN}.zendesk.com/, where MY_SUBDOMAIN is the value of your subdomain.
	Subdomain string `json:"subdomain"`
}

func (s SourceZendeskTalkUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskTalkUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskTalkUpdate) GetCredentials() *SourceZendeskTalkUpdateAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceZendeskTalkUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceZendeskTalkUpdate) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}
