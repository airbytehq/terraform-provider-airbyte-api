// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"encoding/json"
	"fmt"
)

type SourcePolygonStockAPIPolygonStockAPIEnum string

const (
	SourcePolygonStockAPIPolygonStockAPIEnumPolygonStockAPI SourcePolygonStockAPIPolygonStockAPIEnum = "polygon-stock-api"
)

func (e *SourcePolygonStockAPIPolygonStockAPIEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "polygon-stock-api":
		*e = SourcePolygonStockAPIPolygonStockAPIEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePolygonStockAPIPolygonStockAPIEnum: %s", s)
	}
}

// SourcePolygonStockAPI - The values required to configure the source.
type SourcePolygonStockAPI struct {
	// Determines whether or not the results are adjusted for splits. By default, results are adjusted and set to true. Set this to false to get results that are NOT adjusted for splits.
	Adjusted *string `json:"adjusted,omitempty"`
	// Your API ACCESS Key
	APIKey string `json:"apiKey"`
	// The target date for the aggregate window.
	EndDate types.Date `json:"end_date"`
	// The target date for the aggregate window.
	Limit *int64 `json:"limit,omitempty"`
	// The size of the timespan multiplier.
	Multiplier int64 `json:"multiplier"`
	// Sort the results by timestamp. asc will return results in ascending order (oldest at the top), desc will return results in descending order (newest at the top).
	Sort       *string                                  `json:"sort,omitempty"`
	SourceType SourcePolygonStockAPIPolygonStockAPIEnum `json:"sourceType"`
	// The beginning date for the aggregate window.
	StartDate types.Date `json:"start_date"`
	// The exchange symbol that this item is traded under.
	StocksTicker string `json:"stocksTicker"`
	// The size of the time window.
	Timespan string `json:"timespan"`
}
