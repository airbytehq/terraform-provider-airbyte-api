// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Pennylane string

const (
	PennylanePennylane Pennylane = "pennylane"
)

func (e Pennylane) ToPointer() *Pennylane {
	return &e
}
func (e *Pennylane) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pennylane":
		*e = Pennylane(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Pennylane: %v", v)
	}
}

type SourcePennylane struct {
	APIKey     string    `json:"api_key"`
	sourceType Pennylane `const:"pennylane" json:"sourceType"`
	StartTime  time.Time `json:"start_time"`
}

func (s SourcePennylane) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePennylane) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePennylane) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourcePennylane) GetSourceType() Pennylane {
	return PennylanePennylane
}

func (o *SourcePennylane) GetStartTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartTime
}
