// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Tempo string

const (
	TempoTempo Tempo = "tempo"
)

func (e Tempo) ToPointer() *Tempo {
	return &e
}
func (e *Tempo) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tempo":
		*e = Tempo(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Tempo: %v", v)
	}
}

type SourceTempo struct {
	// Tempo API Token. Go to Tempo>Settings, scroll down to Data Access and select API integration.
	APIToken   string `json:"api_token"`
	sourceType Tempo  `const:"tempo" json:"sourceType"`
}

func (s SourceTempo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceTempo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceTempo) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceTempo) GetSourceType() Tempo {
	return TempoTempo
}
