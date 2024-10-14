// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Aircall string

const (
	AircallAircall Aircall = "aircall"
)

func (e Aircall) ToPointer() *Aircall {
	return &e
}
func (e *Aircall) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aircall":
		*e = Aircall(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Aircall: %v", v)
	}
}

type SourceAircall struct {
	// App ID found at settings https://dashboard.aircall.io/integrations/api-keys
	APIID string `json:"api_id"`
	// App token found at settings (Ref- https://dashboard.aircall.io/integrations/api-keys)
	APIToken   string  `json:"api_token"`
	sourceType Aircall `const:"aircall" json:"sourceType"`
	// Date time filter for incremental filter, Specify which date to extract from.
	StartDate time.Time `json:"start_date"`
}

func (s SourceAircall) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAircall) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAircall) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *SourceAircall) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceAircall) GetSourceType() Aircall {
	return AircallAircall
}

func (o *SourceAircall) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
