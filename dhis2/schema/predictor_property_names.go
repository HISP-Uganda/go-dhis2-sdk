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

// PredictorPropertyNames the model 'PredictorPropertyNames'
type PredictorPropertyNames string

// List of PredictorPropertyNames
const (
	PREDICTORPROPERTYNAMES_ACCESS                        PredictorPropertyNames = "access"
	PREDICTORPROPERTYNAMES_ANNUAL_SAMPLE_COUNT           PredictorPropertyNames = "annualSampleCount"
	PREDICTORPROPERTYNAMES_ATTRIBUTE_VALUES              PredictorPropertyNames = "attributeValues"
	PREDICTORPROPERTYNAMES_CODE                          PredictorPropertyNames = "code"
	PREDICTORPROPERTYNAMES_CREATED                       PredictorPropertyNames = "created"
	PREDICTORPROPERTYNAMES_CREATED_BY                    PredictorPropertyNames = "createdBy"
	PREDICTORPROPERTYNAMES_DESCRIPTION                   PredictorPropertyNames = "description"
	PREDICTORPROPERTYNAMES_DISPLAY_DESCRIPTION           PredictorPropertyNames = "displayDescription"
	PREDICTORPROPERTYNAMES_DISPLAY_FORM_NAME             PredictorPropertyNames = "displayFormName"
	PREDICTORPROPERTYNAMES_DISPLAY_NAME                  PredictorPropertyNames = "displayName"
	PREDICTORPROPERTYNAMES_DISPLAY_SHORT_NAME            PredictorPropertyNames = "displayShortName"
	PREDICTORPROPERTYNAMES_FAVORITE                      PredictorPropertyNames = "favorite"
	PREDICTORPROPERTYNAMES_FAVORITES                     PredictorPropertyNames = "favorites"
	PREDICTORPROPERTYNAMES_FORM_NAME                     PredictorPropertyNames = "formName"
	PREDICTORPROPERTYNAMES_GENERATOR                     PredictorPropertyNames = "generator"
	PREDICTORPROPERTYNAMES_HREF                          PredictorPropertyNames = "href"
	PREDICTORPROPERTYNAMES_ID                            PredictorPropertyNames = "id"
	PREDICTORPROPERTYNAMES_LAST_UPDATED                  PredictorPropertyNames = "lastUpdated"
	PREDICTORPROPERTYNAMES_LAST_UPDATED_BY               PredictorPropertyNames = "lastUpdatedBy"
	PREDICTORPROPERTYNAMES_NAME                          PredictorPropertyNames = "name"
	PREDICTORPROPERTYNAMES_ORGANISATION_UNIT_DESCENDANTS PredictorPropertyNames = "organisationUnitDescendants"
	PREDICTORPROPERTYNAMES_ORGANISATION_UNIT_LEVELS      PredictorPropertyNames = "organisationUnitLevels"
	PREDICTORPROPERTYNAMES_OUTPUT                        PredictorPropertyNames = "output"
	PREDICTORPROPERTYNAMES_OUTPUT_COMBO                  PredictorPropertyNames = "outputCombo"
	PREDICTORPROPERTYNAMES_PERIOD_TYPE                   PredictorPropertyNames = "periodType"
	PREDICTORPROPERTYNAMES_PREDICTOR_GROUPS              PredictorPropertyNames = "predictorGroups"
	PREDICTORPROPERTYNAMES_SAMPLE_SKIP_TEST              PredictorPropertyNames = "sampleSkipTest"
	PREDICTORPROPERTYNAMES_SEQUENTIAL_SAMPLE_COUNT       PredictorPropertyNames = "sequentialSampleCount"
	PREDICTORPROPERTYNAMES_SEQUENTIAL_SKIP_COUNT         PredictorPropertyNames = "sequentialSkipCount"
	PREDICTORPROPERTYNAMES_SHARING                       PredictorPropertyNames = "sharing"
	PREDICTORPROPERTYNAMES_SHORT_NAME                    PredictorPropertyNames = "shortName"
	PREDICTORPROPERTYNAMES_TRANSLATIONS                  PredictorPropertyNames = "translations"
)

// All allowed values of PredictorPropertyNames enum
var AllowedPredictorPropertyNamesEnumValues = []PredictorPropertyNames{
	"access",
	"annualSampleCount",
	"attributeValues",
	"code",
	"created",
	"createdBy",
	"description",
	"displayDescription",
	"displayFormName",
	"displayName",
	"displayShortName",
	"favorite",
	"favorites",
	"formName",
	"generator",
	"href",
	"id",
	"lastUpdated",
	"lastUpdatedBy",
	"name",
	"organisationUnitDescendants",
	"organisationUnitLevels",
	"output",
	"outputCombo",
	"periodType",
	"predictorGroups",
	"sampleSkipTest",
	"sequentialSampleCount",
	"sequentialSkipCount",
	"sharing",
	"shortName",
	"translations",
}

func (v *PredictorPropertyNames) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PredictorPropertyNames(value)
	for _, existing := range AllowedPredictorPropertyNamesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PredictorPropertyNames", value)
}

// NewPredictorPropertyNamesFromValue returns a pointer to a valid PredictorPropertyNames
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPredictorPropertyNamesFromValue(v string) (*PredictorPropertyNames, error) {
	ev := PredictorPropertyNames(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PredictorPropertyNames: valid values are %v", v, AllowedPredictorPropertyNamesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PredictorPropertyNames) IsValid() bool {
	for _, existing := range AllowedPredictorPropertyNamesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PredictorPropertyNames value
func (v PredictorPropertyNames) Ptr() *PredictorPropertyNames {
	return &v
}

type NullablePredictorPropertyNames struct {
	value *PredictorPropertyNames
	isSet bool
}

func (v NullablePredictorPropertyNames) Get() *PredictorPropertyNames {
	return v.value
}

func (v *NullablePredictorPropertyNames) Set(val *PredictorPropertyNames) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictorPropertyNames) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictorPropertyNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictorPropertyNames(val *PredictorPropertyNames) *NullablePredictorPropertyNames {
	return &NullablePredictorPropertyNames{value: val, isSet: true}
}

func (v NullablePredictorPropertyNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictorPropertyNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
