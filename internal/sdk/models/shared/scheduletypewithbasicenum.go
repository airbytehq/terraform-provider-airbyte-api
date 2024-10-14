// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ScheduleTypeWithBasicEnum string

const (
	ScheduleTypeWithBasicEnumManual ScheduleTypeWithBasicEnum = "manual"
	ScheduleTypeWithBasicEnumCron   ScheduleTypeWithBasicEnum = "cron"
	ScheduleTypeWithBasicEnumBasic  ScheduleTypeWithBasicEnum = "basic"
)

func (e ScheduleTypeWithBasicEnum) ToPointer() *ScheduleTypeWithBasicEnum {
	return &e
}
func (e *ScheduleTypeWithBasicEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manual":
		fallthrough
	case "cron":
		fallthrough
	case "basic":
		*e = ScheduleTypeWithBasicEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScheduleTypeWithBasicEnum: %v", v)
	}
}
