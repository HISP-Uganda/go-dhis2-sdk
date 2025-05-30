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

// MapService the model 'MapService'
type MapService string

// List of MapService
const (
	MAPSERVICE_WMS            MapService = "WMS"
	MAPSERVICE_TMS            MapService = "TMS"
	MAPSERVICE_XYZ            MapService = "XYZ"
	MAPSERVICE_VECTOR_STYLE   MapService = "VECTOR_STYLE"
	MAPSERVICE_GEOJSON_URL    MapService = "GEOJSON_URL"
	MAPSERVICE_ARCGIS_FEATURE MapService = "ARCGIS_FEATURE"
)

// All allowed values of MapService enum
var AllowedMapServiceEnumValues = []MapService{
	"WMS",
	"TMS",
	"XYZ",
	"VECTOR_STYLE",
	"GEOJSON_URL",
	"ARCGIS_FEATURE",
}

func (v *MapService) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MapService(value)
	for _, existing := range AllowedMapServiceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MapService", value)
}

// NewMapServiceFromValue returns a pointer to a valid MapService
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMapServiceFromValue(v string) (*MapService, error) {
	ev := MapService(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MapService: valid values are %v", v, AllowedMapServiceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MapService) IsValid() bool {
	for _, existing := range AllowedMapServiceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MapService value
func (v MapService) Ptr() *MapService {
	return &v
}

type NullableMapService struct {
	value *MapService
	isSet bool
}

func (v NullableMapService) Get() *MapService {
	return v.value
}

func (v *NullableMapService) Set(val *MapService) {
	v.value = val
	v.isSet = true
}

func (v NullableMapService) IsSet() bool {
	return v.isSet
}

func (v *NullableMapService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapService(val *MapService) *NullableMapService {
	return &NullableMapService{value: val, isSet: true}
}

func (v NullableMapService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
