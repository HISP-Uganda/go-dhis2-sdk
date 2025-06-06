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

// SendStrategy the model 'SendStrategy'
type SendStrategy string

// List of SendStrategy
const (
	SENDSTRATEGY_COLLECTIVE_SUMMARY  SendStrategy = "COLLECTIVE_SUMMARY"
	SENDSTRATEGY_SINGLE_NOTIFICATION SendStrategy = "SINGLE_NOTIFICATION"
)

// All allowed values of SendStrategy enum
var AllowedSendStrategyEnumValues = []SendStrategy{
	"COLLECTIVE_SUMMARY",
	"SINGLE_NOTIFICATION",
}

func (v *SendStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SendStrategy(value)
	for _, existing := range AllowedSendStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SendStrategy", value)
}

// NewSendStrategyFromValue returns a pointer to a valid SendStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSendStrategyFromValue(v string) (*SendStrategy, error) {
	ev := SendStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SendStrategy: valid values are %v", v, AllowedSendStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SendStrategy) IsValid() bool {
	for _, existing := range AllowedSendStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SendStrategy value
func (v SendStrategy) Ptr() *SendStrategy {
	return &v
}

type NullableSendStrategy struct {
	value *SendStrategy
	isSet bool
}

func (v NullableSendStrategy) Get() *SendStrategy {
	return v.value
}

func (v *NullableSendStrategy) Set(val *SendStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableSendStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableSendStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendStrategy(val *SendStrategy) *NullableSendStrategy {
	return &NullableSendStrategy{value: val, isSet: true}
}

func (v NullableSendStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
