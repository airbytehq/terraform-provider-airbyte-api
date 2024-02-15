// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type SourceMailchimpUpdateSchemasAuthType string

const (
	SourceMailchimpUpdateSchemasAuthTypeApikey SourceMailchimpUpdateSchemasAuthType = "apikey"
)

func (e SourceMailchimpUpdateSchemasAuthType) ToPointer() *SourceMailchimpUpdateSchemasAuthType {
	return &e
}

func (e *SourceMailchimpUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "apikey":
		*e = SourceMailchimpUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMailchimpUpdateSchemasAuthType: %v", v)
	}
}

type APIKey struct {
	// Mailchimp API Key. See the <a href="https://docs.airbyte.com/integrations/sources/mailchimp">docs</a> for information on how to generate this key.
	Apikey   string                               `json:"apikey"`
	authType SourceMailchimpUpdateSchemasAuthType `const:"apikey" json:"auth_type"`
}

func (a APIKey) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIKey) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *APIKey) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

func (o *APIKey) GetAuthType() SourceMailchimpUpdateSchemasAuthType {
	return SourceMailchimpUpdateSchemasAuthTypeApikey
}

type SourceMailchimpUpdateAuthType string

const (
	SourceMailchimpUpdateAuthTypeOauth20 SourceMailchimpUpdateAuthType = "oauth2.0"
)

func (e SourceMailchimpUpdateAuthType) ToPointer() *SourceMailchimpUpdateAuthType {
	return &e
}

func (e *SourceMailchimpUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceMailchimpUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMailchimpUpdateAuthType: %v", v)
	}
}

type SourceMailchimpUpdateOAuth20 struct {
	// An access token generated using the above client ID and secret.
	AccessToken string                        `json:"access_token"`
	authType    SourceMailchimpUpdateAuthType `const:"oauth2.0" json:"auth_type"`
	// The Client ID of your OAuth application.
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your OAuth application.
	ClientSecret *string `json:"client_secret,omitempty"`
}

func (s SourceMailchimpUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMailchimpUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMailchimpUpdateOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceMailchimpUpdateOAuth20) GetAuthType() SourceMailchimpUpdateAuthType {
	return SourceMailchimpUpdateAuthTypeOauth20
}

func (o *SourceMailchimpUpdateOAuth20) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *SourceMailchimpUpdateOAuth20) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

type SourceMailchimpUpdateAuthenticationType string

const (
	SourceMailchimpUpdateAuthenticationTypeSourceMailchimpUpdateOAuth20 SourceMailchimpUpdateAuthenticationType = "source-mailchimp-update_OAuth2.0"
	SourceMailchimpUpdateAuthenticationTypeAPIKey                       SourceMailchimpUpdateAuthenticationType = "API Key"
)

type SourceMailchimpUpdateAuthentication struct {
	SourceMailchimpUpdateOAuth20 *SourceMailchimpUpdateOAuth20
	APIKey                       *APIKey

	Type SourceMailchimpUpdateAuthenticationType
}

func CreateSourceMailchimpUpdateAuthenticationSourceMailchimpUpdateOAuth20(sourceMailchimpUpdateOAuth20 SourceMailchimpUpdateOAuth20) SourceMailchimpUpdateAuthentication {
	typ := SourceMailchimpUpdateAuthenticationTypeSourceMailchimpUpdateOAuth20

	return SourceMailchimpUpdateAuthentication{
		SourceMailchimpUpdateOAuth20: &sourceMailchimpUpdateOAuth20,
		Type:                         typ,
	}
}

func CreateSourceMailchimpUpdateAuthenticationAPIKey(apiKey APIKey) SourceMailchimpUpdateAuthentication {
	typ := SourceMailchimpUpdateAuthenticationTypeAPIKey

	return SourceMailchimpUpdateAuthentication{
		APIKey: &apiKey,
		Type:   typ,
	}
}

func (u *SourceMailchimpUpdateAuthentication) UnmarshalJSON(data []byte) error {

	apiKey := new(APIKey)
	if err := utils.UnmarshalJSON(data, &apiKey, "", true, true); err == nil {
		u.APIKey = apiKey
		u.Type = SourceMailchimpUpdateAuthenticationTypeAPIKey
		return nil
	}

	sourceMailchimpUpdateOAuth20 := new(SourceMailchimpUpdateOAuth20)
	if err := utils.UnmarshalJSON(data, &sourceMailchimpUpdateOAuth20, "", true, true); err == nil {
		u.SourceMailchimpUpdateOAuth20 = sourceMailchimpUpdateOAuth20
		u.Type = SourceMailchimpUpdateAuthenticationTypeSourceMailchimpUpdateOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMailchimpUpdateAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceMailchimpUpdateOAuth20 != nil {
		return utils.MarshalJSON(u.SourceMailchimpUpdateOAuth20, "", true)
	}

	if u.APIKey != nil {
		return utils.MarshalJSON(u.APIKey, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceMailchimpUpdate struct {
	CampaignID  *string                              `json:"campaign_id,omitempty"`
	Credentials *SourceMailchimpUpdateAuthentication `json:"credentials,omitempty"`
	// The date from which you want to start syncing data for Incremental streams. Only records that have been created or modified since this date will be synced. If left blank, all data will by synced.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceMailchimpUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMailchimpUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMailchimpUpdate) GetCampaignID() *string {
	if o == nil {
		return nil
	}
	return o.CampaignID
}

func (o *SourceMailchimpUpdate) GetCredentials() *SourceMailchimpUpdateAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceMailchimpUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
