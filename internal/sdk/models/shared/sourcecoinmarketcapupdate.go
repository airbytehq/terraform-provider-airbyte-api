// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DataType - /latest: Latest market ticker quotes and averages for cryptocurrencies and exchanges. /historical: Intervals of historic market data like OHLCV data or data for use in charting libraries. See <a href="https://coinmarketcap.com/api/documentation/v1/#section/Endpoint-Overview">here</a>.
type DataType string

const (
	DataTypeLatest     DataType = "latest"
	DataTypeHistorical DataType = "historical"
)

func (e DataType) ToPointer() *DataType {
	return &e
}
func (e *DataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "latest":
		fallthrough
	case "historical":
		*e = DataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DataType: %v", v)
	}
}

type SourceCoinmarketcapUpdate struct {
	// Your API Key. See <a href="https://coinmarketcap.com/api/documentation/v1/#section/Authentication">here</a>. The token is case sensitive.
	APIKey string `json:"api_key"`
	// /latest: Latest market ticker quotes and averages for cryptocurrencies and exchanges. /historical: Intervals of historic market data like OHLCV data or data for use in charting libraries. See <a href="https://coinmarketcap.com/api/documentation/v1/#section/Endpoint-Overview">here</a>.
	DataType DataType `json:"data_type"`
	// Cryptocurrency symbols. (only used for quotes stream)
	Symbols []string `json:"symbols,omitempty"`
}

func (o *SourceCoinmarketcapUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceCoinmarketcapUpdate) GetDataType() DataType {
	if o == nil {
		return DataType("")
	}
	return o.DataType
}

func (o *SourceCoinmarketcapUpdate) GetSymbols() []string {
	if o == nil {
		return nil
	}
	return o.Symbols
}
