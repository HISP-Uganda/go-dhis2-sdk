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

// IndicatorTypePropertyNames the model 'IndicatorTypePropertyNames'
type IndicatorTypePropertyNames string

// List of IndicatorTypePropertyNames
const (
	INDICATORTYPEPROPERTYNAMES_ACCESS           IndicatorTypePropertyNames = "access"
	INDICATORTYPEPROPERTYNAMES_ATTRIBUTE_VALUES IndicatorTypePropertyNames = "attributeValues"
	INDICATORTYPEPROPERTYNAMES_CODE             IndicatorTypePropertyNames = "code"
	INDICATORTYPEPROPERTYNAMES_CREATED          IndicatorTypePropertyNames = "created"
	INDICATORTYPEPROPERTYNAMES_CREATED_BY       IndicatorTypePropertyNames = "createdBy"
	INDICATORTYPEPROPERTYNAMES_DISPLAY_NAME     IndicatorTypePropertyNames = "displayName"
	INDICATORTYPEPROPERTYNAMES_FACTOR           IndicatorTypePropertyNames = "factor"
	INDICATORTYPEPROPERTYNAMES_FAVORITE         IndicatorTypePropertyNames = "favorite"
	INDICATORTYPEPROPERTYNAMES_FAVORITES        IndicatorTypePropertyNames = "favorites"
	INDICATORTYPEPROPERTYNAMES_HREF             IndicatorTypePropertyNames = "href"
	INDICATORTYPEPROPERTYNAMES_ID               IndicatorTypePropertyNames = "id"
	INDICATORTYPEPROPERTYNAMES_LAST_UPDATED     IndicatorTypePropertyNames = "lastUpdated"
	INDICATORTYPEPROPERTYNAMES_LAST_UPDATED_BY  IndicatorTypePropertyNames = "lastUpdatedBy"
	INDICATORTYPEPROPERTYNAMES_NAME             IndicatorTypePropertyNames = "name"
	INDICATORTYPEPROPERTYNAMES_NUMBER           IndicatorTypePropertyNames = "number"
	INDICATORTYPEPROPERTYNAMES_SHARING          IndicatorTypePropertyNames = "sharing"
	INDICATORTYPEPROPERTYNAMES_TRANSLATIONS     IndicatorTypePropertyNames = "translations"
)

// All allowed values of IndicatorTypePropertyNames enum
var AllowedIndicatorTypePropertyNamesEnumValues = []IndicatorTypePropertyNames{
	"access",
	"attributeValues",
	"code",
	"created",
	"createdBy",
	"displayName",
	"factor",
	"favorite",
	"favorites",
	"href",
	"id",
	"lastUpdated",
	"lastUpdatedBy",
	"name",
	"number",
	"sharing",
	"translations",
}

func (v *IndicatorTypePropertyNames) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IndicatorTypePropertyNames(value)
	for _, existing := range AllowedIndicatorTypePropertyNamesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IndicatorTypePropertyNames", value)
}

// NewIndicatorTypePropertyNamesFromValue returns a pointer to a valid IndicatorTypePropertyNames
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIndicatorTypePropertyNamesFromValue(v string) (*IndicatorTypePropertyNames, error) {
	ev := IndicatorTypePropertyNames(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IndicatorTypePropertyNames: valid values are %v", v, AllowedIndicatorTypePropertyNamesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IndicatorTypePropertyNames) IsValid() bool {
	for _, existing := range AllowedIndicatorTypePropertyNamesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IndicatorTypePropertyNames value
func (v IndicatorTypePropertyNames) Ptr() *IndicatorTypePropertyNames {
	return &v
}

type NullableIndicatorTypePropertyNames struct {
	value *IndicatorTypePropertyNames
	isSet bool
}

func (v NullableIndicatorTypePropertyNames) Get() *IndicatorTypePropertyNames {
	return v.value
}

func (v *NullableIndicatorTypePropertyNames) Set(val *IndicatorTypePropertyNames) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorTypePropertyNames) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorTypePropertyNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorTypePropertyNames(val *IndicatorTypePropertyNames) *NullableIndicatorTypePropertyNames {
	return &NullableIndicatorTypePropertyNames{value: val, isSet: true}
}

func (v NullableIndicatorTypePropertyNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorTypePropertyNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
