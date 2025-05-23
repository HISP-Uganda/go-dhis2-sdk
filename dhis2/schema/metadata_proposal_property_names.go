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

// MetadataProposalPropertyNames the model 'MetadataProposalPropertyNames'
type MetadataProposalPropertyNames string

// List of MetadataProposalPropertyNames
const (
	METADATAPROPOSALPROPERTYNAMES_CHANGE       MetadataProposalPropertyNames = "change"
	METADATAPROPOSALPROPERTYNAMES_COMMENT      MetadataProposalPropertyNames = "comment"
	METADATAPROPOSALPROPERTYNAMES_CREATED      MetadataProposalPropertyNames = "created"
	METADATAPROPOSALPROPERTYNAMES_CREATED_BY   MetadataProposalPropertyNames = "createdBy"
	METADATAPROPOSALPROPERTYNAMES_FINALISED    MetadataProposalPropertyNames = "finalised"
	METADATAPROPOSALPROPERTYNAMES_FINALISED_BY MetadataProposalPropertyNames = "finalisedBy"
	METADATAPROPOSALPROPERTYNAMES_ID           MetadataProposalPropertyNames = "id"
	METADATAPROPOSALPROPERTYNAMES_REASON       MetadataProposalPropertyNames = "reason"
	METADATAPROPOSALPROPERTYNAMES_STATUS       MetadataProposalPropertyNames = "status"
	METADATAPROPOSALPROPERTYNAMES_TARGET       MetadataProposalPropertyNames = "target"
	METADATAPROPOSALPROPERTYNAMES_TARGET_ID    MetadataProposalPropertyNames = "targetId"
	METADATAPROPOSALPROPERTYNAMES_TYPE         MetadataProposalPropertyNames = "type"
)

// All allowed values of MetadataProposalPropertyNames enum
var AllowedMetadataProposalPropertyNamesEnumValues = []MetadataProposalPropertyNames{
	"change",
	"comment",
	"created",
	"createdBy",
	"finalised",
	"finalisedBy",
	"id",
	"reason",
	"status",
	"target",
	"targetId",
	"type",
}

func (v *MetadataProposalPropertyNames) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetadataProposalPropertyNames(value)
	for _, existing := range AllowedMetadataProposalPropertyNamesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetadataProposalPropertyNames", value)
}

// NewMetadataProposalPropertyNamesFromValue returns a pointer to a valid MetadataProposalPropertyNames
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetadataProposalPropertyNamesFromValue(v string) (*MetadataProposalPropertyNames, error) {
	ev := MetadataProposalPropertyNames(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetadataProposalPropertyNames: valid values are %v", v, AllowedMetadataProposalPropertyNamesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetadataProposalPropertyNames) IsValid() bool {
	for _, existing := range AllowedMetadataProposalPropertyNamesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetadataProposalPropertyNames value
func (v MetadataProposalPropertyNames) Ptr() *MetadataProposalPropertyNames {
	return &v
}

type NullableMetadataProposalPropertyNames struct {
	value *MetadataProposalPropertyNames
	isSet bool
}

func (v NullableMetadataProposalPropertyNames) Get() *MetadataProposalPropertyNames {
	return v.value
}

func (v *NullableMetadataProposalPropertyNames) Set(val *MetadataProposalPropertyNames) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataProposalPropertyNames) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataProposalPropertyNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataProposalPropertyNames(val *MetadataProposalPropertyNames) *NullableMetadataProposalPropertyNames {
	return &NullableMetadataProposalPropertyNames{value: val, isSet: true}
}

func (v NullableMetadataProposalPropertyNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataProposalPropertyNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
