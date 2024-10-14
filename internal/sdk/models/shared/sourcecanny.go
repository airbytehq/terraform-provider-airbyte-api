// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Canny string

const (
	CannyCanny Canny = "canny"
)

func (e Canny) ToPointer() *Canny {
	return &e
}
func (e *Canny) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "canny":
		*e = Canny(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Canny: %v", v)
	}
}

type SourceCanny struct {
	// You can find your secret API key in Your Canny Subdomain > Settings > API
	APIKey     string `json:"api_key"`
	sourceType Canny  `const:"canny" json:"sourceType"`
}

func (s SourceCanny) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceCanny) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceCanny) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceCanny) GetSourceType() Canny {
	return CannyCanny
}
