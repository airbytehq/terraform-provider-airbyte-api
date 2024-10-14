// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Rss string

const (
	RssRss Rss = "rss"
)

func (e Rss) ToPointer() *Rss {
	return &e
}
func (e *Rss) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rss":
		*e = Rss(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Rss: %v", v)
	}
}

type SourceRss struct {
	sourceType Rss `const:"rss" json:"sourceType"`
	// RSS Feed URL
	URL string `json:"url"`
}

func (s SourceRss) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRss) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceRss) GetSourceType() Rss {
	return RssRss
}

func (o *SourceRss) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
