// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type GeographyEnum string

const (
	GeographyEnumAuto GeographyEnum = "auto"
	GeographyEnumUs   GeographyEnum = "us"
	GeographyEnumEu   GeographyEnum = "eu"
)

func (e GeographyEnum) ToPointer() *GeographyEnum {
	return &e
}
func (e *GeographyEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		fallthrough
	case "us":
		fallthrough
	case "eu":
		*e = GeographyEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GeographyEnum: %v", v)
	}
}
