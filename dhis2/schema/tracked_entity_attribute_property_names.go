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

// TrackedEntityAttributePropertyNames the model 'TrackedEntityAttributePropertyNames'
type TrackedEntityAttributePropertyNames string

// List of TrackedEntityAttributePropertyNames
const (
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_ACCESS                        TrackedEntityAttributePropertyNames = "access"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_AGGREGATION_TYPE              TrackedEntityAttributePropertyNames = "aggregationType"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_ATTRIBUTE_VALUES              TrackedEntityAttributePropertyNames = "attributeValues"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_CODE                          TrackedEntityAttributePropertyNames = "code"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_CONFIDENTIAL                  TrackedEntityAttributePropertyNames = "confidential"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_CREATED                       TrackedEntityAttributePropertyNames = "created"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_CREATED_BY                    TrackedEntityAttributePropertyNames = "createdBy"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_DESCRIPTION                   TrackedEntityAttributePropertyNames = "description"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_DIMENSION_ITEM                TrackedEntityAttributePropertyNames = "dimensionItem"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_DISPLAY_DESCRIPTION           TrackedEntityAttributePropertyNames = "displayDescription"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_DISPLAY_FORM_NAME             TrackedEntityAttributePropertyNames = "displayFormName"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_DISPLAY_IN_LIST_NO_PROGRAM    TrackedEntityAttributePropertyNames = "displayInListNoProgram"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_DISPLAY_NAME                  TrackedEntityAttributePropertyNames = "displayName"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_DISPLAY_ON_VISIT_SCHEDULE     TrackedEntityAttributePropertyNames = "displayOnVisitSchedule"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_DISPLAY_SHORT_NAME            TrackedEntityAttributePropertyNames = "displayShortName"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_EXPRESSION                    TrackedEntityAttributePropertyNames = "expression"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_FAVORITE                      TrackedEntityAttributePropertyNames = "favorite"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_FAVORITES                     TrackedEntityAttributePropertyNames = "favorites"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_FIELD_MASK                    TrackedEntityAttributePropertyNames = "fieldMask"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_FORM_NAME                     TrackedEntityAttributePropertyNames = "formName"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_GENERATED                     TrackedEntityAttributePropertyNames = "generated"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_HREF                          TrackedEntityAttributePropertyNames = "href"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_ID                            TrackedEntityAttributePropertyNames = "id"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_INHERIT                       TrackedEntityAttributePropertyNames = "inherit"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_LAST_UPDATED                  TrackedEntityAttributePropertyNames = "lastUpdated"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_LAST_UPDATED_BY               TrackedEntityAttributePropertyNames = "lastUpdatedBy"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_LEGEND_SET                    TrackedEntityAttributePropertyNames = "legendSet"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_LEGEND_SETS                   TrackedEntityAttributePropertyNames = "legendSets"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_NAME                          TrackedEntityAttributePropertyNames = "name"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_OPTION_SET                    TrackedEntityAttributePropertyNames = "optionSet"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_OPTION_SET_VALUE              TrackedEntityAttributePropertyNames = "optionSetValue"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_ORGUNIT_SCOPE                 TrackedEntityAttributePropertyNames = "orgunitScope"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_PATTERN                       TrackedEntityAttributePropertyNames = "pattern"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_QUERY_MODS                    TrackedEntityAttributePropertyNames = "queryMods"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_SHARING                       TrackedEntityAttributePropertyNames = "sharing"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_SHORT_NAME                    TrackedEntityAttributePropertyNames = "shortName"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_SKIP_SYNCHRONIZATION          TrackedEntityAttributePropertyNames = "skipSynchronization"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_SORT_ORDER_IN_LIST_NO_PROGRAM TrackedEntityAttributePropertyNames = "sortOrderInListNoProgram"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_SORT_ORDER_IN_VISIT_SCHEDULE  TrackedEntityAttributePropertyNames = "sortOrderInVisitSchedule"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_STYLE                         TrackedEntityAttributePropertyNames = "style"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_TRANSLATIONS                  TrackedEntityAttributePropertyNames = "translations"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_UNIQUE                        TrackedEntityAttributePropertyNames = "unique"
	TRACKEDENTITYATTRIBUTEPROPERTYNAMES_VALUE_TYPE                    TrackedEntityAttributePropertyNames = "valueType"
)

// All allowed values of TrackedEntityAttributePropertyNames enum
var AllowedTrackedEntityAttributePropertyNamesEnumValues = []TrackedEntityAttributePropertyNames{
	"access",
	"aggregationType",
	"attributeValues",
	"code",
	"confidential",
	"created",
	"createdBy",
	"description",
	"dimensionItem",
	"displayDescription",
	"displayFormName",
	"displayInListNoProgram",
	"displayName",
	"displayOnVisitSchedule",
	"displayShortName",
	"expression",
	"favorite",
	"favorites",
	"fieldMask",
	"formName",
	"generated",
	"href",
	"id",
	"inherit",
	"lastUpdated",
	"lastUpdatedBy",
	"legendSet",
	"legendSets",
	"name",
	"optionSet",
	"optionSetValue",
	"orgunitScope",
	"pattern",
	"queryMods",
	"sharing",
	"shortName",
	"skipSynchronization",
	"sortOrderInListNoProgram",
	"sortOrderInVisitSchedule",
	"style",
	"translations",
	"unique",
	"valueType",
}

func (v *TrackedEntityAttributePropertyNames) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrackedEntityAttributePropertyNames(value)
	for _, existing := range AllowedTrackedEntityAttributePropertyNamesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrackedEntityAttributePropertyNames", value)
}

// NewTrackedEntityAttributePropertyNamesFromValue returns a pointer to a valid TrackedEntityAttributePropertyNames
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrackedEntityAttributePropertyNamesFromValue(v string) (*TrackedEntityAttributePropertyNames, error) {
	ev := TrackedEntityAttributePropertyNames(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrackedEntityAttributePropertyNames: valid values are %v", v, AllowedTrackedEntityAttributePropertyNamesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrackedEntityAttributePropertyNames) IsValid() bool {
	for _, existing := range AllowedTrackedEntityAttributePropertyNamesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrackedEntityAttributePropertyNames value
func (v TrackedEntityAttributePropertyNames) Ptr() *TrackedEntityAttributePropertyNames {
	return &v
}

type NullableTrackedEntityAttributePropertyNames struct {
	value *TrackedEntityAttributePropertyNames
	isSet bool
}

func (v NullableTrackedEntityAttributePropertyNames) Get() *TrackedEntityAttributePropertyNames {
	return v.value
}

func (v *NullableTrackedEntityAttributePropertyNames) Set(val *TrackedEntityAttributePropertyNames) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackedEntityAttributePropertyNames) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackedEntityAttributePropertyNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackedEntityAttributePropertyNames(val *TrackedEntityAttributePropertyNames) *NullableTrackedEntityAttributePropertyNames {
	return &NullableTrackedEntityAttributePropertyNames{value: val, isSet: true}
}

func (v NullableTrackedEntityAttributePropertyNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackedEntityAttributePropertyNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
