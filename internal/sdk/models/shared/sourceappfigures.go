// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

// SourceAppfiguresGroupBy - Category term for grouping the search results
type SourceAppfiguresGroupBy string

const (
	SourceAppfiguresGroupByNetwork SourceAppfiguresGroupBy = "network"
	SourceAppfiguresGroupByProduct SourceAppfiguresGroupBy = "product"
	SourceAppfiguresGroupByCountry SourceAppfiguresGroupBy = "country"
	SourceAppfiguresGroupByDate    SourceAppfiguresGroupBy = "date"
)

func (e SourceAppfiguresGroupBy) ToPointer() *SourceAppfiguresGroupBy {
	return &e
}
func (e *SourceAppfiguresGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "network":
		fallthrough
	case "product":
		fallthrough
	case "country":
		fallthrough
	case "date":
		*e = SourceAppfiguresGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAppfiguresGroupBy: %v", v)
	}
}

type Appfigures string

const (
	AppfiguresAppfigures Appfigures = "appfigures"
)

func (e Appfigures) ToPointer() *Appfigures {
	return &e
}
func (e *Appfigures) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "appfigures":
		*e = Appfigures(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Appfigures: %v", v)
	}
}

type SourceAppfigures struct {
	APIKey string `json:"api_key"`
	// Category term for grouping the search results
	GroupBy *SourceAppfiguresGroupBy `default:"product" json:"group_by"`
	// The store which needs to be searched in streams
	SearchStore *string    `default:"apple" json:"search_store"`
	sourceType  Appfigures `const:"appfigures" json:"sourceType"`
	StartDate   time.Time  `json:"start_date"`
}

func (s SourceAppfigures) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAppfigures) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAppfigures) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceAppfigures) GetGroupBy() *SourceAppfiguresGroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

func (o *SourceAppfigures) GetSearchStore() *string {
	if o == nil {
		return nil
	}
	return o.SearchStore
}

func (o *SourceAppfigures) GetSourceType() Appfigures {
	return AppfiguresAppfigures
}

func (o *SourceAppfigures) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
