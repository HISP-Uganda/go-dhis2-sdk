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

// HideEmptyItemStrategy the model 'HideEmptyItemStrategy'
type HideEmptyItemStrategy string

// List of HideEmptyItemStrategy
const (
	HIDEEMPTYITEMSTRATEGY_NONE                    HideEmptyItemStrategy = "NONE"
	HIDEEMPTYITEMSTRATEGY_BEFORE_FIRST            HideEmptyItemStrategy = "BEFORE_FIRST"
	HIDEEMPTYITEMSTRATEGY_AFTER_LAST              HideEmptyItemStrategy = "AFTER_LAST"
	HIDEEMPTYITEMSTRATEGY_BEFORE_FIRST_AFTER_LAST HideEmptyItemStrategy = "BEFORE_FIRST_AFTER_LAST"
	HIDEEMPTYITEMSTRATEGY_ALL                     HideEmptyItemStrategy = "ALL"
)

// All allowed values of HideEmptyItemStrategy enum
var AllowedHideEmptyItemStrategyEnumValues = []HideEmptyItemStrategy{
	"NONE",
	"BEFORE_FIRST",
	"AFTER_LAST",
	"BEFORE_FIRST_AFTER_LAST",
	"ALL",
}

func (v *HideEmptyItemStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HideEmptyItemStrategy(value)
	for _, existing := range AllowedHideEmptyItemStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HideEmptyItemStrategy", value)
}

// NewHideEmptyItemStrategyFromValue returns a pointer to a valid HideEmptyItemStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHideEmptyItemStrategyFromValue(v string) (*HideEmptyItemStrategy, error) {
	ev := HideEmptyItemStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HideEmptyItemStrategy: valid values are %v", v, AllowedHideEmptyItemStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HideEmptyItemStrategy) IsValid() bool {
	for _, existing := range AllowedHideEmptyItemStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HideEmptyItemStrategy value
func (v HideEmptyItemStrategy) Ptr() *HideEmptyItemStrategy {
	return &v
}

type NullableHideEmptyItemStrategy struct {
	value *HideEmptyItemStrategy
	isSet bool
}

func (v NullableHideEmptyItemStrategy) Get() *HideEmptyItemStrategy {
	return v.value
}

func (v *NullableHideEmptyItemStrategy) Set(val *HideEmptyItemStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableHideEmptyItemStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableHideEmptyItemStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHideEmptyItemStrategy(val *HideEmptyItemStrategy) *NullableHideEmptyItemStrategy {
	return &NullableHideEmptyItemStrategy{value: val, isSet: true}
}

func (v NullableHideEmptyItemStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHideEmptyItemStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
