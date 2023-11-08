// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceGoogleSearchConsoleSchemasAuthType string

const (
	SourceGoogleSearchConsoleSchemasAuthTypeService SourceGoogleSearchConsoleSchemasAuthType = "Service"
)

func (e SourceGoogleSearchConsoleSchemasAuthType) ToPointer() *SourceGoogleSearchConsoleSchemasAuthType {
	return &e
}

func (e *SourceGoogleSearchConsoleSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Service":
		*e = SourceGoogleSearchConsoleSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleSchemasAuthType: %v", v)
	}
}

type SourceGoogleSearchConsoleServiceAccountKeyAuthentication struct {
	authType SourceGoogleSearchConsoleSchemasAuthType `const:"Service" json:"auth_type"`
	// The email of the user which has permissions to access the Google Workspace Admin APIs.
	Email string `json:"email"`
	// The JSON key of the service account to use for authorization. Read more <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys">here</a>.
	ServiceAccountInfo string `json:"service_account_info"`
}

func (s SourceGoogleSearchConsoleServiceAccountKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleSearchConsoleServiceAccountKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleSearchConsoleServiceAccountKeyAuthentication) GetAuthType() SourceGoogleSearchConsoleSchemasAuthType {
	return SourceGoogleSearchConsoleSchemasAuthTypeService
}

func (o *SourceGoogleSearchConsoleServiceAccountKeyAuthentication) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *SourceGoogleSearchConsoleServiceAccountKeyAuthentication) GetServiceAccountInfo() string {
	if o == nil {
		return ""
	}
	return o.ServiceAccountInfo
}

type SourceGoogleSearchConsoleAuthType string

const (
	SourceGoogleSearchConsoleAuthTypeClient SourceGoogleSearchConsoleAuthType = "Client"
)

func (e SourceGoogleSearchConsoleAuthType) ToPointer() *SourceGoogleSearchConsoleAuthType {
	return &e
}

func (e *SourceGoogleSearchConsoleAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceGoogleSearchConsoleAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleAuthType: %v", v)
	}
}

type SourceGoogleSearchConsoleOAuth struct {
	// Access token for making authenticated requests. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	AccessToken *string                           `json:"access_token,omitempty"`
	authType    SourceGoogleSearchConsoleAuthType `const:"Client" json:"auth_type"`
	// The client ID of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	ClientID string `json:"client_id"`
	// The client secret of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	ClientSecret string `json:"client_secret"`
	// The token for obtaining a new access token. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceGoogleSearchConsoleOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleSearchConsoleOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleSearchConsoleOAuth) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *SourceGoogleSearchConsoleOAuth) GetAuthType() SourceGoogleSearchConsoleAuthType {
	return SourceGoogleSearchConsoleAuthTypeClient
}

func (o *SourceGoogleSearchConsoleOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceGoogleSearchConsoleOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceGoogleSearchConsoleOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceGoogleSearchConsoleAuthenticationTypeType string

const (
	SourceGoogleSearchConsoleAuthenticationTypeTypeSourceGoogleSearchConsoleOAuth                           SourceGoogleSearchConsoleAuthenticationTypeType = "source-google-search-console_OAuth"
	SourceGoogleSearchConsoleAuthenticationTypeTypeSourceGoogleSearchConsoleServiceAccountKeyAuthentication SourceGoogleSearchConsoleAuthenticationTypeType = "source-google-search-console_Service Account Key Authentication"
)

type SourceGoogleSearchConsoleAuthenticationType struct {
	SourceGoogleSearchConsoleOAuth                           *SourceGoogleSearchConsoleOAuth
	SourceGoogleSearchConsoleServiceAccountKeyAuthentication *SourceGoogleSearchConsoleServiceAccountKeyAuthentication

	Type SourceGoogleSearchConsoleAuthenticationTypeType
}

func CreateSourceGoogleSearchConsoleAuthenticationTypeSourceGoogleSearchConsoleOAuth(sourceGoogleSearchConsoleOAuth SourceGoogleSearchConsoleOAuth) SourceGoogleSearchConsoleAuthenticationType {
	typ := SourceGoogleSearchConsoleAuthenticationTypeTypeSourceGoogleSearchConsoleOAuth

	return SourceGoogleSearchConsoleAuthenticationType{
		SourceGoogleSearchConsoleOAuth: &sourceGoogleSearchConsoleOAuth,
		Type:                           typ,
	}
}

func CreateSourceGoogleSearchConsoleAuthenticationTypeSourceGoogleSearchConsoleServiceAccountKeyAuthentication(sourceGoogleSearchConsoleServiceAccountKeyAuthentication SourceGoogleSearchConsoleServiceAccountKeyAuthentication) SourceGoogleSearchConsoleAuthenticationType {
	typ := SourceGoogleSearchConsoleAuthenticationTypeTypeSourceGoogleSearchConsoleServiceAccountKeyAuthentication

	return SourceGoogleSearchConsoleAuthenticationType{
		SourceGoogleSearchConsoleServiceAccountKeyAuthentication: &sourceGoogleSearchConsoleServiceAccountKeyAuthentication,
		Type: typ,
	}
}

func (u *SourceGoogleSearchConsoleAuthenticationType) UnmarshalJSON(data []byte) error {

	sourceGoogleSearchConsoleServiceAccountKeyAuthentication := new(SourceGoogleSearchConsoleServiceAccountKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &sourceGoogleSearchConsoleServiceAccountKeyAuthentication, "", true, true); err == nil {
		u.SourceGoogleSearchConsoleServiceAccountKeyAuthentication = sourceGoogleSearchConsoleServiceAccountKeyAuthentication
		u.Type = SourceGoogleSearchConsoleAuthenticationTypeTypeSourceGoogleSearchConsoleServiceAccountKeyAuthentication
		return nil
	}

	sourceGoogleSearchConsoleOAuth := new(SourceGoogleSearchConsoleOAuth)
	if err := utils.UnmarshalJSON(data, &sourceGoogleSearchConsoleOAuth, "", true, true); err == nil {
		u.SourceGoogleSearchConsoleOAuth = sourceGoogleSearchConsoleOAuth
		u.Type = SourceGoogleSearchConsoleAuthenticationTypeTypeSourceGoogleSearchConsoleOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceGoogleSearchConsoleAuthenticationType) MarshalJSON() ([]byte, error) {
	if u.SourceGoogleSearchConsoleOAuth != nil {
		return utils.MarshalJSON(u.SourceGoogleSearchConsoleOAuth, "", true)
	}

	if u.SourceGoogleSearchConsoleServiceAccountKeyAuthentication != nil {
		return utils.MarshalJSON(u.SourceGoogleSearchConsoleServiceAccountKeyAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SourceGoogleSearchConsoleValidEnums - An enumeration of dimensions.
type SourceGoogleSearchConsoleValidEnums string

const (
	SourceGoogleSearchConsoleValidEnumsCountry SourceGoogleSearchConsoleValidEnums = "country"
	SourceGoogleSearchConsoleValidEnumsDate    SourceGoogleSearchConsoleValidEnums = "date"
	SourceGoogleSearchConsoleValidEnumsDevice  SourceGoogleSearchConsoleValidEnums = "device"
	SourceGoogleSearchConsoleValidEnumsPage    SourceGoogleSearchConsoleValidEnums = "page"
	SourceGoogleSearchConsoleValidEnumsQuery   SourceGoogleSearchConsoleValidEnums = "query"
)

func (e SourceGoogleSearchConsoleValidEnums) ToPointer() *SourceGoogleSearchConsoleValidEnums {
	return &e
}

func (e *SourceGoogleSearchConsoleValidEnums) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "country":
		fallthrough
	case "date":
		fallthrough
	case "device":
		fallthrough
	case "page":
		fallthrough
	case "query":
		*e = SourceGoogleSearchConsoleValidEnums(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleValidEnums: %v", v)
	}
}

type SourceGoogleSearchConsoleCustomReportConfig struct {
	// A list of dimensions (country, date, device, page, query)
	Dimensions []SourceGoogleSearchConsoleValidEnums `json:"dimensions"`
	// The name of the custom report, this name would be used as stream name
	Name string `json:"name"`
}

func (o *SourceGoogleSearchConsoleCustomReportConfig) GetDimensions() []SourceGoogleSearchConsoleValidEnums {
	if o == nil {
		return []SourceGoogleSearchConsoleValidEnums{}
	}
	return o.Dimensions
}

func (o *SourceGoogleSearchConsoleCustomReportConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// SourceGoogleSearchConsoleDataFreshness - If set to 'final', the returned data will include only finalized, stable data. If set to 'all', fresh data will be included. When using Incremental sync mode, we do not recommend setting this parameter to 'all' as it may cause data loss. More information can be found in our <a href='https://docs.airbyte.com/integrations/source/google-search-console'>full documentation</a>.
type SourceGoogleSearchConsoleDataFreshness string

const (
	SourceGoogleSearchConsoleDataFreshnessFinal SourceGoogleSearchConsoleDataFreshness = "final"
	SourceGoogleSearchConsoleDataFreshnessAll   SourceGoogleSearchConsoleDataFreshness = "all"
)

func (e SourceGoogleSearchConsoleDataFreshness) ToPointer() *SourceGoogleSearchConsoleDataFreshness {
	return &e
}

func (e *SourceGoogleSearchConsoleDataFreshness) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "final":
		fallthrough
	case "all":
		*e = SourceGoogleSearchConsoleDataFreshness(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleDataFreshness: %v", v)
	}
}

type GoogleSearchConsole string

const (
	GoogleSearchConsoleGoogleSearchConsole GoogleSearchConsole = "google-search-console"
)

func (e GoogleSearchConsole) ToPointer() *GoogleSearchConsole {
	return &e
}

func (e *GoogleSearchConsole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google-search-console":
		*e = GoogleSearchConsole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GoogleSearchConsole: %v", v)
	}
}

type SourceGoogleSearchConsole struct {
	Authorization SourceGoogleSearchConsoleAuthenticationType `json:"authorization"`
	// (DEPRCATED) A JSON array describing the custom reports you want to sync from Google Search Console. See our <a href='https://docs.airbyte.com/integrations/sources/google-search-console'>documentation</a> for more information on formulating custom reports.
	CustomReports *string `json:"custom_reports,omitempty"`
	// You can add your Custom Analytics report by creating one.
	CustomReportsArray []SourceGoogleSearchConsoleCustomReportConfig `json:"custom_reports_array,omitempty"`
	// If set to 'final', the returned data will include only finalized, stable data. If set to 'all', fresh data will be included. When using Incremental sync mode, we do not recommend setting this parameter to 'all' as it may cause data loss. More information can be found in our <a href='https://docs.airbyte.com/integrations/source/google-search-console'>full documentation</a>.
	DataState *SourceGoogleSearchConsoleDataFreshness `default:"final" json:"data_state"`
	// UTC date in the format YYYY-MM-DD. Any data created after this date will not be replicated. Must be greater or equal to the start date field. Leaving this field blank will replicate all data from the start date onward.
	EndDate *types.Date `json:"end_date,omitempty"`
	// The URLs of the website property attached to your GSC account. Learn more about properties <a href="https://support.google.com/webmasters/answer/34592?hl=en">here</a>.
	SiteUrls   []string            `json:"site_urls"`
	sourceType GoogleSearchConsole `const:"google-search-console" json:"sourceType"`
	// UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated.
	StartDate *types.Date `default:"2021-01-01" json:"start_date"`
}

func (s SourceGoogleSearchConsole) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleSearchConsole) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleSearchConsole) GetAuthorization() SourceGoogleSearchConsoleAuthenticationType {
	if o == nil {
		return SourceGoogleSearchConsoleAuthenticationType{}
	}
	return o.Authorization
}

func (o *SourceGoogleSearchConsole) GetCustomReports() *string {
	if o == nil {
		return nil
	}
	return o.CustomReports
}

func (o *SourceGoogleSearchConsole) GetCustomReportsArray() []SourceGoogleSearchConsoleCustomReportConfig {
	if o == nil {
		return nil
	}
	return o.CustomReportsArray
}

func (o *SourceGoogleSearchConsole) GetDataState() *SourceGoogleSearchConsoleDataFreshness {
	if o == nil {
		return nil
	}
	return o.DataState
}

func (o *SourceGoogleSearchConsole) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceGoogleSearchConsole) GetSiteUrls() []string {
	if o == nil {
		return []string{}
	}
	return o.SiteUrls
}

func (o *SourceGoogleSearchConsole) GetSourceType() GoogleSearchConsole {
	return GoogleSearchConsoleGoogleSearchConsole
}

func (o *SourceGoogleSearchConsole) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}
