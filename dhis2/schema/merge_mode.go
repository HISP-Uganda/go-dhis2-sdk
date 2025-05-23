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

// MergeMode the model 'MergeMode'
type MergeMode string

// List of MergeMode
const (
	MERGEMODE_REPLACE MergeMode = "REPLACE"
)

// All allowed values of MergeMode enum
var AllowedMergeModeEnumValues = []MergeMode{
	"REPLACE",
}

func (v *MergeMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MergeMode(value)
	for _, existing := range AllowedMergeModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MergeMode", value)
}

// NewMergeModeFromValue returns a pointer to a valid MergeMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMergeModeFromValue(v string) (*MergeMode, error) {
	ev := MergeMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MergeMode: valid values are %v", v, AllowedMergeModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MergeMode) IsValid() bool {
	for _, existing := range AllowedMergeModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MergeMode value
func (v MergeMode) Ptr() *MergeMode {
	return &v
}

type NullableMergeMode struct {
	value *MergeMode
	isSet bool
}

func (v NullableMergeMode) Get() *MergeMode {
	return v.value
}

func (v *NullableMergeMode) Set(val *MergeMode) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeMode) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeMode(val *MergeMode) *NullableMergeMode {
	return &NullableMergeMode{value: val, isSet: true}
}

func (v NullableMergeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
