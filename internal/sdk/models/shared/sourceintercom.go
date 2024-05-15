// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Intercom string

const (
	IntercomIntercom Intercom = "intercom"
)

func (e Intercom) ToPointer() *Intercom {
	return &e
}
func (e *Intercom) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "intercom":
		*e = Intercom(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Intercom: %v", v)
	}
}

type SourceIntercom struct {
	// Access token for making authenticated requests. See the <a href="https://developers.intercom.com/building-apps/docs/authentication-types#how-to-get-your-access-token">Intercom docs</a> for more information.
	AccessToken string `json:"access_token"`
	// Client Id for your Intercom application.
	ClientID *string `json:"client_id,omitempty"`
	// Client Secret for your Intercom application.
	ClientSecret *string  `json:"client_secret,omitempty"`
	sourceType   Intercom `const:"intercom" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceIntercom) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceIntercom) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceIntercom) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceIntercom) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *SourceIntercom) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *SourceIntercom) GetSourceType() Intercom {
	return IntercomIntercom
}

func (o *SourceIntercom) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
