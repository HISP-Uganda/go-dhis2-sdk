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

// NotificationDataType the model 'NotificationDataType'
type NotificationDataType string

// List of NotificationDataType
const (
	NOTIFICATIONDATATYPE_PARAMETERS NotificationDataType = "PARAMETERS"
)

// All allowed values of NotificationDataType enum
var AllowedNotificationDataTypeEnumValues = []NotificationDataType{
	"PARAMETERS",
}

func (v *NotificationDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationDataType(value)
	for _, existing := range AllowedNotificationDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationDataType", value)
}

// NewNotificationDataTypeFromValue returns a pointer to a valid NotificationDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationDataTypeFromValue(v string) (*NotificationDataType, error) {
	ev := NotificationDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationDataType: valid values are %v", v, AllowedNotificationDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationDataType) IsValid() bool {
	for _, existing := range AllowedNotificationDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationDataType value
func (v NotificationDataType) Ptr() *NotificationDataType {
	return &v
}

type NullableNotificationDataType struct {
	value *NotificationDataType
	isSet bool
}

func (v NullableNotificationDataType) Get() *NotificationDataType {
	return v.value
}

func (v *NullableNotificationDataType) Set(val *NotificationDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationDataType(val *NotificationDataType) *NullableNotificationDataType {
	return &NullableNotificationDataType{value: val, isSet: true}
}

func (v NullableNotificationDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
