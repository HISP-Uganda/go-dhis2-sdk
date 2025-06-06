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

// TypeOfNumber the model 'TypeOfNumber'
type TypeOfNumber string

// List of TypeOfNumber
const (
	TYPEOFNUMBER_UNKNOWN           TypeOfNumber = "UNKNOWN"
	TYPEOFNUMBER_INTERNATIONAL     TypeOfNumber = "INTERNATIONAL"
	TYPEOFNUMBER_NATIONAL          TypeOfNumber = "NATIONAL"
	TYPEOFNUMBER_NETWORK_SPECIFIC  TypeOfNumber = "NETWORK_SPECIFIC"
	TYPEOFNUMBER_SUBSCRIBER_NUMBER TypeOfNumber = "SUBSCRIBER_NUMBER"
	TYPEOFNUMBER_ALPHANUMERIC      TypeOfNumber = "ALPHANUMERIC"
	TYPEOFNUMBER_ABBREVIATED       TypeOfNumber = "ABBREVIATED"
)

// All allowed values of TypeOfNumber enum
var AllowedTypeOfNumberEnumValues = []TypeOfNumber{
	"UNKNOWN",
	"INTERNATIONAL",
	"NATIONAL",
	"NETWORK_SPECIFIC",
	"SUBSCRIBER_NUMBER",
	"ALPHANUMERIC",
	"ABBREVIATED",
}

func (v *TypeOfNumber) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeOfNumber(value)
	for _, existing := range AllowedTypeOfNumberEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeOfNumber", value)
}

// NewTypeOfNumberFromValue returns a pointer to a valid TypeOfNumber
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeOfNumberFromValue(v string) (*TypeOfNumber, error) {
	ev := TypeOfNumber(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypeOfNumber: valid values are %v", v, AllowedTypeOfNumberEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeOfNumber) IsValid() bool {
	for _, existing := range AllowedTypeOfNumberEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TypeOfNumber value
func (v TypeOfNumber) Ptr() *TypeOfNumber {
	return &v
}

type NullableTypeOfNumber struct {
	value *TypeOfNumber
	isSet bool
}

func (v NullableTypeOfNumber) Get() *TypeOfNumber {
	return v.value
}

func (v *NullableTypeOfNumber) Set(val *TypeOfNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeOfNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeOfNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeOfNumber(val *TypeOfNumber) *NullableTypeOfNumber {
	return &NullableTypeOfNumber{value: val, isSet: true}
}

func (v NullableTypeOfNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeOfNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
