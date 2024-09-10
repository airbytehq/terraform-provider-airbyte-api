// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceHardcodedRecordsUpdate struct {
	// How many records per stream should be generated
	Count *int64 `default:"1000" json:"count"`
}

func (s SourceHardcodedRecordsUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceHardcodedRecordsUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceHardcodedRecordsUpdate) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}
