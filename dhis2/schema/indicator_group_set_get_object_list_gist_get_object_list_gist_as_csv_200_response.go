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

// IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response - struct for IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response
type IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response struct {
	IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf *IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	ArrayOfIndicatorGroupSet                                                 *[]IndicatorGroupSet
}

// IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf wrapped in IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response
func IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response(v *IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response{
		IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf: v,
	}
}

// []IndicatorGroupSetAsIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns []IndicatorGroupSet wrapped in IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response
func ArrayOfIndicatorGroupSetAsIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response(v *[]IndicatorGroupSet) IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response{
		ArrayOfIndicatorGroupSet: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	if err == nil {
		jsonIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf, _ := json.Marshal(dst.IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
		if string(jsonIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) == "{}" { // empty struct
			dst.IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfIndicatorGroupSet
	err = newStrictDecoder(data).Decode(&dst.ArrayOfIndicatorGroupSet)
	if err == nil {
		jsonArrayOfIndicatorGroupSet, _ := json.Marshal(dst.ArrayOfIndicatorGroupSet)
		if string(jsonArrayOfIndicatorGroupSet) == "{}" { // empty struct
			dst.ArrayOfIndicatorGroupSet = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfIndicatorGroupSet = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		dst.ArrayOfIndicatorGroupSet = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	if src.IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return json.Marshal(&src.IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	}

	if src.ArrayOfIndicatorGroupSet != nil {
		return json.Marshal(&src.ArrayOfIndicatorGroupSet)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return obj.IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	}

	if obj.ArrayOfIndicatorGroupSet != nil {
		return obj.ArrayOfIndicatorGroupSet
	}

	// all schemas are nil
	return nil
}

type NullableIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response struct {
	value *IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response
	isSet bool
}

func (v NullableIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) Get() *IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return v.value
}

func (v *NullableIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) Set(val *IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response(val *IndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) *NullableIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return &NullableIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response{value: val, isSet: true}
}

func (v NullableIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
