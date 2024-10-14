// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type RkiCovid string

const (
	RkiCovidRkiCovid RkiCovid = "rki-covid"
)

func (e RkiCovid) ToPointer() *RkiCovid {
	return &e
}
func (e *RkiCovid) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rki-covid":
		*e = RkiCovid(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RkiCovid: %v", v)
	}
}

type SourceRkiCovid struct {
	sourceType RkiCovid `const:"rki-covid" json:"sourceType"`
	// UTC date in the format 2017-01-25. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}

func (s SourceRkiCovid) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRkiCovid) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceRkiCovid) GetSourceType() RkiCovid {
	return RkiCovidRkiCovid
}

func (o *SourceRkiCovid) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
