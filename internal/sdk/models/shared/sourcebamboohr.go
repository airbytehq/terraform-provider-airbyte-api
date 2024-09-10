// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type BambooHr string

const (
	BambooHrBambooHr BambooHr = "bamboo-hr"
)

func (e BambooHr) ToPointer() *BambooHr {
	return &e
}
func (e *BambooHr) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bamboo-hr":
		*e = BambooHr(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BambooHr: %v", v)
	}
}

type SourceBambooHr struct {
	// Api key of bamboo hr
	APIKey string `json:"api_key"`
	// Comma-separated list of fields to include in custom reports.
	CustomReportsFields *string `json:"custom_reports_fields,omitempty"`
	// If true, the custom reports endpoint will include the default fields defined here: https://documentation.bamboohr.com/docs/list-of-field-names.
	CustomReportsIncludeDefaultFields *bool      `default:"true" json:"custom_reports_include_default_fields"`
	sourceType                        BambooHr   `const:"bamboo-hr" json:"sourceType"`
	StartDate                         *time.Time `json:"start_date,omitempty"`
	// Sub Domain of bamboo hr
	Subdomain string `json:"subdomain"`
}

func (s SourceBambooHr) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceBambooHr) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceBambooHr) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceBambooHr) GetCustomReportsFields() *string {
	if o == nil {
		return nil
	}
	return o.CustomReportsFields
}

func (o *SourceBambooHr) GetCustomReportsIncludeDefaultFields() *bool {
	if o == nil {
		return nil
	}
	return o.CustomReportsIncludeDefaultFields
}

func (o *SourceBambooHr) GetSourceType() BambooHr {
	return BambooHrBambooHr
}

func (o *SourceBambooHr) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceBambooHr) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}
