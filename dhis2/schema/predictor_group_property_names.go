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

// PredictorGroupPropertyNames the model 'PredictorGroupPropertyNames'
type PredictorGroupPropertyNames string

// List of PredictorGroupPropertyNames
const (
	PREDICTORGROUPPROPERTYNAMES_ACCESS           PredictorGroupPropertyNames = "access"
	PREDICTORGROUPPROPERTYNAMES_ATTRIBUTE_VALUES PredictorGroupPropertyNames = "attributeValues"
	PREDICTORGROUPPROPERTYNAMES_CODE             PredictorGroupPropertyNames = "code"
	PREDICTORGROUPPROPERTYNAMES_CREATED          PredictorGroupPropertyNames = "created"
	PREDICTORGROUPPROPERTYNAMES_CREATED_BY       PredictorGroupPropertyNames = "createdBy"
	PREDICTORGROUPPROPERTYNAMES_DESCRIPTION      PredictorGroupPropertyNames = "description"
	PREDICTORGROUPPROPERTYNAMES_DISPLAY_NAME     PredictorGroupPropertyNames = "displayName"
	PREDICTORGROUPPROPERTYNAMES_FAVORITE         PredictorGroupPropertyNames = "favorite"
	PREDICTORGROUPPROPERTYNAMES_FAVORITES        PredictorGroupPropertyNames = "favorites"
	PREDICTORGROUPPROPERTYNAMES_HREF             PredictorGroupPropertyNames = "href"
	PREDICTORGROUPPROPERTYNAMES_ID               PredictorGroupPropertyNames = "id"
	PREDICTORGROUPPROPERTYNAMES_LAST_UPDATED     PredictorGroupPropertyNames = "lastUpdated"
	PREDICTORGROUPPROPERTYNAMES_LAST_UPDATED_BY  PredictorGroupPropertyNames = "lastUpdatedBy"
	PREDICTORGROUPPROPERTYNAMES_NAME             PredictorGroupPropertyNames = "name"
	PREDICTORGROUPPROPERTYNAMES_PREDICTORS       PredictorGroupPropertyNames = "predictors"
	PREDICTORGROUPPROPERTYNAMES_SHARING          PredictorGroupPropertyNames = "sharing"
	PREDICTORGROUPPROPERTYNAMES_TRANSLATIONS     PredictorGroupPropertyNames = "translations"
)

// All allowed values of PredictorGroupPropertyNames enum
var AllowedPredictorGroupPropertyNamesEnumValues = []PredictorGroupPropertyNames{
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
	"predictors",
	"sharing",
	"translations",
}

func (v *PredictorGroupPropertyNames) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PredictorGroupPropertyNames(value)
	for _, existing := range AllowedPredictorGroupPropertyNamesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PredictorGroupPropertyNames", value)
}

// NewPredictorGroupPropertyNamesFromValue returns a pointer to a valid PredictorGroupPropertyNames
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPredictorGroupPropertyNamesFromValue(v string) (*PredictorGroupPropertyNames, error) {
	ev := PredictorGroupPropertyNames(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PredictorGroupPropertyNames: valid values are %v", v, AllowedPredictorGroupPropertyNamesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PredictorGroupPropertyNames) IsValid() bool {
	for _, existing := range AllowedPredictorGroupPropertyNamesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PredictorGroupPropertyNames value
func (v PredictorGroupPropertyNames) Ptr() *PredictorGroupPropertyNames {
	return &v
}

type NullablePredictorGroupPropertyNames struct {
	value *PredictorGroupPropertyNames
	isSet bool
}

func (v NullablePredictorGroupPropertyNames) Get() *PredictorGroupPropertyNames {
	return v.value
}

func (v *NullablePredictorGroupPropertyNames) Set(val *PredictorGroupPropertyNames) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictorGroupPropertyNames) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictorGroupPropertyNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictorGroupPropertyNames(val *PredictorGroupPropertyNames) *NullablePredictorGroupPropertyNames {
	return &NullablePredictorGroupPropertyNames{value: val, isSet: true}
}

func (v NullablePredictorGroupPropertyNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictorGroupPropertyNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
