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

// NotificationTrigger the model 'NotificationTrigger'
type NotificationTrigger string

// List of NotificationTrigger
const (
	NOTIFICATIONTRIGGER_ENROLLMENT                     NotificationTrigger = "ENROLLMENT"
	NOTIFICATIONTRIGGER_COMPLETION                     NotificationTrigger = "COMPLETION"
	NOTIFICATIONTRIGGER_PROGRAM_RULE                   NotificationTrigger = "PROGRAM_RULE"
	NOTIFICATIONTRIGGER_SCHEDULED_DAYS_DUE_DATE        NotificationTrigger = "SCHEDULED_DAYS_DUE_DATE"
	NOTIFICATIONTRIGGER_SCHEDULED_DAYS_INCIDENT_DATE   NotificationTrigger = "SCHEDULED_DAYS_INCIDENT_DATE"
	NOTIFICATIONTRIGGER_SCHEDULED_DAYS_ENROLLMENT_DATE NotificationTrigger = "SCHEDULED_DAYS_ENROLLMENT_DATE"
)

// All allowed values of NotificationTrigger enum
var AllowedNotificationTriggerEnumValues = []NotificationTrigger{
	"ENROLLMENT",
	"COMPLETION",
	"PROGRAM_RULE",
	"SCHEDULED_DAYS_DUE_DATE",
	"SCHEDULED_DAYS_INCIDENT_DATE",
	"SCHEDULED_DAYS_ENROLLMENT_DATE",
}

func (v *NotificationTrigger) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationTrigger(value)
	for _, existing := range AllowedNotificationTriggerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationTrigger", value)
}

// NewNotificationTriggerFromValue returns a pointer to a valid NotificationTrigger
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationTriggerFromValue(v string) (*NotificationTrigger, error) {
	ev := NotificationTrigger(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationTrigger: valid values are %v", v, AllowedNotificationTriggerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationTrigger) IsValid() bool {
	for _, existing := range AllowedNotificationTriggerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationTrigger value
func (v NotificationTrigger) Ptr() *NotificationTrigger {
	return &v
}

type NullableNotificationTrigger struct {
	value *NotificationTrigger
	isSet bool
}

func (v NullableNotificationTrigger) Get() *NotificationTrigger {
	return v.value
}

func (v *NullableNotificationTrigger) Set(val *NotificationTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTrigger(val *NotificationTrigger) *NullableNotificationTrigger {
	return &NullableNotificationTrigger{value: val, isSet: true}
}

func (v NullableNotificationTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
