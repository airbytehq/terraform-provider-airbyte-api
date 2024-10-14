// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceLinkedinPagesUpdateSchemasAuthMethod string

const (
	SourceLinkedinPagesUpdateSchemasAuthMethodAccessToken SourceLinkedinPagesUpdateSchemasAuthMethod = "access_token"
)

func (e SourceLinkedinPagesUpdateSchemasAuthMethod) ToPointer() *SourceLinkedinPagesUpdateSchemasAuthMethod {
	return &e
}
func (e *SourceLinkedinPagesUpdateSchemasAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceLinkedinPagesUpdateSchemasAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesUpdateSchemasAuthMethod: %v", v)
	}
}

type SourceLinkedinPagesUpdateAccessToken struct {
	// The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.
	AccessToken string                                      `json:"access_token"`
	authMethod  *SourceLinkedinPagesUpdateSchemasAuthMethod `const:"access_token" json:"auth_method,omitempty"`
}

func (s SourceLinkedinPagesUpdateAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinPagesUpdateAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinPagesUpdateAccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceLinkedinPagesUpdateAccessToken) GetAuthMethod() *SourceLinkedinPagesUpdateSchemasAuthMethod {
	return SourceLinkedinPagesUpdateSchemasAuthMethodAccessToken.ToPointer()
}

type SourceLinkedinPagesUpdateAuthMethod string

const (
	SourceLinkedinPagesUpdateAuthMethodOAuth20 SourceLinkedinPagesUpdateAuthMethod = "oAuth2.0"
)

func (e SourceLinkedinPagesUpdateAuthMethod) ToPointer() *SourceLinkedinPagesUpdateAuthMethod {
	return &e
}
func (e *SourceLinkedinPagesUpdateAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oAuth2.0":
		*e = SourceLinkedinPagesUpdateAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesUpdateAuthMethod: %v", v)
	}
}

type SourceLinkedinPagesUpdateOAuth20 struct {
	authMethod *SourceLinkedinPagesUpdateAuthMethod `const:"oAuth2.0" json:"auth_method,omitempty"`
	// The client ID of the LinkedIn developer application.
	ClientID string `json:"client_id"`
	// The client secret of the LinkedIn developer application.
	ClientSecret string `json:"client_secret"`
	// The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceLinkedinPagesUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinPagesUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinPagesUpdateOAuth20) GetAuthMethod() *SourceLinkedinPagesUpdateAuthMethod {
	return SourceLinkedinPagesUpdateAuthMethodOAuth20.ToPointer()
}

func (o *SourceLinkedinPagesUpdateOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceLinkedinPagesUpdateOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceLinkedinPagesUpdateOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceLinkedinPagesUpdateAuthenticationType string

const (
	SourceLinkedinPagesUpdateAuthenticationTypeSourceLinkedinPagesUpdateOAuth20     SourceLinkedinPagesUpdateAuthenticationType = "source-linkedin-pages-update_OAuth2.0"
	SourceLinkedinPagesUpdateAuthenticationTypeSourceLinkedinPagesUpdateAccessToken SourceLinkedinPagesUpdateAuthenticationType = "source-linkedin-pages-update_Access token"
)

type SourceLinkedinPagesUpdateAuthentication struct {
	SourceLinkedinPagesUpdateOAuth20     *SourceLinkedinPagesUpdateOAuth20
	SourceLinkedinPagesUpdateAccessToken *SourceLinkedinPagesUpdateAccessToken

	Type SourceLinkedinPagesUpdateAuthenticationType
}

func CreateSourceLinkedinPagesUpdateAuthenticationSourceLinkedinPagesUpdateOAuth20(sourceLinkedinPagesUpdateOAuth20 SourceLinkedinPagesUpdateOAuth20) SourceLinkedinPagesUpdateAuthentication {
	typ := SourceLinkedinPagesUpdateAuthenticationTypeSourceLinkedinPagesUpdateOAuth20

	return SourceLinkedinPagesUpdateAuthentication{
		SourceLinkedinPagesUpdateOAuth20: &sourceLinkedinPagesUpdateOAuth20,
		Type:                             typ,
	}
}

func CreateSourceLinkedinPagesUpdateAuthenticationSourceLinkedinPagesUpdateAccessToken(sourceLinkedinPagesUpdateAccessToken SourceLinkedinPagesUpdateAccessToken) SourceLinkedinPagesUpdateAuthentication {
	typ := SourceLinkedinPagesUpdateAuthenticationTypeSourceLinkedinPagesUpdateAccessToken

	return SourceLinkedinPagesUpdateAuthentication{
		SourceLinkedinPagesUpdateAccessToken: &sourceLinkedinPagesUpdateAccessToken,
		Type:                                 typ,
	}
}

func (u *SourceLinkedinPagesUpdateAuthentication) UnmarshalJSON(data []byte) error {

	var sourceLinkedinPagesUpdateAccessToken SourceLinkedinPagesUpdateAccessToken = SourceLinkedinPagesUpdateAccessToken{}
	if err := utils.UnmarshalJSON(data, &sourceLinkedinPagesUpdateAccessToken, "", true, true); err == nil {
		u.SourceLinkedinPagesUpdateAccessToken = &sourceLinkedinPagesUpdateAccessToken
		u.Type = SourceLinkedinPagesUpdateAuthenticationTypeSourceLinkedinPagesUpdateAccessToken
		return nil
	}

	var sourceLinkedinPagesUpdateOAuth20 SourceLinkedinPagesUpdateOAuth20 = SourceLinkedinPagesUpdateOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceLinkedinPagesUpdateOAuth20, "", true, true); err == nil {
		u.SourceLinkedinPagesUpdateOAuth20 = &sourceLinkedinPagesUpdateOAuth20
		u.Type = SourceLinkedinPagesUpdateAuthenticationTypeSourceLinkedinPagesUpdateOAuth20
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceLinkedinPagesUpdateAuthentication", string(data))
}

func (u SourceLinkedinPagesUpdateAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceLinkedinPagesUpdateOAuth20 != nil {
		return utils.MarshalJSON(u.SourceLinkedinPagesUpdateOAuth20, "", true)
	}

	if u.SourceLinkedinPagesUpdateAccessToken != nil {
		return utils.MarshalJSON(u.SourceLinkedinPagesUpdateAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type SourceLinkedinPagesUpdateAuthentication: all fields are null")
}

// TimeGranularityType - Granularity of the statistics for metrics per time period. Must be either "DAY" or "MONTH"
type TimeGranularityType string

const (
	TimeGranularityTypeDay   TimeGranularityType = "DAY"
	TimeGranularityTypeMonth TimeGranularityType = "MONTH"
)

func (e TimeGranularityType) ToPointer() *TimeGranularityType {
	return &e
}
func (e *TimeGranularityType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DAY":
		fallthrough
	case "MONTH":
		*e = TimeGranularityType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TimeGranularityType: %v", v)
	}
}

type SourceLinkedinPagesUpdate struct {
	Credentials *SourceLinkedinPagesUpdateAuthentication `json:"credentials,omitempty"`
	// Specify the Organization ID
	OrgID string `json:"org_id"`
	// Start date for getting metrics per time period. Must be atmost 12 months before the request date (UTC) and atleast 2 days prior to the request date (UTC). See https://bit.ly/linkedin-pages-date-rules {{"{{"}} "\n" }} {{"{{"}} response.errorDetails }}
	StartDate *time.Time `default:"2023-01-01T00:00:00Z" json:"start_date"`
	// Granularity of the statistics for metrics per time period. Must be either "DAY" or "MONTH"
	TimeGranularityType *TimeGranularityType `default:"DAY" json:"time_granularity_type"`
}

func (s SourceLinkedinPagesUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinPagesUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinPagesUpdate) GetCredentials() *SourceLinkedinPagesUpdateAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceLinkedinPagesUpdate) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *SourceLinkedinPagesUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceLinkedinPagesUpdate) GetTimeGranularityType() *TimeGranularityType {
	if o == nil {
		return nil
	}
	return o.TimeGranularityType
}
