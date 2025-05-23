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

// RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response - struct for RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response
type RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response struct {
	RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	ArrayOfRelationshipType                                                 *[]RelationshipType
}

// RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf wrapped in RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response
func RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response(v *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response {
	return RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response{
		RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf: v,
	}
}

// []RelationshipTypeAsRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns []RelationshipType wrapped in RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response
func ArrayOfRelationshipTypeAsRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response(v *[]RelationshipType) RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response {
	return RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response{
		ArrayOfRelationshipType: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	if err == nil {
		jsonRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf, _ := json.Marshal(dst.RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
		if string(jsonRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) == "{}" { // empty struct
			dst.RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfRelationshipType
	err = newStrictDecoder(data).Decode(&dst.ArrayOfRelationshipType)
	if err == nil {
		jsonArrayOfRelationshipType, _ := json.Marshal(dst.ArrayOfRelationshipType)
		if string(jsonArrayOfRelationshipType) == "{}" { // empty struct
			dst.ArrayOfRelationshipType = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfRelationshipType = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		dst.ArrayOfRelationshipType = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	if src.RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return json.Marshal(&src.RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	}

	if src.ArrayOfRelationshipType != nil {
		return json.Marshal(&src.ArrayOfRelationshipType)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return obj.RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	}

	if obj.ArrayOfRelationshipType != nil {
		return obj.ArrayOfRelationshipType
	}

	// all schemas are nil
	return nil
}

type NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response struct {
	value *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response
	isSet bool
}

func (v NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) Get() *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response {
	return v.value
}

func (v *NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) Set(val *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response(val *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) *NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response {
	return &NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response{value: val, isSet: true}
}

func (v NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
