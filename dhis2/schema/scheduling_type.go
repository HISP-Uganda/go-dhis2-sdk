/*
DHIS2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.42
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"fmt"
)

// SchedulingType the model 'SchedulingType'
type SchedulingType string

// List of SchedulingType
const (
	SCHEDULINGTYPE_CRON        SchedulingType = "CRON"
	SCHEDULINGTYPE_FIXED_DELAY SchedulingType = "FIXED_DELAY"
	SCHEDULINGTYPE_ONCE_ASAP   SchedulingType = "ONCE_ASAP"
)

// All allowed values of SchedulingType enum
var AllowedSchedulingTypeEnumValues = []SchedulingType{
	"CRON",
	"FIXED_DELAY",
	"ONCE_ASAP",
}

func (v *SchedulingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SchedulingType(value)
	for _, existing := range AllowedSchedulingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SchedulingType", value)
}

// NewSchedulingTypeFromValue returns a pointer to a valid SchedulingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSchedulingTypeFromValue(v string) (*SchedulingType, error) {
	ev := SchedulingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SchedulingType: valid values are %v", v, AllowedSchedulingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SchedulingType) IsValid() bool {
	for _, existing := range AllowedSchedulingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SchedulingType value
func (v SchedulingType) Ptr() *SchedulingType {
	return &v
}

type NullableSchedulingType struct {
	value *SchedulingType
	isSet bool
}

func (v NullableSchedulingType) Get() *SchedulingType {
	return v.value
}

func (v *NullableSchedulingType) Set(val *SchedulingType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulingType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulingType(val *SchedulingType) *NullableSchedulingType {
	return &NullableSchedulingType{value: val, isSet: true}
}

func (v NullableSchedulingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
