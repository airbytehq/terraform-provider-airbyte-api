// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Eventbrite string

const (
	EventbriteEventbrite Eventbrite = "eventbrite"
)

func (e Eventbrite) ToPointer() *Eventbrite {
	return &e
}
func (e *Eventbrite) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "eventbrite":
		*e = Eventbrite(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Eventbrite: %v", v)
	}
}

type SourceEventbrite struct {
	// The private token to use for authenticating API requests.
	PrivateToken string     `json:"private_token"`
	sourceType   Eventbrite `const:"eventbrite" json:"sourceType"`
	StartDate    time.Time  `json:"start_date"`
}

func (s SourceEventbrite) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceEventbrite) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceEventbrite) GetPrivateToken() string {
	if o == nil {
		return ""
	}
	return o.PrivateToken
}

func (o *SourceEventbrite) GetSourceType() Eventbrite {
	return EventbriteEventbrite
}

func (o *SourceEventbrite) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
