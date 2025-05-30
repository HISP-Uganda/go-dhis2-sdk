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

// LegendSetGetObjectListGistGetObjectListGistAsCsv200Response - struct for LegendSetGetObjectListGistGetObjectListGistAsCsv200Response
type LegendSetGetObjectListGistGetObjectListGistAsCsv200Response struct {
	LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf *LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	ArrayOfLegendSet                                                 *[]LegendSet
}

// LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsLegendSetGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf wrapped in LegendSetGetObjectListGistGetObjectListGistAsCsv200Response
func LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsLegendSetGetObjectListGistGetObjectListGistAsCsv200Response(v *LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) LegendSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return LegendSetGetObjectListGistGetObjectListGistAsCsv200Response{
		LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf: v,
	}
}

// []LegendSetAsLegendSetGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns []LegendSet wrapped in LegendSetGetObjectListGistGetObjectListGistAsCsv200Response
func ArrayOfLegendSetAsLegendSetGetObjectListGistGetObjectListGistAsCsv200Response(v *[]LegendSet) LegendSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return LegendSetGetObjectListGistGetObjectListGistAsCsv200Response{
		ArrayOfLegendSet: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LegendSetGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	if err == nil {
		jsonLegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf, _ := json.Marshal(dst.LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
		if string(jsonLegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) == "{}" { // empty struct
			dst.LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfLegendSet
	err = newStrictDecoder(data).Decode(&dst.ArrayOfLegendSet)
	if err == nil {
		jsonArrayOfLegendSet, _ := json.Marshal(dst.ArrayOfLegendSet)
		if string(jsonArrayOfLegendSet) == "{}" { // empty struct
			dst.ArrayOfLegendSet = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfLegendSet = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		dst.ArrayOfLegendSet = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LegendSetGetObjectListGistGetObjectListGistAsCsv200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LegendSetGetObjectListGistGetObjectListGistAsCsv200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LegendSetGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	if src.LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return json.Marshal(&src.LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	}

	if src.ArrayOfLegendSet != nil {
		return json.Marshal(&src.ArrayOfLegendSet)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LegendSetGetObjectListGistGetObjectListGistAsCsv200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return obj.LegendSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	}

	if obj.ArrayOfLegendSet != nil {
		return obj.ArrayOfLegendSet
	}

	// all schemas are nil
	return nil
}

type NullableLegendSetGetObjectListGistGetObjectListGistAsCsv200Response struct {
	value *LegendSetGetObjectListGistGetObjectListGistAsCsv200Response
	isSet bool
}

func (v NullableLegendSetGetObjectListGistGetObjectListGistAsCsv200Response) Get() *LegendSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return v.value
}

func (v *NullableLegendSetGetObjectListGistGetObjectListGistAsCsv200Response) Set(val *LegendSetGetObjectListGistGetObjectListGistAsCsv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableLegendSetGetObjectListGistGetObjectListGistAsCsv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableLegendSetGetObjectListGistGetObjectListGistAsCsv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegendSetGetObjectListGistGetObjectListGistAsCsv200Response(val *LegendSetGetObjectListGistGetObjectListGistAsCsv200Response) *NullableLegendSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return &NullableLegendSetGetObjectListGistGetObjectListGistAsCsv200Response{value: val, isSet: true}
}

func (v NullableLegendSetGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegendSetGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
