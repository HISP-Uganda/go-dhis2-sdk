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

// DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response - struct for DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response
type DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response struct {
	DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf *DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	ArrayOfDataEntryForm                                                 *[]DataEntryForm
}

// DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf wrapped in DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response
func DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response(v *DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response {
	return DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response{
		DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf: v,
	}
}

// []DataEntryFormAsDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns []DataEntryForm wrapped in DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response
func ArrayOfDataEntryFormAsDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response(v *[]DataEntryForm) DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response {
	return DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response{
		ArrayOfDataEntryForm: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	if err == nil {
		jsonDataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf, _ := json.Marshal(dst.DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
		if string(jsonDataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) == "{}" { // empty struct
			dst.DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfDataEntryForm
	err = newStrictDecoder(data).Decode(&dst.ArrayOfDataEntryForm)
	if err == nil {
		jsonArrayOfDataEntryForm, _ := json.Marshal(dst.ArrayOfDataEntryForm)
		if string(jsonArrayOfDataEntryForm) == "{}" { // empty struct
			dst.ArrayOfDataEntryForm = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfDataEntryForm = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		dst.ArrayOfDataEntryForm = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	if src.DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return json.Marshal(&src.DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	}

	if src.ArrayOfDataEntryForm != nil {
		return json.Marshal(&src.ArrayOfDataEntryForm)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return obj.DataEntryFormGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	}

	if obj.ArrayOfDataEntryForm != nil {
		return obj.ArrayOfDataEntryForm
	}

	// all schemas are nil
	return nil
}

type NullableDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response struct {
	value *DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response
	isSet bool
}

func (v NullableDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) Get() *DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response {
	return v.value
}

func (v *NullableDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) Set(val *DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response(val *DataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) *NullableDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response {
	return &NullableDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response{value: val, isSet: true}
}

func (v NullableDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataEntryFormGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
