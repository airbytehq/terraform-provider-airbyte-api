// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type ClickupAPI string

const (
	ClickupAPIClickupAPI ClickupAPI = "clickup-api"
)

func (e ClickupAPI) ToPointer() *ClickupAPI {
	return &e
}
func (e *ClickupAPI) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "clickup-api":
		*e = ClickupAPI(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClickupAPI: %v", v)
	}
}

type SourceClickupAPI struct {
	// Every ClickUp API call required authentication. This field is your personal API token. See <a href="https://clickup.com/api/developer-portal/authentication/#personal-token">here</a>.
	APIToken string `json:"api_token"`
	// Include or exclude closed tasks. By default, they are excluded. See <a https://clickup.com/api/clickupreference/operation/GetTasks/#!in=query&path=include_closed&t=request">here</a>.
	IncludeClosedTasks *bool      `default:"false" json:"include_closed_tasks"`
	sourceType         ClickupAPI `const:"clickup-api" json:"sourceType"`
}

func (s SourceClickupAPI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceClickupAPI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceClickupAPI) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceClickupAPI) GetIncludeClosedTasks() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeClosedTasks
}

func (o *SourceClickupAPI) GetSourceType() ClickupAPI {
	return ClickupAPIClickupAPI
}
