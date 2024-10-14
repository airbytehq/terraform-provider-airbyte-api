// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceGuruUpdate struct {
	Password *string `json:"password,omitempty"`
	// Query for searching cards
	SearchCardsQuery *string   `json:"search_cards_query,omitempty"`
	StartDate        time.Time `json:"start_date"`
	// Team ID received through response of /teams streams, make sure about access to the team
	TeamID   *string `json:"team_id,omitempty"`
	Username string  `json:"username"`
}

func (s SourceGuruUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGuruUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGuruUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceGuruUpdate) GetSearchCardsQuery() *string {
	if o == nil {
		return nil
	}
	return o.SearchCardsQuery
}

func (o *SourceGuruUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceGuruUpdate) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *SourceGuruUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
