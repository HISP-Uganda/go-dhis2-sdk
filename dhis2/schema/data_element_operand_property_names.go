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

// DataElementOperandPropertyNames the model 'DataElementOperandPropertyNames'
type DataElementOperandPropertyNames string

// List of DataElementOperandPropertyNames
const (
	DATAELEMENTOPERANDPROPERTYNAMES_ACCESS                 DataElementOperandPropertyNames = "access"
	DATAELEMENTOPERANDPROPERTYNAMES_ATTRIBUTE_OPTION_COMBO DataElementOperandPropertyNames = "attributeOptionCombo"
	DATAELEMENTOPERANDPROPERTYNAMES_ATTRIBUTE_VALUES       DataElementOperandPropertyNames = "attributeValues"
	DATAELEMENTOPERANDPROPERTYNAMES_CATEGORY_OPTION_COMBO  DataElementOperandPropertyNames = "categoryOptionCombo"
	DATAELEMENTOPERANDPROPERTYNAMES_CODE                   DataElementOperandPropertyNames = "code"
	DATAELEMENTOPERANDPROPERTYNAMES_CREATED                DataElementOperandPropertyNames = "created"
	DATAELEMENTOPERANDPROPERTYNAMES_CREATED_BY             DataElementOperandPropertyNames = "createdBy"
	DATAELEMENTOPERANDPROPERTYNAMES_DATA_ELEMENT           DataElementOperandPropertyNames = "dataElement"
	DATAELEMENTOPERANDPROPERTYNAMES_DESCRIPTION            DataElementOperandPropertyNames = "description"
	DATAELEMENTOPERANDPROPERTYNAMES_DISPLAY_DESCRIPTION    DataElementOperandPropertyNames = "displayDescription"
	DATAELEMENTOPERANDPROPERTYNAMES_DISPLAY_FORM_NAME      DataElementOperandPropertyNames = "displayFormName"
	DATAELEMENTOPERANDPROPERTYNAMES_FAVORITE               DataElementOperandPropertyNames = "favorite"
	DATAELEMENTOPERANDPROPERTYNAMES_FAVORITES              DataElementOperandPropertyNames = "favorites"
	DATAELEMENTOPERANDPROPERTYNAMES_FORM_NAME              DataElementOperandPropertyNames = "formName"
	DATAELEMENTOPERANDPROPERTYNAMES_HREF                   DataElementOperandPropertyNames = "href"
	DATAELEMENTOPERANDPROPERTYNAMES_LAST_UPDATED           DataElementOperandPropertyNames = "lastUpdated"
	DATAELEMENTOPERANDPROPERTYNAMES_LAST_UPDATED_BY        DataElementOperandPropertyNames = "lastUpdatedBy"
	DATAELEMENTOPERANDPROPERTYNAMES_LEGEND_SET             DataElementOperandPropertyNames = "legendSet"
	DATAELEMENTOPERANDPROPERTYNAMES_LEGEND_SETS            DataElementOperandPropertyNames = "legendSets"
	DATAELEMENTOPERANDPROPERTYNAMES_QUERY_MODS             DataElementOperandPropertyNames = "queryMods"
	DATAELEMENTOPERANDPROPERTYNAMES_SHARING                DataElementOperandPropertyNames = "sharing"
	DATAELEMENTOPERANDPROPERTYNAMES_TRANSLATIONS           DataElementOperandPropertyNames = "translations"
)

// All allowed values of DataElementOperandPropertyNames enum
var AllowedDataElementOperandPropertyNamesEnumValues = []DataElementOperandPropertyNames{
	"access",
	"attributeOptionCombo",
	"attributeValues",
	"categoryOptionCombo",
	"code",
	"created",
	"createdBy",
	"dataElement",
	"description",
	"displayDescription",
	"displayFormName",
	"favorite",
	"favorites",
	"formName",
	"href",
	"lastUpdated",
	"lastUpdatedBy",
	"legendSet",
	"legendSets",
	"queryMods",
	"sharing",
	"translations",
}

func (v *DataElementOperandPropertyNames) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataElementOperandPropertyNames(value)
	for _, existing := range AllowedDataElementOperandPropertyNamesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataElementOperandPropertyNames", value)
}

// NewDataElementOperandPropertyNamesFromValue returns a pointer to a valid DataElementOperandPropertyNames
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataElementOperandPropertyNamesFromValue(v string) (*DataElementOperandPropertyNames, error) {
	ev := DataElementOperandPropertyNames(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataElementOperandPropertyNames: valid values are %v", v, AllowedDataElementOperandPropertyNamesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataElementOperandPropertyNames) IsValid() bool {
	for _, existing := range AllowedDataElementOperandPropertyNamesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataElementOperandPropertyNames value
func (v DataElementOperandPropertyNames) Ptr() *DataElementOperandPropertyNames {
	return &v
}

type NullableDataElementOperandPropertyNames struct {
	value *DataElementOperandPropertyNames
	isSet bool
}

func (v NullableDataElementOperandPropertyNames) Get() *DataElementOperandPropertyNames {
	return v.value
}

func (v *NullableDataElementOperandPropertyNames) Set(val *DataElementOperandPropertyNames) {
	v.value = val
	v.isSet = true
}

func (v NullableDataElementOperandPropertyNames) IsSet() bool {
	return v.isSet
}

func (v *NullableDataElementOperandPropertyNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataElementOperandPropertyNames(val *DataElementOperandPropertyNames) *NullableDataElementOperandPropertyNames {
	return &NullableDataElementOperandPropertyNames{value: val, isSet: true}
}

func (v NullableDataElementOperandPropertyNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataElementOperandPropertyNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
