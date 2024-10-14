// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Dremio string

const (
	DremioDremio Dremio = "dremio"
)

func (e Dremio) ToPointer() *Dremio {
	return &e
}
func (e *Dremio) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dremio":
		*e = Dremio(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Dremio: %v", v)
	}
}

type SourceDremio struct {
	// API Key that is generated when you authenticate to Dremio API
	APIKey string `json:"api_key"`
	// URL of your Dremio instance
	BaseURL    *string `default:"https://app.dremio.cloud" json:"base_url"`
	sourceType Dremio  `const:"dremio" json:"sourceType"`
}

func (s SourceDremio) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceDremio) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceDremio) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceDremio) GetBaseURL() *string {
	if o == nil {
		return nil
	}
	return o.BaseURL
}

func (o *SourceDremio) GetSourceType() Dremio {
	return DremioDremio
}
