// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Sevenshifts string

const (
	SevenshiftsSevenshifts Sevenshifts = "7shifts"
)

func (e Sevenshifts) ToPointer() *Sevenshifts {
	return &e
}
func (e *Sevenshifts) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "7shifts":
		*e = Sevenshifts(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sevenshifts: %v", v)
	}
}

type Source7shifts struct {
	// Access token to use for authentication. Generate it in the 7shifts Developer Tools.
	AccessToken string      `json:"access_token"`
	sourceType  Sevenshifts `const:"7shifts" json:"sourceType"`
	StartDate   time.Time   `json:"start_date"`
}

func (s Source7shifts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Source7shifts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Source7shifts) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *Source7shifts) GetSourceType() Sevenshifts {
	return SevenshiftsSevenshifts
}

func (o *Source7shifts) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
