// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/types"
)

type SourceAmazonAdsAuthType string

const (
	SourceAmazonAdsAuthTypeOauth20 SourceAmazonAdsAuthType = "oauth2.0"
)

func (e SourceAmazonAdsAuthType) ToPointer() *SourceAmazonAdsAuthType {
	return &e
}
func (e *SourceAmazonAdsAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceAmazonAdsAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsAuthType: %v", v)
	}
}

// SourceAmazonAdsRegion - Region to pull data from (EU/NA/FE). See <a href="https://advertising.amazon.com/API/docs/en-us/info/api-overview#api-endpoints">docs</a> for more details.
type SourceAmazonAdsRegion string

const (
	SourceAmazonAdsRegionNa SourceAmazonAdsRegion = "NA"
	SourceAmazonAdsRegionEu SourceAmazonAdsRegion = "EU"
	SourceAmazonAdsRegionFe SourceAmazonAdsRegion = "FE"
)

func (e SourceAmazonAdsRegion) ToPointer() *SourceAmazonAdsRegion {
	return &e
}
func (e *SourceAmazonAdsRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NA":
		fallthrough
	case "EU":
		fallthrough
	case "FE":
		*e = SourceAmazonAdsRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsRegion: %v", v)
	}
}

type SourceAmazonAdsReportRecordTypes string

const (
	SourceAmazonAdsReportRecordTypesAdGroups      SourceAmazonAdsReportRecordTypes = "adGroups"
	SourceAmazonAdsReportRecordTypesAsins         SourceAmazonAdsReportRecordTypes = "asins"
	SourceAmazonAdsReportRecordTypesAsinsKeywords SourceAmazonAdsReportRecordTypes = "asins_keywords"
	SourceAmazonAdsReportRecordTypesAsinsTargets  SourceAmazonAdsReportRecordTypes = "asins_targets"
	SourceAmazonAdsReportRecordTypesCampaigns     SourceAmazonAdsReportRecordTypes = "campaigns"
	SourceAmazonAdsReportRecordTypesKeywords      SourceAmazonAdsReportRecordTypes = "keywords"
	SourceAmazonAdsReportRecordTypesProductAds    SourceAmazonAdsReportRecordTypes = "productAds"
	SourceAmazonAdsReportRecordTypesTargets       SourceAmazonAdsReportRecordTypes = "targets"
)

func (e SourceAmazonAdsReportRecordTypes) ToPointer() *SourceAmazonAdsReportRecordTypes {
	return &e
}
func (e *SourceAmazonAdsReportRecordTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "adGroups":
		fallthrough
	case "asins":
		fallthrough
	case "asins_keywords":
		fallthrough
	case "asins_targets":
		fallthrough
	case "campaigns":
		fallthrough
	case "keywords":
		fallthrough
	case "productAds":
		fallthrough
	case "targets":
		*e = SourceAmazonAdsReportRecordTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsReportRecordTypes: %v", v)
	}
}

type AmazonAds string

const (
	AmazonAdsAmazonAds AmazonAds = "amazon-ads"
)

func (e AmazonAds) ToPointer() *AmazonAds {
	return &e
}
func (e *AmazonAds) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "amazon-ads":
		*e = AmazonAds(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AmazonAds: %v", v)
	}
}

type SourceAmazonAdsStateFilter string

const (
	SourceAmazonAdsStateFilterEnabled  SourceAmazonAdsStateFilter = "enabled"
	SourceAmazonAdsStateFilterPaused   SourceAmazonAdsStateFilter = "paused"
	SourceAmazonAdsStateFilterArchived SourceAmazonAdsStateFilter = "archived"
)

func (e SourceAmazonAdsStateFilter) ToPointer() *SourceAmazonAdsStateFilter {
	return &e
}
func (e *SourceAmazonAdsStateFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enabled":
		fallthrough
	case "paused":
		fallthrough
	case "archived":
		*e = SourceAmazonAdsStateFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsStateFilter: %v", v)
	}
}

type SourceAmazonAds struct {
	authType *SourceAmazonAdsAuthType `const:"oauth2.0" json:"auth_type,omitempty"`
	// The client ID of your Amazon Ads developer application. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens#retrieve-your-client-id-and-client-secret">docs</a> for more information.
	ClientID string `json:"client_id"`
	// The client secret of your Amazon Ads developer application. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens#retrieve-your-client-id-and-client-secret">docs</a> for more information.
	ClientSecret string `json:"client_secret"`
	// The amount of days to go back in time to get the updated data from Amazon Ads
	LookBackWindow *int64 `default:"3" json:"look_back_window"`
	// Marketplace IDs you want to fetch data for. Note: If Profile IDs are also selected, profiles will be selected if they match the Profile ID OR the Marketplace ID.
	MarketplaceIds []string `json:"marketplace_ids,omitempty"`
	// Profile IDs you want to fetch data for. The Amazon Ads source connector supports only profiles with seller and vendor type, profiles with agency type will be ignored. See <a href="https://advertising.amazon.com/API/docs/en-us/concepts/authorization/profiles">docs</a> for more details. Note: If Marketplace IDs are also selected, profiles will be selected if they match the Profile ID OR the Marketplace ID.
	Profiles []int64 `json:"profiles,omitempty"`
	// Amazon Ads refresh token. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens">docs</a> for more information on how to obtain this token.
	RefreshToken string `json:"refresh_token"`
	// Region to pull data from (EU/NA/FE). See <a href="https://advertising.amazon.com/API/docs/en-us/info/api-overview#api-endpoints">docs</a> for more details.
	Region *SourceAmazonAdsRegion `default:"NA" json:"region"`
	// Optional configuration which accepts an array of string of record types. Leave blank for default behaviour to pull all report types. Use this config option only if you want to pull specific report type(s). See <a href="https://advertising.amazon.com/API/docs/en-us/reporting/v2/report-types">docs</a> for more details
	ReportRecordTypes []SourceAmazonAdsReportRecordTypes `json:"report_record_types,omitempty"`
	sourceType        AmazonAds                          `const:"amazon-ads" json:"sourceType"`
	// The Start date for collecting reports, should not be more than 60 days in the past. In YYYY-MM-DD format
	StartDate *types.Date `json:"start_date,omitempty"`
	// Reflects the state of the Display, Product, and Brand Campaign streams as enabled, paused, or archived. If you do not populate this field, it will be ignored completely.
	StateFilter []SourceAmazonAdsStateFilter `json:"state_filter,omitempty"`
}

func (s SourceAmazonAds) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAmazonAds) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAmazonAds) GetAuthType() *SourceAmazonAdsAuthType {
	return SourceAmazonAdsAuthTypeOauth20.ToPointer()
}

func (o *SourceAmazonAds) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceAmazonAds) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceAmazonAds) GetLookBackWindow() *int64 {
	if o == nil {
		return nil
	}
	return o.LookBackWindow
}

func (o *SourceAmazonAds) GetMarketplaceIds() []string {
	if o == nil {
		return nil
	}
	return o.MarketplaceIds
}

func (o *SourceAmazonAds) GetProfiles() []int64 {
	if o == nil {
		return nil
	}
	return o.Profiles
}

func (o *SourceAmazonAds) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceAmazonAds) GetRegion() *SourceAmazonAdsRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *SourceAmazonAds) GetReportRecordTypes() []SourceAmazonAdsReportRecordTypes {
	if o == nil {
		return nil
	}
	return o.ReportRecordTypes
}

func (o *SourceAmazonAds) GetSourceType() AmazonAds {
	return AmazonAdsAmazonAds
}

func (o *SourceAmazonAds) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceAmazonAds) GetStateFilter() []SourceAmazonAdsStateFilter {
	if o == nil {
		return nil
	}
	return o.StateFilter
}
