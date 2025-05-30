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

// OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response - struct for OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response
type OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response struct {
	OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf *OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	ArrayOfOrganisationUnitGroupSet                                                 *[]OrganisationUnitGroupSet
}

// OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf wrapped in OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response
func OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response(v *OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response{
		OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf: v,
	}
}

// []OrganisationUnitGroupSetAsOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns []OrganisationUnitGroupSet wrapped in OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response
func ArrayOfOrganisationUnitGroupSetAsOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response(v *[]OrganisationUnitGroupSet) OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response{
		ArrayOfOrganisationUnitGroupSet: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	if err == nil {
		jsonOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf, _ := json.Marshal(dst.OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
		if string(jsonOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) == "{}" { // empty struct
			dst.OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfOrganisationUnitGroupSet
	err = newStrictDecoder(data).Decode(&dst.ArrayOfOrganisationUnitGroupSet)
	if err == nil {
		jsonArrayOfOrganisationUnitGroupSet, _ := json.Marshal(dst.ArrayOfOrganisationUnitGroupSet)
		if string(jsonArrayOfOrganisationUnitGroupSet) == "{}" { // empty struct
			dst.ArrayOfOrganisationUnitGroupSet = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfOrganisationUnitGroupSet = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		dst.ArrayOfOrganisationUnitGroupSet = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	if src.OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return json.Marshal(&src.OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	}

	if src.ArrayOfOrganisationUnitGroupSet != nil {
		return json.Marshal(&src.ArrayOfOrganisationUnitGroupSet)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return obj.OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	}

	if obj.ArrayOfOrganisationUnitGroupSet != nil {
		return obj.ArrayOfOrganisationUnitGroupSet
	}

	// all schemas are nil
	return nil
}

type NullableOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response struct {
	value *OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response
	isSet bool
}

func (v NullableOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) Get() *OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return v.value
}

func (v *NullableOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) Set(val *OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response(val *OrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) *NullableOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return &NullableOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response{value: val, isSet: true}
}

func (v NullableOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisationUnitGroupSetGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
