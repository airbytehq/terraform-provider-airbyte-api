// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceBasecampUpdate struct {
	AccountID           float64   `json:"account_id"`
	ClientID            string    `json:"client_id"`
	ClientRefreshToken2 string    `json:"client_refresh_token_2"`
	ClientSecret        string    `json:"client_secret"`
	StartDate           time.Time `json:"start_date"`
}

func (s SourceBasecampUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceBasecampUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceBasecampUpdate) GetAccountID() float64 {
	if o == nil {
		return 0.0
	}
	return o.AccountID
}

func (o *SourceBasecampUpdate) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceBasecampUpdate) GetClientRefreshToken2() string {
	if o == nil {
		return ""
	}
	return o.ClientRefreshToken2
}

func (o *SourceBasecampUpdate) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceBasecampUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
