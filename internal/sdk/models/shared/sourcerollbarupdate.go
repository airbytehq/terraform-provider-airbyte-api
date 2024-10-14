// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceRollbarUpdate struct {
	AccountAccessToken string    `json:"account_access_token"`
	ProjectAccessToken string    `json:"project_access_token"`
	StartDate          time.Time `json:"start_date"`
}

func (s SourceRollbarUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRollbarUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceRollbarUpdate) GetAccountAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccountAccessToken
}

func (o *SourceRollbarUpdate) GetProjectAccessToken() string {
	if o == nil {
		return ""
	}
	return o.ProjectAccessToken
}

func (o *SourceRollbarUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
