// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Lob string

const (
	LobLob Lob = "lob"
)

func (e Lob) ToPointer() *Lob {
	return &e
}
func (e *Lob) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "lob":
		*e = Lob(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Lob: %v", v)
	}
}

type SourceLob struct {
	// API key to use for authentication. You can find your account's API keys in your Dashboard Settings at https://dashboard.lob.com/settings/api-keys.
	APIKey string `json:"api_key"`
	// Max records per page limit
	Limit      *string   `default:"50" json:"limit"`
	sourceType Lob       `const:"lob" json:"sourceType"`
	StartDate  time.Time `json:"start_date"`
}

func (s SourceLob) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLob) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLob) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceLob) GetLimit() *string {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *SourceLob) GetSourceType() Lob {
	return LobLob
}

func (o *SourceLob) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
