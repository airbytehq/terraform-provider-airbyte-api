// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Instatus string

const (
	InstatusInstatus Instatus = "instatus"
)

func (e Instatus) ToPointer() *Instatus {
	return &e
}
func (e *Instatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "instatus":
		*e = Instatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Instatus: %v", v)
	}
}

type SourceInstatus struct {
	// Instatus REST API key
	APIKey     string   `json:"api_key"`
	sourceType Instatus `const:"instatus" json:"sourceType"`
}

func (s SourceInstatus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceInstatus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceInstatus) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceInstatus) GetSourceType() Instatus {
	return InstatusInstatus
}
