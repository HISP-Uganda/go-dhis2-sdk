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

// IndicatorGroupSetPropertyNames the model 'IndicatorGroupSetPropertyNames'
type IndicatorGroupSetPropertyNames string

// List of IndicatorGroupSetPropertyNames
const (
	INDICATORGROUPSETPROPERTYNAMES_ACCESS           IndicatorGroupSetPropertyNames = "access"
	INDICATORGROUPSETPROPERTYNAMES_ATTRIBUTE_VALUES IndicatorGroupSetPropertyNames = "attributeValues"
	INDICATORGROUPSETPROPERTYNAMES_CODE             IndicatorGroupSetPropertyNames = "code"
	INDICATORGROUPSETPROPERTYNAMES_COMPULSORY       IndicatorGroupSetPropertyNames = "compulsory"
	INDICATORGROUPSETPROPERTYNAMES_CREATED          IndicatorGroupSetPropertyNames = "created"
	INDICATORGROUPSETPROPERTYNAMES_CREATED_BY       IndicatorGroupSetPropertyNames = "createdBy"
	INDICATORGROUPSETPROPERTYNAMES_DESCRIPTION      IndicatorGroupSetPropertyNames = "description"
	INDICATORGROUPSETPROPERTYNAMES_DISPLAY_NAME     IndicatorGroupSetPropertyNames = "displayName"
	INDICATORGROUPSETPROPERTYNAMES_FAVORITE         IndicatorGroupSetPropertyNames = "favorite"
	INDICATORGROUPSETPROPERTYNAMES_FAVORITES        IndicatorGroupSetPropertyNames = "favorites"
	INDICATORGROUPSETPROPERTYNAMES_HREF             IndicatorGroupSetPropertyNames = "href"
	INDICATORGROUPSETPROPERTYNAMES_ID               IndicatorGroupSetPropertyNames = "id"
	INDICATORGROUPSETPROPERTYNAMES_INDICATOR_GROUPS IndicatorGroupSetPropertyNames = "indicatorGroups"
	INDICATORGROUPSETPROPERTYNAMES_LAST_UPDATED     IndicatorGroupSetPropertyNames = "lastUpdated"
	INDICATORGROUPSETPROPERTYNAMES_LAST_UPDATED_BY  IndicatorGroupSetPropertyNames = "lastUpdatedBy"
	INDICATORGROUPSETPROPERTYNAMES_NAME             IndicatorGroupSetPropertyNames = "name"
	INDICATORGROUPSETPROPERTYNAMES_SHARING          IndicatorGroupSetPropertyNames = "sharing"
	INDICATORGROUPSETPROPERTYNAMES_SHORT_NAME       IndicatorGroupSetPropertyNames = "shortName"
	INDICATORGROUPSETPROPERTYNAMES_TRANSLATIONS     IndicatorGroupSetPropertyNames = "translations"
)

// All allowed values of IndicatorGroupSetPropertyNames enum
var AllowedIndicatorGroupSetPropertyNamesEnumValues = []IndicatorGroupSetPropertyNames{
	"access",
	"attributeValues",
	"code",
	"compulsory",
	"created",
	"createdBy",
	"description",
	"displayName",
	"favorite",
	"favorites",
	"href",
	"id",
	"indicatorGroups",
	"lastUpdated",
	"lastUpdatedBy",
	"name",
	"sharing",
	"shortName",
	"translations",
}

func (v *IndicatorGroupSetPropertyNames) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IndicatorGroupSetPropertyNames(value)
	for _, existing := range AllowedIndicatorGroupSetPropertyNamesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IndicatorGroupSetPropertyNames", value)
}

// NewIndicatorGroupSetPropertyNamesFromValue returns a pointer to a valid IndicatorGroupSetPropertyNames
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIndicatorGroupSetPropertyNamesFromValue(v string) (*IndicatorGroupSetPropertyNames, error) {
	ev := IndicatorGroupSetPropertyNames(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IndicatorGroupSetPropertyNames: valid values are %v", v, AllowedIndicatorGroupSetPropertyNamesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IndicatorGroupSetPropertyNames) IsValid() bool {
	for _, existing := range AllowedIndicatorGroupSetPropertyNamesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IndicatorGroupSetPropertyNames value
func (v IndicatorGroupSetPropertyNames) Ptr() *IndicatorGroupSetPropertyNames {
	return &v
}

type NullableIndicatorGroupSetPropertyNames struct {
	value *IndicatorGroupSetPropertyNames
	isSet bool
}

func (v NullableIndicatorGroupSetPropertyNames) Get() *IndicatorGroupSetPropertyNames {
	return v.value
}

func (v *NullableIndicatorGroupSetPropertyNames) Set(val *IndicatorGroupSetPropertyNames) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorGroupSetPropertyNames) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorGroupSetPropertyNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorGroupSetPropertyNames(val *IndicatorGroupSetPropertyNames) *NullableIndicatorGroupSetPropertyNames {
	return &NullableIndicatorGroupSetPropertyNames{value: val, isSet: true}
}

func (v NullableIndicatorGroupSetPropertyNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorGroupSetPropertyNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
