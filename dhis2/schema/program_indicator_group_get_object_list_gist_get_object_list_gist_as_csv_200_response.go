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

// ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response - struct for ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response
type ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response struct {
	ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf *ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	ArrayOfProgramIndicatorGroup                                                 *[]ProgramIndicatorGroup
}

// ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf wrapped in ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response
func ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response(v *ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response {
	return ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response{
		ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf: v,
	}
}

// []ProgramIndicatorGroupAsProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns []ProgramIndicatorGroup wrapped in ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response
func ArrayOfProgramIndicatorGroupAsProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response(v *[]ProgramIndicatorGroup) ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response {
	return ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response{
		ArrayOfProgramIndicatorGroup: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	if err == nil {
		jsonProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf, _ := json.Marshal(dst.ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
		if string(jsonProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) == "{}" { // empty struct
			dst.ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfProgramIndicatorGroup
	err = newStrictDecoder(data).Decode(&dst.ArrayOfProgramIndicatorGroup)
	if err == nil {
		jsonArrayOfProgramIndicatorGroup, _ := json.Marshal(dst.ArrayOfProgramIndicatorGroup)
		if string(jsonArrayOfProgramIndicatorGroup) == "{}" { // empty struct
			dst.ArrayOfProgramIndicatorGroup = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfProgramIndicatorGroup = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		dst.ArrayOfProgramIndicatorGroup = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	if src.ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return json.Marshal(&src.ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	}

	if src.ArrayOfProgramIndicatorGroup != nil {
		return json.Marshal(&src.ArrayOfProgramIndicatorGroup)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return obj.ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	}

	if obj.ArrayOfProgramIndicatorGroup != nil {
		return obj.ArrayOfProgramIndicatorGroup
	}

	// all schemas are nil
	return nil
}

type NullableProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response struct {
	value *ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response
	isSet bool
}

func (v NullableProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) Get() *ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response {
	return v.value
}

func (v *NullableProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) Set(val *ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response(val *ProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) *NullableProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response {
	return &NullableProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response{value: val, isSet: true}
}

func (v NullableProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramIndicatorGroupGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
