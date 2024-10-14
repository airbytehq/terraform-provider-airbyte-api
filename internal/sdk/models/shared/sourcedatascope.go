// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Datascope string

const (
	DatascopeDatascope Datascope = "datascope"
)

func (e Datascope) ToPointer() *Datascope {
	return &e
}
func (e *Datascope) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "datascope":
		*e = Datascope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Datascope: %v", v)
	}
}

type SourceDatascope struct {
	// API Key
	APIKey     string    `json:"api_key"`
	sourceType Datascope `const:"datascope" json:"sourceType"`
	// Start date for the data to be replicated
	StartDate string `json:"start_date"`
}

func (s SourceDatascope) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceDatascope) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceDatascope) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceDatascope) GetSourceType() Datascope {
	return DatascopeDatascope
}

func (o *SourceDatascope) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
