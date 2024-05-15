// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Delighted string

const (
	DelightedDelighted Delighted = "delighted"
)

func (e Delighted) ToPointer() *Delighted {
	return &e
}
func (e *Delighted) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "delighted":
		*e = Delighted(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Delighted: %v", v)
	}
}

type SourceDelighted struct {
	// A Delighted API key.
	APIKey string `json:"api_key"`
	// The date from which you'd like to replicate the data
	Since      time.Time `json:"since"`
	sourceType Delighted `const:"delighted" json:"sourceType"`
}

func (s SourceDelighted) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceDelighted) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceDelighted) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceDelighted) GetSince() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Since
}

func (o *SourceDelighted) GetSourceType() Delighted {
	return DelightedDelighted
}