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

// LoginResponseStatus the model 'LoginResponseStatus'
type LoginResponseStatus string

// List of LoginResponseStatus
const (
	LOGINRESPONSESTATUS_SUCCESS                         LoginResponseStatus = "SUCCESS"
	LOGINRESPONSESTATUS_ACCOUNT_DISABLED                LoginResponseStatus = "ACCOUNT_DISABLED"
	LOGINRESPONSESTATUS_ACCOUNT_LOCKED                  LoginResponseStatus = "ACCOUNT_LOCKED"
	LOGINRESPONSESTATUS_ACCOUNT_EXPIRED                 LoginResponseStatus = "ACCOUNT_EXPIRED"
	LOGINRESPONSESTATUS_PASSWORD_EXPIRED                LoginResponseStatus = "PASSWORD_EXPIRED"
	LOGINRESPONSESTATUS_EMAIL_TWO_FACTOR_CODE_SENT      LoginResponseStatus = "EMAIL_TWO_FACTOR_CODE_SENT"
	LOGINRESPONSESTATUS_INCORRECT_TWO_FACTOR_CODE_TOTP  LoginResponseStatus = "INCORRECT_TWO_FACTOR_CODE_TOTP"
	LOGINRESPONSESTATUS_INCORRECT_TWO_FACTOR_CODE_EMAIL LoginResponseStatus = "INCORRECT_TWO_FACTOR_CODE_EMAIL"
	LOGINRESPONSESTATUS_REQUIRES_TWO_FACTOR_ENROLMENT   LoginResponseStatus = "REQUIRES_TWO_FACTOR_ENROLMENT"
)

// All allowed values of LoginResponseStatus enum
var AllowedLoginResponseStatusEnumValues = []LoginResponseStatus{
	"SUCCESS",
	"ACCOUNT_DISABLED",
	"ACCOUNT_LOCKED",
	"ACCOUNT_EXPIRED",
	"PASSWORD_EXPIRED",
	"EMAIL_TWO_FACTOR_CODE_SENT",
	"INCORRECT_TWO_FACTOR_CODE_TOTP",
	"INCORRECT_TWO_FACTOR_CODE_EMAIL",
	"REQUIRES_TWO_FACTOR_ENROLMENT",
}

func (v *LoginResponseStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoginResponseStatus(value)
	for _, existing := range AllowedLoginResponseStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoginResponseStatus", value)
}

// NewLoginResponseStatusFromValue returns a pointer to a valid LoginResponseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoginResponseStatusFromValue(v string) (*LoginResponseStatus, error) {
	ev := LoginResponseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LoginResponseStatus: valid values are %v", v, AllowedLoginResponseStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoginResponseStatus) IsValid() bool {
	for _, existing := range AllowedLoginResponseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LoginResponseStatus value
func (v LoginResponseStatus) Ptr() *LoginResponseStatus {
	return &v
}

type NullableLoginResponseStatus struct {
	value *LoginResponseStatus
	isSet bool
}

func (v NullableLoginResponseStatus) Get() *LoginResponseStatus {
	return v.value
}

func (v *NullableLoginResponseStatus) Set(val *LoginResponseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginResponseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginResponseStatus(val *LoginResponseStatus) *NullableLoginResponseStatus {
	return &NullableLoginResponseStatus{value: val, isSet: true}
}

func (v NullableLoginResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
