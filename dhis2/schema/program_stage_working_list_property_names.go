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

// ProgramStageWorkingListPropertyNames the model 'ProgramStageWorkingListPropertyNames'
type ProgramStageWorkingListPropertyNames string

// List of ProgramStageWorkingListPropertyNames
const (
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_ACCESS                       ProgramStageWorkingListPropertyNames = "access"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_ATTRIBUTE_VALUES             ProgramStageWorkingListPropertyNames = "attributeValues"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_CODE                         ProgramStageWorkingListPropertyNames = "code"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_CREATED                      ProgramStageWorkingListPropertyNames = "created"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_CREATED_BY                   ProgramStageWorkingListPropertyNames = "createdBy"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_DESCRIPTION                  ProgramStageWorkingListPropertyNames = "description"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_DISPLAY_DESCRIPTION          ProgramStageWorkingListPropertyNames = "displayDescription"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_DISPLAY_NAME                 ProgramStageWorkingListPropertyNames = "displayName"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_FAVORITE                     ProgramStageWorkingListPropertyNames = "favorite"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_FAVORITES                    ProgramStageWorkingListPropertyNames = "favorites"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_HREF                         ProgramStageWorkingListPropertyNames = "href"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_ID                           ProgramStageWorkingListPropertyNames = "id"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_LAST_UPDATED                 ProgramStageWorkingListPropertyNames = "lastUpdated"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_LAST_UPDATED_BY              ProgramStageWorkingListPropertyNames = "lastUpdatedBy"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_NAME                         ProgramStageWorkingListPropertyNames = "name"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_PROGRAM                      ProgramStageWorkingListPropertyNames = "program"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_PROGRAM_STAGE                ProgramStageWorkingListPropertyNames = "programStage"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_PROGRAM_STAGE_QUERY_CRITERIA ProgramStageWorkingListPropertyNames = "programStageQueryCriteria"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_SHARING                      ProgramStageWorkingListPropertyNames = "sharing"
	PROGRAMSTAGEWORKINGLISTPROPERTYNAMES_TRANSLATIONS                 ProgramStageWorkingListPropertyNames = "translations"
)

// All allowed values of ProgramStageWorkingListPropertyNames enum
var AllowedProgramStageWorkingListPropertyNamesEnumValues = []ProgramStageWorkingListPropertyNames{
	"access",
	"attributeValues",
	"code",
	"created",
	"createdBy",
	"description",
	"displayDescription",
	"displayName",
	"favorite",
	"favorites",
	"href",
	"id",
	"lastUpdated",
	"lastUpdatedBy",
	"name",
	"program",
	"programStage",
	"programStageQueryCriteria",
	"sharing",
	"translations",
}

func (v *ProgramStageWorkingListPropertyNames) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProgramStageWorkingListPropertyNames(value)
	for _, existing := range AllowedProgramStageWorkingListPropertyNamesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProgramStageWorkingListPropertyNames", value)
}

// NewProgramStageWorkingListPropertyNamesFromValue returns a pointer to a valid ProgramStageWorkingListPropertyNames
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProgramStageWorkingListPropertyNamesFromValue(v string) (*ProgramStageWorkingListPropertyNames, error) {
	ev := ProgramStageWorkingListPropertyNames(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProgramStageWorkingListPropertyNames: valid values are %v", v, AllowedProgramStageWorkingListPropertyNamesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProgramStageWorkingListPropertyNames) IsValid() bool {
	for _, existing := range AllowedProgramStageWorkingListPropertyNamesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProgramStageWorkingListPropertyNames value
func (v ProgramStageWorkingListPropertyNames) Ptr() *ProgramStageWorkingListPropertyNames {
	return &v
}

type NullableProgramStageWorkingListPropertyNames struct {
	value *ProgramStageWorkingListPropertyNames
	isSet bool
}

func (v NullableProgramStageWorkingListPropertyNames) Get() *ProgramStageWorkingListPropertyNames {
	return v.value
}

func (v *NullableProgramStageWorkingListPropertyNames) Set(val *ProgramStageWorkingListPropertyNames) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramStageWorkingListPropertyNames) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramStageWorkingListPropertyNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramStageWorkingListPropertyNames(val *ProgramStageWorkingListPropertyNames) *NullableProgramStageWorkingListPropertyNames {
	return &NullableProgramStageWorkingListPropertyNames{value: val, isSet: true}
}

func (v NullableProgramStageWorkingListPropertyNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramStageWorkingListPropertyNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
