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

// FileResourceDomain the model 'FileResourceDomain'
type FileResourceDomain string

// List of FileResourceDomain
const (
	FILERESOURCEDOMAIN_DATA_VALUE         FileResourceDomain = "DATA_VALUE"
	FILERESOURCEDOMAIN_PUSH_ANALYSIS      FileResourceDomain = "PUSH_ANALYSIS"
	FILERESOURCEDOMAIN_DOCUMENT           FileResourceDomain = "DOCUMENT"
	FILERESOURCEDOMAIN_MESSAGE_ATTACHMENT FileResourceDomain = "MESSAGE_ATTACHMENT"
	FILERESOURCEDOMAIN_USER_AVATAR        FileResourceDomain = "USER_AVATAR"
	FILERESOURCEDOMAIN_ORG_UNIT           FileResourceDomain = "ORG_UNIT"
	FILERESOURCEDOMAIN_ICON               FileResourceDomain = "ICON"
	FILERESOURCEDOMAIN_JOB_DATA           FileResourceDomain = "JOB_DATA"
)

// All allowed values of FileResourceDomain enum
var AllowedFileResourceDomainEnumValues = []FileResourceDomain{
	"DATA_VALUE",
	"PUSH_ANALYSIS",
	"DOCUMENT",
	"MESSAGE_ATTACHMENT",
	"USER_AVATAR",
	"ORG_UNIT",
	"ICON",
	"JOB_DATA",
}

func (v *FileResourceDomain) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileResourceDomain(value)
	for _, existing := range AllowedFileResourceDomainEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileResourceDomain", value)
}

// NewFileResourceDomainFromValue returns a pointer to a valid FileResourceDomain
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileResourceDomainFromValue(v string) (*FileResourceDomain, error) {
	ev := FileResourceDomain(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileResourceDomain: valid values are %v", v, AllowedFileResourceDomainEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileResourceDomain) IsValid() bool {
	for _, existing := range AllowedFileResourceDomainEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileResourceDomain value
func (v FileResourceDomain) Ptr() *FileResourceDomain {
	return &v
}

type NullableFileResourceDomain struct {
	value *FileResourceDomain
	isSet bool
}

func (v NullableFileResourceDomain) Get() *FileResourceDomain {
	return v.value
}

func (v *NullableFileResourceDomain) Set(val *FileResourceDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableFileResourceDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableFileResourceDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileResourceDomain(val *FileResourceDomain) *NullableFileResourceDomain {
	return &NullableFileResourceDomain{value: val, isSet: true}
}

func (v NullableFileResourceDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileResourceDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
