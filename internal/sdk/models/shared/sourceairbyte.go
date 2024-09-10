// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Airbyte string

const (
	AirbyteAirbyte Airbyte = "airbyte"
)

func (e Airbyte) ToPointer() *Airbyte {
	return &e
}
func (e *Airbyte) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "airbyte":
		*e = Airbyte(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Airbyte: %v", v)
	}
}

type SourceAirbyte struct {
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	sourceType   Airbyte   `const:"airbyte" json:"sourceType"`
	StartDate    time.Time `json:"start_date"`
}

func (s SourceAirbyte) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAirbyte) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAirbyte) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceAirbyte) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceAirbyte) GetSourceType() Airbyte {
	return AirbyteAirbyte
}

func (o *SourceAirbyte) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
