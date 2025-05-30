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

// ValidationRuleGroupPropertyNames the model 'ValidationRuleGroupPropertyNames'
type ValidationRuleGroupPropertyNames string

// List of ValidationRuleGroupPropertyNames
const (
	VALIDATIONRULEGROUPPROPERTYNAMES_ACCESS           ValidationRuleGroupPropertyNames = "access"
	VALIDATIONRULEGROUPPROPERTYNAMES_ATTRIBUTE_VALUES ValidationRuleGroupPropertyNames = "attributeValues"
	VALIDATIONRULEGROUPPROPERTYNAMES_CODE             ValidationRuleGroupPropertyNames = "code"
	VALIDATIONRULEGROUPPROPERTYNAMES_CREATED          ValidationRuleGroupPropertyNames = "created"
	VALIDATIONRULEGROUPPROPERTYNAMES_CREATED_BY       ValidationRuleGroupPropertyNames = "createdBy"
	VALIDATIONRULEGROUPPROPERTYNAMES_DESCRIPTION      ValidationRuleGroupPropertyNames = "description"
	VALIDATIONRULEGROUPPROPERTYNAMES_DISPLAY_NAME     ValidationRuleGroupPropertyNames = "displayName"
	VALIDATIONRULEGROUPPROPERTYNAMES_FAVORITE         ValidationRuleGroupPropertyNames = "favorite"
	VALIDATIONRULEGROUPPROPERTYNAMES_FAVORITES        ValidationRuleGroupPropertyNames = "favorites"
	VALIDATIONRULEGROUPPROPERTYNAMES_HREF             ValidationRuleGroupPropertyNames = "href"
	VALIDATIONRULEGROUPPROPERTYNAMES_ID               ValidationRuleGroupPropertyNames = "id"
	VALIDATIONRULEGROUPPROPERTYNAMES_LAST_UPDATED     ValidationRuleGroupPropertyNames = "lastUpdated"
	VALIDATIONRULEGROUPPROPERTYNAMES_LAST_UPDATED_BY  ValidationRuleGroupPropertyNames = "lastUpdatedBy"
	VALIDATIONRULEGROUPPROPERTYNAMES_NAME             ValidationRuleGroupPropertyNames = "name"
	VALIDATIONRULEGROUPPROPERTYNAMES_SHARING          ValidationRuleGroupPropertyNames = "sharing"
	VALIDATIONRULEGROUPPROPERTYNAMES_TRANSLATIONS     ValidationRuleGroupPropertyNames = "translations"
	VALIDATIONRULEGROUPPROPERTYNAMES_VALIDATION_RULES ValidationRuleGroupPropertyNames = "validationRules"
)

// All allowed values of ValidationRuleGroupPropertyNames enum
var AllowedValidationRuleGroupPropertyNamesEnumValues = []ValidationRuleGroupPropertyNames{
	"access",
	"attributeValues",
	"code",
	"created",
	"createdBy",
	"description",
	"displayName",
	"favorite",
	"favorites",
	"href",
	"id",
	"lastUpdated",
	"lastUpdatedBy",
	"name",
	"sharing",
	"translations",
	"validationRules",
}

func (v *ValidationRuleGroupPropertyNames) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ValidationRuleGroupPropertyNames(value)
	for _, existing := range AllowedValidationRuleGroupPropertyNamesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ValidationRuleGroupPropertyNames", value)
}

// NewValidationRuleGroupPropertyNamesFromValue returns a pointer to a valid ValidationRuleGroupPropertyNames
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewValidationRuleGroupPropertyNamesFromValue(v string) (*ValidationRuleGroupPropertyNames, error) {
	ev := ValidationRuleGroupPropertyNames(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ValidationRuleGroupPropertyNames: valid values are %v", v, AllowedValidationRuleGroupPropertyNamesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ValidationRuleGroupPropertyNames) IsValid() bool {
	for _, existing := range AllowedValidationRuleGroupPropertyNamesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ValidationRuleGroupPropertyNames value
func (v ValidationRuleGroupPropertyNames) Ptr() *ValidationRuleGroupPropertyNames {
	return &v
}

type NullableValidationRuleGroupPropertyNames struct {
	value *ValidationRuleGroupPropertyNames
	isSet bool
}

func (v NullableValidationRuleGroupPropertyNames) Get() *ValidationRuleGroupPropertyNames {
	return v.value
}

func (v *NullableValidationRuleGroupPropertyNames) Set(val *ValidationRuleGroupPropertyNames) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationRuleGroupPropertyNames) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationRuleGroupPropertyNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationRuleGroupPropertyNames(val *ValidationRuleGroupPropertyNames) *NullableValidationRuleGroupPropertyNames {
	return &NullableValidationRuleGroupPropertyNames{value: val, isSet: true}
}

func (v NullableValidationRuleGroupPropertyNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationRuleGroupPropertyNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
