// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// NonBreakingSchemaUpdatesBehaviorEnumNoDefault - Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
type NonBreakingSchemaUpdatesBehaviorEnumNoDefault string

const (
	NonBreakingSchemaUpdatesBehaviorEnumNoDefaultIgnore            NonBreakingSchemaUpdatesBehaviorEnumNoDefault = "ignore"
	NonBreakingSchemaUpdatesBehaviorEnumNoDefaultDisableConnection NonBreakingSchemaUpdatesBehaviorEnumNoDefault = "disable_connection"
	NonBreakingSchemaUpdatesBehaviorEnumNoDefaultPropagateColumns  NonBreakingSchemaUpdatesBehaviorEnumNoDefault = "propagate_columns"
	NonBreakingSchemaUpdatesBehaviorEnumNoDefaultPropagateFully    NonBreakingSchemaUpdatesBehaviorEnumNoDefault = "propagate_fully"
)

func (e NonBreakingSchemaUpdatesBehaviorEnumNoDefault) ToPointer() *NonBreakingSchemaUpdatesBehaviorEnumNoDefault {
	return &e
}

func (e *NonBreakingSchemaUpdatesBehaviorEnumNoDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ignore":
		fallthrough
	case "disable_connection":
		fallthrough
	case "propagate_columns":
		fallthrough
	case "propagate_fully":
		*e = NonBreakingSchemaUpdatesBehaviorEnumNoDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NonBreakingSchemaUpdatesBehaviorEnumNoDefault: %v", v)
	}
}
