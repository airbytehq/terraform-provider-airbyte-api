// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceLinkedinPagesSchemasAuthMethod string

const (
	SourceLinkedinPagesSchemasAuthMethodAccessToken SourceLinkedinPagesSchemasAuthMethod = "access_token"
)

func (e SourceLinkedinPagesSchemasAuthMethod) ToPointer() *SourceLinkedinPagesSchemasAuthMethod {
	return &e
}
func (e *SourceLinkedinPagesSchemasAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceLinkedinPagesSchemasAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesSchemasAuthMethod: %v", v)
	}
}

type SourceLinkedinPagesAccessToken struct {
	// The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.
	AccessToken string                                `json:"access_token"`
	authMethod  *SourceLinkedinPagesSchemasAuthMethod `const:"access_token" json:"auth_method,omitempty"`
}

func (s SourceLinkedinPagesAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinPagesAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinPagesAccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceLinkedinPagesAccessToken) GetAuthMethod() *SourceLinkedinPagesSchemasAuthMethod {
	return SourceLinkedinPagesSchemasAuthMethodAccessToken.ToPointer()
}

type SourceLinkedinPagesAuthMethod string

const (
	SourceLinkedinPagesAuthMethodOAuth20 SourceLinkedinPagesAuthMethod = "oAuth2.0"
)

func (e SourceLinkedinPagesAuthMethod) ToPointer() *SourceLinkedinPagesAuthMethod {
	return &e
}
func (e *SourceLinkedinPagesAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oAuth2.0":
		*e = SourceLinkedinPagesAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesAuthMethod: %v", v)
	}
}

type SourceLinkedinPagesOAuth20 struct {
	authMethod *SourceLinkedinPagesAuthMethod `const:"oAuth2.0" json:"auth_method,omitempty"`
	// The client ID of the LinkedIn developer application.
	ClientID string `json:"client_id"`
	// The client secret of the LinkedIn developer application.
	ClientSecret string `json:"client_secret"`
	// The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceLinkedinPagesOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinPagesOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinPagesOAuth20) GetAuthMethod() *SourceLinkedinPagesAuthMethod {
	return SourceLinkedinPagesAuthMethodOAuth20.ToPointer()
}

func (o *SourceLinkedinPagesOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceLinkedinPagesOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceLinkedinPagesOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceLinkedinPagesAuthenticationType string

const (
	SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesOAuth20     SourceLinkedinPagesAuthenticationType = "source-linkedin-pages_OAuth2.0"
	SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAccessToken SourceLinkedinPagesAuthenticationType = "source-linkedin-pages_Access token"
)

type SourceLinkedinPagesAuthentication struct {
	SourceLinkedinPagesOAuth20     *SourceLinkedinPagesOAuth20
	SourceLinkedinPagesAccessToken *SourceLinkedinPagesAccessToken

	Type SourceLinkedinPagesAuthenticationType
}

func CreateSourceLinkedinPagesAuthenticationSourceLinkedinPagesOAuth20(sourceLinkedinPagesOAuth20 SourceLinkedinPagesOAuth20) SourceLinkedinPagesAuthentication {
	typ := SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesOAuth20

	return SourceLinkedinPagesAuthentication{
		SourceLinkedinPagesOAuth20: &sourceLinkedinPagesOAuth20,
		Type:                       typ,
	}
}

func CreateSourceLinkedinPagesAuthenticationSourceLinkedinPagesAccessToken(sourceLinkedinPagesAccessToken SourceLinkedinPagesAccessToken) SourceLinkedinPagesAuthentication {
	typ := SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAccessToken

	return SourceLinkedinPagesAuthentication{
		SourceLinkedinPagesAccessToken: &sourceLinkedinPagesAccessToken,
		Type:                           typ,
	}
}

func (u *SourceLinkedinPagesAuthentication) UnmarshalJSON(data []byte) error {

	var sourceLinkedinPagesAccessToken SourceLinkedinPagesAccessToken = SourceLinkedinPagesAccessToken{}
	if err := utils.UnmarshalJSON(data, &sourceLinkedinPagesAccessToken, "", true, true); err == nil {
		u.SourceLinkedinPagesAccessToken = &sourceLinkedinPagesAccessToken
		u.Type = SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAccessToken
		return nil
	}

	var sourceLinkedinPagesOAuth20 SourceLinkedinPagesOAuth20 = SourceLinkedinPagesOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceLinkedinPagesOAuth20, "", true, true); err == nil {
		u.SourceLinkedinPagesOAuth20 = &sourceLinkedinPagesOAuth20
		u.Type = SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesOAuth20
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceLinkedinPagesAuthentication", string(data))
}

func (u SourceLinkedinPagesAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceLinkedinPagesOAuth20 != nil {
		return utils.MarshalJSON(u.SourceLinkedinPagesOAuth20, "", true)
	}

	if u.SourceLinkedinPagesAccessToken != nil {
		return utils.MarshalJSON(u.SourceLinkedinPagesAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type SourceLinkedinPagesAuthentication: all fields are null")
}

type LinkedinPages string

const (
	LinkedinPagesLinkedinPages LinkedinPages = "linkedin-pages"
)

func (e LinkedinPages) ToPointer() *LinkedinPages {
	return &e
}
func (e *LinkedinPages) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "linkedin-pages":
		*e = LinkedinPages(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LinkedinPages: %v", v)
	}
}

// SourceLinkedinPagesTimeGranularityType - Granularity of the statistics for metrics per time period. Must be either "DAY" or "MONTH"
type SourceLinkedinPagesTimeGranularityType string

const (
	SourceLinkedinPagesTimeGranularityTypeDay   SourceLinkedinPagesTimeGranularityType = "DAY"
	SourceLinkedinPagesTimeGranularityTypeMonth SourceLinkedinPagesTimeGranularityType = "MONTH"
)

func (e SourceLinkedinPagesTimeGranularityType) ToPointer() *SourceLinkedinPagesTimeGranularityType {
	return &e
}
func (e *SourceLinkedinPagesTimeGranularityType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DAY":
		fallthrough
	case "MONTH":
		*e = SourceLinkedinPagesTimeGranularityType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesTimeGranularityType: %v", v)
	}
}

type SourceLinkedinPages struct {
	Credentials *SourceLinkedinPagesAuthentication `json:"credentials,omitempty"`
	// Specify the Organization ID
	OrgID      string        `json:"org_id"`
	sourceType LinkedinPages `const:"linkedin-pages" json:"sourceType"`
	// Start date for getting metrics per time period. Must be atmost 12 months before the request date (UTC) and atleast 2 days prior to the request date (UTC). See https://bit.ly/linkedin-pages-date-rules {{"{{"}} "\n" }} {{"{{"}} response.errorDetails }}
	StartDate *time.Time `default:"2023-01-01T00:00:00Z" json:"start_date"`
	// Granularity of the statistics for metrics per time period. Must be either "DAY" or "MONTH"
	TimeGranularityType *SourceLinkedinPagesTimeGranularityType `default:"DAY" json:"time_granularity_type"`
}

func (s SourceLinkedinPages) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinPages) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinPages) GetCredentials() *SourceLinkedinPagesAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceLinkedinPages) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *SourceLinkedinPages) GetSourceType() LinkedinPages {
	return LinkedinPagesLinkedinPages
}

func (o *SourceLinkedinPages) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceLinkedinPages) GetTimeGranularityType() *SourceLinkedinPagesTimeGranularityType {
	if o == nil {
		return nil
	}
	return o.TimeGranularityType
}
