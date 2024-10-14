// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Pendo string

const (
	PendoPendo Pendo = "pendo"
)

func (e Pendo) ToPointer() *Pendo {
	return &e
}
func (e *Pendo) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pendo":
		*e = Pendo(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Pendo: %v", v)
	}
}

type SourcePendo struct {
	APIKey     string `json:"api_key"`
	sourceType Pendo  `const:"pendo" json:"sourceType"`
}

func (s SourcePendo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePendo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePendo) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourcePendo) GetSourceType() Pendo {
	return PendoPendo
}
