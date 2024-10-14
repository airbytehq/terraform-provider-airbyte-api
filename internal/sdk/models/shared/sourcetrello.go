// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Trello string

const (
	TrelloTrello Trello = "trello"
)

func (e Trello) ToPointer() *Trello {
	return &e
}
func (e *Trello) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "trello":
		*e = Trello(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Trello: %v", v)
	}
}

type SourceTrello struct {
	// IDs of the boards to replicate data from. If left empty, data from all boards to which you have access will be replicated. Please note that this is not the 8-character ID in the board's shortLink (URL of the board). Rather, what is required here is the 24-character ID usually returned by the API
	BoardIds []string `json:"board_ids,omitempty"`
	// Trello API key. See the <a href="https://developer.atlassian.com/cloud/trello/guides/rest-api/authorization/#using-basic-oauth">docs</a> for instructions on how to generate it.
	Key        string `json:"key"`
	sourceType Trello `const:"trello" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
	// Trello API token. See the <a href="https://developer.atlassian.com/cloud/trello/guides/rest-api/authorization/#using-basic-oauth">docs</a> for instructions on how to generate it.
	Token string `json:"token"`
}

func (s SourceTrello) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceTrello) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceTrello) GetBoardIds() []string {
	if o == nil {
		return nil
	}
	return o.BoardIds
}

func (o *SourceTrello) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *SourceTrello) GetSourceType() Trello {
	return TrelloTrello
}

func (o *SourceTrello) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceTrello) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}
