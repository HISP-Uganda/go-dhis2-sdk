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

// CompletenessMethod the model 'CompletenessMethod'
type CompletenessMethod string

// List of CompletenessMethod
const (
	COMPLETENESSMETHOD_ALL_DATAVALUE          CompletenessMethod = "ALL_DATAVALUE"
	COMPLETENESSMETHOD_AT_LEAST_ONE_DATAVALUE CompletenessMethod = "AT_LEAST_ONE_DATAVALUE"
	COMPLETENESSMETHOD_DO_NOT_MARK_COMPLETE   CompletenessMethod = "DO_NOT_MARK_COMPLETE"
)

// All allowed values of CompletenessMethod enum
var AllowedCompletenessMethodEnumValues = []CompletenessMethod{
	"ALL_DATAVALUE",
	"AT_LEAST_ONE_DATAVALUE",
	"DO_NOT_MARK_COMPLETE",
}

func (v *CompletenessMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompletenessMethod(value)
	for _, existing := range AllowedCompletenessMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompletenessMethod", value)
}

// NewCompletenessMethodFromValue returns a pointer to a valid CompletenessMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompletenessMethodFromValue(v string) (*CompletenessMethod, error) {
	ev := CompletenessMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompletenessMethod: valid values are %v", v, AllowedCompletenessMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompletenessMethod) IsValid() bool {
	for _, existing := range AllowedCompletenessMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompletenessMethod value
func (v CompletenessMethod) Ptr() *CompletenessMethod {
	return &v
}

type NullableCompletenessMethod struct {
	value *CompletenessMethod
	isSet bool
}

func (v NullableCompletenessMethod) Get() *CompletenessMethod {
	return v.value
}

func (v *NullableCompletenessMethod) Set(val *CompletenessMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletenessMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletenessMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletenessMethod(val *CompletenessMethod) *NullableCompletenessMethod {
	return &NullableCompletenessMethod{value: val, isSet: true}
}

func (v NullableCompletenessMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletenessMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
