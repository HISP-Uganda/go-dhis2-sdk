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

// DashboardItemType the model 'DashboardItemType'
type DashboardItemType string

// List of DashboardItemType
const (
	DASHBOARDITEMTYPE_VISUALIZATION       DashboardItemType = "VISUALIZATION"
	DASHBOARDITEMTYPE_EVENT_VISUALIZATION DashboardItemType = "EVENT_VISUALIZATION"
	DASHBOARDITEMTYPE_EVENT_CHART         DashboardItemType = "EVENT_CHART"
	DASHBOARDITEMTYPE_MAP                 DashboardItemType = "MAP"
	DASHBOARDITEMTYPE_EVENT_REPORT        DashboardItemType = "EVENT_REPORT"
	DASHBOARDITEMTYPE_USERS               DashboardItemType = "USERS"
	DASHBOARDITEMTYPE_REPORTS             DashboardItemType = "REPORTS"
	DASHBOARDITEMTYPE_RESOURCES           DashboardItemType = "RESOURCES"
	DASHBOARDITEMTYPE_TEXT                DashboardItemType = "TEXT"
	DASHBOARDITEMTYPE_MESSAGES            DashboardItemType = "MESSAGES"
	DASHBOARDITEMTYPE_APP                 DashboardItemType = "APP"
)

// All allowed values of DashboardItemType enum
var AllowedDashboardItemTypeEnumValues = []DashboardItemType{
	"VISUALIZATION",
	"EVENT_VISUALIZATION",
	"EVENT_CHART",
	"MAP",
	"EVENT_REPORT",
	"USERS",
	"REPORTS",
	"RESOURCES",
	"TEXT",
	"MESSAGES",
	"APP",
}

func (v *DashboardItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DashboardItemType(value)
	for _, existing := range AllowedDashboardItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DashboardItemType", value)
}

// NewDashboardItemTypeFromValue returns a pointer to a valid DashboardItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDashboardItemTypeFromValue(v string) (*DashboardItemType, error) {
	ev := DashboardItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DashboardItemType: valid values are %v", v, AllowedDashboardItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DashboardItemType) IsValid() bool {
	for _, existing := range AllowedDashboardItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardItemType value
func (v DashboardItemType) Ptr() *DashboardItemType {
	return &v
}

type NullableDashboardItemType struct {
	value *DashboardItemType
	isSet bool
}

func (v NullableDashboardItemType) Get() *DashboardItemType {
	return v.value
}

func (v *NullableDashboardItemType) Set(val *DashboardItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardItemType(val *DashboardItemType) *NullableDashboardItemType {
	return &NullableDashboardItemType{value: val, isSet: true}
}

func (v NullableDashboardItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
