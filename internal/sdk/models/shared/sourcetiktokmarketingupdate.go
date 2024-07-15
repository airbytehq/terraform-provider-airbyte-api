// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/types"
)

type SourceTiktokMarketingUpdateSchemasAuthType string

const (
	SourceTiktokMarketingUpdateSchemasAuthTypeSandboxAccessToken SourceTiktokMarketingUpdateSchemasAuthType = "sandbox_access_token"
)

func (e SourceTiktokMarketingUpdateSchemasAuthType) ToPointer() *SourceTiktokMarketingUpdateSchemasAuthType {
	return &e
}
func (e *SourceTiktokMarketingUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sandbox_access_token":
		*e = SourceTiktokMarketingUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTiktokMarketingUpdateSchemasAuthType: %v", v)
	}
}

type SandboxAccessToken struct {
	// The long-term authorized access token.
	AccessToken string `json:"access_token"`
	// The Advertiser ID which generated for the developer's Sandbox application.
	AdvertiserID string                                      `json:"advertiser_id"`
	authType     *SourceTiktokMarketingUpdateSchemasAuthType `const:"sandbox_access_token" json:"auth_type,omitempty"`
}

func (s SandboxAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SandboxAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SandboxAccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SandboxAccessToken) GetAdvertiserID() string {
	if o == nil {
		return ""
	}
	return o.AdvertiserID
}

func (o *SandboxAccessToken) GetAuthType() *SourceTiktokMarketingUpdateSchemasAuthType {
	return SourceTiktokMarketingUpdateSchemasAuthTypeSandboxAccessToken.ToPointer()
}

type SourceTiktokMarketingUpdateAuthType string

const (
	SourceTiktokMarketingUpdateAuthTypeOauth20 SourceTiktokMarketingUpdateAuthType = "oauth2.0"
)

func (e SourceTiktokMarketingUpdateAuthType) ToPointer() *SourceTiktokMarketingUpdateAuthType {
	return &e
}
func (e *SourceTiktokMarketingUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceTiktokMarketingUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTiktokMarketingUpdateAuthType: %v", v)
	}
}

type SourceTiktokMarketingUpdateOAuth20 struct {
	// Long-term Authorized Access Token.
	AccessToken string `json:"access_token"`
	// The Advertiser ID to filter reports and streams. Let this empty to retrieve all.
	AdvertiserID *string `json:"advertiser_id,omitempty"`
	// The Developer Application App ID.
	AppID    string                               `json:"app_id"`
	authType *SourceTiktokMarketingUpdateAuthType `const:"oauth2.0" json:"auth_type,omitempty"`
	// The Developer Application Secret.
	Secret string `json:"secret"`
}

func (s SourceTiktokMarketingUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceTiktokMarketingUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceTiktokMarketingUpdateOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceTiktokMarketingUpdateOAuth20) GetAdvertiserID() *string {
	if o == nil {
		return nil
	}
	return o.AdvertiserID
}

func (o *SourceTiktokMarketingUpdateOAuth20) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *SourceTiktokMarketingUpdateOAuth20) GetAuthType() *SourceTiktokMarketingUpdateAuthType {
	return SourceTiktokMarketingUpdateAuthTypeOauth20.ToPointer()
}

func (o *SourceTiktokMarketingUpdateOAuth20) GetSecret() string {
	if o == nil {
		return ""
	}
	return o.Secret
}

type SourceTiktokMarketingUpdateAuthenticationMethodType string

const (
	SourceTiktokMarketingUpdateAuthenticationMethodTypeSourceTiktokMarketingUpdateOAuth20 SourceTiktokMarketingUpdateAuthenticationMethodType = "source-tiktok-marketing-update_OAuth2.0"
	SourceTiktokMarketingUpdateAuthenticationMethodTypeSandboxAccessToken                 SourceTiktokMarketingUpdateAuthenticationMethodType = "Sandbox Access Token"
)

// SourceTiktokMarketingUpdateAuthenticationMethod - Authentication method
type SourceTiktokMarketingUpdateAuthenticationMethod struct {
	SourceTiktokMarketingUpdateOAuth20 *SourceTiktokMarketingUpdateOAuth20
	SandboxAccessToken                 *SandboxAccessToken

	Type SourceTiktokMarketingUpdateAuthenticationMethodType
}

func CreateSourceTiktokMarketingUpdateAuthenticationMethodSourceTiktokMarketingUpdateOAuth20(sourceTiktokMarketingUpdateOAuth20 SourceTiktokMarketingUpdateOAuth20) SourceTiktokMarketingUpdateAuthenticationMethod {
	typ := SourceTiktokMarketingUpdateAuthenticationMethodTypeSourceTiktokMarketingUpdateOAuth20

	return SourceTiktokMarketingUpdateAuthenticationMethod{
		SourceTiktokMarketingUpdateOAuth20: &sourceTiktokMarketingUpdateOAuth20,
		Type:                               typ,
	}
}

func CreateSourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken(sandboxAccessToken SandboxAccessToken) SourceTiktokMarketingUpdateAuthenticationMethod {
	typ := SourceTiktokMarketingUpdateAuthenticationMethodTypeSandboxAccessToken

	return SourceTiktokMarketingUpdateAuthenticationMethod{
		SandboxAccessToken: &sandboxAccessToken,
		Type:               typ,
	}
}

func (u *SourceTiktokMarketingUpdateAuthenticationMethod) UnmarshalJSON(data []byte) error {

	var sandboxAccessToken SandboxAccessToken = SandboxAccessToken{}
	if err := utils.UnmarshalJSON(data, &sandboxAccessToken, "", true, true); err == nil {
		u.SandboxAccessToken = &sandboxAccessToken
		u.Type = SourceTiktokMarketingUpdateAuthenticationMethodTypeSandboxAccessToken
		return nil
	}

	var sourceTiktokMarketingUpdateOAuth20 SourceTiktokMarketingUpdateOAuth20 = SourceTiktokMarketingUpdateOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceTiktokMarketingUpdateOAuth20, "", true, true); err == nil {
		u.SourceTiktokMarketingUpdateOAuth20 = &sourceTiktokMarketingUpdateOAuth20
		u.Type = SourceTiktokMarketingUpdateAuthenticationMethodTypeSourceTiktokMarketingUpdateOAuth20
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceTiktokMarketingUpdateAuthenticationMethod", string(data))
}

func (u SourceTiktokMarketingUpdateAuthenticationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceTiktokMarketingUpdateOAuth20 != nil {
		return utils.MarshalJSON(u.SourceTiktokMarketingUpdateOAuth20, "", true)
	}

	if u.SandboxAccessToken != nil {
		return utils.MarshalJSON(u.SandboxAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type SourceTiktokMarketingUpdateAuthenticationMethod: all fields are null")
}

type SourceTiktokMarketingUpdate struct {
	// The attribution window in days.
	AttributionWindow *int64 `default:"3" json:"attribution_window"`
	// Authentication method
	Credentials *SourceTiktokMarketingUpdateAuthenticationMethod `json:"credentials,omitempty"`
	// The date until which you'd like to replicate data for all incremental streams, in the format YYYY-MM-DD. All data generated between start_date and this date will be replicated. Not setting this option will result in always syncing the data till the current date.
	EndDate *types.Date `json:"end_date,omitempty"`
	// Set to active if you want to include deleted data in report based streams and Ads, Ad Groups and Campaign streams.
	IncludeDeleted *bool `default:"false" json:"include_deleted"`
	// The Start Date in format: YYYY-MM-DD. Any data before this date will not be replicated. If this parameter is not set, all data will be replicated.
	StartDate *types.Date `default:"2016-09-01" json:"start_date"`
}

func (s SourceTiktokMarketingUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceTiktokMarketingUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceTiktokMarketingUpdate) GetAttributionWindow() *int64 {
	if o == nil {
		return nil
	}
	return o.AttributionWindow
}

func (o *SourceTiktokMarketingUpdate) GetCredentials() *SourceTiktokMarketingUpdateAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceTiktokMarketingUpdate) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceTiktokMarketingUpdate) GetIncludeDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeleted
}

func (o *SourceTiktokMarketingUpdate) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}
