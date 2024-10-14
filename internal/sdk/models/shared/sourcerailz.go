// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Railz string

const (
	RailzRailz Railz = "railz"
)

func (e Railz) ToPointer() *Railz {
	return &e
}
func (e *Railz) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "railz":
		*e = Railz(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Railz: %v", v)
	}
}

type SourceRailz struct {
	// Client ID (client_id)
	ClientID string `json:"client_id"`
	// Secret key (secret_key)
	SecretKey  string `json:"secret_key"`
	sourceType Railz  `const:"railz" json:"sourceType"`
	// Start date
	StartDate string `json:"start_date"`
}

func (s SourceRailz) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRailz) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceRailz) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceRailz) GetSecretKey() string {
	if o == nil {
		return ""
	}
	return o.SecretKey
}

func (o *SourceRailz) GetSourceType() Railz {
	return RailzRailz
}

func (o *SourceRailz) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
